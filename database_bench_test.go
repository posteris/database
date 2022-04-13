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

func BenchmarkGetAllowedDB(b *testing.B) {
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
		b.Run(tt.name, func(b *testing.B) {
			got := getAllowedDB()

			sort.Strings(got)
			sort.Strings(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				b.Errorf("GetAllowedDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_getDatabaseType(b *testing.B) {
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
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			if tt.name != "none" {
				os.Setenv("DATABASE_TYPE", tt.name)
			} else {
				os.Unsetenv("DATABASE_TYPE")
			}

			if got := getDatabaseType(); got != tt.want {
				b.Errorf("getDatabaseType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkNew(b *testing.B) {
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
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			_, err := new(tt.args.dbType, tt.args.dsn, tt.args.config)

			if (err != nil) != tt.wantErr {
				b.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func BenchmarkNew_with_migrations(b *testing.B) {
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
			name: "postgres-successfull",
			args: args{
				dbType: "postgres",
				dsn:    "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: false,
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
			name: "mysql-successfull",
			args: args{
				dbType: "mysql",
				dsn:    "root:mysql@tcp(localhost:3306)/mysql?parseTime=true",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			db, err := new(tt.args.dbType, tt.args.dsn, tt.args.config)

			if (err != nil) != tt.wantErr {
				b.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				migration_error := db.AutoMigrate(&someTest{})
				if migration_error != nil {
					b.Errorf("Unable to migrate model'%v'", &someTest{})
					return
				}

				test := &someTest{
					Name: uuid.NewString(),
				}

				db.Save(test)

				test_search := someTest{}
				db.Last(&test_search)

				if test_search.ID == 0 {
					b.Errorf("error ID == 0, want == %d", test_search.ID)
					return
				}

				if test_search.Name != test.Name {
					b.Errorf("error Name == %s, want == %s", test_search.Name, test.Name)
					return
				}
			}
		})
	}
}
