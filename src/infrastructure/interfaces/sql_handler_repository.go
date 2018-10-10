package interfaces

import (
	"database/sql"
)

// SQLHandler ...
type SQLHandler interface {
	Prepare(string) (*sql.Stmt, error)
}
