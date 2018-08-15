package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"apibuilder-server/helper"
	"apibuilder-server/app"
	"reflect"
	"strings"
)

var Com ControllerInterface

type ControllerInterface interface {
	Update(c *gin.Context)
	Create(c *gin.Context)
	Info(c *gin.Context)
	List(c *gin.Context)
	Delete(c *gin.Context)
	IsRestRoutePk() bool
	RouteName() string
	ParentNode() ControllerInterface
}

type Controller struct {
	ParentController ControllerInterface
	Rester           RestInterface
	RestModel        func() model.ResourceInterface
	RestModelSlice   func() interface{} //https://golang.org/doc/faq#convert_slice_of_interface
	*EmptyRest
}

func (this *Controller) ParentNode() ControllerInterface {
	return this.ParentController
}
func (this *Controller) IsRestRoutePk() bool {
	return false
}
func (this *Controller) RouteName() string {
	obj := this.RestModel()
	f := reflect.TypeOf(obj)
	if f.Kind() == reflect.Ptr {
		f = f.Elem()
	}
	return strings.ToLower(f.Name())
}

func (this *Controller) Create(c *gin.Context) {
	obj := this.RestModel()
	err := c.BindJSON(obj)
	if err != nil {
		panic(JsonTypeError(err))
	}
	this.Rester.BeforeRest(c, obj)
	this.Rester.BeforeCreate(c, obj)
	info := model.Create(obj)
	this.Rester.AfterCreate(c, obj)
	this.Rester.AfterRest(c, obj)
	helper.ReturnSuccess(c, http.StatusCreated, info)
}

func (this *Controller) Info(c *gin.Context) {
	obj := this.RestModel()
	id, _ := strconv.Atoi(c.Param(GetRouteID(this.Rester)))
	info := model.ByID(obj, id)
	helper.ReturnSuccess(c, http.StatusOK, info)
}

func (this *Controller) List(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		panic(err)
	}
	obj := this.RestModelSlice()
	condition := this.Rester.ListCondition(c)
	condition = helper.MapUrlQuery(condition, c.Request.URL.Query(), this.RestModel())
	model.FindListWhereMap(obj, condition, "id desc", page, app.Config.PerPage)
	helper.ReturnSuccess(c, http.StatusOK, obj)
}

func (this *Controller) Update(c *gin.Context) {
	obj := this.RestModel()
	err := c.BindJSON(obj)
	if err != nil {
		panic(JsonTypeError(err))
	}
	this.Rester.BeforeRest(c, obj)
	this.Rester.BeforeUpdate(c, obj)
	condition := this.Rester.UpdateCondition(c, GetRouteID(this.Rester))
	info := model.UpdateWhere(condition, obj)
	this.Rester.AfterUpdate(c, obj)
	this.Rester.AfterRest(c, obj)

	helper.ReturnSuccess(c, http.StatusOK, info)
}

func (this *Controller) Delete(c *gin.Context) {
	obj := this.RestModel()
	id, _ := strconv.Atoi(c.Param(GetRouteID(this.Rester)))
	model.Delete(obj, id)
	helper.ReturnSuccess(c, http.StatusOK, gin.H{"id": id})
}
func BuildRoute(controller ControllerInterface) (path, resourceName, routeId string) {
	if controller.ParentNode() != nil {
		pp,pr,pi := BuildRoute(controller.ParentNode())
		path =  pp + "/" + pr + "/:" + pi
	}
	resourceName = controller.RouteName() + "s"
	routeId = GetRouteID(controller)
	return
}
func GetRouteID(controller ControllerInterface) (routeId string) {
	routeId = "id"
	if controller.IsRestRoutePk() {
		routeId = controller.RouteName() + "_id"
	}
	return
}
