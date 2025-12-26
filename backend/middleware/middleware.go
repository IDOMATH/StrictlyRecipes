package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type Middleware func(handler http.Handler) http.Handler

func CreateStack(stack ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(stack) - 1; i >= 0; i-- {
			mw := stack[i]
			next = mw(next)
		}

		return next
	}
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Println(r.Method, r.URL, time.Since(start))
	})
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Not Authenticated")
		next.ServeHTTP(w, r)
	})
}
