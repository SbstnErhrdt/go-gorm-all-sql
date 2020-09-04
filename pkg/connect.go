package pkg

import (
	"errors"
	"github.com/SbstnErhrdt/go-gorm-all-sql/pkg/mysql"
	"github.com/SbstnErhrdt/go-gorm-all-sql/pkg/postgres"
	"github.com/SbstnErhrdt/go-gorm-all-sql/pkg/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
	"time"
)
// Connects to a database with the default config
func ConnectToDatabase() (client *gorm.DB, err error) {
	config := gorm.Config{}
	return connect(&config)
}

// Connects to a database with a specific config
func ConnectToDatabaseWithConfig(config *gorm.Config) (client *gorm.DB, err error) {
	return connect(config)
}

// Extracts the sql type environment flag from the environment
// if there is no flag present the fallback is MYSQL
func getSQlType() string {
	// Get env variables
	sqlType := os.Getenv("SQL_TYPE")
	if len(sqlType) == 0 {
		sqlType = "MYSQL"
	}
	return strings.ToUpper(sqlType)
}

var FailedConnectionCounter = 10 // amount of retries to connect to the database

// Tries to connect to the database
// Does this recursively until there is a connection
// or until the counter is 0
func connect(config *gorm.Config) (client *gorm.DB, err error) {
	switch getSQlType() {
	case "MYSQL":
		client, err = mysql.ConnectToMysql(config)
	case "POSTGRES", "POSTGRESQL":
		client, err = postgres.ConnectToPostgres(config)
	case "SQLITE", "LITE", "SQLLITE":
		client, err = sqlite.ConnectToSQLite(config)
	default:
		err = errors.New("can not find this sql type")
		return
	}
	// if there are problems with connecting to the database
	// retry
	if err != nil {
		// If we can not connect we try again in 10 seconds
		log.Println("InitMysqlClient: error:", err)
		FailedConnectionCounter--
		// Base case
		if FailedConnectionCounter == 0 {
			// If for multiple consecutive attempts its not possible to connect
			// panic
			log.Fatal("InitSQLClient: FATAL", err)
			return
		}
		log.Println("InitMysqlClient: retry")
		time.Sleep(time.Second * 5)
		// Retry again
		return connect(config)
	}
	return
}
