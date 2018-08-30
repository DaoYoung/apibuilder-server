package endpoint

import "apibuilder-server/model"

type TeamTaskController struct {
	Controller
}
func (action *TeamTaskController) model() model.ResourceInterface {
	return &(model.TeamTask{})
}
func (action *TeamTaskController) modelSlice() interface{} {
	return &[]model.TeamTask{}
}
func (action TeamTaskController) Rester() (actionPtr *TeamTaskController) {
	action.init(&action)
	return  &action
}