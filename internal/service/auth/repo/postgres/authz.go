package postgresql

func (s *storage) ConnString() string {
	return s.Config().ConnString()
}
