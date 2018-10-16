package userscontroller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rikiya/go-clean/src/adapter/errorlog"
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
	errorlog.ErrorStatus(w, err, http.StatusBadRequest)
	err = c.Interactor.Store(u)
	errorlog.ErrorStatus(w, err, http.StatusInternalServerError)
	log.Println("Created User!!")
}

// FindAll ...
func (c *UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	users, err := c.Interactor.Index()
	errorlog.ErrorStatus(w, err, http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(users)
	errorlog.ErrorStatus(w, err, http.StatusInternalServerError)
}

// Update ...
func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
	id, err := strconv.Atoi(vars["id"])
	// errorlog.ErrorStatus(w, err, http.StatusBadRequest)

	u := domain.User{}
	err = json.NewDecoder(r.Body).Decode(&u)
	errorlog.ErrorStatus(w, err, http.StatusBadRequest)
	err = c.Interactor.Update(id, u)
	errorlog.ErrorStatus(w, err, http.StatusInternalServerError)
	log.Println("Updated User!!")
}

// Delete ...
func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	// errorlog.ErrorStatus(w, err, http.StatusBadRequest)
	err = c.Interactor.Delete(id)
	errorlog.ErrorStatus(w, err, http.StatusInternalServerError)
	log.Println("Deleted User!!")
}
