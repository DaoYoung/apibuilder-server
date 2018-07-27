// Comet服务
//
// liamylian
// 2018/03/09

package core

import (
	"hlj-comet/core/conn"
	"hlj-comet/core/protocol"
	"time"
	"github.com/nats-io/nats"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"errors"
	"fmt"
)

const (
	SubjectSubscribe   = "subscribe"
	SubjectUnsubscribe = "unsubscribe"
	SubjectSetRemote   = "set-remote"
	SubjectPeerJoin    = "peer-join"
	SubjectPeerQuery   = "peer-query"
	SubjectPeerOnline  = "peer-online"

	ErrorCodeBadSubscribe      = 200
	ErrorCodeSubscribeFailed   = 201
	ErrorCodeBadUnsubscribe    = 210
	ErrorCodeUnsubscribeFailed = 211

	ErrorCodeUnsupportSubject = 1000
)

var comet *Comet

func InitComet(ec *nats.EncodedConn, cometOption CometOption) {
	msger := protocol.NewMessager(ec)
	comet = NewComet("comet", msger, cometOption)
}

type CometOption struct {
	MsgBufSize   int
	MsgTrunkSize int
	GcInterval   int
}

type Comet struct {
	Pool    *conn.Pool
	service string
	msger   *protocol.Messager
	option  CometOption
}

func GetComet() *Comet {
	return comet
}

func NewComet(serviceName string, msger *protocol.Messager, cometOption CometOption) *Comet {
	pool := conn.NewPool(serviceName + "pool")
	return &Comet{
		Pool:    pool,
		service: serviceName,
		msger:   msger,
		option:  cometOption,
	}
}

func (c *Comet) AddPeer(peer *conn.Peer) error {
	// 广播用户上线消息
	c.msger.Publish(protocol.MakeClusterRoute(c.service, SubjectPeerJoin), protocol.Message{
		Id:      uuid.Must(uuid.NewV4()).String(),
		Service: c.service,
		Subject: SubjectPeerJoin,
		Remote:  peer.Remote,
	})

	// 给用户自动订阅comet无中转消息
	peer.Subscribe(c.msger, c.service+"."+peer.Uid+".*")

	// todo to-remove
	peer.Subscribe(c.msger, "css.user."+peer.Remote.Get("id").String())

	// 启动收发协程
	go func() {
		peerOut := make(chan *protocol.Message, c.option.MsgBufSize)
		peer.Receive(peerOut)
		for {
			select {
			case <-peer.DestroyedChan():
				close(peerOut)
				return
			case msg := <-peerOut:
				msg.Remote = peer.Remote
				switch msg.Service {
				case c.service:
					if err := c.serveClient(peer, msg); err != nil {
						msg := protocol.Message{
							Service: c.service,
							Subject: msg.Service + "." + protocol.SubjectAckSuffix,
							Payload: protocol.AckPayload{
								Code: ErrorCodeUnsupportSubject,
								Msg:  err.Error(),
							},
						}
						peer.Send(&msg)
					}
				default:
					c.msger.Publish(protocol.MakeClusterRoute(msg.Service, msg.Subject), *msg) // 转发给相应服务
				}
			}
		}
		close(peerOut)
	}()
	c.Pool.Add(peer)
	return nil
}

func (c *Comet) Run() {
	c.serve()

	// 定时清理连接池
	ticker := time.NewTicker(time.Duration(c.option.GcInterval) * time.Second)
	for {
		select {
		case <-ticker.C:
			c.Pool.Range(func(peer *conn.Peer) bool {
				// do nothing, let p.Range remove destroyed peers
				return true
			})
		}
	}
}

func (c *Comet) serve() {
	// 订阅服务
	c.msger.HandleFunc(
		protocol.MakeClusterRoute(c.service, SubjectSubscribe),
		func(w protocol.Response, r *protocol.Message) {
			// todo to remove
			w.Write(protocol.AckPayload{Data: 0})
			return

			data := protocol.Message{}
			r.LoadPayload(&data)
			subscribed := 0
			for _, peer := range c.Pool.GetPeers(data.Remote.ToMap(), 10) {
				if peer.IsSubscribed(data.Subject) {
					continue // 防止重复订阅
				}
				if err := peer.Subscribe(c.msger, data.Subject); err != nil {
					w.Write(protocol.AckPayload{
						Code: ErrorCodeSubscribeFailed,
						Data: subscribed,
					})
					return
				}
				subscribed++
			}

			w.Write(protocol.AckPayload{Data: subscribed})
		},
	)
	// 取消订阅服务
	c.msger.HandleFunc(
		protocol.MakeClusterRoute(c.service, SubjectUnsubscribe),
		func(w protocol.Response, r *protocol.Message) {
			data := protocol.Message{}
			r.LoadPayload(&data)

			unsubscribed := 0
			for _, peer := range c.Pool.GetPeers(data.Remote.ToMap(), 10) {
				if err := peer.UnSubscribe(data.Subject); err != nil {
					w.Write(protocol.AckPayload{Code: ErrorCodeUnsubscribeFailed, Data: unsubscribed})
					return
				}
				unsubscribed++
			}

			w.Write(protocol.AckPayload{Data: unsubscribed})
			return
		},
	)
	// 在线用户查询服务
	c.msger.HandleFunc(protocol.MakeClusterRoute(c.service, SubjectPeerQuery), func(response protocol.Response, msg *protocol.Message) {

		if msg.Service == c.service && msg.Subject == SubjectPeerQuery {
			peerExtra := protocol.Values{}
			if err := msg.LoadPayload(&peerExtra); err != nil {
				log.Warn("[Comet] SubjectPeerQuery: %s", err.Error())
				return
			}

			peers := c.Pool.GetPeers(peerExtra.ToMap(), 0)
			peersExtra := make([]map[string]string, 0)
			for _, peer := range peers {
				peersExtra = append(peersExtra, peer.Remote.ToMap())
			}

			response.Write(peersExtra)
		}
	})
}

func (c *Comet) serveClient(peer *conn.Peer, msg *protocol.Message) error {
	switch msg.Subject {
	case SubjectSetRemote:
		extras := make(map[string]string)
		if err := msg.LoadPayload(&extras); err != nil {
			return err
		}

		peer.Remote.Sets(extras)
	case SubjectSubscribe:
		subject, ok := msg.Payload.(string)
		if !ok || subject == "" {
			return errors.New(fmt.Sprintf("Bad Subject: %v", subject))
		}

		cometSubject := c.service + "." + subject
		if peer.IsSubscribed(subject) {
			return nil
		}

		return peer.Subscribe(c.msger, cometSubject)
	case SubjectUnsubscribe:
		subject, ok := msg.Payload.(string)
		if !ok || subject == "" {
			return errors.New(fmt.Sprintf("Bad Subject: %v", subject))
		}

		cometSubject := c.service + "." + subject
		return peer.UnSubscribe(cometSubject)
	default:
		c.msger.Publish(protocol.MakeRoute(msg.Service, msg.Subject), *msg) // 无中转发送点对点消息
	}

	return nil
}
