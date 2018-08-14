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
func (action UserController) Rester() ControllerInterface {
	actionPtr := &action
	action.Controller.Rester = actionPtr
	action.Controller.RestModel = func() model.ResourceInterface { return &(model.User{}) }
	action.Controller.RestModelSlice = func() interface{} { return &[]model.User{} }
	return  actionPtr
}
func Profile(c *gin.Context)  {
	user := model.GetUserFromToken(c)
	helper.ReturnSuccess(c, http.StatusOK, user)
}

