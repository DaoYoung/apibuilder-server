package handler

import (
	"log"
	"hlj-comet/app"
	"net/http"
	"golang.org/x/net/websocket"
	"hlj-rest/rest"
	mid "hlj-rest/middleware"
	"hlj-comet/model"
	"hlj-comet/core/protocol"
	"hlj-comet/handler/api"
	"hlj-comet/handler/envelope"
	"hlj-comet/handler/middleware"
	"strconv"
	"hlj-comet/core/conn"
)

var cors = mid.CorsMiddleware{
	AccessControlExposeHeaders:    []string{"Content-Type"},
	AccessControlAllowCredentials: true,
	AllowedHeaders:                []string{"Origin", "Content-Type", "Accept", "login-as"},
	AllowedMethods:                []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS", "CONNECT", "TRACE"},
	AccessControlMaxAge:           86400,
	OriginValidator: func(origin string, request *rest.Request) bool {
		return true
	},
}

var auth = middleware.HljAuthMiddleware{
	Authenticator: func(userToken string) (interface{}, bool) {
		if userToken == "" {
			return nil, false
		}
		user := model.GetUserByToken(userToken)
		if user == nil {
			return nil, false
		}
		return user, true
	},
}

var authWs = middleware.HljAuthMiddleware{
	Authenticator: auth.Authenticator,
	Unauthenticated: func(w rest.ResponseWriter, r *rest.Request) {
		websocket.Handler(func(ws *websocket.Conn) {
			websocket.JSON.Send(ws, protocol.Message{
				Service: "ws",
				Subject: conn.SubjectAuth,
				Payload: protocol.AckPayload{
					Code: 401,
					Msg:  "登录失败",
				},
			})
			ws.Close()
		}).ServeHTTP(w.(http.ResponseWriter), r.Request)
	},
}

func Serve() {
	routes := getRoutes()
	if router, err := rest.MakeRouter(routes...); err != nil {
		log.Fatal(err)
	} else {
		rest.SetDefaultEnvelope(&envelope.WeddingEnvelope{})
		http.Handle("/", http.FileServer(http.Dir("public")))
		http.Handle("/api/", http.StripPrefix("/api", router.Use(&cors)))

		if app.Config.HttpsPemPath != "" && app.Config.HttpsKeyPath != "" {
			go func() {
				log.Fatal(http.ListenAndServeTLS(
					":"+strconv.FormatUint(app.Config.HttpsPort, 10),
					app.Config.HttpsPemPath,
					app.Config.HttpsKeyPath,
					nil),
				)
			}()
		}
		log.Fatal(http.ListenAndServe(":"+strconv.FormatUint(app.Config.Port, 10), nil))
	}
}

func getRoutes() []*rest.Route {

	return []*rest.Route{
		rest.GetFunc("/peers", api.GetPeers),
		rest.GetFunc("/ws", authWs.MiddlewareFunc(api.Ws)),
		rest.GetFunc("/stats", api.GetStat),
	}
}
