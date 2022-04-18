# database

[![CI/CD](https://github.com/posteris/database/workflows/CI/badge.svg)](https://github.com/posteris/database/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/posteris/database.svg)](https://pkg.go.dev/github.com/posteris/database)
[![GitHub license](https://badgen.net/github/license/posteris/database)](https://github.com/posteris/database/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/posteris/database.svg)](https://GitHub.com/posteris/database/releases/)

This lib was created to help a microservice environment creation preventing the same process from being done many times, and their objective is to create a simple way to chose the database just passing the type and the database DSN.

There ara two kind of this lib usage, the first one is use the function __New__ and pass all required parameters, and the other one is using the function __FromEnv__, that returns an database connection based on environment variables.

The follow code shows how to use the __FromEnv__ function. Note for this function, you should set the __DATABASE_TYPE__ and __DATABASE_DSN__ environment variables, otherwise the lib will try to create a __sqlite__ database and connect to.

The available databases are:
* clickhouse
* mssql
* mysql
* postgres
* __sqlite__ _(Default)_

```go
//this function accept nil or a custom *gorm.Config object pointer
dbInstance, err := database.FromEnv(nil) 
```

The follow code shows how to use the __New__ function.

```go 
dbType := "postgres"
dbDSN  := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"

//this function accept nil for a *gorm.Config object pointer
dbInstance, err := database.New(dbType, dbDSN, nil)
```

Another way to user this lib is add new instance to a factory map; in the next sub section you can see the Factory map usage.

### Factory Map


```go
//create or get factory
factory := database.GetFactory()

//create new instance based at environment variables and store it at the instance map
factory.AddFromEnv("some-instance-name", nil)
```


```go
//create or get factory
factory := database.GetFactory()

//create new instance and store it at the instance map
factory.AddNew("some-instance-name", "sqlite", "./database.db", nil)
```

```go
//create or get factory
factory := database.GetFactory()

dbInstance, err := database.FromEnv(nil)
if err == nil {
    log.Fatal("error message")
}

factory.SetInstance("some-instance-name", dbInstance)
```

```go
//obtains the instance from factory map
instance := factory.GetInstance("some-instance-name")
```


## Software Quality
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=posteris_database&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=posteris_database)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=posteris_database&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=posteris_database)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=posteris_database&metric=coverage)](https://sonarcloud.io/summary/new_code?id=posteris_database)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=posteris_database&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=posteris_database)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=posteris_database&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=posteris_database)


This lib use [Sonarcloud](https://sonarcloud.io/) to help understend the code quality and security.

In conjunction with [Sonarcloud](https://sonarcloud.io/), this lib uses [Horusec](https://horusec.io/) which blocks CI/CD in any vulnerability incidence


## Benchmark

Thinking in the software quality, the __benchmark regression__ was created. It's can be viewed at the link bellow.

[Performance Regeression](https://posteris.github.io/database/dev/bench/)