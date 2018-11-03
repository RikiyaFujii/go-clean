# go-clean-architecture

## example

### adapter
It describes the processing of controller for calling by routing.

```go
package userscontroller

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/rikiya/go-clean/src/adapter/errorlog"
    "github.com/rikiya/go-clean/src/entity"
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
    u := entity.User{}
    err := json.NewDecoder(r.Body).Decode(&u)
    errorlog.ErrorStatus(w, err, http.StatusBadRequest)
    err = c.Interactor.Store(u)
    errorlog.ErrorStatus(w, err, http.StatusInternalServerError)
    log.Println("Created User!!")
}
```

```go
package errorlog

import (
    "log"
    "net/http"
)

// ErrorStatus ...
func ErrorStatus(w http.ResponseWriter, err error, status int) {
    if err != nil {
        log.Println("Error: ", err)
        w.WriteHeader(status)
        return
    }
}
```

### db
It is a layer for migration and has DDL statements and migration files.

```sql
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `users`(
    `id` INT (11) NOT NULL AUTO_INCREMENT,
    `first_name` VARCHAR(256) NOT NULL,
    `last_name` VARCHAR(256) NOT NULL,
    PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `users`;
```

```yaml
development:
    driver: mysql
    open: root:@/clean_go?parseTime=true&charset=utf8mb4&interpolateParams=true
```

### entity
A data structure describing business rules.

```go
package entity

// User ...
type User struct {
    ID        int    `json:"id" db:"id"`
    FirstName string `json:"first_name" db:"first_name"`
    LastName  string `json:"last_name" db:"last_name"`
}
```

### external
Call controller on layer describing API routing.

```go
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
}
```

### infrastructure
Definition of interface handled by Handler, Model layer (impl), Model for accessing DB.
```go
package interfaces

import (
    "database/sql"
)

// SQLHandler ...
type SQLHandler interface {
    Prepare(string) (*sql.Stmt, error)
}
```

```go
package interfaces

import (
    "github.com/rikiya/go-clean/src/entity"
)

// UserRepository ...
type UserRepository interface {
    Store(entity.User) error
}
```

```go
package database

import (
    "database/sql"

    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

// SQLHandler ...
type SQLHandler struct {
    Conn *sqlx.DB
}

// NewSQLHandler ...
func NewSQLHandler() *SQLHandler {
    ctx, err := sqlx.Open("mysql", "root:@/clean_go?parseTime=true&charset=utf8mb4&interpolateParams=true")
    if err != nil {
        panic(err)
    }
    sqlHandler := new(SQLHandler)
    sqlHandler.Conn = ctx
    return sqlHandler
}

// Prepare ...
func (s *SQLHandler) Prepare(query string) (*sql.Stmt, error) {
    result, err := s.Conn.Prepare(query)
    if err != nil {
        return nil, err
    }
    return result, nil
}
```

```go
package user

import (
    "github.com/rikiya/go-clean/src/entity"
    "github.com/rikiya/go-clean/src/infrastructure/database"
)

// UserImpl ...
type UserImpl struct {
    database.SQLHandler
}

// Store ...
func (ui *UserImpl) Store(u entity.User) error {
    ctx := database.NewSQLHandler()
    res, err := ctx.Prepare(
        `INSERT INTO users (first_name, last_name) VALUES(?, ?)`,
    )
    if err != nil {
        return err
    }
    defer res.Close()

    _, err = res.Exec(u.FirstName, u.LastName)
    return err
}
```

### usecase
Includes application-specific business rules. Encapsulate and implement all system use cases.

```go
package usecase

import (
    "github.com/rikiya/go-clean/src/entity"
    "github.com/rikiya/go-clean/src/infrastructure/interfaces"
)

// UserInteractor ...
type UserInteractor struct {
    UserRepository interfaces.UserRepository
}

// Store ...
func (ui *UserInteractor) Store(u entity.User) error {
    err := ui.UserRepository.Store(u)
    if err != nil {
        return err
    }
    return nil
}
```