package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"errors"
)
type BaseAction struct {
	ModFunc *model.BaseFunc
}

func JsonDecode(c *gin.Context, obj interface{}){
	err := c.BindJSON(obj)
	if err != nil {
		panic(JsonTypeError(err))
	}
}
func (ba *BaseAction) Info(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	info := ba.ModFunc.ByID(id)
	c.JSON(http.StatusOK, info)
}
func (ba *BaseAction) List(c *gin.Context) {
	list := ba.ModFunc.FindList()
	c.JSON(http.StatusOK, list)
}
func (ba *BaseAction) Create(c *gin.Context) {
	err := c.BindJSON(ba.ModFunc.Mod)
	if err != nil {
		panic(JsonTypeError(err))
	}
	info := ba.ModFunc.Create(ba.ModFunc.Mod)
	c.JSON(http.StatusCreated, info)
}
func (ba *BaseAction) Update(c *gin.Context) {
	err := c.BindJSON(ba.ModFunc.Mod)
	if err != nil {
		panic(JsonTypeError(err))
	}
	id, _ := strconv.Atoi(c.Param("id"))
	info := ba.ModFunc.Update(id, ba.ModFunc.Mod)
	c.JSON(http.StatusOK, info)
}
func (ba *BaseAction) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ba.ModFunc.Delete(id)
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (ba *BaseAction) CrudService(funcName string) func(c *gin.Context) {
	if ba.ModFunc.Mod == nil{
		panic(model.NotExistDaoError(errors.New("model not exist ")))
	}
	switch funcName {
	case "info":
		return ba.Info
	case "create":
		return ba.Create
	case "update":
		return ba.Update
	case "delete":
		return ba.Delete
	default:
		return ba.List
	}

}
