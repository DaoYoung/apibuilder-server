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
	"errors"
)


type ControllerInterface interface {
	Update(c *gin.Context)
	Create(c *gin.Context)
	Info(c *gin.Context)
	List(c *gin.Context)
	Delete(c *gin.Context)
	IsRestRoutePk() bool //false id
	RouteName() string //rewrite resource name in route url
	ParentNode() ControllerInterface

	init(r ControllerInterface)
	model() model.ResourceInterface
	modelSlice() interface{}
	parentController() ControllerInterface
	beforeDelete(c *gin.Context, m model.ResourceInterface, id int)
	afterDelete(c *gin.Context, m model.ResourceInterface, id int)
	beforeCreate(c *gin.Context, m model.ResourceInterface)
	afterCreate(c *gin.Context, m model.ResourceInterface)
	beforeUpdate(c *gin.Context, old model.ResourceInterface, new model.ResourceInterface)
	updateCondition(c *gin.Context, pk string) map[string]interface{}
	afterUpdate(c *gin.Context, old model.ResourceInterface, new model.ResourceInterface)
	listCondition(c *gin.Context) map[string]interface{}

}

type Controller struct {
	InfoFields   []string
	ListFields   []string
	ParentController ControllerInterface
	Rester           ControllerInterface
	RestModel        func() model.ResourceInterface
	RestModelSlice   func() interface{} //https://golang.org/doc/faq#convert_slice_of_interface
	*EmptyRest
}
func (this *Controller) getInfoFields(){

}
func (this *Controller) init(r ControllerInterface){
	if r == nil {
		panic(NOContentError(errors.New("param r: is not a controller")))
	}
	this.Rester = r
	this.RestModel = r.model
	this.RestModelSlice = r.modelSlice
	this.ParentController = r.parentController()
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
	this.Rester.beforeCreate(c, obj)
	info := model.Create(obj)
	this.Rester.afterCreate(c, info)
	helper.ReturnSuccess(c, http.StatusCreated, info)
}

func (this *Controller) Info(c *gin.Context) {
	obj := this.RestModel()
	id, _ := strconv.Atoi(c.Param(GetRouteID(this.Rester)))
	model.ByID(obj, id, this.InfoFields...)
	helper.ReturnSuccess(c, http.StatusOK, obj)
}

func (this *Controller) List(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		panic(err)
	}
	obj := this.RestModelSlice()
	condition := this.Rester.listCondition(c)
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
	condition := this.Rester.updateCondition(c, GetRouteID(this.Rester))
	if val, ok := condition["id"]; ok {
		old := this.RestModel()
		model.ByID(old, val.(int))
		CheckupdateCondition(old, condition)
		this.Rester.beforeUpdate(c, old, obj)
		info := model.Update(val.(int), obj)
		this.Rester.afterUpdate(c, old, info)
		helper.ReturnSuccess(c, http.StatusOK, info)
	}else {
		panic(NOChangeError(errors.New("can't find data to update")))
	}
}
func (this *Controller) Delete(c *gin.Context) {
	obj := this.RestModel()
	id, _ := strconv.Atoi(c.Param(GetRouteID(this.Rester)))
	this.Rester.beforeDelete(c,obj, id)
	model.Delete(obj, id)
	this.Rester.afterDelete(c,obj, id)
	helper.ReturnSuccess(c, http.StatusOK, gin.H{"id": id})
}
