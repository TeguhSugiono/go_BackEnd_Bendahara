package master_kategori_uang

import (
	"errors"
	"net/http"
	"time"

	"rest_api_bendahara/helper"
	"rest_api_bendahara/master_group_kategori"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ShowKategoriUang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var master []Tbl_kategori_uangs
	//db.Find(&master)
	db.Where("flag_aktif = ?", 0).Find(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", FormatShowData(master))
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
			response := helper.APIResponse("Error Validation ...", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		var error_binding []string
		error_binding = append(error_binding, err.Error())
		errorMessage := gin.H{"errors": error_binding}
		response := helper.APIResponse("Error Validation ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//time.Now().Format("2006-01-02 15:04:05")
	var datenows string = time.Now().UTC().Format("2006-01-02 15:04:05")
	date := "2006-01-02 15:04:05"
	datenowx, err := time.Parse(date, datenows)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors, "date": datenowx}
		response := helper.APIResponse("Wrong Date Format ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser")
	data := Tbl_kategori_uangs{
		Kd_group:    dataInput.Kd_group,
		Nm_kategori: dataInput.Nm_kategori,
		Created_by:  currentUser.(string),
		Created_on:  datenowx,
		Flag_aktif:  0,
	}
	// var data models.Tbl_jenis_trans
	// data.Proses_uang = dataInput.Proses_uang
	// data.Created_on = datenow
	// data.Created_by = currentUser.(string)
	// data.Flag_aktif = 0
	//err = db.Omit("Edited_on", "Edited_by").Create(&data).Error
	err = db.Omit("Edited_on", "Edited_by").Create(&data).Error
	if err != nil {
		response := helper.APIResponse("Save Data Failed ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Raw SQL
	var nm_kategori string
	var str_nm_kategori string

	//str_nm_kategori := ""
	//var jmldata int
	//db.Raw("select COUNT(*) as 'jmldata'  from tbl_kategori_uangs where flag_aktif = ?", 0).Scan(&jmldata)

	nomor := 0
	rows, _ := db.Raw("select nm_kategori from tbl_kategori_uangs where flag_aktif = ? and kd_group=?", 0, dataInput.Kd_group).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&nm_kategori)
		if nomor == 0 {
			str_nm_kategori = nm_kategori
		} else {
			str_nm_kategori = str_nm_kategori + "," + nm_kategori
		}
		nomor++
	}

	var dataMasterKategoris master_group_kategori.Tbl_group_kategoris
	data_group_kategoris := DataKdGroup{
		Nm_header: str_nm_kategori,
	}
	db.Model(&dataMasterKategoris).Updates(data_group_kategoris)

	response := helper.APIResponse("Save Data Successfully ...", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func UpdateKategoriUang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var dataMaster Tbl_kategori_uangs
	if err := db.Where("kd_kategori = ? and flag_aktif=? ", c.Param("kdkategori"), 0).First(&dataMaster).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Not Found ..."}
		response := helper.APIResponse("Update Data Failed ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Validate input
	var input KategoriUangInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errors := helper.FormatValidationError(err)
			errorMessage := gin.H{"errors": errors}
			response := helper.APIResponse("Error Validation ...", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		var error_binding []string
		error_binding = append(error_binding, err.Error())
		errorMessage := gin.H{"errors": error_binding}
		response := helper.APIResponse("Error Validation ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var datenows string = time.Now().UTC().Format("2006-01-02 15:04:05")
	date := "2006-01-02 15:04:05"
	datenowx, err := time.Parse(date, datenows)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors, "date": datenowx}
		response := helper.APIResponse("Wrong Date Format ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser")

	data := Tbl_kategori_uangs{
		Kd_group:    input.Kd_group,
		Nm_kategori: input.Nm_kategori,
		Edited_by:   currentUser.(string),
		Edited_on:   datenowx,
		Flag_aktif:  0,
	}

	// var data models.Tbl_jenis_trans
	// data.Proses_uang = input.Proses_uang
	// data.Edited_by = currentUser.(string)
	// data.Flag_aktif = 0

	err = db.Model(&dataMaster).Updates(data).Error
	//err = db.Model(&dataMaster).Omit("Created_on", "Created_by").Updates(&data).Error
	//err = db.Model(&dataMaster).Omit("Created_on").Updates(data)
	if err != nil {
		response := helper.APIResponse("Update Data Failed ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Update Data Successfully ...", http.StatusOK, "success", dataMaster)
	c.JSON(http.StatusOK, response)
}

func DeleteKategoriUang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var dataMaster Tbl_kategori_uangs
	if err := db.Where("kd_kategori = ? and flag_aktif=?", c.Param("kdkategori"), 0).First(&dataMaster).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Not Found ..."}
		response := helper.APIResponse("Delete Data Failed ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var datenows string = time.Now().UTC().Format("2006-01-02 15:04:05")
	date := "2006-01-02 15:04:05"
	datenowx, err := time.Parse(date, datenows)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors, "date": datenowx}
		response := helper.APIResponse("Wrong Date Format ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser")

	data := Tbl_kategori_uangs{
		Edited_by:  currentUser.(string),
		Edited_on:  datenowx,
		Flag_aktif: 9,
	}

	// var data models.Tbl_jenis_trans
	// data.Proses_uang = input.Proses_uang
	// data.Edited_by = currentUser.(string)
	// data.Flag_aktif = 0

	err = db.Model(&dataMaster).Updates(data).Error
	//err = db.Model(&dataMaster).Omit("Created_on", "Created_by").Updates(&data).Error
	//err = db.Model(&dataMaster).Omit("Created_on").Updates(data)
	if err != nil {
		response := helper.APIResponse("Delete Data Failed ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Delete Data Successfully ...", http.StatusOK, "success", dataMaster)
	c.JSON(http.StatusOK, response)
}
