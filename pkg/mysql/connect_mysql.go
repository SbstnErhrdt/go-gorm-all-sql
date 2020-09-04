package mysql

import (
	"fmt"
	"github.com/SbstnErhrdt/go-gorm-all-sql/pkg/environement"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var requiredEnvironmentVariablesForMySQL = []string{
	"SQL_HOST",
	"SQL_USER",
	"SQL_PASSWORD",
	"SQL_PORT",
	"SQL_DATABASE",
}

func ConnectToMysql(config *gorm.Config) (client *gorm.DB, err error) {
	// check env variables
	environement.CheckEnvironmentVariables(requiredEnvironmentVariablesForMySQL)
	// env variables
	host := os.Getenv("SQL_HOST")
	user := os.Getenv("SQL_USER")
	pw := os.Getenv("SQL_PASSWORD")
	port := os.Getenv("SQL_PORT")
	dbName := os.Getenv("SQL_DATABASE")
	// build connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pw, host, port, dbName)
	// connect to db
	client, err = gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		return nil, err
	}
	log.Println("MySQL Client: connected", err)
	return
}
