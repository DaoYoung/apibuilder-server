package main

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/handler"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"apibuilder-server/app"
)


func main() {
	if err := app.InitConfig("config.yml"); err != nil {
		log.Fatal(err)
	}
	if err := app.InitDb(); err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	handler.Serve(r)
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}