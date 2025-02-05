package service

// DeleteUserByID gets all users using an instance of UserService
// (implementation of the interface UserServicePort)
func (s *UserService) DeleteUserByID(id int) error {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteUserByID(user.ID)
}
