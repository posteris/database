package database

import (
	"io/ioutil"
	"log"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestNew_with_migrations(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	type args struct {
		dbType string
		dsn    string
		config *gorm.Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "PostgreSQL",
			args: args{
				dbType: "postgres",
				dsn:    "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: false,
		},
		{
			name: "SQLite",
			args: args{
				dbType: "sqlite",
				dsn:    "./test.db",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: false,
		},
		{
			name: "MySQL",
			args: args{
				dbType: "mysql",
				dsn:    "root:mysql@tcp(localhost:3306)/mysql?parseTime=true",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: false,
		},
		{
			name: "Clickhouse",
			args: args{
				dbType: "clickhouse",
				dsn:    "tcp://localhost:9000?database=default",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: false,
		},
		{
			name: "MSSQL",
			args: args{
				dbType: "mssql",
				dsn:    "sqlserver://sa:Adm1n123@localhost:1433?database=master",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			database, err := New(test.args.dbType, test.args.dsn, test.args.config)

			if (err != nil) != test.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, test.wantErr)
				return
			}

			if err == nil {
				//nolint
				database.AutoMigrate(&someTest{})
			}
		})
	}
}
