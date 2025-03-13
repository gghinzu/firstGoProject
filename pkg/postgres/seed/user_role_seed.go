package seed

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func RoleSeed(db *gorm.DB) error {
	roles := []entity.UserRole{
		{RoleId: uuid.New().String(), Name: enum.User},
		{RoleId: uuid.New().String(), Name: enum.Admin},
	}

	for _, role := range roles {
		var count int64
		db.Model(&entity.UserRole{}).Where("name = ?", role.Name).Count(&count)
		if count == 0 {
			if err := db.Create(&role).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
