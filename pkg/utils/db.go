package utils

import (
	"database/sql"
	"fmt"

	"github.com/benmorehouse/traveler/config"
	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

// InitDatabase initializes and returns a database connection
func InitDatabase() (*sql.DB, error) {
	config := config.DefaultConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/traveler?parseTime=true", config.DBUser, config.DBPass)
	// Open a new connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Optional: Verify the connection is successful by pinging the database
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
