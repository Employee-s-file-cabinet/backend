package auth

func (s *service) DataSourceName() string {
	return s.authRepository.ConnString()
}
