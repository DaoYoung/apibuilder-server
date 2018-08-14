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
func (action *UserController) GetRestModel() model.ResourceInterface{
	return &(model.User{})
}
func (action *UserController) GetRestModelSlice() interface{}{
	return &[]model.User{}
}
func (action *UserController) GetRester() *UserController {
	return action
}

func Profile(c *gin.Context)  {
	user := model.GetUserFromToken(c)
	helper.ReturnSuccess(c, http.StatusOK, user)
}

