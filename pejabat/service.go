package pejabat

type Service interface {
	GetByUnitKerjaID(unitKerjaID int, kategori string) ([]MasterPejabat, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetByUnitKerjaID(unitKerjaID int, kategori string) ([]MasterPejabat, error) {
	masterPejabat, err := s.repository.FindByUnitKerjaID(unitKerjaID, kategori)
	if err != nil {
		return masterPejabat, err
	}

	return masterPejabat, nil
}
