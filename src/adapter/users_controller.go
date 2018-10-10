package adapter

import (
	"encoding/json"
	"log"
	"net/http"

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
