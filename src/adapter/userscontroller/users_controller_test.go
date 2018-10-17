package userscontroller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rikiya/go-clean/src/adapter/userscontroller"
	"github.com/rikiya/go-clean/src/entity"
	"github.com/rikiya/go-clean/src/infrastructure/database"
)

func TestCreate(t *testing.T) {
	usersController := userscontroller.NewUserController(*database.NewSQLHandler())
	newUser := entity.User{
		FirstName: "Fujii",
		LastName:  "Rikiya",
	}
	b, err := json.Marshal(newUser)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("POST", "http://localhost:8080/users", bytes.NewBuffer(b))
	res := httptest.NewRecorder()
	usersController.Create(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("invalid code: %d", res.Code)
	}
}

func TestFindAll(t *testing.T) {
	usersController := userscontroller.NewUserController(*database.NewSQLHandler())
	req := httptest.NewRequest("GET", "http://localhost:8080/users", nil)
	res := httptest.NewRecorder()
	usersController.FindAll(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("invalid code: %d", res.Code)
	}
}

func TestUpdate(t *testing.T) {
	usersController := userscontroller.NewUserController(*database.NewSQLHandler())
	user := entity.User{
		FirstName: "Fujii",
		LastName:  "rikiya",
	}
	b, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("PUT", "http://localhost:8080/users/2", bytes.NewBuffer(b))
	res := httptest.NewRecorder()
	usersController.Update(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("invalid code: %d", res.Code)
	}
}

func TestDelete(t *testing.T) {
	usersController := userscontroller.NewUserController(*database.NewSQLHandler())
	req := httptest.NewRequest("DELETE", "http://localhost:8080/users/3", nil)
	res := httptest.NewRecorder()
	usersController.Delete(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("invalid code: %d", res.Code)
	}
}
