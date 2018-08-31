package endpoint

import "apibuilder-server/model"

type TeamController struct {
	Controller
}

func (action *TeamController) model() model.ResourceInterface {
	return &(model.Team{})
}
func (action *TeamController) modelSlice() interface{} {
	return &[]model.Team{}
}
func (action TeamController) Rester() (actionPtr *TeamController) {
	action.init(&action)
	return &action
}