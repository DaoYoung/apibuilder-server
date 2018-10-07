package endpoint

import (
	"apibuilder-server/model"
	"github.com/gin-gonic/gin"
	"errors"
	"strconv"
)

type ProxyChannelController struct {
	Controller
}

func (action *ProxyChannelController) model() model.ResourceInterface {
	return &(model.ProxyChannel{})
}
func (action *ProxyChannelController) modelSlice() interface{} {
	return &[]model.ProxyChannel{}
}
func (action ProxyChannelController) Rester() (actionPtr *ProxyChannelController) {
	action.init(&action)
	return &action
}
func (action *ProxyChannelController) beforeCreate(c *gin.Context, m model.ResourceInterface) {
	user := model.GetUserFromToken(c)
	m.(*model.ProxyChannel).AuthorId = user.ID
	proxy := &(model.Proxy{})
	model.ByID(proxy, m.(*model.ProxyChannel).ProxyId)
	if proxy.HoldChannelId > 0{
		panic(ForbidError(errors.New(strconv.Itoa(proxy.Port) + " was testing.")))
	}
}
func (action *ProxyChannelController) afterCreate(c *gin.Context, m model.ResourceInterface) {
	proxy := &(model.Proxy{})
	proxy.HoldChannelId = m.(*model.ProxyChannel).ID
	model.Update(m.(*model.ProxyChannel).ProxyId, proxy)
	//todo update with HoldChannelId
	go Proxy(proxy.Port,m.(*model.ProxyChannel).ID)
}