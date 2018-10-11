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
	ctx := database.NewSQLHandler()
	res, err := ctx.Prepare(
		`INSERT INTO users (first_name, last_name) VALUES(?, ?)`,
	)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	return res.Exec(u.FirstName, u.LastName)
}

// Index ...
func (ui *UserImpl) Index() (users []domain.User, err error) {
	ctx := database.NewSQLHandler()
	if err = ctx.Select(&users, `SELECT * FROM users`); err != nil {
		return nil, err
	}
	return users, nil
}

// Update ...
func (ui *UserImpl) Update(u domain.User) (sql.Result, error) {
	ctx := database.NewSQLHandler()
	res, err := ctx.Prepare(
		`UPDATE users SET first_name = ?, last_name = ? WHERE id = ?`,
	)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	return res.Exec(u.FirstName, u.LastName, u.ID)
}

// Delete ...
func (ui *UserImpl) Delete(id int) (sql.Result, error) {
	ctx := database.NewSQLHandler()
	res, err := ctx.Prepare(
		`DELETE FROM users where id = ?`,
	)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	return res.Exec(id)
}
