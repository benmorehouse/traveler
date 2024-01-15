package utils

import (
	"database/sql"
)

// InitDatabase will return a single database connection
func InitDatabase() (*sql.DB, error) {
	return sql.Open("mysql", "user:password@/dbname")
}
