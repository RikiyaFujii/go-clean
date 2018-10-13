package adapter_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rikiya/go-clean/src/adapter"
	"github.com/rikiya/go-clean/src/domain"
	"github.com/rikiya/go-clean/src/infrastructure/database"
)

func TestCreate(t *testing.T) {
	usersController := adapter.NewUserController(*database.NewSQLHandler())
	newUser := domain.User{
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
	usersController := adapter.NewUserController(*database.NewSQLHandler())
	req := httptest.NewRequest("GET", "http://localhost:8080/users", nil)
	res := httptest.NewRecorder()
	usersController.FindAll(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("invalid code: %d", res.Code)
	}
}
