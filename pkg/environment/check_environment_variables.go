package environment

import (
	"errors"
	"log"
	"os"
)

func CheckEnvironmentVariables(requiredEnvironmentVariablesForSQL []string) {
	// Check necessary env variables
	// init err array
	var errs []error
	// iterate over necessary os vars
	for _, v := range requiredEnvironmentVariablesForSQL {
		if len(os.Getenv(v)) == 0 {
			msg := "Environment variable " + v + " is missing. Please add the variable to your environment to connect to the database"
			errs = append(errs, errors.New(msg))
			log.Println(msg)
		}
	}
	// panic if necessary env variable is missing
	if len(errs) != 0 {
		panic("Missing environment variables")
	}
}
