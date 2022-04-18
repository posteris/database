package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestGetFactory(t *testing.T) {
	instance := GetFactory()

	assert.NotNil(t, instance)
	assert.NotNil(t, instance.instance)
}

func TestFactory_SetInstance(t *testing.T) {
	//nolint
	os.Setenv(DatabaseTypeLabel, "sqlite")
	//nolint
	os.Setenv(DatabaseDsnLabel, "./6ea4bc80-d7d2-447b-b2a3-1b47efbc9e8f.db")

	factory := GetFactory()

	gormInstance, err := NewFromEnv(nil)
	if err != nil {
		t.Errorf("Connection Error: %v", err)
	}

	factory.SetInstance("test", gormInstance)

	assert.NotEmpty(t, factory.instance)
}

func TestFactory_GetInstance(t *testing.T) {
	os.Unsetenv(DatabaseTypeLabel)
	os.Unsetenv(DatabaseDsnLabel)

	factory := GetFactory()

	gormInstance, err := NewFromEnv(nil)
	if err != nil {
		t.Errorf("Connection Error: %v", err)
	}

	factory.SetInstance("test", gormInstance)

	instance := factory.GetInstance("test")
	assert.NotNil(t, instance)

	notInstance := factory.GetInstance("not-exist")
	assert.Nil(t, notInstance)
}

func TestFactory_AddFromEnv(t *testing.T) {
	type args struct {
		key    string
		dbType string
		dsn    string
		config *gorm.Config
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "SQLite",
			args: args{
				key:    "f060d2a2-5169-4e27-b0ba-29d8d9a52fa2",
				dbType: "sqlite",
				dsn:    "./f060d2a2-5169-4e27-b0ba-29d8d9a52fa2.db",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			want: true,
		},
		{
			name: "SQLite",
			args: args{
				key:    "975c39a1-fa95-4c2b-924c-baf67d6d7763",
				dbType: "sqlite",
				dsn:    "/usr/bin/f060d2a2-5169-4e27-b0ba-29d8d9a52fa2.db",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			factory := GetFactory()
			if got := factory.AddNew(test.args.key, test.args.dbType, test.args.dsn, test.args.config); got != test.want {
				t.Errorf("Factory.AddNew() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestFactory_AddNew(t *testing.T) {
	type args struct {
		key    string
		dbType string
		dsn    string
		config *gorm.Config
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "None",
			args: args{
				key:    "efb5fe25-8923-4217-ab5c-bc96c1749780",
				dbType: "",
				dsn:    "",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			want: true,
		},
		{
			name: "SQLite",
			args: args{
				key:    "f060d2a2-5169-4e27-b0ba-29d8d9a52fa2",
				dbType: "sqlite",
				dsn:    "./f060d2a2-5169-4e27-b0ba-29d8d9a52fa2.db",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			want: true,
		},
		{
			name: "SQLiteError",
			args: args{
				key:    "975c39a1-fa95-4c2b-924c-baf67d6d7763",
				dbType: "sqlite",
				dsn:    "/usr/bin/f060d2a2-5169-4e27-b0ba-29d8d9a52fa2.db",
				config: &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			},
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.args.dbType == "" {
				os.Unsetenv(DatabaseTypeLabel)
			} else {
				os.Setenv(DatabaseTypeLabel, test.args.dbType)
			}

			if test.args.dsn == "" {
				os.Unsetenv(DatabaseDsnLabel)
			} else {
				os.Setenv(DatabaseDsnLabel, test.args.dsn)
			}

			factory := GetFactory()

			if got := factory.AddFromEnv(test.args.key, test.args.config); got != test.want {
				t.Errorf("Factory.AddFromEnv() = %v, want %v", got, test.want)
			}
		})
	}
}
