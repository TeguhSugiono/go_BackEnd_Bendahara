package master_kategori_uang

type ReturnData struct {
	Kd_kategori int    `json:"kd_kategori"`
	Nm_kategori string `json:"nm_kategori"`
}

func FormatJenisTranss(table Tbl_kategori_uangs) ReturnData {

	arraydata := ReturnData{}
	arraydata.Kd_kategori = table.Kd_kategori
	arraydata.Nm_kategori = table.Nm_kategori

	return arraydata
}

func FormatShowData(table []Tbl_kategori_uangs) []ReturnData {
	arraydata := []ReturnData{}
	for _, resultdata := range table {
		arraytemporary := FormatJenisTranss(resultdata)
		arraydata = append(arraydata, arraytemporary)
	}

	return arraydata
}
