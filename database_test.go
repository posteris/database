package database

import (
	"reflect"
	"sort"
	"testing"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func TestGetAllowedDB(t *testing.T) {
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

type someTest struct {
	*gorm.Model
	Name string `gorm"not null" json:"name"`
}

func TestNew(t *testing.T) {
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
				config: &gorm.Config{},
			},
			wantErr: true,
		},
		{
			name: "postgres-successfull",
			args: args{
				dbType: "postgres",
				dsn:    "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable",
				config: &gorm.Config{},
			},
			wantErr: false,
		},
		{
			name: "postgres-error",
			args: args{
				dbType: "postgres",
				dsn:    "host=not-a-host port=5432 user=postgres password=postgres dbname=postgres sslmode=disable",
				config: &gorm.Config{},
			},
			wantErr: true,
		},
		{
			name: "sqlite-successfull",
			args: args{
				dbType: "sqlite",
				dsn:    "./test.db",
				config: &gorm.Config{},
			},
			wantErr: false,
		},
		{
			name: "sqlite-error",
			args: args{
				dbType: "sqlite",
				dsn:    "/root/test.db", //permission denied
				config: &gorm.Config{},
			},
			wantErr: true,
		},
		{
			name: "mysql-successfull",
			args: args{
				dbType: "mysql",
				dsn:    "root:mysql@tcp(localhost:3306)/mysql",
				config: &gorm.Config{},
			},
			wantErr: false,
		},
		{
			name: "mysql-error",
			args: args{
				dbType: "mysql",
				dsn:    "no-metter-anymore",
				config: &gorm.Config{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.args.dbType, tt.args.dsn, tt.args.config)

			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNew_with_migrations(t *testing.T) {
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
				config: &gorm.Config{},
			},
			wantErr: false,
		},
		{
			name: "sqlite-successfull",
			args: args{
				dbType: "sqlite",
				dsn:    "./test.db",
				config: &gorm.Config{},
			},
			wantErr: false,
		},
		{
			name: "mysql-successfull",
			args: args{
				dbType: "mysql",
				dsn:    "root:mysql@tcp(localhost:3306)/mysql",
				config: &gorm.Config{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := New(tt.args.dbType, tt.args.dsn, tt.args.config)

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
				db.First(&test_search)

				if test_search.ID == 0 || test_search.Name != test.Name {
					t.Errorf("element search error")
					return
				}
			}
		})
	}
}
