package database

import (
	"io/ioutil"
	"log"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func BenchmarkConnection(b *testing.B) {
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
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			_, err := New(test.args.dbType, test.args.dsn, test.args.config)

			if (err != nil) != test.wantErr {
				b.Errorf("New() error = %v, wantErr %v", err, test.wantErr)
				return
			}
		})
	}
}
