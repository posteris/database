package database

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/google/uuid"
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
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			database, err := New(test.args.dbType, test.args.dsn, test.args.config)

			if (err != nil) != test.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, test.wantErr)
				return
			}

			if err == nil {
				migration_error := database.AutoMigrate(&someTest{})
				if migration_error != nil {
					t.Errorf("Unable to migrate model'%v'", &someTest{})
					return
				}

				test := &someTest{
					Name: uuid.NewString(),
				}

				database.Save(test)

				test_search := someTest{}
				database.Last(&test_search)

				if test_search.ID == 0 {
					t.Errorf("error ID == 0, want == %d", test_search.ID)
					return
				}

				if test_search.Name != test.Name {
					t.Errorf("error Name == %s, want == %s", test_search.Name, test.Name)
					return
				}
			}
		})
	}
}
