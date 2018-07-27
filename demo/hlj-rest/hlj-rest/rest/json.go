package rest

import(
	"github.com/liamylian/jsontime"
	"github.com/json-iterator/go"
)

var json = jsontime.ConfigWithCustomTimeFormat

func GetJson() jsoniter.API {
	return json
}

func SetJson(api jsoniter.API) {
	json = api
}
