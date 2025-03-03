package service

func (s *UserService) DeleteUserByID(id string) error {
	user, err := s.GetUserByID(id)
	if err != nil {
		return err
	} else {
		return s.repo.DeleteUserByID(user.ID)
	}
}
