package report_uangkeluar

type ParamReportPRA struct {
	Tgl_document1 string `form:"tgl_document1" json:"tgl_document1" binding:"required"`
	Tgl_document2 string `form:"tgl_document2" json:"tgl_document2" binding:"required"`
	No_document   string `form:"no_document" json:"no_document"`
}

type GetDataHeaderPRA struct {
	Nm_group     string
	Nm_kategori  string
	No_document  string
	Tgl_document string
	Total_biaya  float64
	Total_bayar  float64
	Sisa_biaya   float64
	Keterangan   string
	Detail       interface{}
}

type GetDataDetailPRA struct {
	Jml_bayar         float64
	Keterangan_detail string
}

type ParamReportPRAACT struct {
	Tgl_document1 string `form:"tgl_document1" json:"tgl_document1" binding:"required"`
	Tgl_document2 string `form:"tgl_document2" json:"tgl_document2" binding:"required"`
	No_document   string `form:"no_document" json:"no_document"`
	Tgl_bayar1    string `form:"tgl_bayar1" json:"tgl_bayar1"`
	Tgl_bayar2    string `form:"tgl_bayar2" json:"tgl_bayar2"`
	Kd_pembayaran string `form:"kd_pembayaran" json:"kd_pembayaran"`
}

type GetDataHeaderPRAACT struct {
	Nm_group     string
	Nm_kategori  string
	No_document  string
	Tgl_document string
	Total_biaya  float64
	Total_bayar  float64
	Sisa_biaya   float64
	Keterangan   string
	Detail       interface{}
}

type GetDataDetailPRAACT struct {
	Pos_uang_masuk    string
	Tgl_bayar         string
	Jml_bayar         float64
	Keterangan_detail string
	Tipe_pembayaran   string
}

type ParamReportACT struct {
	Tgl_document1 string `form:"tgl_document1" json:"tgl_document1" binding:"required"`
	Tgl_document2 string `form:"tgl_document2" json:"tgl_document2" binding:"required"`
	No_document   string `form:"no_document" json:"no_document"`
	Tgl_bayar1    string `form:"tgl_bayar1" json:"tgl_bayar1"`
	Tgl_bayar2    string `form:"tgl_bayar2" json:"tgl_bayar2"`
	Kd_pembayaran string `form:"kd_pembayaran" json:"kd_pembayaran"`
}

type GetDataHeaderACT struct {
	Nm_group     string
	Nm_kategori  string
	No_document  string
	Tgl_document string
	Total_biaya  float64
	Total_bayar  float64
	Sisa_biaya   float64
	Keterangan   string
	Detail       interface{}
}

type GetDataDetailACT struct {
	Pos_uang_masuk    string
	Tgl_bayar         string
	Jml_bayar         float64
	Keterangan_detail string
	Tipe_pembayaran   string
}
