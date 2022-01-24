package security

import "database/sql"

type AuthContext struct {
	dbConn *sql.DB
}

func NewAuthContext(db *sql.DB) *AuthContext {
	return &AuthContext{
		dbConn: db,
	}
}
