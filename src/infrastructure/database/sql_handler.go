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
	conn, err := sqlx.Open("mysql", "root:@/clean_go?parseTime=true&charset=utf8mb4&interpolateParams=true")
	if err != nil {
		panic(err)
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn
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
