package seed

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/helper"
	"firstGoProject/pkg/config"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

func Seed(db *gorm.DB) error {
	configuration, errC := config.LoadConfig()
	if errC != nil {
		log.Fatal("cannot load config:", errC)
	}

	var count int64
	if db.Table("users").Count(&count); count == 0 {
		var adminRole *entity.UserRole
		err := db.Table("user_roles").Where("name = ?", enum.Admin).First(&adminRole).Error
		if err != nil {
			return err
		}

		hashedPassword, errP := helper.EncryptPassword(configuration.SeedPassword)
		if errP != nil {
			return errP
		}

		return db.Create(&entity.User{
			ID:        uuid.New(),
			Email:     configuration.SeedMail,
			Password:  string(hashedPassword),
			Name:      configuration.SeedName,
			Surname:   configuration.SeedSurname,
			Age:       0,
			Gender:    "",
			Education: "",
			Status:    0,
			RoleID:    adminRole.RoleId,
		}).Error
	}
	return nil
}
