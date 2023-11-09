package master_conf_biaya_kategori

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"rest_api_bendahara/helper"
	"rest_api_bendahara/table_data"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ListBiayaKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var master []ListData

	sql := " SELECT a.kd_biaya_kategori,a.kd_kategori,b.nm_kategori,a.jml_biaya " +
		" FROM tbl_biaya_masuk_keluars a " +
		" INNER JOIN tbl_kategori_uangs b on a.kd_kategori=b.kd_kategori "

	db.Raw(sql).Scan(&master)
	
	
	if len(master) == 0 {
		SetArrayDataMasuk := []ListData{}
		response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayDataMasuk)
		c.JSON(http.StatusOK, response)
		return
	}
	
	response := helper.APIResponse("List Data ...", http.StatusOK, "success", master)
	c.JSON(http.StatusOK, response)
}

func ShowBiayaKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var master []ListData

	sql := " SELECT a.kd_biaya_kategori,a.kd_kategori,b.nm_kategori,a.jml_biaya " +
		" FROM tbl_biaya_masuk_keluars a " +
		" INNER JOIN tbl_kategori_uangs b on a.kd_kategori=b.kd_kategori where a.kd_biaya_kategori <> '' "

	if s := c.Query("search"); s != "" {
		if len(c.Query("search")) >= 3 {
			sql = fmt.Sprintf("%s and b.nm_kategori LIKE '%%%s%%' ", sql, s)
		}
	}

	if sort := c.Query("sort"); sort != "" {
		sql = fmt.Sprintf("%s ORDER BY b.nm_kategori %s", sql, sort)
	} else {
		sql = fmt.Sprintf("%s ORDER BY b.nm_kategori %s", sql, "desc")
	}

	page := c.Query("page")
	perPage := c.Query("perpage")

	intpage, err := strconv.Atoi(page)
	if err != nil {
		response := helper.APIResponse("Format Page Salah ...", http.StatusUnprocessableEntity, "error", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	intperPage, err := strconv.Atoi(perPage)
	if err != nil {
		response := helper.APIResponse("Format Perpage Salah ...", http.StatusUnprocessableEntity, "error", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var total int64

	db.Raw(sql).Count(&total)

	sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, intperPage, (intpage-1)*intperPage)
	db.Raw(sql).Scan(&master)

	CompTableData := table_data.TableData{
		Total:     total,
		Page:      intpage,
		Last_page: int(math.Ceil(float64(total) / float64(intperPage))),
	}


	if len(master) == 0 {
		SetArrayDataMasuk := []ListData{}
		response := helper.APIResponseTable("List Data ...", http.StatusOK, "success", "", CompTableData, SetArrayDataMasuk)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponseTable("List Data ...", http.StatusOK, "success", "", CompTableData, master)
	c.JSON(http.StatusOK, response)

}

func InsertBiayaKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramInputBiayaKategori ParamInputBiayaKategori
	if err := c.ShouldBindJSON(&paramInputBiayaKategori); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errors := helper.FormatValidationError(err)
			errorMessage := gin.H{"errors": errors}
			response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		var error_binding []string
		error_binding = append(error_binding, err.Error())
		errorMessage := gin.H{"errors": error_binding}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var dataMaster table_data.Tbl_kategori_uangs
	if err := db.Where("kd_kategori = ? and flag_aktif=? ", paramInputBiayaKategori.Kd_kategori, 0).First(&dataMaster).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
		response := helper.APIResponse("Simpan Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var dataConf table_data.Tbl_biaya_masuk_keluars
	checkBiaya := db.Select("*").Where("kd_kategori = ?", paramInputBiayaKategori.Kd_kategori).Find(&dataConf)
	if checkBiaya.RowsAffected > 0 {
		errorMessage := gin.H{"errors": "Data Konfigurasi Sudah Ada ..."}
		response := helper.APIResponse("Simpan Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := table_data.Tbl_biaya_masuk_keluars{
		Kd_kategori: paramInputBiayaKategori.Kd_kategori,
		Jml_biaya:   paramInputBiayaKategori.Jml_biaya,
	}

	err := db.Create(&data).Error
	if err != nil {
		response := helper.APIResponse("Simpan Data Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Simpan Data Sukses ...", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func UpdateBiayaKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramInputBiayaKategori ParamInputBiayaKategori
	if err := c.ShouldBindJSON(&paramInputBiayaKategori); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errors := helper.FormatValidationError(err)
			errorMessage := gin.H{"errors": errors}
			response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		var error_binding []string
		error_binding = append(error_binding, err.Error())
		errorMessage := gin.H{"errors": error_binding}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var dataMaster table_data.Tbl_kategori_uangs
	if err := db.Where("kd_kategori = ? and flag_aktif=? ", paramInputBiayaKategori.Kd_kategori, 0).First(&dataMaster).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
		response := helper.APIResponse("Simpan Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var int_kd_kategori_old int
	db.Raw("select kd_kategori from tbl_biaya_masuk_keluars where kd_biaya_kategori =? ", c.Param("kdbiayakategori")).Scan(&int_kd_kategori_old)

	var dataConf table_data.Tbl_biaya_masuk_keluars
	if paramInputBiayaKategori.Kd_kategori != int_kd_kategori_old {
		checkBiaya := db.Select("*").Where("kd_kategori = ?", paramInputBiayaKategori.Kd_kategori).Find(&dataConf)
		if checkBiaya.RowsAffected > 0 {
			errorMessage := gin.H{"errors": "Data Konfigurasi Sudah Ada ..."}
			response := helper.APIResponse("Simpan Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}

	err := db.Raw("UPDATE tbl_biaya_masuk_keluars SET kd_kategori = ?,jml_biaya=? WHERE kd_biaya_kategori = ? ", paramInputBiayaKategori.Kd_kategori, paramInputBiayaKategori.Jml_biaya, c.Param("kdbiayakategori")).Scan(&dataConf).Error
	if err != nil {
		response := helper.APIResponse("Update Data Ke tbl_biaya_masuk_keluars Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Update Data Sukses ...", http.StatusOK, "success", paramInputBiayaKategori)
	c.JSON(http.StatusOK, response)
}

func DeleteBiayaKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var dataConf table_data.Tbl_biaya_masuk_keluars
	checkBiaya := db.Select("*").Where("kd_biaya_kategori = ?", c.Param("kdbiayakategori")).Find(&dataConf)
	if checkBiaya.RowsAffected == 0 {
		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
		response := helper.APIResponse("Delete Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := db.Where("kd_biaya_kategori = ?", c.Param("kdbiayakategori")).Delete(&dataConf).Error
	if err != nil {
		response := helper.APIResponse("Delete Data Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Delete Data Sukses ...", http.StatusOK, "success", dataConf)
	c.JSON(http.StatusOK, response)

}
