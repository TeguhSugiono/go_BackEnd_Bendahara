package master_conf_biaya_kategori

type ParamInputBiayaKategori struct {
	Kd_kategori int     `form:"kd_kategori" json:"kd_kategori" binding:"required,number"`
	Jml_biaya   float64 `form:"jml_biaya" json:"jml_biaya" binding:"required,number"`
}
