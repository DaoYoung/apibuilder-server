package main

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/handler"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"apibuilder-server/app"
	"apibuilder-server/handler/endpoint"
)


func main() {
	if err := app.InitConfig("config.yml"); err != nil {
		log.Fatal(err)
	}
	if err := app.InitDb(); err != nil {
		log.Fatal(err)
	}
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(endpoint.Recovery())
	handler.Serve(r)
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}