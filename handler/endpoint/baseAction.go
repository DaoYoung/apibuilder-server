package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
)
type BaseAction struct {
	ModFunc *model.BaseFunc
}

func (ba *BaseAction) Info(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	info, err := ba.ModFunc.ByID(id)
	if err == nil {
		c.JSON(http.StatusOK, info)
	}else {
		c.JSON(http.StatusOK, err)
	}
}
func (ba *BaseAction) List(c *gin.Context) {
	list,err := ba.ModFunc.FindList()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, list)
	}
}
func (ba *BaseAction) Create(c *gin.Context) {
	err := c.BindJSON(ba.ModFunc.Mod)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	info, err := ba.ModFunc.Create(ba.ModFunc.Mod)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusCreated, info)
	}
}
func (ba *BaseAction) Update(c *gin.Context) {
	err := c.BindJSON(ba.ModFunc.Mod)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	id, _ := strconv.Atoi(c.Param("id"))
	info,err := ba.ModFunc.Update(id, ba.ModFunc.Mod)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusResetContent, info)
	}
}
func (ba *BaseAction) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_,err := ba.ModFunc.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusNoContent, gin.H{"id": id})
	}
}

func (ba *BaseAction) CrudService(funcName string) func(c *gin.Context) {
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
