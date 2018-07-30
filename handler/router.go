package handler

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/handler/endpoint"
	)

func Serve(engine *gin.Engine) {

	mkdoc := engine.Group("/doc")
	{
		mkdoc.GET("/api/:id", endpoint.ApiAction("info"))
		mkdoc.GET("/api/", endpoint.ApiAction("list"))
		mkdoc.PUT("/api/:id", endpoint.ApiAction("update"))
		mkdoc.POST("/api/", endpoint.ApiAction("create"))
		mkdoc.DELETE("/api/:id", endpoint.ApiAction("delete"))

		mkdoc.POST("/api/:id/publish", endpoint.PublishApi)//发布接口
		mkdoc.POST("/api/:id/commit", endpoint.CommitApi)//在已发布接口基础上修改结构
		mkdoc.POST("/api/:id/rebuild", endpoint.RebuildApi)//重构
		mkdoc.POST("/api/:id/note", endpoint.NoteApi)//注释接口参数
		mkdoc.POST("/api/:id/render", endpoint.RenderApi)//对接
		//mkdoc.POST("/api/test", endpoint.NoteApi)//测试
		//mkdoc.POST("/api/handover", endpoint.NoteApi)//离职交接责任人

		mkdoc.POST("/task/:id/translate", endpoint.NoteApi)//变更任务对接人，进度
		mkdoc.POST("/task/:id/test", endpoint.NoteApi)//测试

		mkdoc.GET("/module/", endpoint.ModuleAction("list"))
		mkdoc.GET("/module/:id", endpoint.ModuleAction("info"))
		mkdoc.PUT("/module/:id", endpoint.ModuleAction("update"))
		mkdoc.POST("/module/", endpoint.ModuleAction("create"))
		mkdoc.DELETE("/module/:id", endpoint.ModuleAction("delete"))
	}

}