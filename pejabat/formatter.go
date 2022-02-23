package pejabat

type PejabatFormatter struct {
	JabatanID int    `json:"jabatan_id"`
	JabatanNM string `json:"jabatan_nm"`
	Nama      string `json:"nama"`
	Kategori  string `json:"kategori"`
	Nip       string `json:"nip"`
}

func FormatPejabat(pejabat MasterPejabat) PejabatFormatter {
	return PejabatFormatter{
		JabatanID: pejabat.JabatanID,
		JabatanNM: pejabat.JabatanNM,
		Nama:      pejabat.Nama,
		Kategori:  pejabat.Kategori,
		Nip:       pejabat.Nip,
	}
}

func FormatMultiplePejabat(masterPejabat []MasterPejabat) []PejabatFormatter {
	var pejabatFormatter []PejabatFormatter
	for _, pejabat := range masterPejabat {
		pejabatFormatter = append(pejabatFormatter, FormatPejabat(pejabat))
	}
	return pejabatFormatter
}
