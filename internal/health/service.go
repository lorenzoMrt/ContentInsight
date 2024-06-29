package health

type Service interface {
	Health() string
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Health() string {
	return "healthy"
}
