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

		mkdoc.GET("/module/", endpoint.ModuleAction("list"))
		mkdoc.GET("/module/:id", endpoint.ModuleAction("info"))
		mkdoc.PUT("/module/:id", endpoint.ModuleAction("update"))
		mkdoc.POST("/module/", endpoint.ModuleAction("create"))
		mkdoc.DELETE("/module/:id", endpoint.ModuleAction("delete"))
	}

}