package endpoint

import (
	"apibuilder-server/model"
)

type ModuleController struct {
	Controller
}
func (action ModuleController) Rester() ControllerInterface {
	action.Controller.Rester = &action
	action.Controller.RestModel = func() model.ResourceInterface { return &(model.Module{}) }
	action.Controller.RestModelSlice = func() interface{} { return &[]model.Module{} }
	return  &action
}

