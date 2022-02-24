package notifikasi

type Service interface {
	CreateNotifikasi(input NotifikasiInput, createdBY int) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateNotifikasi(input NotifikasiInput, createdBy int) (bool, error) {
	notifikasi := MasterNotifikasi{}
	notifikasi.Data = input.Data
	notifikasi.UserTujuanID = input.UserTujuanID
	notifikasi.CreatedBY = createdBy

	_, err := s.repository.Save(&notifikasi)
	if err != nil {
		return false, err
	}

	return true, nil
}
