package pejabat

type MasterPejabat struct {
	JabatanID int    `json:"jabatan_id"`
	JabatanNM string `json:"jabatan_nm"`
	Nama      string `json:"nama"`
	EselonID  int    `json:"eselon_id"`
}
