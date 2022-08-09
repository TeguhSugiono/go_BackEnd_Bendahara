package master_jenis_trans

type JenisTransInput struct {
	Proses_uang string `form:"proses_uang" binding:"required"`
	Flag_aktif  int    `form:"flag_aktif"`
	Created_on  string `form:"created_on"`
	Created_by  string `form:"created_by"`
	Edited_on   string `form:"edited_on"`
	Edited_by   string `form:"edited_by"`
}

type TableData struct {
	Total     int64 `json:"total"`
	Page      int   `json:"page"`
	Last_page int   `json:"last_page"`
}
