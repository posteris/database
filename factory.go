package database

import (
	"gorm.io/gorm"
)

type Factory struct {
	instance map[string]*gorm.DB
}

var factoryInstance *Factory

func GetFactory() *Factory {
	if factoryInstance == nil {
		factoryInstance = &Factory{
			instance: make(map[string]*gorm.DB),
		}
	}

	return factoryInstance
}

//GetInstance function to obtains the instance from wallet based on
//a providede key. In case of key not exist, this method will return
//nil
func (factory *Factory) GetInstance(key string) *gorm.DB {
	dbInstance := factory.instance[key]

	if dbInstance != nil {
		return dbInstance
	}

	return nil
}

//SetInstance function to add new gorm instance to the wallet
func (factory *Factory) SetInstance(key string, value *gorm.DB) bool {
	factory.instance[key] = value

	return true
}

//AddFromEnv function that add a new instance based on a environment
//variable
func (factory *Factory) AddFromEnv(key string, config *gorm.Config) bool {
	dbInstance, err := NewFromEnv(config)

	if err != nil {
		return false
	}

	return factory.SetInstance(key, dbInstance)
}

//AddNew function that add a new instance based on passed parameters
func (factory *Factory) AddNew(key string, dbType string, dsn string, config *gorm.Config) bool {
	dbInstance, err := New(dbType, dsn, config)

	if err != nil {
		return false
	}

	return factory.SetInstance(key, dbInstance)
}
