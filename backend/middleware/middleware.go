package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/IDOMATH/StrictlyRecipes/repository"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Use(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		handler = m(handler)
	}
	return handler
}

func CreateStack(stack ...Middleware) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		for i := len(stack) - 1; i >= 0; i-- {
			mw := stack[i]
			next = mw(next)
		}

		return next
	}
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Println(r.Method, r.URL, time.Since(start))
	})
}

func Authenticate(repo *repository.Repository) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Make and hook up session storage to repository
			t, found := repo.Session.Get(r.Header.Get("cheetauth"))
			if !found {
				fmt.Println("NOT AUTHENTICATED")
				// Potentially do some rerouting if the endpoint is protected
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(r.Header.Get("cheetauth")))
				return
			}
			id, err := strconv.Atoi(t)
			if err != nil {
				fmt.Println("error converting token id to int")
				// Potentially do some rerouting if the endpoint is protected
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(r.Header.Get("cheetauth")))
				return

			}
			fmt.Println(id)

			next(w, r)
		}
	}
}
