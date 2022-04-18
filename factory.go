package database

import "gorm.io/gorm"

var instance = make(map[string]*gorm.DB)

//GetInstance function to obtains the instance from wallet based on
//a providede key. In case of key not exist, this method will return
//nil
func GetInstance(key string) *gorm.DB {
	dbInstance := instance[key]

	if dbInstance != nil {
		return dbInstance
	}

	return nil
}

//SetInstance function to add new gorm instance to the wallet
func SetInstance(key string, value *gorm.DB) bool {
	instance[key] = value

	return true
}

//AddFromEnv function that add a new instance based on a environment
//variable
func AddFromEnv(key string, config *gorm.Config) bool {
	dbInstance, err := EnvInstance(config)

	if err != nil {
		return false
	}

	return SetInstance(key, dbInstance)
}

//AddNew function that add a new instance based on passed parameters
func AddNew(key string, dbType string, dsn string, config *gorm.Config) bool {
	dbInstance, err := New(dbType, dsn, config)

	if err != nil {
		return false
	}

	return SetInstance(key, dbInstance)
}
