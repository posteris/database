package database

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"sort"
	"testing"

	"gorm.io/gorm"
)

func TestGetAllowedDB(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	tests := []struct {
		name string
		want []string
	}{
		{
			name: "all-databases",
			want: []string{"postgres", "mysql", "sqlite", "clickhouse", "mssql"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getAllowedDB()

			sort.Strings(got)
			sort.Strings(test.want)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("GetAllowedDB() = %v, want %v", got, test.want)
			}
		})
	}
}

func Test_getDatabaseType(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	os.Unsetenv("DATABASE_TYPE")

	tests := []struct {
		name string
		want string
	}{
		{
			name: "none",
			want: "sqlite",
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
			name: "clickhouse",
			want: "clickhouse",
		},
		{
			name: "mssql",
			want: "mssql",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.name != "none" {
				os.Setenv("DATABASE_TYPE", test.name)
			}

			if got := getDatabaseType(); got != test.want {
				t.Errorf("getDatabaseType() = %v, want %v", got, test.want)
			}
		})
	}
}

type someTest struct {
	*gorm.Model
	Name string `json:"name"`
}
