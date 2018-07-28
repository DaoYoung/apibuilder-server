package handler

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/handler/endpoint"
	)

func Serve(engine *gin.Engine) {

	mkdoc := engine.Group("/doc")
	{
		mkdoc.GET("/api/:id", endpoint.GetApiInfo)
		mkdoc.GET("/api/", endpoint.GetApiList)
		mkdoc.PUT("/api/:id", endpoint.UpdateApi)
		mkdoc.POST("/api/", endpoint.CreateApi)
		mkdoc.DELETE("/api/:id", endpoint.DeleteApi)
	}

}