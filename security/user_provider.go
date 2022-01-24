package security

import (
	"database/sql"
	"gftp/server"
)

type SQLiteUserProvider struct {
	dbConn *sql.DB
}

func NewSQLiteUserProvider(dbConn *sql.DB) *SQLiteUserProvider {
	return &SQLiteUserProvider{
		dbConn: dbConn,
	}
}

func (p *SQLiteUserProvider) GetUser(username string) (*server.User, error) {
	ret := p.dbConn.QueryRow("select root from users where username = ?", username)
	var root string
	ret.Scan(&root)
	return &server.User{RootPath: root}, nil
}

func (p *SQLiteUserProvider) Exists(username string) (bool, error) {
	ret := p.dbConn.QueryRow("select count(*) from users where username = ?", username)
	var count int

	if err := ret.Scan(&count); err == sql.ErrNoRows {
		return false, err
	}

	return count == 1, nil
}
