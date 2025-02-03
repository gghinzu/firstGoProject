package service

import (
	"firstGoProject/internal/domain/entity"
)

func (s *UserService) InsertNewUser(newUser *entity.UserDTO) error {
	converted, err := ConvertToUser(newUser)
	if err != nil {
		return err
	}
	return s.repo.InsertNewUser(converted)
}

func ConvertToUser(dto *entity.UserDTO) (*entity.User, error) {
	user := &entity.User{
		ID:        dto.ID,
		Name:      dto.Name,
		Surname:   dto.Surname,
		Age:       dto.Age,
		Gender:    dto.Gender,
		Education: dto.Education,
	}
	return user, nil
}
