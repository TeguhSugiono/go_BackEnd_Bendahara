package master_conf_spp_ppdb

type ParamInput struct {
	Kd_group    int `form:"kd_group" binding:"number,required"`
	Kd_kategori int `form:"kd_group" binding:"number,required"`
}
