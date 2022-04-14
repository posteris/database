package database

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"sort"
	"testing"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestGetAllowedDB(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	tests := []struct {
		name string
		want []string
	}{
		{
			name: "all-databases",
			want: []string{"postgres", "mysql", "sqlite"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getAllowedDB()

			sort.Strings(got)
			sort.Strings(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllowedDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDatabaseType(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	tests := []struct {
		name string
		want string
	}{
		{
			name: "none",
			want: "postgres",
		},
		{
			name: "postgres",
			want: "postgres",
		},
		{
			name: "sqlite",
			want: "sqlite",
		},
		{
			name: "mysql",
			want: "mysql",
		},
		{
			name: "oracle",
			want: "oracle",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name != "none" {
				os.Setenv("DATABASE_TYPE", tt.name)
			}

			if got := getDatabaseType(); got != tt.want {
				t.Errorf("getDatabaseType() = %v, want %v", got, tt.want)
			}
		})
	}
}

type someTest struct {
	*gorm.Model
	Name string `json:"name"`
}

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
			name: "mysql-error",
			args: args{
				dbType: "mysql",
				dsn:    "no-metter-anymore",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: true,
		},
		{
			name: "oracle-successfull",
			args: args{
				dbType: "oracle",
				dsn:    "system/oracle@localhost:1521/xe",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: false,
		},
		{
			name: "oracle-error",
			args: args{
				dbType: "oracle",
				dsn:    "no-metter-anymore",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := new(tt.args.dbType, tt.args.dsn, tt.args.config)

			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

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
			name: "Oracle",
			args: args{
				dbType: "oracle",
				dsn:    "system/oracle@localhost:1521/xe",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := new(tt.args.dbType, tt.args.dsn, tt.args.config)

			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				migration_error := db.AutoMigrate(&someTest{})
				if migration_error != nil {
					t.Errorf("Unable to migrate model'%v'", &someTest{})
					return
				}

				test := &someTest{
					Name: uuid.NewString(),
				}

				db.Save(test)

				test_search := someTest{}
				db.Last(&test_search)

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
