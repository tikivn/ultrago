package u_middleware

import "github.com/go-chi/httplog"

var DefaultOptions = HttpLogOptions{
	Options: httplog.DefaultOptions,
	Body:    true,
}

type HttpLogOptions struct {
	httplog.Options
	//Log request and response bodies or not
	Body bool
}

func NewHttpLogConfig(opts ...HttpLogOptions) {
	if len(opts) > 0 {
		httplog.Configure(opts[0].Options)
	} else {
		httplog.Configure(DefaultOptions.Options)
	}
}
