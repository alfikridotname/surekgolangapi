package unitkerja

type Service interface {
	GetAll(unitKerjaID int) ([]MasterUnitKerja, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll(unitKerjaID int) ([]MasterUnitKerja, error) {
	masterUnitKerja, err := s.repository.FindAll(unitKerjaID)
	if err != nil {
		return masterUnitKerja, err
	}

	return masterUnitKerja, nil
}
