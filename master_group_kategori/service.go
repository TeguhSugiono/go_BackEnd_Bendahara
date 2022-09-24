package master_group_kategori

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

func InsertGroupKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var dataInput GroupKategoriInput
	err := c.ShouldBindJSON(&dataInput)
	if err != nil {
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

	var dataMaster table_data.Tbl_jenis_trans
	if err := db.Where("kd_jenis = ? and flag_aktif=? ", dataInput.Kd_jenis, 0).First(&dataMaster).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
		response := helper.APIResponse("Simpan Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
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
	data := table_data.Tbl_group_kategoris{
		Kd_jenis:   dataInput.Kd_jenis,
		Nm_group:   dataInput.Nm_group,
		Created_by: currentUser.(string),
		Created_on: datenowx,
		Flag_aktif: 0,
	}

	err = db.Omit("Edited_on", "Edited_by", "Nm_header").Create(&data).Error
	if err != nil {
		response := helper.APIResponse("Simpan Data Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Simpan Data Sukses ...", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func ListGroupKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var master []ListData
	//db.Where("flag_aktif = ?", 0).Find(&master)

	sql := "  SELECT a.*,b.proses_uang FROM tbl_group_kategoris as a " +
		" inner join tbl_jenis_trans as b on a.kd_jenis=b.kd_jenis  " +
		" where a.flag_aktif=0 and b.flag_aktif=0  "

	db.Raw(sql).Scan(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", FormatGroupKategori(master))
	c.JSON(http.StatusOK, response)
}

func ShowGroupKategori(c *gin.Context) {
	//pagination terdiri dari
	//1. page (halaman kebarapa)
	//2. perpage (jumlah perhalaman yang ditampilkan)
	//3. search (data yang dicari berdasarkan setiap coulumn)
	//4. ordering (sorting data berdasarkan asc dan desc)

	db := c.MustGet("db").(*gorm.DB)

	var master []ListData

	sql := "  SELECT a.*,b.proses_uang FROM tbl_group_kategoris as a " +
		" inner join tbl_jenis_trans as b on a.kd_jenis=b.kd_jenis  " +
		" where a.flag_aktif=0 and b.flag_aktif=0  "
	//sql = fmt.Sprintf("%s inner join tbl_jenis_trans as b on a.kd_jenis=b.kd_jenis ", sql)
	//sql = fmt.Sprintf("%s where a.flag_aktif=0 and b.flag_aktif=0 ", sql)
	//sql = sql + " inner join tbl_jenis_trans as b on a.kd_jenis=b.kd_jenis "
	//sql = sql + " where a.flag_aktif=0 and b.flag_aktif=0  "

	if s := c.Query("search"); s != "" {
		if len(c.Query("search")) >= 3 {
			// CompTableData := TableData{
			// 	Total:     0,
			// 	Page:      0,
			// 	Last_page: 0,
			// }
			// response := helper.APIResponseTable("List Data ...", http.StatusOK, "success", "", CompTableData, FormatJenisTrans(master))
			// c.JSON(http.StatusOK, response)
			// return
			sql = fmt.Sprintf("%s and a.nm_group LIKE '%%%s%%' ", sql, s)
		}
	}

	if filter := c.Query("filter"); filter != "" {
		sql = fmt.Sprintf("%s and b.proses_uang = '%s'", sql, filter)
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

	CompTableData := TableData{
		Total:     total,
		Page:      intpage,
		Last_page: int(math.Ceil(float64(total) / float64(intperPage))),
	}

	response := helper.APIResponseTable("List Data ...", http.StatusOK, "success", "", CompTableData, FormatGroupKategori(master))
	c.JSON(http.StatusOK, response)

	// return c.JSON(http.StatusOK.Map{
	// 	"data":      products,
	// 	"total":     total,
	// 	"page":      page,
	// 	"last_page": math.Ceil(float64(total / int64(perPage))),
	// })
}

func UpdateGroupKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input GroupKategoriInput
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

	var dataUtama table_data.Tbl_group_kategoris
	if err := db.Where("kd_group = ? and flag_aktif=? ", c.Param("kdgroup"), 0).First(&dataUtama).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
		response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var dataMaster table_data.Tbl_jenis_trans
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
		errorMessage := gin.H{"errors": errors, "tgl": datenowx}
		response := helper.APIResponse("Format Tanggal Salah ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser")

	data := table_data.Tbl_group_kategoris{
		Kd_jenis:   input.Kd_jenis,
		Nm_group:   input.Nm_group,
		Edited_by:  currentUser.(string),
		Edited_on:  datenowx,
		Flag_aktif: 0,
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

func DeleteGroupKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var dataMaster table_data.Tbl_group_kategoris
	if err := db.Where("kd_group = ? and flag_aktif=? ", c.Param("kdgroup"), 0).First(&dataMaster).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
		response := helper.APIResponse("Delete Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//========= Cek Data
	//tbl_group_kategoris
	var dataCek table_data.Tbl_kategori_uangs
	result := db.Where("kd_group = ? and flag_aktif=? ", dataMaster.Kd_group, 0).First(&dataCek)
	if result.RowsAffected > 0 {
		errorMessage := gin.H{"errors": "Data Sudah Terpakai Di Master Kategori Uang ..."}
		response := helper.APIResponse("Delete Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	//========= Cek Data

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

	data := table_data.Tbl_group_kategoris{
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
