package main

import (
	"log"
	"net/http"
	"hlj-rest/rest"
)

type MyEnvelope struct{}

func (e *MyEnvelope) Error(w rest.ResponseWriter, code int, msg string, httpCode int) error {
	ret := struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}{code, msg}

	w.WriteHeader(httpCode)
	return w.WriteJson(ret)
}

func (e *MyEnvelope) WriteData(w rest.ResponseWriter, data interface{}, meta map[string]interface{}, httpCode int) error {
	w.WriteHeader(httpCode)
	return w.WriteJson(data)
}

func main() {
	rest.SetDefaultEnvelope(&MyEnvelope{})

	route := rest.GetFunc("/envelope", func(w rest.ResponseWriter, r *rest.Request) {
		ok := r.QueryParam("ok").Default("1").Bool()
		if ok {
			rest.Ok(w, "Everything is fine", nil)
		} else {
			rest.Error(w, 500, "not ok", 500)
		}
	})
	router, err := rest.MakeRouter(route)

	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/api/", http.StripPrefix("/api", router))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
