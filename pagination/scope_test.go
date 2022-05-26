package pagination

import (
	"os"
	"testing"

	"github.com/google/uuid"
	conndata "github.com/posteris/ci-database-starter/conn-data"
	"github.com/posteris/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type testPaginate struct {
	Test string `gorm:"string"`
}

func TestPaginate(t *testing.T) {
	//obtains the connection database parameters from the ci-database-starter
	databases := conndata.GetTestData()

	for _, dbData := range databases {
		t.Run(dbData.Name, func(t *testing.T) {
			os.Setenv(database.DatabaseTypeLabel, dbData.Type)
			os.Setenv(database.DatabaseDsnLabel, dbData.DSN)

			db, err := database.NewFromEnv(&gorm.Config{})
			if err != nil {
				t.Errorf("Connection error %v", err)
			}

			db.Migrator().DropTable(testPaginate{})
			db.AutoMigrate(testPaginate{})

			for count := 0; count < 100; count++ {
				db.Create(&testPaginate{
					Test: uuid.NewString(),
				})
			}

			type args struct {
				value      interface{}
				pagination *Pagination
				page       int
			}
			tests := []struct {
				name string
				args args
				want func(db *gorm.DB) *gorm.DB
			}{
				{
					name: "page-1",
					args: args{
						value: []*testPaginate{},
						pagination: &Pagination{
							Limit: 10,
						},
						page: 10,
					},
				},
			}
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					db.Scopes(Paginate(tt.args.value, tt.args.pagination, db)).Find(&tt.args.value)
					tt.args.pagination.Rows = tt.args.value

					assert.Equal(t, tt.args.page, tt.args.pagination.TotalPages)
				})
			}
		})
	}
}
