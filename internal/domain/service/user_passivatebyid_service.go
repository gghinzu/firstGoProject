package service

import (
	"errors"
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
)

func (s *UserService) PassivateUserByID(id int, updatedUser *entity.ActivatePassivateUserDTO) error {
	converted, err := ActivatePassivateConvertToUser(updatedUser)
	if err != nil {
		return err
	}
	var user, errID = s.GetUserByID(id)
	if errID != nil {
		return err
	}
	if user.Status != enum.Passive {
		return s.repo.PassivateUserByID(user.ID, converted)
	}
	return errors.New("user is already passive")
}
