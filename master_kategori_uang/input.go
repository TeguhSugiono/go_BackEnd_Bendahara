package master_kategori_uang

type KategoriUangInput struct {
	Kd_group    int    `form:"kd_group" json:"kd_group" binding:"required,number"`
	Nm_kategori string `form:"nm_kategori" json:"nm_kategori" binding:"required"`
	Nm_detail   string `form:"nm_detail" json:"nm_detail" `
	Flag_aktif  int    `form:"flag_aktif" json:"flag_aktif"`
	Created_on  string `form:"created_on" json:"created_on"`
	Created_by  string `form:"created_by" json:"created_by"`
	Edited_on   string `form:"edited_on" json:"edited_on"`
	Edited_by   string `form:"edited_by" json:"edited_by"`
}

type DataKdGroup struct {
	Nm_header string `form:"nm_header" json:"nm_header"`
}
