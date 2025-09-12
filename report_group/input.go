package report_group

type ParamData struct {
	Tgl_bayar1         string `form:"tgl_bayar1" json:"tgl_bayar1"`
	Tgl_bayar2         string `form:"tgl_bayar2" json:"tgl_bayar2"`
	Kd_group           string `form:"kd_group" json:"kd_group"`
	Kd_kategori        string `form:"kd_kategori" json:"kd_kategori"`
	Kd_post_uang_masuk string `form:"kd_post_uang_masuk" json:"kd_post_uang_masuk"`
}

type GroupUang struct {
	DataUang interface{}
}

type JenisGroupUang struct {
	JenisUang  string
	DetailData interface{}
}

type GroupUangMasuk struct {
	Kd_group int
	Nm_group string
	Kategori interface{}
}

type GroupUangMasukDetail struct {
	Kd_kategori int
	Nm_kategori string
	Total_bayar float64
	DetailBayar interface{}
}

type DetailBayar struct {
	Tgl_bayar       string
	Jml_bayar       float64
	Tipe_pembayaran string
	Data_ID         string
	Data_Name       string
	Keterangan      string
}

type DetailBayarIn struct {
	Nm_group        string
	Nm_kategori     string
	Jml_bayar       float64
	Tipe_pembayaran string
	Tgl_bayar       string
	Nik             string
	Nm_siswa        string
	Keterangan      string
}

type DetailBayarOut struct {
	Nm_group        string
	Nm_kategori     string
	Jml_bayar       float64
	Tipe_pembayaran string
	Pos_uang_masuk  string
	No_document     string
	Tgl_document    string
	Tgl_bayar       string
	Keterangan      string
}

// type DetailBayarOut struct {
// 	Tgl_bayar         string
// 	Jml_bayar         float64
// 	Tipe_pembayaran   string
// 	Pos_uang_masuk    string
// 	Data_no_document  string
// 	Data_tgl_document string
// 	Keterangan        string
// }

type NewLoop1 struct {
	Kd_kategori int
	Nm_kategori string
	Kd_group    int
}

type NewLoop2 struct {
	Total_bayar float64
	Kd_kategori int
}

type NewLoop3 struct {
	Tgl_bayar       string
	Jml_bayar       float64
	Tipe_pembayaran string
	Kd_kategori     int
	Data_ID         string
	Data_Name       string
	Keterangan      string
}

type NewLoop3a struct {
	Tgl_bayar         string
	Jml_bayar         float64
	Tipe_pembayaran   string
	Pos_uang_masuk    string
	Kd_kategori       int
	Data_no_document  string
	Data_tgl_document string
	Keterangan        string
}
