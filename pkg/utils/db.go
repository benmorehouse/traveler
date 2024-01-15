package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// InitDatabase will return a single database connection
func InitDatabase() (*sql.DB, error) {
	return sql.Open("mysql", "user:password@/dbname")
}
