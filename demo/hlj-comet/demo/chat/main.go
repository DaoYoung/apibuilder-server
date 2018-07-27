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
	ServiceName = "chat"
	SubjectMsg  = "msg"

	PeerJoinQueue   = "peer_join"
	SubjectMsgQueue = "msg"
)

type Message struct {
	Id       int    `json:"id"`
	Content  string `json:"content"`
	UserId   int    `json:"user_id"`
	ToUserId int    `json:"to_user_id"`
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
		protocol.MakeClusterRoute("comet", "peer-join"),
		PeerJoinQueue,
		func(route string, raw protocol.Message) {
			remoteId := raw.Remote.Get("id").Int()
			if remoteId <= 0 {
				fmt.Printf("remote id is invalid: %d", remoteId)
				return
			}

			remote := protocol.Values{}
			remote.Set("id", raw.Remote.Get("id").String())
			msg := protocol.Message{
				Service: "comet",
				Subject: "subscribe",
				Payload: protocol.Message{
					Service: ServiceName,
					Subject: ServiceName + "." + SubjectMsg + "." + strconv.Itoa(remoteId),
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
		protocol.MakeClusterRoute(ServiceName, SubjectMsg),
		SubjectMsgQueue,
		func(route string, msg protocol.Message) {
			// 只是中转一下
			chatMsg := Message{}
			err := msg.LoadPayload(&chatMsg)
			if err != nil {
				log.Printf("%v", err.Error())
				return
			}
			msger.Publish(protocol.MakeRoute(ServiceName, SubjectMsg+"."+strconv.Itoa(chatMsg.UserId)), msg)
			msger.Publish(protocol.MakeRoute(ServiceName, SubjectMsg+"."+strconv.Itoa(chatMsg.ToUserId)), msg)
		},
	)
	time.Sleep(time.Hour)
}
