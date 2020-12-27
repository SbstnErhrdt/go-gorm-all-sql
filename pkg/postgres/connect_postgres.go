package postgres

import (
	"fmt"
	"github.com/SbstnErhrdt/go-gorm-all-sql/pkg/environment"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var requiredEnvironmentVariablesForPostGres = []string{
	"SQL_HOST",
	"SQL_USER",
	"SQL_PASSWORD",
	"SQL_PORT",
	"SQL_DATABASE",
}

// Creates a connection to a postgres database
func ConnectToPostgres(config *gorm.Config) (client *gorm.DB, err error) {
	// check env variables
	environment.CheckEnvironmentVariables(requiredEnvironmentVariablesForPostGres)
	// required env variables
	host := os.Getenv("SQL_HOST")
	user := os.Getenv("SQL_USER")
	pw := os.Getenv("SQL_PASSWORD")
	port := os.Getenv("SQL_PORT")
	dbName := os.Getenv("SQL_DATABASE")
	// Connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, dbName, pw)
	// check if optional env variable ssl is present
	ssl := os.Getenv("SQL_SSL")
	if len(ssl) > 0 {
		// extend dsn
		dsn = fmt.Sprintf("%s sslmode=%s", dsn, ssl)
	}
	// Connect to db
	client, err = gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Println("Postgres Client: error:", err)
		return nil, err
	}
	log.Println("Postgres Client: connected")
	return
}
