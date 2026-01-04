// Package middleware: writing own middlewares from scratch using standard packages
// Logging middleware
package middleware

import (
	"log/slog"
	"net/http"
	"slices"
	"time"
)

type Middleware func(http.Handler)http.Handler

func MiddlewareChain(next http.Handler, m ...Middleware)http.Handler{
	wrapped:=next
	for _, middleware:= range slices.Backward(m){
		wrapped=middleware(wrapped)
	}
	return wrapped
}

func LoggerMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start:=time.Now()
		var (
			ip = r.RemoteAddr
			method=r.Method
			url=r.URL.String()
			proto=r.Proto
		)
		userAttrs:=slog.Group("user", "ip",ip)
		requestAttrs:=slog.Group("request", "method", method, "url", url,"proto", proto)
		next.ServeHTTP(w, r)
		duration:=time.Since(start)
		time:=slog.Group("time","response time", duration)
		slog.Info("request received", userAttrs, requestAttrs, time)
		
	})
}
