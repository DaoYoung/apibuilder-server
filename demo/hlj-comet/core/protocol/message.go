// 消息定义
//
// liamylian
// 2018/03/09

package protocol

const (
	ClusterPrefix = "cluster" // 服务器间消息前缀

	SubjectAckSuffix = "ack" // 通用响应后缀
)

// 生成集群路由（消息只发布给服务器）
func MakeClusterRoute(service, subject string) string {
	return ClusterPrefix + "." + service + "." + subject
}

// 生成路由
func MakeRoute(service, subject string) string {
	return service + "." + subject
}

// 消息格式
type Message struct {
	Id          string      `json:"id,omitempty"`      // 消息id
	Service     string      `json:"service"`           // 服务
	Subject     string      `json:"subject"`           // 内容
	Payload     interface{} `json:"payload"`           // 数据
	Remote      Values      `json:"remote,omitempty"`  // 发送者数据
	ReplyId     string      `json:"replyId,omitempty"` // 回复的消息id
	CurrentTime int64       `json:"current_time"`      // 服务器时间
}

// 响应数据
type AckPayload struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// 从字节流生成消息
func FromBytes(bytes []byte) (Message, error) {
	var msg Message
	err := json.Unmarshal(bytes, &msg)

	return msg, err
}

// 将消息数据载入对象
func (m *Message) LoadPayload(dst interface{}) error {
	bytes, err := json.Marshal(m.Payload)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, dst)
}

// 响应接口
type Response interface {
	Set(service, subject string)
	Write(data interface{}) error
}

// 处理接口
type Handler interface {
	Serve(responseWriter Response, request *Message)
}

// 闭包处理实现
type HandlerFunc func(responseWriter Response, request *Message)

func (h HandlerFunc) Serve(responseWriter Response, request *Message) {
	h(responseWriter, request)
}
