package database

import (
	"io/ioutil"
	"log"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestNew(t *testing.T) {
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
			name: "unknown",
			args: args{
				dbType: "unknown",
				dsn:    "no-metter-anymore",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: true,
		},
		{
			name: "postgres-successfull",
			args: args{
				dbType: "postgres",
				dsn:    "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: false,
		},
		{
			name: "postgres-successfull-with-nil-config",
			args: args{
				dbType: "postgres",
				dsn:    "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable",
				config: nil,
			},
			wantErr: false,
		},
		{
			name: "postgres-error",
			args: args{
				dbType: "postgres",
				dsn:    "host=not-a-host port=5432 user=postgres password=postgres dbname=postgres sslmode=disable",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: true,
		},
		{
			name: "sqlite-successfull",
			args: args{
				dbType: "sqlite",
				dsn:    "./test.db",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: false,
		},
		{
			name: "sqlite-successfull-with-nil-config",
			args: args{
				dbType: "sqlite",
				dsn:    "./test.db",
				config: nil,
			},
			wantErr: false,
		},
		{
			name: "sqlite-error",
			args: args{
				dbType: "sqlite",
				dsn:    "/root/test.db", //permission denied
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: true,
		},
		{
			name: "mysql-successfull",
			args: args{
				dbType: "mysql",
				dsn:    "root:mysql@tcp(localhost:3306)/mysql?parseTime=true",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: false,
		},
		{
			name: "mysql-successfull-with-nil-config",
			args: args{
				dbType: "mysql",
				dsn:    "root:mysql@tcp(localhost:3306)/mysql?parseTime=true",
				config: nil,
			},
			wantErr: false,
		},
		{
			name: "mysql-error",
			args: args{
				dbType: "mysql",
				dsn:    "no-metter-anymore",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: true,
		},
		{
			name: "clickhouse-successfull",
			args: args{
				dbType: "clickhouse",
				dsn:    "tcp://localhost:9000?database=default",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: false,
		},
		{
			name: "clickhouse-successfull-with-nil-config",
			args: args{
				dbType: "clickhouse",
				dsn:    "tcp://localhost:9000?database=default",
				config: nil,
			},
			wantErr: false,
		},
		{
			name: "clickhouse-error",
			args: args{
				dbType: "clickhouse",
				dsn:    "no-metter-anymore",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: true,
		},
		{
			name: "mssql-successfull",
			args: args{
				dbType: "mssql",
				dsn:    "sqlserver://sa:Adm1n123@localhost:1433?database=master",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: false,
		},
		{
			name: "mssql-successfull-with-nil-config",
			args: args{
				dbType: "mssql",
				dsn:    "sqlserver://sa:Adm1n123@localhost:1433?database=master",
				config: nil,
			},
			wantErr: false,
		},
		{
			name: "mssql-error",
			args: args{
				dbType: "mssql",
				dsn:    "no-metter-anymore",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := New(test.args.dbType, test.args.dsn, test.args.config)

			if (err != nil) != test.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, test.wantErr)
				return
			}
		})
	}
}
