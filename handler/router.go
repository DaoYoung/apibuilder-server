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
		mkdoc.POST("/model/:id/note", endpoint.NoteModel) //注释接口参数
		mkdoc.GET("/model/:id/note", endpoint.NoteModelDetail)
		//mkdoc.POST("/api/:id/render", endpoint.RenderApi)//对接
		//mkdoc.POST("/api/test", endpoint.NoteApi)//测试
		//mkdoc.POST("/api/handover", endpoint.NoteApi)//离职交接责任人
		//mkdoc.POST("/task/:id/translate", endpoint.NoteApi)//变更任务对接人，进度
		//mkdoc.POST("/task/:id/test", endpoint.NoteApi)//测试
		curdRoutes(mkdoc,  endpoint.ApiController{}.Rester())
		curdRoutes(mkdoc,  endpoint.ModuleController{}.Rester())
		curdRoutes(mkdoc,  endpoint.ModelController{}.Rester())
		curdRoutes(mkdoc,  endpoint.ModelMapController{}.Rester(), "list", "create", "update", "delete")
	}


	engine.POST("/auth/login", middleware.AuthMiddleware.LoginHandler)
	engine.POST("/auth/reg", endpoint.UserController{}.Rester().Create)
	auth := engine.Group("/auth")
	auth.Use(middleware.AuthHandlerFunc)
	{
		auth.GET("/refresh_token", middleware.AuthMiddleware.RefreshHandler)
		auth.GET("/profile", endpoint.Profile)
	}

	admin := engine.Group("/admin")
	admin.Use(middleware.AuthHandlerFunc)
	{
		curdRoutes(admin, endpoint.ContainerController{}.Rester())
		curdRoutes(admin,  endpoint.ContainerParamController{}.Rester(), "list", "create", "update", "delete")
		curdRoutes(admin, endpoint.ContainerDeployController{}.Rester(), "list", "create", "update", "delete")
	}
	task := engine.Group("/workflow")
	task.Use(middleware.AuthHandlerFunc)
	{
		curdRoutes(task, endpoint.TaskController{}.Rester())

	}
	//todo user Permission
}

func curdRoutes(group *gin.RouterGroup, controller endpoint.ControllerInterface, actions ...string) {
	path, resourceName, routeId := endpoint.BuildRoute(controller)
	if len(actions) == 0 || helper.Contains(actions, "list") {
		group.GET(path+"/"+resourceName, controller.List)
	}
	if len(actions) == 0 || helper.Contains(actions, "info") {
		group.GET(path+"/"+resourceName+"/:"+routeId, controller.Info)
	}
	if len(actions) == 0 || helper.Contains(actions, "update") {
		group.PUT(path+"/"+resourceName+"/:"+routeId, controller.Update)
	}
	if len(actions) == 0 || helper.Contains(actions, "create") {
		group.POST(path+"/"+resourceName, controller.Create)
	}
	if len(actions) == 0 || helper.Contains(actions, "delete") {
		group.DELETE(path+"/"+resourceName+"/:"+routeId, controller.Delete)
	}
}
