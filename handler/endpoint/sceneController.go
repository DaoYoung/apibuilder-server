package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
	"apibuilder-server/helper"
		"net/http"
	"errors"
	"log"
	"strings"
			)

type SceneController struct {
	Controller
}
func (action *SceneController) model() model.ResourceInterface {
	return &(model.Scene{})
}
func (action *SceneController) modelSlice() interface{} {
	return &[]model.Scene{}
}
func (action SceneController) Rester() (actionPtr *SceneController) {
	action.init(&action)
	return  &action
}
func (action *SceneController) beforeCreate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.Scene).AuthorId = user.ID
	m.(*model.Scene).SerialNo = helper.NewUuid()
}

func (action *SceneController) Match(c *gin.Context) {
	ip := strings.Split(c.Request.RemoteAddr, ":")[0]
	s := c.GetHeader("User-Agent")
	log.Print(ip, s)
	serialNo := c.Param("serial_no")
	obj := action.RestModel()
	condition := make(map[string]interface{})
	condition["serial_no"] = serialNo
	model.Find(obj, condition)
	if obj.(*model.Scene).Status == model.SceneStatusMatch {
		panic(NOChangeError(errors.New("serial_no: " + serialNo + " was match yet.")))
	}
	scene := action.RestModel().(*model.Scene)
	scene.Status = model.SceneStatusMatch
	model.Update(obj.(*model.Scene).ID, scene)
	helper.ReturnSuccess(c, http.StatusOK, obj)
}
