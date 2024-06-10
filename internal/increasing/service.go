package increasing

type ContentCounterService struct{}

func NewContentCounterService() ContentCounterService {
	return ContentCounterService{}
}

func (s ContentCounterService) Increase(id string) error {
	return nil
}
