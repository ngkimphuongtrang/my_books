package httpkit

import (
	"fmt"
	"log"
	"time"
)

type Middleware interface {
	Handle(ctx *RequestContext)
}

type RequestTimeMiddleware struct{}

func (m *RequestTimeMiddleware) Handle(ctx *RequestContext) {
	startTime := time.Now()
	ctx.Next()
	runTime := time.Since(startTime)

	log.Println(
		ctx.GetContext(),
		"remote_addr", ctx.Request.RemoteAddr,
		"method", ctx.Request.Method,
		"path", ctx.Request.URL.Path,
		"status_code", ctx.Response.StatusCode,
		"response_verdict", ctx.Response.Verdict,
		"response_message", fmt.Sprintf("%q", ctx.Response.Message),
		"runtime", runTime,
	)
}
