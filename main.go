package main

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/handler"
	"apibuilder-server/handler/endpoint"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"apibuilder-server/app"
	"github.com/appleboy/gin-jwt"
	"apibuilder-server/handler/middleware"
)

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID": claims["id"],
		"text":   "Hello World.",
	})
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}
func main() {
	if err := app.InitConfig("config.yml"); err != nil {
		log.Fatal(err)
	}
	if err := app.InitDb(); err != nil {
		log.Fatal(err)
	}
	middleware.InitJWT()
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