package seed

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	var count int64
	if db.Table("users").Count(&count); count == 0 {
		var adminRole *entity.UserRole
		err := db.Table("user_roles").Where("name = ?", enum.Admin).First(&adminRole).Error
		if err != nil {
			return err
		}

		return db.Create(&entity.User{
			ID:        uuid.New(),
			Email:     "admin@example.com",
			Password:  "123456",
			Name:      "seed",
			Surname:   "seed",
			Age:       0,
			Gender:    "",
			Education: "",
			Status:    0,
			RoleID:    adminRole.RoleId,
		}).Error
	}
	return nil
}
