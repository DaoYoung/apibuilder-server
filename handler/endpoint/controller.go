package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"errors"
)
type ControllerInterface interface {
	CrudService(str string) func(c *gin.Context)
}

type Controller struct {
	Res model.Resource
}

func (this *Controller) Info(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	info := this.Res.InitDao().ByID(id)
	ReturnSuccess(c, http.StatusOK, info)
}
func (this *Controller) List(c *gin.Context) {
	list := this.Res.InitDao().FindList()
	ReturnSuccess(c, http.StatusOK, list)
}
func (this *Controller) Create(c *gin.Context) {
	err := c.BindJSON(this.Res)
	if err != nil {
		panic(JsonTypeError(err))
	}
	info := this.Res.InitDao().Create(this.Res)
	ReturnSuccess(c, http.StatusCreated, info)
}
func (this *Controller) Update(c *gin.Context) {
	obj := this.Res.UpdateStruct()
	if obj == nil {
		panic(ForbidError(errors.New("forbid to update model")))
	}
	err := c.BindJSON(obj)
	if err != nil {
		panic(JsonTypeError(err))
	}
	id, _ := strconv.Atoi(c.Param("id"))
	info := this.Res.InitDao().Update(id, obj)
	ReturnSuccess(c, http.StatusOK, info)
}
func (this *Controller) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	this.Res.InitDao().Delete(id)
	ReturnSuccess(c, http.StatusOK, gin.H{"id": id})
}

func (this *Controller) DaoService(funcName string) func(c *gin.Context) {
	if this.Res.InitDao() == nil{
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
type JsonError struct {
	Errors interface{} `json:"errors"`
}

func ReturnSuccess(c *gin.Context, code int, data interface{}) {
	js := new(JsonSuccess)
	js.Data = data
	c.JSON(code, js)
}

func ReturnError(c *gin.Context, code int, err interface{}) {
	js := new(JsonError)
	js.Errors = err
	c.JSON(code, js)
}