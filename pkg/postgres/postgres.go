package postgres

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/pkg/postgres/seed"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connection() *gorm.DB {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func Migrate(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&entity.User{}, &entity.UserRole{})
	if err != nil {
		log.Fatalf("auto migration failed: %s", err)
	}
}

func Seed(db *gorm.DB) {
	err := seed.RoleSeed(db)
	if err != nil {
		log.Fatalf("role seeding failed: %v\n", err)
	}

	err = seed.AdminSeed(db)
	if err != nil {
		log.Fatalf("user seeding failed: %v\n", err)
	}
}
