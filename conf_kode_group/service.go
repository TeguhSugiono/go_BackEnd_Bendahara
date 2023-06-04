package conf_kode_group

import (
	"net/http"
	"rest_api_bendahara/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListGroup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramIdJenis ParamIdJenis
	if err := c.ShouldBindJSON(&paramIdJenis); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var master []ReturnData
	sql := "  SELECT a.kd_group,a.nm_group FROM tbl_group_kategoris as a " +
		" inner join tbl_jenis_trans as b on a.kd_jenis=b.kd_jenis  " +
		" where a.flag_aktif=0 and b.flag_aktif=0  " +
		" and b.kd_jenis in(" + paramIdJenis.Kd_jenis + ") "

	db.Raw(sql).Scan(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", master)
	c.JSON(http.StatusOK, response)
}

func ListKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramIdGroup ParamIdGroup
	if err := c.ShouldBindJSON(&paramIdGroup); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var master []ReturnDataKategori
	sql := " SELECT a.*,b.nm_group FROM tbl_kategori_uangs as a  " +
		" INNER JOIN tbl_group_kategoris b on a.kd_group=b.kd_group " +
		" where a.flag_aktif=0 and b.flag_aktif=0 " +
		" and b.kd_group in (" + paramIdGroup.Kd_group + ") "

	db.Raw(sql).Scan(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", master)
	c.JSON(http.StatusOK, response)
}
