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

	// if paramCreateCode.Namahalaman == "biaya-eksternal" || paramCreateCode.Namahalaman == "biaya-operasional" || paramCreateCode.Namahalaman == "perencanaan-biaya" {

	// } else {
	// response := helper.APIResponse("Parameter Nama Halaman Salah ...", http.StatusBadRequest, "success", paramCreateCode.Namahalaman)
	// c.JSON(http.StatusOK, response)
	// return
	// }

	// if paramCreateCode.Kode == "BE" || paramCreateCode.Kode == "BO" || paramCreateCode.Kode == "PB" {

	// } else {
	// 	response := helper.APIResponse("Parameter Kode Salah ...", http.StatusBadRequest, "success", paramCreateCode.Kode)
	// 	c.JSON(http.StatusOK, response)
	// 	return
	// }

	if paramCreateCode.Namahalaman == "biaya-eksternal" && paramCreateCode.Kode == "BE" {

	} else if paramCreateCode.Namahalaman == "biaya-operasional" && paramCreateCode.Kode == "BO" {

	} else if paramCreateCode.Namahalaman == "perencanaan-biaya" && paramCreateCode.Kode == "PB" {

	} else {
		response := helper.APIResponse("Parameter Kode dan Nama Halaman Salah ...", http.StatusBadRequest, "success", "kode => "+paramCreateCode.Kode+" namahalaman => "+paramCreateCode.Namahalaman)
		c.JSON(http.StatusOK, response)
		return
	}

	var tahunBulan TahunBulan
	sql := " SELECT DATE_FORMAT(CURRENT_DATE,'%Y') as 'tahun',DATE_FORMAT(CURRENT_DATE,'%m') as 'bulan' "
	db.Raw(sql).Scan(&tahunBulan)

	var tbl_runcode_documents table_data.Tbl_runcode_documents
	checkUser := db.Select("*").Where("tahun = ? and bulan=? and kode=? "+
		" and namahalaman=?", tahunBulan.Tahun, tahunBulan.Bulan, paramCreateCode.Kode, paramCreateCode.Namahalaman).Find(&tbl_runcode_documents)
	if checkUser.RowsAffected > 0 {

		var GetAvailableNomor string
		db.Raw("SELECT generate_nomor FROM tbl_runcode_documents where kode=? "+
			" and tahun=? and bulan=? "+
			" and namahalaman=?", paramCreateCode.Kode, tahunBulan.Tahun, tahunBulan.Bulan, paramCreateCode.Namahalaman).Scan(&GetAvailableNomor)

		//Cek jika nomor sudah digenerate tapi belum dipakai maka pakai yang sudah ada
		if paramCreateCode.Namahalaman == "biaya-eksternal" {
			var tbl_trans_uang_masuk_lain_headers table_data.Tbl_trans_uang_masuk_lain_headers
			checkNomor := db.Select("*").Where("no_document=?", GetAvailableNomor).Find(&tbl_trans_uang_masuk_lain_headers)
			if checkNomor.RowsAffected == 0 {
				response := helper.APIResponse("Generate Kode Document Berhasil ...", http.StatusOK, "success", GetAvailableNomor)
				c.JSON(http.StatusOK, response)
				return
			}
		}

		if paramCreateCode.Namahalaman == "biaya-operasional" {
			var tbl_trans_uang_keluar_act_headers table_data.Tbl_trans_uang_keluar_act_headers
			checkNomor := db.Select("*").Where("no_document=?", GetAvailableNomor).Find(&tbl_trans_uang_keluar_act_headers)
			if checkNomor.RowsAffected == 0 {
				response := helper.APIResponse("Generate Kode Document Berhasil ...", http.StatusOK, "success", GetAvailableNomor)
				c.JSON(http.StatusOK, response)
				return
			}
		}

		if paramCreateCode.Namahalaman == "perencanaan-biaya" {
			var tbl_trans_uang_keluar_pra_headers table_data.Tbl_trans_uang_keluar_pra_headers
			checkNomor := db.Select("*").Where("no_document=?", GetAvailableNomor).Find(&tbl_trans_uang_keluar_pra_headers)
			if checkNomor.RowsAffected == 0 {
				response := helper.APIResponse("Generate Kode Document Berhasil ...", http.StatusOK, "success", GetAvailableNomor)
				c.JSON(http.StatusOK, response)
				return
			}
		}

		//End Cek jika nomor sudah digenerate tapi belum dipakai maka pakai yang sudah ada

		var nomor int
		db.Raw("SELECT nomor FROM tbl_runcode_documents where kode=? "+
			" and tahun=? and bulan=? and namahalaman=?", paramCreateCode.Kode, tahunBulan.Tahun, tahunBulan.Bulan, paramCreateCode.Namahalaman).Scan(&nomor)
		nomor = nomor + 1

		CodeGenerate := tahunBulan.Tahun + "" + tahunBulan.Bulan + "-" + paramCreateCode.Kode + "" + helper.StrPad(strconv.Itoa(nomor), 4, "0", "LEFT")

		err := db.Raw("update tbl_runcode_documents set nomor=?,generate_nomor=? where kode=? "+
			" and tahun=? and bulan=? ", nomor, CodeGenerate, paramCreateCode.Kode, tahunBulan.Tahun, tahunBulan.Bulan).Scan(&tbl_runcode_documents).Error
		if err != nil {
			response := helper.APIResponse("Generate Kode Document Gagal ...", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.APIResponse("Generate Kode Document Berhasil ...", http.StatusOK, "success", CodeGenerate)
		c.JSON(http.StatusOK, response)
		return
	} else {
		CodeGenerate := tahunBulan.Tahun + "" + tahunBulan.Bulan + "-" + paramCreateCode.Kode + "" + helper.StrPad("1", 4, "0", "LEFT")

		data := table_data.Tbl_runcode_documents{
			Kode:           paramCreateCode.Kode,
			Tahun:          tahunBulan.Tahun,
			Bulan:          tahunBulan.Bulan,
			Nomor:          1,
			Generate_nomor: CodeGenerate,
			Namahalaman:    paramCreateCode.Namahalaman,
		}
		err := db.Create(&data).Error
		if err != nil {
			response := helper.APIResponse("Generate Kode Document Gagal ...", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.APIResponse("Generate Kode Document Berhasil ...", http.StatusOK, "success", CodeGenerate)
		c.JSON(http.StatusOK, response)
		return
	}

}
