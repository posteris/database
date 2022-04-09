package database

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//dialectSelector type
type dialectSelector func(dsn string) gorm.Dialector

//dialectors GORM database dialector map.
var dialectors map[string]dialectSelector = map[string]dialectSelector{
	"postgres": postgres.Open,
	"mysql":    mysql.Open,
	"sqlite":   sqlite.Open,
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

//New function to start a new connection based on a database type and a dsn connection
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
