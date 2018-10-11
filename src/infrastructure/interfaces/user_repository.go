package interfaces

import (
	"github.com/rikiya/go-clean/src/domain"
)

// UserRepository ...
type UserRepository interface {
	Store(domain.User) error
	Index() ([]domain.User, error)
	Update(domain.User) error
	Delete(id int) error
}
