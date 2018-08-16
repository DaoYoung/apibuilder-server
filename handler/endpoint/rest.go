package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
	"errors"
	"strconv"
	"reflect"
	"apibuilder-server/helper"
)

type EmptyRest struct {}
func (this EmptyRest)  Rester() {
	panic(NOContentError(errors.New("can't find func:Rester in your controller")))
}
func (this EmptyRest)  modelSlice() interface{} {
	panic(NOContentError(errors.New("can't find func:modelSlice in your controller")))
}
func (this *EmptyRest) model() model.ResourceInterface{
	panic(NOContentError(errors.New("can't find func:model in your controller")))
}
func (this *EmptyRest) parentController() ControllerInterface { return  nil}
func (this *EmptyRest) beforeCreate(c *gin.Context, m model.ResourceInterface) {}
func (this *EmptyRest) afterCreate(c *gin.Context, m model.ResourceInterface) {}
func (this *EmptyRest) beforeUpdate(c *gin.Context, old model.ResourceInterface, new model.ResourceInterface) {}
func (this *EmptyRest) afterUpdate(c *gin.Context, old model.ResourceInterface, new model.ResourceInterface) {}
func (this *EmptyRest) beforeDelete(c *gin.Context, m model.ResourceInterface, id int) {}
func (this *EmptyRest) afterDelete(c *gin.Context, m model.ResourceInterface, id int) {}
func (this *EmptyRest) listCondition(c *gin.Context) map[string]interface{} {
	return make(map[string]interface{})
}
func (this *EmptyRest) updateCondition(c *gin.Context, pk string) map[string]interface{} {
	condition := make(map[string]interface{})
	id, err := strconv.Atoi(c.Param(pk))
	if err!=nil {
		panic(NOContentError(errors.New("can't Update without ID")))
	}
	condition["id"] = id
	return condition
}



func BuildRoute(controller ControllerInterface) (path, resourceName, routeId string) {
	if controller.ParentNode() != nil {
		pp, pr, pi := BuildRoute(controller.ParentNode())
		path = pp + "/" + pr + "/:" + pi
	}
	resourceName = controller.RouteName() + "s"
	routeId = GetRouteID(controller)
	return
}
func GetRouteID(controller ControllerInterface) (routeId string) {
	routeId = "id"
	if controller.IsRestRoutePk() {
		routeId = controller.RouteName() + "_id"
	}
	return
}
func CheckupdateCondition(m model.ResourceInterface, condition map[string]interface{}) {
	v := reflect.ValueOf(m)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for key, val := range condition {
		old := v.FieldByName(helper.CamelString(key))
		switch old.Kind() {
		case reflect.String:
			if old.String() != val {
				panic(ForbidError(errors.New("forbid update by field:" + key)))
			}
			break
		case reflect.Int:
			if old.Int() != int64(val.(int)) {
				panic(ForbidError(errors.New("forbid update by field:" + key)))
			}
			break
		default:
			panic(ForbidError(errors.New("forbid update by field type:" + old.Kind().String())))
		}
	}
}
