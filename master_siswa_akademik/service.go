package master_siswa_akademik

import (
	"net/http"
	"rest_api_bendahara/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListSiswaAkademik(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var master []ListData
	//db.Find(&master)
	sql := " SELECT nis,nm_siswa FROM tbl_siswa " +
		" WHERE flag_siswa = 0 AND status_siswa NOT IN ('Tidak Aktif') order by nm_siswa "

	db.Raw(sql).Scan(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", FormatShowData(master))
	c.JSON(http.StatusOK, response)
}
