package adapter

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rikiya/go-clean/src/domain"

	"github.com/rikiya/go-clean/src/infrastructure/database"
	"github.com/rikiya/go-clean/src/infrastructure/user"
	"github.com/rikiya/go-clean/src/usecase"
)

// UserController ...
type UserController struct {
	Interactor usecase.UserInteractor
}

// NewUserController ...
func NewUserController(sqlHandler database.SQLHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &user.UserImpl{
				SQLHandler: sqlHandler,
			},
		},
	}
}

// Create ...
func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	u := domain.User{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(500)
		return
	}
	if err := c.Interactor.Store(u); err != nil {
		w.WriteHeader(500)
		return
	}
	log.Println("Created User!!")
}

// FindAll ...
func (c *UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	users, err := c.Interactor.Index()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.WriteHeader(500)
		return
	}
}

// Update ...
func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	u := domain.User{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(500)
		return
	}
	if err := c.Interactor.Update(u); err != nil {
		w.WriteHeader(500)
		return
	}
	log.Println("Updated User!!")
}

// Delete ...
func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(404)
		return
	}
	if err := c.Interactor.Delete(id); err != nil {
		w.WriteHeader(500)
		return
	}
	log.Println("Deleted User!!")
}
