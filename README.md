# Gorm - Connect to all databases (Work in Progress)

This is a generic connector interface for the following databases:

* MySQL
* Postgres
* SQLite

With this connector, it is easy to run your sql based golang software,
written with the gorm.io framework,
on different databases, by only changing one environment variable.

Just set `SQL_TYPE` to either `MYSQL`, `POSTGRES`, `SQLITE` and you are good to go. 

## Installation

```
go get github.com/SbstnErhrdt/go-gorm-all-sql
```


## Todo
* Tests
* Docker compose environment

## MySQL
Env variables
```
SQL_HOST=
SQL_USER=
SQL_PASSWORD=
SQL_PORT=
SQL_DATABASE=
```

## Postgres
```
SQL_HOST=
SQL_USER=
SQL_PASSWORD=
SQL_PORT=
SQL_DATABASE=
SQL_SSL= // optional
```
## SQLite
```
SQL_DATABASE=
```

## Examples
```
ConnectToDatabase()
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