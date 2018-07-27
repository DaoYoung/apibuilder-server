package main

import (
	"net/http"
	"log"
	"hlj-rest/rest"
	"hlj-rest/middleware"
)

type AuthMiddleware struct{}

func (auth *AuthMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {
	return func(w rest.ResponseWriter, r *rest.Request) {
		r.Env["USER"] = r.QueryParam("name").Default("anonymous").String()

		handler(w, r)
	}
}

func Login(w rest.ResponseWriter, r *rest.Request) {
	user := r.Env["USER"]
	name, _ := user.(string)

	rest.Ok(w, "Hi, " + name, nil)
}

func main() {
	auth := &AuthMiddleware{}
	routes := []*rest.Route{
		rest.Get("/login", rest.HandlerFunc(Login).Use(auth)),
	}

	router, _ := rest.MakeRouter(routes...)

	http.Handle("/api/", http.StripPrefix("/api", router.Use(middleware.DefaultDevStack...)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
