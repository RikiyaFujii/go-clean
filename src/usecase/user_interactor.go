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
	err := ui.UserRepository.Store(u)
	if err != nil {
		return err
	}
	return nil
}

// Index ...
func (ui *UserInteractor) Index() ([]domain.User, error) {
	users, err := ui.UserRepository.Index()
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

// Update ...
func (ui *UserInteractor) Update(u domain.User) error {
	err := ui.UserRepository.Update(u)
	if err != nil {
		return err
	}
	return nil
}

// Delete ...
func (ui *UserInteractor) Delete(id int) error {
	err := ui.UserRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
