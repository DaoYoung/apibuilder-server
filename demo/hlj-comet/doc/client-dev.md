# 客户端开发

## WebSocket服务
客户端收到服务器ping消息，必须回复pong消息(payload为ping消息的payload)，否则会被下线。客户端也可主动ping服务器，服务器也会回复pong消息。
+ Ping消息
```json
{
  "service": "ws",
  "subject": "ping",
  "payload": 1520991260
}
```
+ Pong消息
```json
{
  "service": "ws",
  "subject": "pong",
  "payload": 1520991260
}
```

## Comet服务
+ 修改连接信息
```json
{
  "service": "comet",
  "subject": "set-remote",
  "payload": {
    "{key}": "{value}"
  }
}
```

## 其他服务
各类服务的消息均会使用以下格式进行传递。客户端接收消息时需要根据service和subject进行分发处理；给对应服务发送消息时也需带上service和subject。
+ 服务消息
```json
{
  "service": "{service}",
  "subject": "{subject}",
  "payload": "{payload}"
}
```
+ 服务回复消息
```json
{
  "service": "{service}",
  "subject": "{service}",
  "payload": {
    "code": 0,
    "msg": "",
    "data": "{data}"
  }
}
```