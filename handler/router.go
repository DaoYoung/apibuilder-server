package handler

import (
	"log"
	"apibuilder-server/demo/hlj-comet/handler/envelope"
	"net/http"
	"apibuilder-server/demo/hlj-comet/app"
	"strconv"
	"apibuilder-server/demo/hlj-comet/handler/api"
	"apibuilder-server/demo/hlj-rest/hlj-rest/rest"
	"github.com/gin-gonic/gin"
)

func Serve(engine *gin.Engine) {
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
		rest.GetFunc("/*Action/*Fun", api.User),
	}
}