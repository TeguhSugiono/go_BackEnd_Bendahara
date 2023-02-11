package report_histori

import (
	"net/http"
	"rest_api_bendahara/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListNikNis(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var listdata []ListData
	sql := "  SELECT nis,nik,nm_siswa from tbl_siswa " +
		" where status_siswa NOT IN ('Tidak Aktif') and flag_siswa=0 "

	db.Raw(sql).Scan(&listdata)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", listdata)
	c.JSON(http.StatusOK, response)
}

func ReportHistori(c *gin.Context) {
	//db := c.MustGet("db").(*gorm.DB)

	var paramSearch ParamSearch
	if err := c.ShouldBindJSON(&paramSearch); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	SetArrayData := []DetailHistori{}
	arraydata := DetailHistori{}

	arraydata.Nis = paramSearch.Nis
	arraydata.Nik = paramSearch.Nik
	arraydata.Nm_siswa = paramSearch.Nm_siswa

	//arraydata.Detail = listDataPPDBDetail
	SetArrayData = append(SetArrayData, arraydata)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
	c.JSON(http.StatusOK, response)

}
