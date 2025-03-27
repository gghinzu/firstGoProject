package postgres

import (
	"errors"
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/dto"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetUserByID(id string) (*dto.UserResponseDTO, error) {
	var user entity.User
	err := r.db.Preload("Role").Where("id = ?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	userDTO := &dto.UserResponseDTO{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Surname:   user.Surname,
		Age:       user.Age,
		Gender:    user.Gender,
		Education: user.Education.String(),
		Status:    user.Status,
		Role:      string(user.Role.Name),
	}
	return userDTO, nil
}

func (r *UserRepository) GetUserByIDRaw(id string) (*entity.User, error) {
	var user entity.User
	err := r.db.Preload("Role").Where("id = ?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) DeleteUserByID(id string) error {
	err := r.db.Delete(&entity.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUserByID(id string, updatedUser *entity.User) error {
	updateData := map[string]interface{}{
		"email":                    updatedUser.Email,
		"password":                 updatedUser.Password,
		"name":                     updatedUser.Name,
		"surname":                  updatedUser.Surname,
		"age":                      updatedUser.Age,
		"gender":                   updatedUser.Gender,
		"education":                updatedUser.Education,
		"status":                   updatedUser.Status,
		"verification_code":        updatedUser.VerificationCode,
		"verification_code_expiry": updatedUser.VerificationCodeExpiry,
	}
	err := r.db.Model(&entity.User{}).Where("id = ?", id).Updates(updateData).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUserStatusByID(id string, userStatus enum.UserStatus) error {
	err := r.db.Model(&entity.User{}).Where("id = ?", id).Update("status", userStatus).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FilterUser(info dto.FilterDTO) (*[]dto.UserResponseDTO, error) {
	var users []entity.User
	query := r.paginate(int(info.Limit), int(info.Page)).Preload("Role").
		Select("id", "email", "name", "surname", "age", "gender", "education", "status", "role_id")
	if info.Name != nil {
		query = query.Where("name ILIKE ?", info.Name)
	}
	if info.Surname != nil {
		query = query.Where("surname ILIKE ?", info.Surname)
	}
	if info.Education != nil {
		query = query.Where("education ILIKE ?", info.Education)
	}
	if info.Gender != nil {
		query = query.Where("gender ILIKE ?", info.Gender)
	}
	if info.Age != nil {
		query = query.Where("age = ?", info.Age)
	}
	if info.Status != nil {
		query = query.Where("status = ?", info.Status)
	}
	if info.Order != nil {
		query = query.Order(*info.Order)
	}
	err := query.Find(&users).Error
	if err != nil {
		return nil, err
	}
	userDTOs := make([]dto.UserResponseDTO, len(users))
	for i, user := range users {
		userDTOs[i] = dto.UserResponseDTO{
			ID:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			Surname:   user.Surname,
			Age:       user.Age,
			Gender:    user.Gender,
			Education: user.Education.String(),
			Status:    user.Status,
			Role:      string(user.Role.Name),
		}
	}
	return &userDTOs, nil
}

func (r *UserRepository) Register(info *entity.User) error {
	err := r.db.Create(&info).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserRoleByRoleName(roleName string) (*entity.UserRole, error) {
	var userRole *entity.UserRole
	err := r.db.Model(&entity.UserRole{}).Where("name = ?", roleName).First(&userRole).Error
	if err != nil {
		return nil, err
	}
	return userRole, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user *entity.User
	err := r.db.Preload("Role").Where("lower(email) = lower(?)", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetRoleByUserInfo(user *entity.User) (enum.UserRole, error) {
	var userRole entity.UserRole
	err := r.db.Model(&entity.UserRole{}).Where("role_id = ?", user.RoleID).First(&userRole).Error
	if err != nil {
		return "", err
	}
	return userRole.Name, nil
}

func (r *UserRepository) paginate(limit int, offset int) *gorm.DB {
	return r.db.Limit(limit).Offset((offset - 1) * limit)
}
