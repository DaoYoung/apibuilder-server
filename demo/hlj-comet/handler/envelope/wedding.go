package envelope

import (
	"hlj-rest/rest"
	"time"
)

type weddingStructure struct {
	Status struct {
		RetCode int    `json:"RetCode"`
		Msg     string `json:"msg"`
	} `json:"status"`
	Data        interface{} `json:"data"`
	CurrentTime int64       `json:"current_time"`
}

type WeddingEnvelope struct{}

func (e *WeddingEnvelope) Error(w rest.ResponseWriter, code int, msg string, httpCode int) error {
	ret := new(weddingStructure)
	ret.Status.RetCode = code
	ret.Status.Msg = msg
	ret.CurrentTime = time.Now().Unix()

	w.WriteHeader(httpCode)
	return w.WriteJson(ret)
}

func (e *WeddingEnvelope) WriteData(w rest.ResponseWriter, data interface{}, meta map[string]interface{}, httpCode int) error {
	ret := new(weddingStructure)
	ret.Data = data
	ret.CurrentTime = time.Now().Unix()

	w.WriteHeader(httpCode)
	return w.WriteJson(ret)
}
