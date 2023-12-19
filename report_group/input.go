package report_group

type ParamData struct {
	Tgl_bayar1    string `form:"tgl_bayar1" json:"tgl_bayar1"`
	Tgl_bayar2    string `form:"tgl_bayar2" json:"tgl_bayar2"`
	Kd_pembayaran int    `form:"kd_pembayaran" json:"kd_pembayaran"`
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
}

type DetailBayarOut struct {
	Tgl_bayar       string
	Jml_bayar       float64
	Tipe_pembayaran string
	Pos_uang_masuk  string
}

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
}

type NewLoop3a struct {
	Tgl_bayar       string
	Jml_bayar       float64
	Tipe_pembayaran string
	Pos_uang_masuk  string
	Kd_kategori     int
}
