package master_sub_kategori_uang

type SKategoriUangInput struct {
	Kd_kategori     int    `form:"kd_kategori" json:"kd_kategori" binding:"required,number"`
	Nm_sub_kategori string `form:"nm_sub_kategori" json:"nm_sub_kategori" binding:"required"`
	Flag_aktif      int    `form:"flag_aktif" json:"flag_aktif"`
	Created_on      string `form:"created_on" json:"created_on"`
	Created_by      string `form:"created_by" json:"created_by"`
	Edited_on       string `form:"edited_on" json:"edited_on"`
	Edited_by       string `form:"edited_by" json:"edited_by"`
}

type DataSKategoriUang struct {
	Nm_header string `form:"nm_header" json:"nm_header"`
}
