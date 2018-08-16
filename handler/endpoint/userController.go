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
func (action *UserController) model() model.ResourceInterface {
	return &(model.User{})
}
func (action *UserController) modelSlice() interface{} {
	return &[]model.User{}
}
func (action UserController) Rester() (actionPtr *UserController) {
	action.init(&action)
	return  &action
}
func Profile(c *gin.Context)  {
	user := model.GetUserFromToken(c)
	helper.ReturnSuccess(c, http.StatusOK, user)
}

