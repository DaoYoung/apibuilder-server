package endpoint

import (
	"apibuilder-server/model"
		)

type ProxyController struct {
	Controller
}

func (action *ProxyController) model() model.ResourceInterface {
	return &(model.Proxy{})
}
func (action *ProxyController) modelSlice() interface{} {
	return &[]model.Proxy{}
}
func (action ProxyController) Rester() (actionPtr *ProxyController) {
	action.init(&action)
	return &action
}
