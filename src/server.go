package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rikiya/go-clean/src/external"
)

func main() {
	r := mux.NewRouter()
	external.Router(r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
	}
	srv.ListenAndServe()
}
