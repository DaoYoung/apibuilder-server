package rest

// Middleware defines the interface that objects must implement in order to wrap a HandlerFunc and
// be used in the middleware stack.
type Middleware interface {
	MiddlewareFunc(handler HandlerFunc) HandlerFunc
}

// MiddlewareSimple is an adapter type that makes it easy to write a Middleware with a simple
// function. eg: api.Use(rest.MiddlewareSimple(func(h HandlerFunc) Handlerfunc { ... }))
type MiddlewareSimple func(handler HandlerFunc) HandlerFunc

// MiddlewareFunc makes MiddlewareSimple implement the Middleware interface.
func (ms MiddlewareSimple) MiddlewareFunc(handler HandlerFunc) HandlerFunc {
	return ms(handler)
}
