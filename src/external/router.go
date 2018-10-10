package external

import (
	"github.com/gorilla/mux"
	"github.com/rikiya/go-clean/src/adapter"
	"github.com/rikiya/go-clean/src/infrastructure/database"
)

// Router ...
func Router(r *mux.Router) {
	usersController := adapter.NewUserController(*database.NewSQLHandler())
	r.HandleFunc("/create", usersController.Create).Methods("POST")
}
