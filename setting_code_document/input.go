package setting_code_document

type ParamCreateCode struct {
	Kode string `form:"kode" json:"kode" binding:"required"`
}

type TahunBulan struct {
	Tahun string
	Bulan string
}
