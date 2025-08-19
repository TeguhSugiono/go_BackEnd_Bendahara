package master_kategori_uang

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"rest_api_bendahara/helper"
	"rest_api_bendahara/table_data"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ListKategoriUang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var master []ListData

	sql := "SELECT a.kd_kategori,a.nm_kategori,a.nm_detail,b.proses_uang,a.kd_jenis " +
		" FROM tbl_kategori_uangs a " +
		" INNER JOIN tbl_jenis_trans b on a.kd_jenis=b.kd_jenis " +
		" where a.flag_aktif=0 and b.flag_aktif=0"

	db.Raw(sql).Scan(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", FormatShowData(master))
	c.JSON(http.StatusOK, response)
}

func ListKategoriUangId(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramData SearchKategoriUang
	if err := c.ShouldBindJSON(&paramData); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	var master []ListData

	sql := "SELECT a.kd_kategori,a.nm_kategori,a.nm_detail,b.proses_uang,a.kd_jenis " +
		" FROM tbl_kategori_uangs a " +
		" INNER JOIN tbl_jenis_trans b on a.kd_jenis=b.kd_jenis " +
		" where a.flag_aktif=0 and b.flag_aktif=0"

	if paramData.Kd_jenis != "" {
		sql = fmt.Sprintf("%s and a.kd_jenis = '%s'", sql, paramData.Kd_jenis)
	}

	db.Raw(sql).Scan(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", FormatShowData(master))
	c.JSON(http.StatusOK, response)
}

func ShowKategoriUang(c *gin.Context) {

	//pagination terdiri dari
	//1. page (halaman kebarapa)
	//2. perpage (jumlah perhalaman yang ditampilkan)
	//3. search (data yang dicari berdasarkan setiap coulumn)
	//4. ordering (sorting data berdasarkan asc dan desc)

	db := c.MustGet("db").(*gorm.DB)

	var master []ListData

	sql := "SELECT a.kd_kategori,a.nm_kategori,a.nm_detail,b.proses_uang,a.kd_jenis " +
		" FROM tbl_kategori_uangs a " +
		" INNER JOIN tbl_jenis_trans b on a.kd_jenis=b.kd_jenis " +
		" where a.flag_aktif=0 and b.flag_aktif=0"

	if s := c.Query("search"); s != "" {

		if len(c.Query("search")) >= 3 {
			sql = fmt.Sprintf("%s and a.nm_kategori LIKE '%%%s%%' ", sql, s)
		}

	}

	if sort := c.Query("sort"); sort != "" {
		sql = fmt.Sprintf("%s ORDER BY a.created_on %s", sql, sort)
	} else {
		sql = fmt.Sprintf("%s ORDER BY a.created_on %s", sql, "desc")
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

	response := helper.APIResponseTable("List Data ...", http.StatusOK, "success", "", CompTableData, FormatShowData(master))
	c.JSON(http.StatusOK, response)

}

func InsertKategoriUang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var dataInput KategoriUangInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
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

	var datenows string = time.Now().UTC().Format("2006-01-02 15:04:05")
	date := "2006-01-02 15:04:05"
	datenowx, err := time.Parse(date, datenows)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors, "tgl": datenowx}
		response := helper.APIResponse("Format Tanggal Salah ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser")
	data := table_data.Tbl_kategori_uangs{
		Kd_jenis:    dataInput.Kd_jenis,
		Nm_kategori: dataInput.Nm_kategori,
		Created_by:  currentUser.(string),
		Created_on:  datenowx,
		Flag_aktif:  0,
	}

	err = db.Omit("Edited_on", "Edited_by", "Nm_detail").Create(&data).Error
	if err != nil {
		response := helper.APIResponse("Simpan Data Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Simpan Data Sukses ...", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func UpdateKategoriUang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input KategoriUangInput
	if err := c.ShouldBindJSON(&input); err != nil {
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

	var dataUtama table_data.Tbl_kategori_uangs
	if err := db.Where("kd_kategori = ? and flag_aktif=? ", c.Param("kdkategori"), 0).First(&dataUtama).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
		response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var dataMaster table_data.Tbl_group_kategoris
	if err := db.Where("kd_jenis = ? and flag_aktif=? ", input.Kd_jenis, 0).First(&dataMaster).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
		response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var datenows string = time.Now().UTC().Format("2006-01-02 15:04:05")
	date := "2006-01-02 15:04:05"
	datenowx, err := time.Parse(date, datenows)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors, "date": datenowx}
		response := helper.APIResponse("Tanggal Format Salah ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser")

	data := table_data.Tbl_kategori_uangs{
		Kd_jenis:    input.Kd_jenis,
		Nm_kategori: input.Nm_kategori,
		Edited_by:   currentUser.(string),
		Edited_on:   datenowx,
		Flag_aktif:  0,
	}

	err = db.Model(&dataUtama).Updates(data).Error

	if err != nil {
		response := helper.APIResponse("Update Data Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Update Data Sukses ...", http.StatusOK, "success", dataUtama)
	c.JSON(http.StatusOK, response)
}

func DeleteKategoriUang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var dataMaster table_data.Tbl_kategori_uangs
	if err := db.Where("kd_kategori = ? and flag_aktif=?", c.Param("kdkategori"), 0).First(&dataMaster).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
		response := helper.APIResponse("Delete Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var dataCek table_data.Tbl_sub_kategori_uangs
	result := db.Where("kd_kategori = ? and flag_aktif=? ", dataMaster.Kd_kategori, 0).First(&dataCek)
	if result.RowsAffected > 0 {
		errorMessage := gin.H{"errors": "Data Sudah Terpakai Di Master Sub Kategori Uang ..."}
		response := helper.APIResponse("Delete Data Gagal \n Data Sudah Terpakai Di Master Sub Kategori Uang PPDB...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var dataCek2 table_data.Tbl_biaya_masuk_keluars
	result1 := db.Where("kd_kategori = ? ", dataMaster.Kd_kategori).First(&dataCek2)
	if result1.RowsAffected > 0 {
		errorMessage := gin.H{"errors": "Data Sudah Terpakai Di Master Biaya Kategori ..."}
		response := helper.APIResponse("Delete Data Gagal ... \n Data Sudah Terpakai Di Master Biaya Kategori ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var datenows string = time.Now().UTC().Format("2006-01-02 15:04:05")
	date := "2006-01-02 15:04:05"
	datenowx, err := time.Parse(date, datenows)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors, "tgl": datenowx}
		response := helper.APIResponse("Format Tanggal Salah ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser")

	data := table_data.Tbl_kategori_uangs{
		Edited_by:  currentUser.(string),
		Edited_on:  datenowx,
		Flag_aktif: 9,
	}

	err = db.Model(&dataMaster).Updates(data).Error

	if err != nil {
		response := helper.APIResponse("Delete Data Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Delete Data Sukses ...", http.StatusOK, "success", dataMaster)
	c.JSON(http.StatusOK, response)
}
