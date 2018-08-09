package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"errors"
	"reflect"
)
type ControllerInterface interface {
	CrudService(str string) func(c *gin.Context)
}

type Controller struct {
	TableName string
	Res model.Resource
	ResSlice interface{} //https://golang.org/doc/faq#convert_slice_of_interface
}
func (this *Controller) CrudService(str string) func(c *gin.Context)  {
	panic(ForbidError(errors.New("not support model curd")))
}

func (this *Controller) Info(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	info := model.ByID(this.Res, id)
	ReturnSuccess(c, http.StatusOK, info)
}
func (this *Controller) List(c *gin.Context) {
	condition := make(map[string]interface{})
	model.FindListWhereMap(this.ResSlice, condition)
	ReturnSuccess(c, http.StatusOK, this.ResSlice)
}
func clone(i interface{}) interface{} {
	// Wrap argument to reflect.Value, dereference it and return back as interface{}
	return reflect.Indirect(reflect.ValueOf(i)).Interface()
}
func (this *Controller) Create(c *gin.Context) {
	err := c.BindJSON(this.Res)
	if err != nil {
		panic(JsonTypeError(err))
	}
	info := model.CreateNew(this.TableName, this.Res)
	ReturnSuccess(c, http.StatusCreated, info)

	//
	//obj := clone(this.Res)
	//log.Printf("11 %+v", obj)
	//err := c.BindJSON(&obj)
	//if err != nil {
	//	panic(JsonTypeError(err))
	//}
	//log.Printf("%+v", obj)
	//info := model.CreateNew(this.TableName, obj)
	//ReturnSuccess(c, http.StatusCreated, info)
}

func (this *Controller) Update(c *gin.Context) {
	//obj := this.Res.UpdateStruct()
	//if obj == nil {
	//	panic(ForbidError(errors.New("forbid to update model")))
	//}
	err := c.BindJSON(this.Res)
	if err != nil {
		panic(JsonTypeError(err))
	}
	id, _ := strconv.Atoi(c.Param("id"))
	info := model.Update(id, this.Res)
	ReturnSuccess(c, http.StatusOK, info)
}
func (this *Controller) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	model.Delete(this.Res, id)
	ReturnSuccess(c, http.StatusOK, gin.H{"id": id})
}

func (this *Controller) DaoService(funcName string) func(c *gin.Context) {
	if this.Res == nil{
		panic(model.NotExistDaoError(errors.New("model not exist ")))
	}
	switch funcName {
	case "info":
		return this.Info
	case "create":
		return this.Create
	case "update":
		return this.Update
	case "delete":
		return this.Delete
	default:
		return this.List
	}

}


type JsonSuccess struct {
	Data interface{} `json:"data"`
}

func ReturnSuccess(c *gin.Context, code int, data interface{}) {
	js := new(JsonSuccess)
	js.Data = data
	c.JSON(code, js)
}