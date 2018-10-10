package usecase

import (
	"github.com/rikiya/go-clean/src/domain"
	"github.com/rikiya/go-clean/src/infrastructure/interfaces"
)

// UserInteractor ...
type UserInteractor struct {
	UserRepository interfaces.UserRepository
}

// Store ...
func (ui *UserInteractor) Store(u domain.User) error {
	_, err := ui.UserRepository.Store(u)
	if err != nil {
		return err
	}
	return nil
}
