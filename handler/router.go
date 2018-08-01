package handler

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/handler/endpoint"
)

func Serve(engine *gin.Engine) {
	mkdoc := engine.Group("/doc")
	{
		mkdoc.POST("/api/:id/publish", endpoint.PublishApi) //发布接口
		mkdoc.POST("/api/:id/commit", endpoint.CommitApi)   //在已发布接口基础上修改结构
		mkdoc.POST("/api/:id/rebuild", endpoint.RebuildApi) //重构
		mkdoc.POST("/api/:id/note", endpoint.NoteApi)       //注释接口参数
		//mkdoc.POST("/api/:id/render", endpoint.RenderApi)//对接
		//mkdoc.POST("/api/test", endpoint.NoteApi)//测试
		//mkdoc.POST("/api/handover", endpoint.NoteApi)//离职交接责任人
		//mkdoc.POST("/task/:id/translate", endpoint.NoteApi)//变更任务对接人，进度
		//mkdoc.POST("/task/:id/test", endpoint.NoteApi)//测试
		curdRoutes(mkdoc, "api", endpoint.ApiController{})
		curdRoutes(mkdoc, "module", endpoint.ModuleController{})
	}
}

func curdRoutes(group *gin.RouterGroup, resourceName string, controller endpoint.ControllerInterface) {
	group.GET("/"+resourceName+"/", controller.CrudService("list"))
	group.GET("/"+resourceName+"/:id", controller.CrudService("info"))
	group.PUT("/"+resourceName+"/:id", controller.CrudService("update"))
	group.POST("/"+resourceName+"/", controller.CrudService("create"))
	group.DELETE("/"+resourceName+"/:id", controller.CrudService("delete"))
}
