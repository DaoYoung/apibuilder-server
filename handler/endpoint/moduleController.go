package endpoint

import (
	"apibuilder-server/model"
)

type ModuleController struct {
	Controller
}
func (action *ModuleController) GetRestModel() model.ResourceInterface{
	return &(model.Module{})
}
func (action *ModuleController) GetRestModelSlice() interface{}{
	return &[]model.Module{}
}
func (action *ModuleController) GetRester() *ModuleController {
	return action
}

