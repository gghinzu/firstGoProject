package service

import (
	"firstGoProject/internal/domain/entity"
)

// CreateUser gets its own DTO, converts it to entity and sends the data to the repository
// uses the instance of UserService (it connects the service with the repo)
func (s *UserService) CreateUser(newUser *entity.CreateUserDTO) error {
	converted, err := CreateConvertToUser(newUser)
	if err != nil {
		return err
	}
	return s.repo.CreateUser(converted)
}

func CreateConvertToUser(dto *entity.CreateUserDTO) (*entity.User, error) {
	user := &entity.User{
		Name:      dto.Name,
		Surname:   dto.Surname,
		Age:       dto.Age,
		Gender:    dto.Gender,
		Education: dto.Education,
	}
	return user, nil
}
