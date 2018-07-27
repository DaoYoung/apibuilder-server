package rest

import "net/http"

// Handler defines the handler interface. It is the go-rest equivalent of http.Handler.
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

// HandlerFunc defines the handler function. It is the go-rest equivalent of http.HandlerFunc.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls h(w, r).
func (h HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	h(w, r)
}

// Use calls the MiddlewareFunc methods in order and returns an HandlerFunc ready to be executed.
func (h HandlerFunc) Use(middlewares ... Middleware) HandlerFunc {
	wrapped := h
	count := len(middlewares)
	for i := 0; i < count; i++ {
		wrapped = middlewares[i].MiddlewareFunc(wrapped)
	}
	return wrapped
}

// adapterHandleFunc make the transition between net/http and go-json-rest objects.
// It intanciates the rest.Request and rest.ResponseWriter, ...
func adapterHandlerFunc(h HandlerFunc) http.HandlerFunc {

	return func(origWriter http.ResponseWriter, origRequest *http.Request) {

		// instantiate the rest objects
		request := &Request{
			Request:    origRequest,
			PathParams: nil,
			Env:        map[string]interface{}{},
		}

		writer := &responseWriter{
			origWriter,
			false,
		}

		// call the wrapped handler
		h(writer, request)
	}
}
