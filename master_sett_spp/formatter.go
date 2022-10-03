package master_sett_spp

// func FormatShowData(table []ListData) []ReturnData {
// 	arraydata := []ReturnData{}
// 	for _, resultdata := range table {
// 		arraytemporary := FormatTampungData(resultdata)
// 		arraydata = append(arraydata, arraytemporary)
// 	}

// 	return arraydata
// }

// type ReturnData struct {
// 	Kd_periode_spp int     `json:"kd_periode_spp"`
// 	Kd_sett_spp    int     `json:"kd_sett_spp"`
// 	Nm_kelas       string  `json:"nm_kelas"`
// 	Biaya_spp      float64 `json:"biaya_spp"`
// 	Keterangan     string  `json:"keterangan"`
// 	Kd_bulan       string  `json:"kd_bulan"`
// 	Tahun          int     `json:"tahun"`
// 	Tahun_akademik string  `json:"tahun_akademik"`
// }

// type ListData struct {
// 	Kd_periode_spp int
// 	Kd_sett_spp    int
// 	Nm_kelas       string
// 	Biaya_spp      float64
// 	Keterangan     string
// 	Kd_bulan       string
// 	Tahun          int
// 	Tahun_akademik string
// }

// func FormatTampungData(table ListData) ReturnData {
// 	arraydata := ReturnData{}
// 	arraydata.Kd_periode_spp = table.Kd_periode_spp
// 	arraydata.Kd_sett_spp = table.Kd_sett_spp
// 	arraydata.Nm_kelas = table.Nm_kelas
// 	arraydata.Biaya_spp = table.Biaya_spp
// 	arraydata.Keterangan = table.Keterangan
// 	arraydata.Kd_bulan = table.Kd_bulan
// 	arraydata.Tahun = table.Tahun
// 	arraydata.Tahun_akademik = table.Tahun_akademik
// 	return arraydata
// }
