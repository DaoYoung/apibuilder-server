package rest

import (
	"strings"
)

// Route defines a route as consumed by the router. It can be instantiated directly, or using one
// of the shortcut methods: rest.Get, rest.Post, rest.Put, rest.Patch and rest.Delete.
type Route struct {
	// Any HTTP method. It will be used as uppercase to avoid common mistakes.
	HttpMethod string

	// A string like "/resource/:id.json".
	// Placeholders supported are:
	// :paramName that matches any char to the first '/' or '.'
	// #paramName that matches any char to the first '/'
	// *paramName that matches everything to the end of the string
	// (placeholder names must be unique per PathExp)
	PathExp string

	// Code that will be executed when this route is taken.
	Handler Handler
}

// MakePath generates the path corresponding to this Route and the provided path parameters.
// This is used for reverse route resolution.
func (route *Route) MakePath(pathParams map[string]string) string {
	path := route.PathExp
	for paramName, paramValue := range pathParams {
		paramPlaceholder := ":" + paramName
		relaxedPlaceholder := "#" + paramName
		splatPlaceholder := "*" + paramName
		r := strings.NewReplacer(paramPlaceholder, paramValue, splatPlaceholder, paramValue, relaxedPlaceholder, paramValue)
		path = r.Replace(path)
	}
	return path
}

// Head is a shortcut method that instantiates a HEAD route. See the Route object the parameters definitions.
// Equivalent to &Route{"HEAD", pathExp, handler}
func Head(pathExp string, handler Handler) *Route {
	return &Route{
		HttpMethod: "HEAD",
		PathExp:    pathExp,
		Handler:    handler,
	}
}

// HeadFunc is a shortcut method that instantiates a HEAD route. See the Route object the parameters definitions.
func HeadFunc(pathExp string, handlerFunc HandlerFunc) *Route {
	return Head(pathExp, handlerFunc)
}

// Get is a shortcut method that instantiates a GET route. See the Route object the parameters definitions.
// Equivalent to &Route{"GET", pathExp, handler}
func Get(pathExp string, handler Handler) *Route {
	return &Route{
		HttpMethod: "GET",
		PathExp:    pathExp,
		Handler:    handler,
	}
}

// GetFunc is a shortcut method that instantiates a GET route. See the Route object the parameters definitions.
func GetFunc(pathExp string, handlerFunc HandlerFunc) *Route {
	return Get(pathExp, handlerFunc)
}

// Post is a shortcut method that instantiates a POST route. See the Route object the parameters definitions.
// Equivalent to &Route{"POST", pathExp, handler}
func Post(pathExp string, handler Handler) *Route {
	return &Route{
		HttpMethod: "POST",
		PathExp:    pathExp,
		Handler:    handler,
	}
}

// PostFunc is a shortcut method that instantiates a POST route. See the Route object the parameters definitions.
func PostFunc(pathExp string, handlerFunc HandlerFunc) *Route {
	return Post(pathExp, handlerFunc)
}

// Put is a shortcut method that instantiates a PUT route.  See the Route object the parameters definitions.
// Equivalent to &Route{"PUT", pathExp, handler}
func Put(pathExp string, handler Handler) *Route {
	return &Route{
		HttpMethod: "PUT",
		PathExp:    pathExp,
		Handler:    handler,
	}
}

// PutFunc is a shortcut method that instantiates a PUT route.  See the Route object the parameters definitions.
func PutFunc(pathExp string, handlerFunc HandlerFunc) *Route {
	return Put(pathExp, handlerFunc)
}

// Patch is a shortcut method that instantiates a PATCH route.  See the Route object the parameters definitions.
// Equivalent to &Route{"PATCH", pathExp, handler}
func Patch(pathExp string, handler Handler) *Route {
	return &Route{
		HttpMethod: "PATCH",
		PathExp:    pathExp,
		Handler:    handler,
	}
}

// PatchFunc is a shortcut method that instantiates a PATCH route.  See the Route object the parameters definitions.
func PatchFunc(pathExp string, handlerFunc HandlerFunc) *Route {
	return Patch(pathExp, handlerFunc)
}

// Delete is a shortcut method that instantiates a DELETE route. Equivalent to &Route{"DELETE", pathExp, handler}
func Delete(pathExp string, handler Handler) *Route {
	return &Route{
		HttpMethod: "DELETE",
		PathExp:    pathExp,
		Handler:    handler,
	}
}
// DeleteFunc is a shortcut method that instantiates a DELETE route. Equivalent to &Route{"DELETE", pathExp, handler}
func DeleteFunc(pathExp string, handlerFunc HandlerFunc) *Route {
	return Delete(pathExp, handlerFunc)
}

// Options is a shortcut method that instantiates an OPTIONS route.  See the Route object the parameters definitions.
// Equivalent to &Route{"OPTIONS", pathExp, handler}
func Options(pathExp string, handler Handler) *Route {
	return &Route{
		HttpMethod: "OPTIONS",
		PathExp:    pathExp,
		Handler:    handler,
	}
}

// OptionsFunc is a shortcut method that instantiates an OPTIONS route.  See the Route object the parameters definitions.
func OptionsFunc(pathExp string, handlerFunc HandlerFunc) *Route {
	return Options(pathExp, handlerFunc)
}