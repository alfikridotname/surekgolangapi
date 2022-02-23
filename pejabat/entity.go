package pejabat

type MasterPejabat struct {
	JabatanID int    `json:"jabatan_id"`
	JabatanNM string `json:"jabatan_nm"`
	Nama      string `json:"nama"`
	Kategori  string `json:"kategori"`
	Nip       string `json:"nip"`
	EselonID  int    `json:"eselon_id"`
}
