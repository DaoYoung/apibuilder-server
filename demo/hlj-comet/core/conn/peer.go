// WebSocket连接节点
//
// liamylian
// 2018/03/09

package conn

import (
	"golang.org/x/net/websocket"
	"hlj-comet/core/protocol"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
	"github.com/nats-io/nats"
	"github.com/satori/go.uuid"
)

const (
	ServiceWs   = "ws"
	SubjectPing = "ping"
	SubjectPong = "pong"
	SubjectAuth = "auth"
)

// 连接节点选项
type PeerOption struct {
	BufSize      int
	PingTry      int
	PingInterval time.Duration
	ReadDeadLine time.Duration
}

// 连接节点
type Peer struct {
	Uid           string
	ws            *websocket.Conn
	in            chan *protocol.Message
	out           chan<- *protocol.Message
	destroyed     chan struct{}
	Remote        protocol.Values `json:"remote"`
	option        PeerOption
	pings         int
	pongs         int
	subscriptions []*nats.Subscription
	once          *sync.Once
}

// 新建连接节点
func NewPeer(ws *websocket.Conn, option PeerOption) (*Peer) {
	if option.BufSize <= 0 {
		option.BufSize = 1
	}
	if option.PingTry <= 0 {
		option.PingTry = 1
	}
	if option.PingInterval <= 0 {
		option.PingInterval = time.Minute
	}
	if option.ReadDeadLine < 0 {
		option.ReadDeadLine = time.Minute
	}

	var peer = &Peer{
		Uid:       uuid.Must(uuid.NewV4()).String(),
		ws:        ws,
		in:        make(chan *protocol.Message, option.BufSize),
		destroyed: make(chan struct{}),
		once:      new(sync.Once),
		option:    option,
	}
	peer.Remote.Set("uid", peer.Uid)

	return peer
}

// 从节点接收数据，外部的out管道不能在节点销毁前关闭
func (p *Peer) Receive(out chan<- *protocol.Message) error {
	if p.out != nil {
		return errors.New("peer's out chan is occupied")
	}
	p.out = out
	return nil
}

// 给节点发送收据
func (p *Peer) Send(msg *protocol.Message) (err error) {
	defer func() {
		// 当peer.Send()和peer.Destroy()同时调用时，可能会panic
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("recoverd from send: %v", e))
		}
	}()
	if p.IsDestroyed() {
		return errors.New("already off")
	}

	p.in <- msg
	return nil
}

// 运行节点
func (p *Peer) Run() {
	go p.send()
	p.receive()
}

// 销毁节点
func (p *Peer) Destroy() {
	p.once.Do(func() {
		close(p.in)
		close(p.destroyed)
		for _, sub := range p.subscriptions {
			sub.Unsubscribe()
		}
	})
}

// 是否已销毁
func (p *Peer) IsDestroyed() bool {
	select {
	case <-p.destroyed:
		return true
	default:
		return false
	}
}

// 是否销毁信道
func (p *Peer) DestroyedChan() <-chan struct{} {
	return p.destroyed
}

// 添加订阅
func (p *Peer) Subscribe(msger *protocol.Messager, subject string) error {
	if p.IsSubscribed(subject) {
		return nil
	}
	subscription, err := msger.Subscribe(subject, func(route string, msg protocol.Message) {
		p.Send(&msg) // 引用了外层的currentPeer
	})
	if err != nil {
		return err
	}

	p.subscriptions = append(p.subscriptions, subscription)
	return nil
}

// 取消订阅
func (p *Peer) UnSubscribe(subject string) error {
	for i, subscription := range p.subscriptions {
		if subscription.Subject == subject {
			if err := subscription.Unsubscribe(); err != nil {
				return err
			}
			p.subscriptions = append(p.subscriptions[:i], p.subscriptions[i+1:]...)
			return nil
		}
	}

	return nil
}

// 是否已订阅事件
func (p *Peer) IsSubscribed(subject string) bool {
	for _, subscription := range p.subscriptions {
		if subscription.Subject == subject {
			return true
		}
	}

	return false
}

// 接收数据协程
func (p *Peer) receive() {
	defer func() {
		if err := recover(); err != nil {
			// 目前不做任何处理，因为只有out通道被关闭引起panic
		}
		p.Destroy()
	}()

	for {
		if p.IsDestroyed() {
			return
		}

		p.ws.SetReadDeadline(time.Now().Add(time.Minute))
		msg := protocol.Message{}
		err := websocket.JSON.Receive(p.ws, &msg)
		if err != nil {
			if e, ok := err.(net.Error); ok && e.Timeout() {
				continue
			}
			return // WebSocket已关闭
		}

		if msg.Service == ServiceWs {
			switch msg.Subject {
			case SubjectPing:
				err := websocket.JSON.Send(p.ws, protocol.Message{
					Id:          uuid.Must(uuid.NewV4()).String(),
					Service:     ServiceWs,
					Subject:     SubjectPong,
					Payload:     msg.Payload,
					ReplyId:     msg.Id,
					CurrentTime: time.Now().Unix(),
				})
				if err != nil {
					return // WebSocket已关闭
				}
			case SubjectPong:
				p.pongs++
			default:
				p.out <- &msg // 客户端自定义点对点消息
			}
		} else if p.out != nil {
			p.out <- &msg // 如果peer.out被关闭将会panic
		}
	}
}

// 发送数据协程
func (p *Peer) send() {
	ticker := time.NewTicker(p.option.PingInterval)
	defer func() {
		ticker.Stop()
		p.Destroy()
	}()

	for {
		select {
		case <-ticker.C:
			if p.pings-p.pongs > p.option.PingTry {
				return // ping失败
			}

			err := websocket.JSON.Send(p.ws, protocol.Message{
				Id:          uuid.Must(uuid.NewV4()).String(),
				Service:     ServiceWs,
				Subject:     SubjectPing,
				Payload:     time.Now().Unix(),
				CurrentTime: time.Now().Unix(),
			})
			if err != nil {
				return // WebSocket已关闭
			}
			p.pings++
		case msg, ok := <-p.in:
			if !ok {
				return // 管道已关闭
			}
			err := websocket.JSON.Send(p.ws, msg)
			if err != nil {
				return
			}
		case <-p.destroyed:
			return // 节点已销毁
		}
	}
}
