package database

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	//default database type. It was set as postgres because it was the first one that I've tried
	defaultDatabase string = "postgres"
	//Postgres DNF. by default it's point to localhost
	defaultDnf string = "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
)

//dialectSelector type to help select the database dialect
type dialectSelector func(dsn string) gorm.Dialector

//dialectors GORM database dialector map.
var dialectors map[string]dialectSelector = map[string]dialectSelector{
	"postgres": postgres.Open,
	"mysql":    mysql.Open,
	"sqlite":   sqlite.Open,
}

//getEnv function to obtains the environment data or the default fallback
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

//getDatabaseType function to get the database type from the environment var
func getDatabaseType() string {
	return getEnv("DATABASE_TYPE", defaultDatabase)
}

//getDatabaseDsn function to get the database DSN from the environment var
func getDatabaseDsn() string {
	return getEnv("DATABASE_DSN", defaultDnf)
}

//getAllowedDB function to return all database implementations possibilities.
//this function base itself in the dialects map, and use a key stractor to obtains
//the options array.
//
//Now it's just used to show the error message when the system try to start with a
//unknown database type.
func getAllowedDB() []string {
	var keys []string

	for key := range dialectors {
		keys = append(keys, key)
	}

	return keys
}

//NewInstance function that returns a new gorm.DB based on environment variables. Once it was
//not set, the default one is used. By default the selected one is postgres pointed to localhost
func NewInstance(config *gorm.Config) (*gorm.DB, error) {
	dbType := getDatabaseType()
	dbDsn := getDatabaseDsn()

	//create GORM database connection
	if config == nil {
		config = &gorm.Config{}
	}

	return new(dbType, dbDsn, config)
}

//new function to start a new connection based on a database type and a dsn connection
//string. Once connection is created this function returns an gorm database object.
//There is also the opportunity to send by parameter the GORM configuration. If it
//was not provided, the empty one will be created.
func new(dbType string, dsn string, config *gorm.Config) (*gorm.DB, error) {
	//obtain the GORM database dialector function
	dial := dialectors[dbType]

	//return error em case of GORM database dialector function does not exist.
	if dial == nil {
		err := fmt.Sprintf(
			"Database %s doesn't allowed yet! allowed types: %v",
			dbType,
			getAllowedDB(),
		)

		return nil, errors.New(err)
	}

	return gorm.Open(dial(dsn), config)
}
