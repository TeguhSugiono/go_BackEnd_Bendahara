package master_group_kategori

type GroupKategoriFormatter struct {
	Kd_jenis  int    `json:"kd_jenis"`
	Kd_group  int    `json:"kd_group"`
	Nm_group  string `json:"nm_group"`
	Nm_header string `json:"nm_header"`
	Nm_detail string `json:"nm_detail"`
}

func FormatGroupKategoris(table Tbl_group_kategoris) GroupKategoriFormatter {

	arraydata := GroupKategoriFormatter{}
	arraydata.Kd_jenis = table.Kd_jenis
	arraydata.Kd_group = table.Kd_group
	arraydata.Nm_group = table.Nm_group
	arraydata.Nm_header = table.Nm_header
	arraydata.Nm_detail = table.Nm_detail

	return arraydata
}

func FormatGroupKategori(table []Tbl_group_kategoris) []GroupKategoriFormatter {
	arraydata := []GroupKategoriFormatter{}
	for _, resultdata := range table {
		arraytemporary := FormatGroupKategoris(resultdata)
		arraydata = append(arraydata, arraytemporary)
	}

	return arraydata
}
