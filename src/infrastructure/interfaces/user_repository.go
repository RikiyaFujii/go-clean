package interfaces

import (
	"database/sql"

	"github.com/rikiya/go-clean/src/domain"
)

// UserRepository ...
type UserRepository interface {
	Store(domain.User) (sql.Result, error)
}
