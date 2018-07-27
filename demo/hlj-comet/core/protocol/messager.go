// Nats消息处理封装
//
// liamylian
// 2018/03/09

package protocol

import (
	"github.com/satori/go.uuid"
	"github.com/nats-io/nats"
	"log"
	"time"
)

// 消息处理回调
type SubscribeHandler func(route string, message Message)

// 消息处理对象
type Messager struct {
	conn *nats.EncodedConn
}

// 新建消息处理对象
func NewMessager(conn *nats.EncodedConn) *Messager {
	return &Messager{
		conn: conn,
	}
}

// 发布消息
func (m *Messager) Publish(route string, message Message) error {
	message.CurrentTime = time.Now().Unix()
	return m.conn.Publish(route, message)
}

// 订阅消息
func (m *Messager) Subscribe(route string, handler SubscribeHandler) (*nats.Subscription, error) {
	return m.conn.Subscribe(route, func(m *nats.Msg) {
		msg := Message{}
		err := json.Unmarshal(m.Data, &msg)
		if err != nil {
			log.Printf("protocol.Subscribe error: %s", err.Error())
			return
		}
		handler(m.Subject, msg)
	})
}

// 订阅队列消息（相同消息，同一队列只有一个会收到）
func (m *Messager) QueueSubscribe(route, queue string, handler SubscribeHandler) (*nats.Subscription, error) {
	return m.conn.QueueSubscribe(route, queue, func(m *nats.Msg) {
		msg := Message{}
		err := json.Unmarshal(m.Data, &msg)
		if err != nil {
			log.Printf("protocol.Subscribe error: %s", err.Error())
			return
		}
		handler(m.Subject, msg)
	})
}

// 发送请求消息，需自行订阅reply来接收回复
func (m *Messager) SendCall(route, reply string, message Message) error {
	message.CurrentTime = time.Now().Unix()
	return m.conn.PublishRequest(route, reply, message)
}

// 发送请求消息，并返回第一条回复消息。如果有多条回复的情况，使用SendCall，并订阅reply。
func (m *Messager) Call(route string, msg Message, timeout time.Duration) (Message, error) {
	msg.CurrentTime = time.Now().Unix()
	resp := Message{}
	err := m.conn.Request(route, msg, &resp, timeout)

	return resp, err
}

// 处理请求消息
func (m *Messager) Handle(route string, handler Handler) error {
	m.conn.Subscribe(route, func(msg *nats.Msg) {
		req := Message{}
		err := json.Unmarshal(msg.Data, &req)
		if err != nil {
			log.Printf("protocol.Subscribe error: %s", err.Error())
			return
		}
		resp := &response{
			reply:   msg.Reply,
			msger:   m,
			service: req.Service,
			subject: req.Subject,
			replyId: req.Id,
		}
		handler.Serve(resp, &req)
	})

	return nil
}

// 处理队列请求消息（相同消息，同一队列只有一个会收到）
func (m *Messager) QueueHandle(route, queue string, handler Handler) error {
	m.conn.QueueSubscribe(route, queue, func(msg *nats.Msg) {
		req := Message{}
		err := json.Unmarshal(msg.Data, &req)
		if err != nil {
			log.Printf("protocol.Subscribe error: %s", err.Error())
			return
		}
		resp := &response{
			reply:   msg.Reply,
			msger:   m,
			service: req.Service,
			subject: req.Subject,
			replyId: req.Id,
		}
		handler.Serve(resp, &req)
	})

	return nil
}

// Handler快捷调用
func (m *Messager) HandleFunc(route string, handler HandlerFunc) error {
	return m.Handle(route, handler)
}

// Handler快捷调用
func (m *Messager) QueueHandleFunc(route, queue string, handler HandlerFunc) error {
	return m.QueueHandle(route, queue, handler)
}

// 响应实现类
type response struct {
	msger   *Messager
	reply   string
	service string
	subject string
	replyId string
}

func (r *response) Set(service, subject string) {
	r.service = service
	r.subject = subject
}

func (r *response) Write(data interface{}) error {
	msg := Message{
		Id:          uuid.Must(uuid.NewV4()).String(),
		ReplyId:     r.replyId,
		Service:     r.service,
		Subject:     r.subject + "." + SubjectAckSuffix,
		Payload:     data,
		CurrentTime: time.Now().Unix(),
	}

	r.msger.Publish(r.reply, msg)
	return nil
}
