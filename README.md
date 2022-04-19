# database

[![CI/CD](https://github.com/posteris/database/workflows/CI/badge.svg)](https://github.com/posteris/database/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/posteris/database.svg)](https://pkg.go.dev/github.com/posteris/database)
[![GitHub license](https://badgen.net/github/license/posteris/database)](https://github.com/posteris/database/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/posteris/database.svg)](https://GitHub.com/posteris/database/releases/)

This lib was created to help a microservice environment creation, always preventing the same process from being done many times.

The main objective is to create a simple way to chose the database just passing the type and DSN.

focused on your goal, this library is ready to use with the follow databases (_database type_).

* clickhouse
* mssql
* mysql
* postgres
* __sqlite__ _(Default)_


:pencil2: __Note 1:__ _This lib have __NO__ intention to replace __GORM__ or other library!_ 

:pencil2: __Note 2:__ _There is an overweight in this lib due to all database drivers import._ 

:pencil2: __Note 3:__ _This lib want to facilitate the database changes without change the code._

## Usage

For this library, there are two types of usage, which are described below:

* __Core Library:__ Allows to create the connection based on _database type_ and _DSN_.
* __Factory Map:__ Creates a factory map that able to store many _*gorm.DB_ instances and store it referencing them by a unique key.

### Core Library

The core library is formed by 2 functions that creates and return the _*gorm.DB_ instance. 

The __New__ function is generic and allows user to create the instance passing the _database type_ and the _DSN_ by parameter. You also able to send the _*gorm.Config_ object or nil.

Follow you can see an usage example of __New__ function. In this example the _database type_ selected is _sqlite_ and the _DSN_ is _./database.db_.

```go 
db, err := database.New(database.SQLite, "./database.db", nil)
```

The other function to obtains and return the _*gorm.DB_ object is __NewFromEnv__, that able to create the instance based on the environment variables __DATABASE_TYPE__ and __DATABASE_DSN__. 

In case of environment vars not present, the system will choose _database.SQLite_ as database type and _database.db_ as _DSN_.

You can see an usage example below:

```go
db, err := database.FromEnv(nil) 
```

### Factory Map

Before start to use the factory we need to get a factory instance executing the follow command:

```go
factory := database.GetFactory()
```

This code will genarate a factory that able users to store many _*gorm.DB_ instances, always referencing them by an unique key. 

The factory object has three ways to create or insert a _*gorm.DB_ instance and one to recovery it.

To make an existent instance to the factory you just neet to call the __SetInstance__ function as showed below.

```go
db, err := database.FromEnv(nil) 

if err == nil {
    log.Fatal("error message")
}

factory.SetInstance("instance-name", db)
```

Beyond the __SetInstance__ function, the factory has two ways to crete and insert the instance in the factory map automatically.

The function __AddNew__ create and store the _*gorm.DB_ instance to the factory map using the provided unique key. e.g:

```go
factory.AddNew(
    "instance-name", 
    database.SQLite, 
    "./database.db", 
    nil,
)
```

Like __AddNew__, you can use the __AddFromEnv__ to create and add the _*gorm.DB_ instance to the factory map, but this time based on environment variables.


```go
factory.AddFromEnv("instance-name", nil)
```

:pencil2: __Note:__ _this function follow the same env vars of __NewFromEnv__ function._

To obtains the _*gorm.DB_ instance from the factory map you can use the __GetInstance__ function like showed below.

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