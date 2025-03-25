package seed

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/helper"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func AdminSeed(db *gorm.DB) error {
	var count int64

	if db.Model(&entity.User{}).Count(&count); count == 0 {
		var adminRole *entity.UserRole

		err := db.Model(&entity.UserRole{}).Where("name = ?", enum.Admin).First(&adminRole).Error
		if err != nil {
			return err
		}

		hashedPassword, err := helper.EncryptPassword("123456789")
		if err != nil {
			return err
		}

		age := time.Now().Year() - 2002

		return db.Create(&entity.User{
			ID:        uuid.New().String(),
			Email:     "admin@example.com",
			Password:  string(hashedPassword),
			Name:      "seed",
			Surname:   "seed",
			Age:       age,
			Gender:    enum.NotSpecified,
			Education: enum.None,
			Status:    1,
			RoleID:    adminRole.RoleId,
		}).Error
	}

	return nil
}
