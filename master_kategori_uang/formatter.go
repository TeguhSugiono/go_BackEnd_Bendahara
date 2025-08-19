package master_kategori_uang

type ReturnData struct {
	Kd_jenis    int    `json:"kd_jenis"`
	Proses_uang string `json:"proses_uang"`
	Kd_kategori int    `json:"kd_kategori"`
	Nm_kategori string `json:"nm_kategori"`
	Nm_detail   string `json:"nm_detail"`
}

func FormatJenisTranss(table ListData) ReturnData {
	arraydata := ReturnData{}
	arraydata.Kd_kategori = table.Kd_kategori
	arraydata.Nm_kategori = table.Nm_kategori
	arraydata.Nm_detail = table.Nm_detail
	arraydata.Proses_uang = table.Proses_uang
	arraydata.Kd_jenis = table.Kd_jenis
	return arraydata
}

func FormatShowData(table []ListData) []ReturnData {
	arraydata := []ReturnData{}
	for _, resultdata := range table {
		arraytemporary := FormatJenisTranss(resultdata)
		arraydata = append(arraydata, arraytemporary)
	}

	return arraydata
}

type ListData struct {
	Kd_kategori int
	Nm_kategori string
	Nm_detail   string
	Proses_uang string
	Kd_jenis    int
}
