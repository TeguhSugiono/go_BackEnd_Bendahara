package master_sumber_dana

type ParamData struct {
	Kd_group string `form:"kd_group" json:"kd_group"`
}

type DanaMasuk struct {
	TotalBayar float64
	Kd_group   int
	Nm_group   string
}
type DanaKeluar struct {
	TotalBayar float64
	Kd_group   int
	Nm_group   string
}
