package master_conf_spp_ppdb

func FormatShowData(table []ListData) []ReturnData {
	arraydata := []ReturnData{}
	for _, resultdata := range table {
		arraytemporary := FormatTampungData(resultdata)
		arraydata = append(arraydata, arraytemporary)
	}

	return arraydata
}

type ReturnData struct {
	Id_link     int    `json:"id_link"`
	Link_name   string `json:"link_name"`
	Kd_group    int    `json:"kd_group"`
	Nm_group    string `json:"nm_group"`
	Kd_kategori int    `json:"kd_kategori"`
	Nm_kategori string `json:"nm_kategori"`
}

type ListData struct {
	Id_link     int
	Link_name   string
	Kd_group    int
	Nm_group    string
	Kd_kategori int
	Nm_kategori string
}

func FormatTampungData(table ListData) ReturnData {
	arraydata := ReturnData{}
	arraydata.Id_link = table.Id_link
	arraydata.Link_name = table.Link_name
	arraydata.Kd_group = table.Kd_group
	arraydata.Nm_group = table.Nm_group
	arraydata.Kd_kategori = table.Kd_kategori
	arraydata.Nm_kategori = table.Nm_kategori
	return arraydata
}
