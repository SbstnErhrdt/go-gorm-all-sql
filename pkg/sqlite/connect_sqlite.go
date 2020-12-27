package sqlite

import (
	"github.com/SbstnErhrdt/go-gorm-all-sql/pkg/environment"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

var requiredEnvironmentVariablesForSQLite = []string{
	"SQL_DATABASE",
}

func ConnectToSQLite(config *gorm.Config) (client *gorm.DB, err error) {
	// check env variables
	environment.CheckEnvironmentVariables(requiredEnvironmentVariablesForSQLite)
	dbName := os.Getenv("SQL_DATABASE")
	// connect to database
	client, err = gorm.Open(sqlite.Open(dbName), config)
	if err != nil {
		log.Println("SQLite Client: error:", err)
		return nil, err
	}
	log.Println("SQLite Client: connected")
	return
}
