package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rikiya/go-clean/src/external"
)

func main() {
	r := mux.NewRouter()
	external.Router(r)
	http.ListenAndServe("localhost:8080", r)
}
