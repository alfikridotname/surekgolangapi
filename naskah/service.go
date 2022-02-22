package naskah

type Service interface {
	GetAll() ([]MasterNaskah, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]MasterNaskah, error) {
	masterNaskah, err := s.repository.FindAll()
	if err != nil {
		return masterNaskah, err
	}

	return masterNaskah, nil
}
