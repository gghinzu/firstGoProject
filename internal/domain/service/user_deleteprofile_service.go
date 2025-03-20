package service

func (s *UserService) DeleteProfile(id string) error {
	err := s.repo.DeleteUserByID(id)

	if err != nil {
		return err
	}

	return nil
}
