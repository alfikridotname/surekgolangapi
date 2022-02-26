package surat

type BuatSuratInput struct {
	KategoriPenerimaID int    `json:"kategori_penerima_id" binding:"required"`
	NaskahID           int    `json:"naskah_id" binding:"required"`
	SignerID           int    `json:"signer_id" binding:"required"`
	AtasNamaID         int    `json:"atas_nama_id" binding:"required"`
	PenerimaID         string `json:"penerima_id" binding:"required"`
	UnitKerjaTujuan    string `json:"unit_kerja_tujuan" binding:"required"`
	KeamananID         int    `json:"keamanan_id" binding:"required"`
	KecepatanID        int    `json:"kecepatan_id" binding:"required"`
	Tgl                string `json:"tgl" binding:"required"`
	Perihal            string `json:"perihal" binding:"required"`
	Tembusan           string `json:"tembusan"`
	Pemeriksa          string `json:"pemeriksa"`
	Isi                string `json:"isi" binding:"required"`
	CreatedBy          int    `json:"created_by"`
	UpdatedBy          int    `json:"updated_by"`
	CreatedNIP         string
	JabatanID          int
	UnitKerjaID        int
}
