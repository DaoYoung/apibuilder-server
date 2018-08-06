package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
	"net/http"
)

type UserController struct {
	Controller
}

func (action UserController) CrudService(str string) func(c *gin.Context)  {
	actionPtr := &action
	actionPtr.Res = &(model.User{})
	actionPtr.ResSlice = &[]model.User{}
	return actionPtr.Controller.DaoService(str)
}

func Profile(c *gin.Context)  {
	user := model.GetUserFromToken(c)
	ReturnSuccess(c, http.StatusOK, user)
}

