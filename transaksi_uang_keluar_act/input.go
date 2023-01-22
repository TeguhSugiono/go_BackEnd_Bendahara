package transaksi_uang_keluar_act

type ParamChangeKategori struct {
	Kd_group string `form:"kd_group" json:"kd_group" binding:"required"`
}

type ParamInputTransaksi struct {
	Kd_group        int     `form:"kd_group" json:"kd_group" binding:"required"`
	Kd_kategori     int     `form:"kd_kategori" json:"kd_kategori" binding:"required"`
	Kd_trans_keluar int     `form:"Kd_trans_keluar" json:"Kd_trans_keluar"`
	No_document     string  `form:"no_document" json:"no_document" binding:"required"`
	Tgl_document    string  `form:"tgl_document" json:"tgl_document" binding:"required"`
	Total_biaya     float64 `form:"total_biaya" json:"total_biaya" binding:"number"`
	Total_bayar     float64 `form:"total_bayar" json:"total_bayar"`
	Sisa_biaya      float64 `form:"sisa_biaya" json:"sisa_biaya"`
	Keterangan      string  `form:"keterangan" json:"keterangan"`
	Flag_aktif      int     `form:"flag_aktif" json:"flag_aktif"`
	Created_on      string  `form:"created_on" json:"created_on"`
	Created_by      string  `form:"created_by" json:"created_by"`
	Edited_on       string  `form:"edited_on" json:"edited_on"`
	Edited_by       string  `form:"edited_by" json:"edited_by"`
}

type GetBiayaAndSisa struct {
	Kd_trans_keluar int
	Kd_group        int
	Nm_group        string
	Kd_kategori     int
	Nm_kategori     string
	No_document     string
	Tgl_document    string
	Total_biaya     float64
	Total_bayar     float64
	Sisa_biaya      float64
	Keterangan      string
	Detail          interface{}
}

type GetDataUmSiswa struct {
	Kd_trans_keluar_detail int
	Seqno                  int
	Jml_bayar              float64
	Keterangan             string
}

type ParamInputTransaksiEdit struct {
	Kd_group     int     `form:"kd_group" json:"kd_group" binding:"required"`
	Kd_kategori  int     `form:"kd_kategori" json:"kd_kategori" binding:"required"`
	No_document  string  `form:"no_document" json:"no_document" binding:"required"`
	Tgl_document string  `form:"tgl_document" json:"tgl_document" binding:"required"`
	Total_biaya  float64 `form:"total_biaya" json:"total_biaya" binding:"required,number"`
	Total_bayar  float64 `form:"total_bayar" json:"total_bayar"`
	Sisa_biaya   float64 `form:"sisa_biaya" json:"sisa_biaya"`
	Keterangan   string  `form:"keterangan" json:"keterangan"`
	Flag_aktif   int     `form:"flag_aktif" json:"flag_aktif"`
	Created_on   string  `form:"created_on" json:"created_on"`
	Created_by   string  `form:"created_by" json:"created_by"`
	Edited_on    string  `form:"edited_on" json:"edited_on"`
	Edited_by    string  `form:"edited_by" json:"edited_by"`
}

type ParamEditUmSiswaDetail struct {
	Jml_bayar  float64 `form:"jml_bayar" json:"jml_bayar" binding:"number"`
	Keterangan string  `form:"keterangan" json:"keterangan" binding:"required"`
	Edited_on  string  `form:"edited_on" json:"edited_on"`
	Edited_by  string  `form:"edited_by" json:"edited_by"`
}

type ParamAddDetail struct {
	Kd_trans_keluar int `form:"kd_trans_keluar" json:"kd_trans_keluar"`
}

type ParamChangeSiswa struct {
	Tgl_document1 string `form:"tgl_document1" json:"tgl_document1" binding:"required"`
	Tgl_document2 string `form:"tgl_document2" json:"tgl_document2" binding:"required"`
	No_document   string `form:"no_document" json:"no_document"`
}

type ParamDeleteUmLainDetail struct {
	Kd_trans_keluar        string `form:"kd_trans_keluar" json:"kd_trans_keluar" binding:"required"`
	Kd_trans_keluar_detail string `form:"kd_trans_keluar_detail" json:"kd_trans_keluar_detail" binding:"required"`
}
