package database

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"

	"gorm.io/gorm"
)

const (
	//default database type. It was set as sqlite because it was the first one that I've tried
	defaultDatabase string = "sqlite"
	//sqlite file
	defaultDsn string = "database.db"
)

//dialectSelector type to help select the database dialect
type dialectSelector func(dsn string) gorm.Dialector

//dialectors GORM database dialector map.
var dialectors map[string]dialectSelector = map[string]dialectSelector{
	"clickhouse": clickhouse.Open,
	"mssql":      sqlserver.Open,
	"mysql":      mysql.Open,
	"postgres":   postgres.Open,
	"sqlite":     sqlite.Open,
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
	return getEnv("DATABASE_DSN", defaultDsn)
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

//EnvInstance function that returns a new gorm.DB based on environment variables. Once it was
//not set, the default one is used. By default the selected one is postgres pointed to localhost
func EnvInstance(config *gorm.Config) (*gorm.DB, error) {
	dbType := getDatabaseType()
	dbDsn := getDatabaseDsn()

	return New(dbType, dbDsn, config)
}

//New function to start a New connection based on a database type and a dsn connection
//string. Once connection is created this function returns an gorm database object.
//There is also the opportunity to send by parameter the GORM configuration. If it
//was not provided, the empty one will be created.
func New(dbType string, dsn string, config *gorm.Config) (*gorm.DB, error) {
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

	//create GORM database connection
	if config == nil {
		config = &gorm.Config{}
	}

	return gorm.Open(dial(dsn), config)
}
