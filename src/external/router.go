package external

import (
	"github.com/gorilla/mux"
	"github.com/rikiya/go-clean/src/adapter/userscontroller"
	"github.com/rikiya/go-clean/src/infrastructure/database"
)

// Router ...
func Router(r *mux.Router) {
	usersController := userscontroller.NewUserController(*database.NewSQLHandler())
	r.HandleFunc("/users", usersController.Create).Methods("POST")
	r.HandleFunc("/users", usersController.FindAll).Methods("GET")
	r.HandleFunc("/users/{id}", usersController.Update).Methods("PUT")
	r.HandleFunc("/users/{id}", usersController.Delete).Methods("DELETE")
}
