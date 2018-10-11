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
	err := json.NewDecoder(r.Body).Decode(&u)
	ErrorStatus(w, err, 500)
	err = c.Interactor.Store(u)
	ErrorStatus(w, err, 500)
	log.Println("Created User!!")
}

// FindAll ...
func (c *UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	users, err := c.Interactor.Index()
	ErrorStatus(w, err, 500)
	err = json.NewEncoder(w).Encode(users)
	ErrorStatus(w, err, 500)
}

// Update ...
func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	u := domain.User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	ErrorStatus(w, err, 500)
	err = c.Interactor.Update(u)
	ErrorStatus(w, err, 500)
	log.Println("Updated User!!")
}

// Delete ...
func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	ErrorStatus(w, err, 404)
	err = c.Interactor.Delete(id)
	ErrorStatus(w, err, 500)
	log.Println("Deleted User!!")
}
