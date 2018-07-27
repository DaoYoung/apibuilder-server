package main

import (
	"log"
	"net/http"
	"hlj-rest/rest"
)

func main() {
	route := rest.GetFunc("/hello", func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": "Hello World!"})
	})
	router, err := rest.MakeRouter(route)

	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}
