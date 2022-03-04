# Gorm - Connect to all databases (Work in Progress)

This is a generic connector interface for the following databases:

* MySQL
* Postgres
* SQLite

With this connector, it is easy to run your sql based golang software, written with the gorm.io framework, on different
databases, by only changing one environment variable.

Just set `SQL_TYPE` to either `MYSQL`, `POSTGRES`, `SQLITE` and you are good to go.

## Installation

```shell
go get github.com/SbstnErhrdt/go-gorm-all-sql
```

In a development setting you can make use of the `env` package to load variables from a .env file into the ENV.

```shell
go get github.com/SbstnErhrdt/env
```

## Examples

```go
package main

import (
	"github.com/SbstnErhrdt/go-gorm-all-sql/pkg/sql"
)

func main() {
	// load env
	env.LoadEnvFiles() // reads the .env file in the working directory
	// connect to DB
	db, err := sql.ConnectToDatabase()
}
```

## Todo

* Tests
* Docker compose environment

## MySQL

Env variables

```
SQL_TYPE=MYSQL
SQL_HOST=mysql.server.com
SQL_PORT=3306
SQL_USER=sql_user
SQL_PASSWORD=xxxxxx
SQL_DATABASE=test
```

## Postgres

```
SQL_TYPE=POSTGRES
SQL_HOST=postgres.server.com
SQL_PORT=5432
SQL_USER=sql_user
SQL_PASSWORD=xxxxxx
SQL_DATABASE=test
SQL_SSL= // optional
```

## SQLite

```
SQL_TYPE=SQLITE
SQL_DATABASE=data.db
```

## Logger

to add the logrus logger, add the following environment flag

```
SQL_LOGGER=LOGRUS
```

## Functionality

* Works with different SQL dialects
* Checks if necessary environment variables are present
* Retries multiple times, if there is no connection available

## Test

### MySQL

```
docker-compose run --rm
```
