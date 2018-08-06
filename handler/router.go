package handler

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/handler/endpoint"
	"apibuilder-server/helper"
	"apibuilder-server/handler/middleware"
)

func Serve(engine *gin.Engine) {
	mkdoc := engine.Group("/doc")
	{
		mkdoc.POST("/api/:id/publish", endpoint.PublishApi) //发布接口
		mkdoc.POST("/api/:id/commit", endpoint.CommitApi)   //在已发布接口基础上修改结构
		mkdoc.POST("/api/:id/rebuild", endpoint.RebuildApi) //重构
		mkdoc.POST("/api/:id/note", endpoint.NoteApi)       //注释接口参数
		mkdoc.GET("/api/:id/note", endpoint.NoteApiDetail)
		mkdoc.POST("/model/:id/note", endpoint.NoteModel)       //注释接口参数
		mkdoc.GET("/model/:id/note", endpoint.NoteModelDetail)
		//mkdoc.POST("/api/:id/render", endpoint.RenderApi)//对接
		//mkdoc.POST("/api/test", endpoint.NoteApi)//测试
		//mkdoc.POST("/api/handover", endpoint.NoteApi)//离职交接责任人
		//mkdoc.POST("/task/:id/translate", endpoint.NoteApi)//变更任务对接人，进度
		//mkdoc.POST("/task/:id/test", endpoint.NoteApi)//测试
		curdRoutes(mkdoc, "api", endpoint.ApiController{})
		curdRoutes(mkdoc, "module", endpoint.ModuleController{})
		curdRoutes(mkdoc, "model", endpoint.ModelController{})
		curdRoutes(mkdoc, "/model/:id/map", endpoint.ModelMapController{}, "list", "create", "update", "delete")
	}

	engine.POST("/auth/login", middleware.AuthMiddleware.LoginHandler)
	engine.POST("/auth/reg", endpoint.UserController{}.CrudService("create"))
	auth := engine.Group("/auth")
	auth.Use(middleware.AuthMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", middleware.AuthMiddleware.RefreshHandler)
		auth.GET("/profile", endpoint.Profile)
	}
	//todo user Permission
}

func curdRoutes(group *gin.RouterGroup, resourceName string, controller endpoint.ControllerInterface, actions ...string) {
	if len(actions) == 0 || helper.Contains(actions, "list") {
		group.GET("/"+resourceName+"/", controller.CrudService("list"))
	}
	if len(actions) == 0 || helper.Contains(actions, "info") {
		group.GET("/"+resourceName+"/:id", controller.CrudService("info"))
	}
	if len(actions) == 0 || helper.Contains(actions, "update") {
		group.PUT("/"+resourceName+"/:id", controller.CrudService("update"))
	}
	if len(actions) == 0 || helper.Contains(actions, "create") {
		group.POST("/"+resourceName+"/", controller.CrudService("create"))
	}
	if len(actions) == 0 || helper.Contains(actions, "delete") {
		group.DELETE("/"+resourceName+"/:id", controller.CrudService("delete"))
	}
}
