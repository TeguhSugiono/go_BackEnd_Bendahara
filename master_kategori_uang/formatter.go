package master_kategori_uang

type ReturnData struct {
	Kd_group    int    `json:"kd_group"`
	Nm_group    string `json:"nm_group"`
	Kd_kategori int    `json:"kd_kategori"`
	Nm_kategori string `json:"nm_kategori"`
	Nm_detail   string `json:"nm_detail"`
	Proses_uang   string `json:"proses_uang"`
}

func FormatJenisTranss(table ListData) ReturnData {

	arraydata := ReturnData{}
	arraydata.Kd_group = table.Kd_group
	arraydata.Nm_group = table.Nm_group
	arraydata.Kd_kategori = table.Kd_kategori
	arraydata.Nm_kategori = table.Nm_kategori
	arraydata.Nm_detail = table.Nm_detail
	arraydata.Proses_uang = table.Proses_uang
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
	Kd_group    int
	Nm_group    string
	Kd_kategori int
	Nm_kategori string
	Nm_detail   string
	Proses_uang string
}
