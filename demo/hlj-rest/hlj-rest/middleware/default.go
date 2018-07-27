package middleware

import(
	"hlj-rest/rest"
)

// Defines a stack of middlewares convenient for development. Among other things:
// console friendly logging, JSON indentation, error stack strace in the response.
var DefaultDevStack = []rest.Middleware{
	&ContentTypeCheckerMiddleware{},
	&RecoverMiddleware{
		EnableResponseStackTrace: true,
	},
	&PoweredByMiddleware{},
	&RecorderMiddleware{},
	&TimerMiddleware{},
	&AccessLogApacheMiddleware{},
}

// Defines a stack of middlewares convenient for production. Among other things:
// Apache CombinedLogFormat logging, gzip compression.
var DefaultProdStack = []rest.Middleware{
	&ContentTypeCheckerMiddleware{},
	&GzipMiddleware{},
	&RecoverMiddleware{},
	&PoweredByMiddleware{},
	&RecorderMiddleware{},
	&TimerMiddleware{},
	&AccessLogApacheMiddleware{
		Format: CombinedLogFormat,
	},
}

// Defines a stack of middlewares that should be common to most of the middleware stacks.
var DefaultCommonStack = []rest.Middleware{
	&RecoverMiddleware{},
	&PoweredByMiddleware{},
	&RecorderMiddleware{},
	&TimerMiddleware{},
}