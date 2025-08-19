package master_group_kategori

type GroupKategoriFormatter struct {
	Kd_jenis    int    `json:"kd_jenis"`
	Proses_uang string `json:"proses_uang"`
	Kd_group    int    `json:"kd_group"`
	Nm_group    string `json:"nm_group"`
	Nm_header   string `json:"nm_header"`
}

func FormatGroupKategoris(table ListData) GroupKategoriFormatter {

	arraydata := GroupKategoriFormatter{}
	arraydata.Kd_jenis = table.Kd_jenis
	arraydata.Proses_uang = table.Proses_uang
	arraydata.Kd_group = table.Kd_group
	arraydata.Nm_group = table.Nm_group
	arraydata.Nm_header = table.Nm_header

	return arraydata
}

func FormatGroupKategori(table []ListData) []GroupKategoriFormatter {
	arraydata := []GroupKategoriFormatter{}
	for _, resultdata := range table {
		arraytemporary := FormatGroupKategoris(resultdata)
		arraydata = append(arraydata, arraytemporary)
	}

	return arraydata
}

type ListData struct {
	Kd_jenis    int
	Proses_uang string
	Kd_group    int
	Nm_group    string
	Nm_header   string
}
