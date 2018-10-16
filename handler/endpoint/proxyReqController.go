package endpoint

import (
	"apibuilder-server/model"
		)

type ProxyReqController struct {
	Controller
}

func (action *ProxyReqController) model() model.ResourceInterface {
	return &(model.Proxy{})
}
func (action *ProxyReqController) modelSlice() interface{} {
	return &[]model.Proxy{}
}
func (action ProxyReqController) Rester() (actionPtr *ProxyReqController) {
	action.init(&action)
	return &action
}
