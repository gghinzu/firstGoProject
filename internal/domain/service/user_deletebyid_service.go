package service

// DeleteUserByID gets an id and deletes the user with this id, from database
// uses the instance of UserService (it connects the service with the repo)
func (s *UserService) DeleteUserByID(id int) error {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteUserByID(user.ID)
}
