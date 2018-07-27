package main

import (
	"log"
	"hlj-comet/app"
	"hlj-comet/core/protocol"
	"time"
	"strconv"
	"fmt"
)

const (
	ServiceName = "chatroom"
	SubjectRoom = "room"

	SubjectJoinRoom = "join-room"
	SubjectJoinQueue   = "peer-join"
	SubjectMsgQueue = "msg"
)

type Message struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
	RoomId  int    `json:"room_id"`
}

func main() {
	if err := app.InitConfig("config.yml"); err != nil {
		log.Fatal(err)
	}
	if err := app.InitEc(); err != nil {
		log.Fatal(err)
	}
	msger := protocol.NewMessager(app.Ec)

	msger.QueueSubscribe(
		protocol.MakeClusterRoute(ServiceName, SubjectJoinRoom),
		SubjectJoinQueue,
		func(route string, req protocol.Message) {
			var roomId int
			if err := req.LoadPayload(&roomId); err != nil || roomId <= 0 {
				fmt.Printf("invalid args: %v", req.Payload)
				return
			}

			remote := protocol.Values{}
			remote.Set("id", req.Remote.Get("id").String())
			msg := protocol.Message{
				Service: "comet",
				Subject: "subscribe",
				Payload: protocol.Message{
					Service: ServiceName,
					Subject: ServiceName + "." + SubjectRoom + "." +strconv.Itoa(roomId),
					Remote:  remote,
				},
			}
			resp, err := msger.Call(protocol.MakeClusterRoute("comet", "subscribe"), msg, 5*time.Second)
			if err != nil {
				fmt.Printf("subscribe failed: %v", err.Error())
				return
			}

			ack := protocol.AckPayload{}
			resp.LoadPayload(&ack)
			if ack.Code != 0 {
				fmt.Printf("subscribe failed: %s, %v", ack.Msg, ack.Data)
				return
			}
		},
	)

	msger.QueueSubscribe(
		protocol.MakeClusterRoute(ServiceName, SubjectRoom),
		SubjectMsgQueue,
		func(route string, msg protocol.Message) {
			// 只是中转一下
			roomMsg := Message{}
			msg.LoadPayload(&roomMsg)
			room := strconv.Itoa(roomMsg.RoomId)
			msger.Publish(protocol.MakeRoute(ServiceName, SubjectRoom+"."+room), msg)
		},
	)
	time.Sleep(time.Hour)
}
