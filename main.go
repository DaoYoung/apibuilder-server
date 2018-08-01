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
	//// Disable Console Color, you don't need console color when writing the logs to file.
	//gin.DisableConsoleColor()
	//
	//// Logging to a file.
	//f, _ := os.Create("gin.log")
	////gin.DefaultWriter = io.MultiWriter(f)
	//
	//// Use the following code if you need to write the logs to file and console at the same time.
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(endpoint.CatchErrors())
	handler.Serve(r)
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}