package setting_code_document

import (
	"net/http"
	"rest_api_bendahara/helper"
	"rest_api_bendahara/table_data"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCodeDocument(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramCreateCode ParamCreateCode
	if err := c.ShouldBindJSON(&paramCreateCode); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var tahunBulan TahunBulan
	sql := " SELECT DATE_FORMAT(CURRENT_DATE,'%Y') as 'tahun',DATE_FORMAT(CURRENT_DATE,'%m') as 'bulan' "
	db.Raw(sql).Scan(&tahunBulan)

	var tbl_runcode_documents table_data.Tbl_runcode_documents
	checkUser := db.Select("*").Where("tahun = ? and bulan=?", tahunBulan.Tahun, tahunBulan.Bulan).Find(&tbl_runcode_documents)
	if checkUser.RowsAffected > 0 {

		var nomor int
		db.Raw("SELECT nomor FROM tbl_runcode_documents where kode=? and tahun=? and bulan=?", paramCreateCode.Kode, tahunBulan.Tahun, tahunBulan.Bulan).Scan(&nomor)
		nomor = nomor + 1

		CodeGenerate := tahunBulan.Tahun + "" + tahunBulan.Bulan + "-" + paramCreateCode.Kode + "" + helper.StrPad(strconv.Itoa(nomor), 4, "0", "LEFT")

		err := db.Raw("update tbl_runcode_documents set nomor=? where kode=? "+
			" and tahun=? and bulan=? ", nomor, paramCreateCode.Kode, tahunBulan.Tahun, tahunBulan.Bulan).Scan(&tbl_runcode_documents).Error
		if err != nil {
			response := helper.APIResponse("Genearate Kode Document Gagal ...", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.APIResponse("Genearate Kode Document Berhasil ...", http.StatusOK, "success", CodeGenerate)
		c.JSON(http.StatusOK, response)
		return
	} else {
		data := table_data.Tbl_runcode_documents{
			Kode:  paramCreateCode.Kode,
			Tahun: tahunBulan.Tahun,
			Bulan: tahunBulan.Bulan,
			Nomor: 1,
		}
		err := db.Create(&data).Error
		if err != nil {
			response := helper.APIResponse("Genearate Kode Document Gagal ...", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		CodeGenerate := tahunBulan.Tahun + "" + tahunBulan.Bulan + "-" + paramCreateCode.Kode + "" + helper.StrPad("1", 4, "0", "LEFT")

		response := helper.APIResponse("Genearate Kode Document Berhasil ...", http.StatusOK, "success", CodeGenerate)
		c.JSON(http.StatusOK, response)
		return
	}

}
