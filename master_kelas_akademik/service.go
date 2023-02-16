package master_kelas_akademik

import (
	"net/http"
	"rest_api_bendahara/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListKelasAkademik(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var master []ListData
	//db.Find(&master)
	sql := " SELECT  nm_kelas as 'id_kelas', " +
		" nm_kelas as 'nm_kelas' FROM tbl_kelas " +
		" where flag_kelas = 0  " +
		" ORDER BY id_kelas "

	db.Raw(sql).Scan(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", FormatShowData(master))
	c.JSON(http.StatusOK, response)
}
