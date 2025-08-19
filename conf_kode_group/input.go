package conf_kode_group

type ParamIdJenis struct {
	Kd_jenis string `form:"kd_jenis" json:"kd_jenis" binding:"required"`
}

type ReturnData struct {
	Kd_group int
	Nm_group string
}

type ParamIdGroup struct {
	Kd_group string `form:"kd_group" json:"kd_group" binding:"required"`
}

type ReturnDataKategori struct {
	Kd_kategori int
	Nm_kategori string
}
