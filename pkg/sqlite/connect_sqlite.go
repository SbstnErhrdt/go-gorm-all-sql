package sqlite

import (
	"github.com/SbstnErhrdt/go-gorm-all-sql/pkg/environment"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
		log.Error("SQLite Client: error:", err)
		return nil, err
	}
	log.Info("SQLite Client: connected")
	return
}
