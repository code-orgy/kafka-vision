package router

import (
	"kafka-vision/env"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// set up initial webserver
func Start() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("kafka vision api is running"))
	})

	if err := http.ListenAndServe(env.Port, r); err != nil {
		log.Fatal(err)
	}
}
