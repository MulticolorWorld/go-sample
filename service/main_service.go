package service

type MainService struct {
}

func NewMainService() *MainService {
	return &MainService{}
}

func (s MainService) Hello() string {
	return "hello world"
}
