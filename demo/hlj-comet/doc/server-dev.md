# 服务器开发

## Comet服务
当客户端上线时，comet服务会广播用户上线通知。其他服务也可通过在线用户查询消息，获取当前在线的用户。
+ 用户上线广播消息(cluster.comet.peer-join)
```json
{
  "service": "comet",
  "subject": "peer-join",
  "payload": null,
  "remote":  "{peer}"
}
```
+ 在线用户查询消息(cluster.comet.peer-query)
```json
{
  "service": "comet",
  "subject": "peer-query",
  "payload": {
    "{field}": "{value}"
  }
}
```
+ 在线用户消息(cluster.comet.peer-online)
```json
{
  "service": "comet",
  "subject": "peer-online",
  "payload": null,
  "remote":  "{peer}"
}
```
Comet服务还支持对在线用户进行消息订阅和取消
+ 给在线用户添加订阅(cluster.comet.subscribe)
```json
{
  "service": "comet",
  "subject": "subscribe",
  "payload": {
    "service": "{service}",
    "subject": "{subject}",
    "remote": {
      "{field}": "{value}"
    }
  }
}
```
+ 给在线用户取消订阅(cluster.comet.unsubscribe)
```json
{
  "service": "comet",
  "subject": "unsubscribe",
  "payload": {
    "service": "{service}",
    "subject": "{subject}",
    "remote": {
      "{field}": "{value}"
    }
  }
}
```

## 服务开发
每一个服务都需支持`subscribe`和`unsubscribe`类型消息。
当客户端订阅或取消服务时，相应的服务需接收服务变更消息，然后给comet服务发送服务变更消息。只有这样客户端才能成功接收或取消该服务的消息。
+ 订阅服务消息(cluster.{service}.subscribe)
```json
{
  "service": "{service}",
  "subject": "subscribe",
  "payload": {
    "service": "{service}",
    "subject": "{subject}",
    "remote": "{remote}"
  }
}
```
+ 取消订阅消息(cluster.{service}.unsubscribe)
```json
{
  "service": "{service}",
  "subject": "unsubscribe",
  "payload": {
    "service": "{service}",
    "subject": "{subject}",
    "remote": "{remote}"
  }
}
```
