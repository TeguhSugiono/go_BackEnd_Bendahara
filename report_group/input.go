package report_group

type ParamData struct {
	Tgl_bayar1 string `form:"tgl_bayar1" json:"tgl_bayar1"`
	Tgl_bayar2 string `form:"tgl_bayar2" json:"tgl_bayar2"`
}

type GroupUang struct {
	DataUang interface{}
}

type JenisGroupUang struct {
	JenisUang  string
	DetailData interface{}
}

type GroupUangMasuk struct {
	Nm_group string
	Kategori interface{}
}

type GroupUangMasukDetail struct {
	Nm_kategori string
	Total_bayar float64
}
