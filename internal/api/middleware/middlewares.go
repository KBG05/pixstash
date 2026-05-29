// Package middleware: writing own middlewares from scratch using standard packages
// Logging middleware
package middleware

import (
	"fmt"
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

		next.ServeHTTP(w, r)
		duration:=time.Since(start)
		logLine := fmt.Sprintf("%s %s %s %s %v", ip, proto, method, url, duration)

        // Use the string as the main message
        slog.Info(logLine)
	
	})
}
