package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
	"net/http"
	"apibuilder-server/helper"
)

type UserController struct {
	Controller
}

func (action UserController) CrudService(str string) func(c *gin.Context)  {
	actionPtr := &action
	actionPtr.GetResModel = func() model.Resource { return &(model.User{}) }
	actionPtr.GetResSlice = func() interface{} { return &[]model.User{} }
	return actionPtr.Controller.DaoService(str)
}

func Profile(c *gin.Context)  {
	user := model.GetUserFromToken(c)
	helper.ReturnSuccess(c, http.StatusOK, user)
}

