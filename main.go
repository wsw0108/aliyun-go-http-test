package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			qualifier := r.Header.Get("x-fc-qualifier")
			if qualifier != "" && !strings.EqualFold(qualifier, "LATEST") {
				parts := strings.Split(qualifier, "_")
				prefix := "/" + parts[0]
				http.StripPrefix(prefix, next).ServeHTTP(w, r)
			} else {
				next.ServeHTTP(w, r)
			}
		}
		return http.HandlerFunc(fn)
	})

	r.Get("/headers", func(w http.ResponseWriter, r *http.Request) {
		r.Header.WriteSubset(w, nil)
	})

	r.Get("/v1/headers", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("v1\n"))
		r.Header.WriteSubset(w, nil)
	})

	log.Fatalln(http.ListenAndServe(":9000", r))
}
