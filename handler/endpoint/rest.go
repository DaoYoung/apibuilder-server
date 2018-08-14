package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
	"errors"
)

type RestInterface interface {
	ControllerInterface
	Rester() ControllerInterface
	BeforeRest(c *gin.Context, m model.ResourceInterface)
	AfterRest(c *gin.Context, m model.ResourceInterface)
	BeforeCreate(c *gin.Context, m model.ResourceInterface)
	AfterCreate(c *gin.Context, m model.ResourceInterface)
	BeforeUpdate(c *gin.Context, m model.ResourceInterface)
	AfterUpdate(c *gin.Context, m model.ResourceInterface)
	ListCondition(c *gin.Context) map[string]interface{}

}

type EmptyRest struct {}
func (this *EmptyRest) GetRestModel() model.ResourceInterface{
	panic(NOContentError(errors.New("can't find func:GetRestModel in your controller")))
	return nil
}
func (this *EmptyRest) GetRestModelSlice() interface{}{
	panic(NOContentError(errors.New("can't find func:GetRestModelSlice in your controller")))
	return nil
}
func (this *EmptyRest) Rester() ControllerInterface{
	panic(NOContentError(errors.New("can't find func:Rester in your controller")))
}
func (this *EmptyRest) BeforeCreate(c *gin.Context, m model.ResourceInterface) {}
func (this *EmptyRest) AfterCreate(c *gin.Context, m model.ResourceInterface) {}
func (this *EmptyRest) BeforeUpdate(c *gin.Context, m model.ResourceInterface) {}
func (this *EmptyRest) AfterUpdate(c *gin.Context, m model.ResourceInterface) {}
func (this *EmptyRest) BeforeRest(c *gin.Context, m model.ResourceInterface) {}
func (this *EmptyRest) AfterRest(c *gin.Context, m model.ResourceInterface) {}
func (this *EmptyRest) ListCondition(c *gin.Context) map[string]interface{} {
	return make(map[string]interface{})
}