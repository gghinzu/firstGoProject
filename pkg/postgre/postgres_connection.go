package postgre

import (
	"database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connection method returns a postgres connection using GORM
func Connection() *gorm.DB {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable"

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return gormDB
}
