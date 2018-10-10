package user

import (
	"database/sql"

	"github.com/rikiya/go-clean/src/domain"
	"github.com/rikiya/go-clean/src/infrastructure/database"
)

// UserImpl ...
type UserImpl struct {
	database.SQLHandler
}

// Store ...
func (ui *UserImpl) Store(u domain.User) (sql.Result, error) {
	cnn := database.NewSQLHandler()
	res, err := cnn.Prepare(
		`INSERT INTO users (first_name, last_name) VALUES(?, ?)`,
	)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	return res.Exec(u.FirstName, u.LastName)
}
