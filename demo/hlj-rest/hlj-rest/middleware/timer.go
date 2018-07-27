package middleware

import (
	"hlj-rest/rest"
	"time"
)

// TimerMiddleware computes the elapsed time spent during the execution of the wrapped handler.
// The result is available to the wrapping handlers as request.Env["ELAPSED_TIME"].(*time.Duration),
// and as request.Env["START_TIME"].(*time.Time)
type TimerMiddleware struct{}

// MiddlewareFunc makes TimerMiddleware implement the Middleware interface.
func (mw *TimerMiddleware) MiddlewareFunc(h rest.HandlerFunc) rest.HandlerFunc {
	return func(w rest.ResponseWriter, r *rest.Request) {

		start := time.Now()
		r.Env["START_TIME"] = &start

		// call the handler
		h(w, r)

		end := time.Now()
		elapsed := end.Sub(start)
		r.Env["ELAPSED_TIME"] = &elapsed
	}
}
