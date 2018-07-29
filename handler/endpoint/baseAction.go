package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
		)

type BaseAction struct {
	Mod *model.BaseFunc
}

func (ba *BaseAction) Info(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	info, err := ba.Mod.ByID(id)
	if err == nil {
		c.JSON(http.StatusOK, info)
	}else {
		c.JSON(http.StatusOK, err)
	}
}
func (ba *BaseAction) List(c *gin.Context) {
	list := ba.Mod.FindList()
	c.JSON(http.StatusOK, list)
}
func (ba *BaseAction) Create(c *gin.Context) {
	err := c.BindJSON(ba.Mod.Mod)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	info := ba.Mod.Create(ba.Mod.Mod)
	c.JSON(http.StatusCreated, info)
}
func (ba *BaseAction) Update(c *gin.Context) {
	err := c.BindJSON(ba.Mod.Mod)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	id, _ := strconv.Atoi(c.Param("id"))
	info := ba.Mod.Update(id, ba.Mod.Mod)
	c.JSON(http.StatusResetContent, info)
}
func (ba *BaseAction) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ba.Mod.Delete(id)
	c.JSON(http.StatusNoContent, gin.H{"id": id})
}

func CurdAction(ba *BaseAction, funcName string) func(c *gin.Context) {

	switch funcName {
	case "info":
		return ba.Info
	case "list":
		return ba.List
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
