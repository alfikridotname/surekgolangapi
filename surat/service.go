package surat

type Service interface {
	CreateSurat(input BuatSuratInput) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateSurat(input BuatSuratInput) (bool, error) {
	masterSurat := MasterSurat{}
	masterSurat.KategoriPenerimaID = input.KategoriPenerimaID
	masterSurat.NaskahID = input.NaskahID
	masterSurat.SignerID = input.SignerID
	masterSurat.AtasNamaID = input.AtasNamaID
	masterSurat.KeamananID = input.KeamananID
	masterSurat.KecepatanID = input.KecepatanID
	masterSurat.Tgl = input.Tgl
	masterSurat.Perihal = input.Perihal
	masterSurat.Isi = input.Isi
	masterSurat.CreatedBy = input.CreatedBy
	masterSurat.UpdatedBy = input.UpdatedBy

	_, err := s.repository.Save(masterSurat, input.Tembusan, input.PenerimaID, input.Pemeriksa, input.UnitKerjaID)
	if err != nil {
		return false, err
	}

	return true, nil
}
