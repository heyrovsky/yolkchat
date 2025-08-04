package users

type Service struct {
	repo Interface
}

func NewService(r Interface) Interface {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateUserProfile(username string) error {
	panic(0)
}

func (s *Service) DeleteUserProfile(username string) error {
	panic(0)
}
