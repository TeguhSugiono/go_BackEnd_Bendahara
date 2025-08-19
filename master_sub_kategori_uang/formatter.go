package master_sub_kategori_uang

type ReturnData struct {
	Kd_kategori     int    `json:"kd_kategori"`
	Nm_kategori     string `json:"nm_kategori"`
	Kd_sub_kategori int    `json:"kd_sub_kategori"`
	Nm_sub_kategori string `json:"nm_sub_kategori"`
}

func ArrayShowData(table ListData) ReturnData {

	arraydata := ReturnData{}
	arraydata.Kd_kategori = table.Kd_kategori
	arraydata.Nm_kategori = table.Nm_kategori
	arraydata.Kd_sub_kategori = table.Kd_sub_kategori
	arraydata.Nm_sub_kategori = table.Nm_sub_kategori

	return arraydata
}

func FormatShowData(table []ListData) []ReturnData {
	arraydata := []ReturnData{}
	for _, resultdata := range table {
		arraytemporary := ArrayShowData(resultdata)
		arraydata = append(arraydata, arraytemporary)
	}

	return arraydata
}

type ListData struct {
	Kd_kategori     int
	Nm_kategori     string
	Kd_sub_kategori int
	Nm_sub_kategori string
}
