package api

import (
	"net/http"
	"hlj-comet/handler/middleware"
	"golang.org/x/net/websocket"
	"hlj-rest/rest"
	"hlj-comet/core"
	"hlj-comet/core/conn"
	"hlj-comet/app"
	"time"
	"hlj-comet/core/protocol"
)

var json = app.Json

func Ws(w rest.ResponseWriter, r *rest.Request) {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		user := middleware.RemoteRoleUser(r)
		if user == nil {
			websocket.JSON.Send(ws, protocol.Message{
				Service: "ws",
				Subject: conn.SubjectAuth,
				Payload: protocol.AckPayload{
					Code: 401,
					Msg:  "登录失败",
				},
			})
			return
		}

		peer := conn.NewPeer(ws, conn.PeerOption{
			BufSize:      app.Config.Peer.BufSize,
			PingTry:      app.Config.Peer.PingTry,
			PingInterval: time.Duration(app.Config.Peer.PingInterval) * time.Second,
			ReadDeadLine: time.Duration(app.Config.Peer.ReadDeadLine) * time.Second,
		})

		values := protocol.Values{}
		bytes, _ := json.Marshal(user)
		if err := json.Unmarshal(bytes, &values); err == nil {
			peer.Remote.Sets(values.ToMap())
		}

		comet := core.GetComet()
		if err := comet.AddPeer(peer); err != nil {
			websocket.JSON.Send(ws, err.Error())
			return
		}
		peer.Run()
	}).ServeHTTP(w.(http.ResponseWriter), r.Request)
}
