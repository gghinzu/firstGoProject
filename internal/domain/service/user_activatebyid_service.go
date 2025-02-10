package service

import (
	"errors"
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
)

func (s *UserService) ActivateUserByID(id int, updatedUser *entity.ActivatePassivateUserDTO) error {
	converted, err := ActivatePassivateConvertToUser(updatedUser)
	if err != nil {
		return err
	}
	var user, errID = s.GetUserByID(id)
	if errID != nil {
		return err
	}
	if user.Status != enum.Active {
		return s.repo.ActivateUserByID(user.ID, converted)
	}
	return errors.New("user is already active")
}

func ActivatePassivateConvertToUser(dto *entity.ActivatePassivateUserDTO) (*entity.User, error) {
	user := &entity.User{
		Status: dto.Status,
	}
	return user, nil
}
