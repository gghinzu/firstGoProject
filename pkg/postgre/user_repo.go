package postgre

import (
	"errors"
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRepository is used since Go does not have explicit 'implements' keyword, an empty struct is used to implicitly implement interfaces
type UserRepository struct {
	db *gorm.DB
}

type paginate struct {
	limit int
	page  int
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

// GetAllUsers for displaying all the users
func (r *UserRepository) GetAllUsers() (*[]entity.User, error) {
	var userList *[]entity.User
	err := r.db.Preload("Role").Order("name").Find(&userList).Error
	if err != nil {
		return nil, err
	}
	return userList, nil
}

// GetUserByID to get a specific user's details
func (r *UserRepository) GetUserByID(id uuid.UUID) (*entity.User, error) {
	var user *entity.User
	// record count control should be done
	err := r.db.Where("id = ?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return user, nil
}

// DeleteUserByID to delete a specific user by the given id
func (r *UserRepository) DeleteUserByID(id uuid.UUID) error {
	err := r.db.Delete(&entity.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateUserByID to update a specific user's info by the given id
func (r *UserRepository) UpdateUserByID(id uuid.UUID, updatedUser *entity.User) error {
	err := r.db.Where("id = ?", id).Save(&updatedUser).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateUser to create a new user
func (r *UserRepository) CreateUser(newUser *entity.User) error {
	err := r.db.Create(&newUser).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateUserStatusByID updates users' statuses
func (r *UserRepository) UpdateUserStatusByID(id uuid.UUID, userStatus enum.UserStatus) error {
	err := r.db.Model(&entity.User{}).Where("id = ?", id).Update("status", userStatus).Error
	if err != nil {
		return err
	}
	return nil
}

// SearchUser searches for a specific value in users' info and lists users
func (r *UserRepository) SearchUser(info entity.SearchUserDTO) (*[]entity.User, error) {
	var users *[]entity.User
	query := r.db

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
	if info.Limit != nil && info.Page != nil {
		query = query.Scopes(newPaginate(*info.Limit, *info.Page).paginatedResult).Find(&users)
	}

	err := query.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func newPaginate(limit int, page int) *paginate {
	return &paginate{limit: limit, page: page}
}

func (p *paginate) paginatedResult(db *gorm.DB) *gorm.DB {
	offset := (p.page - 1) * p.limit
	return db.Offset(offset).Limit(p.limit)
}
