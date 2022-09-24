package master_group_kategori

type GroupKategoriInput struct {
	Kd_jenis   int    `json:"kd_jenis" form:"kd_jenis" binding:"required,number"`
	Nm_group   string `json:"nm_group" form:"nm_group" binding:"required"`
	Nm_header  string `json:"nm_header" form:"nm_header"`
	Flag_aktif int    `json:"flag_aktif" form:"flag_aktif"`
	Created_on string `json:"created_on" form:"created_on"`
	Created_by string `json:"created_by" form:"created_by"`
	Edited_on  string `json:"edited_on" form:"edited_on"`
	Edited_by  string `json:"edited_by" form:"edited_by"`
}

type TableData struct {
	Total     int64 `json:"total"`
	Page      int   `json:"page"`
	Last_page int   `json:"last_page"`
}
