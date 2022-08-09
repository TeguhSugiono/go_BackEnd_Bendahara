package master_group_kategori

import (
	"errors"
	"net/http"
	"time"

	"rest_api_bendahara/helper"
	"rest_api_bendahara/master_jenis_trans"

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

	var dataMaster master_jenis_trans.Tbl_jenis_trans
	if err := db.Where("kd_jenis = ? and flag_aktif=? ", dataInput.Kd_jenis, 0).First(&dataMaster).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Not Found ..."}
		response := helper.APIResponse("Update Data Failed...", http.StatusUnprocessableEntity, "error", errorMessage)
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
	data := Tbl_group_kategoris{
		Kd_jenis:   dataInput.Kd_jenis,
		Nm_group:   dataInput.Nm_group,
		Created_by: currentUser.(string),
		Created_on: datenowx,
		Flag_aktif: 0,
	}

	err = db.Omit("Edited_on", "Edited_by", "Nm_header", "Nm_detail").Create(&data).Error
	if err != nil {
		response := helper.APIResponse("Save Data Failed ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Save Data Successfully ...", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func ShowGroupKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var master []Tbl_group_kategoris
	db.Where("flag_aktif = ?", 0).Find(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", FormatGroupKategori(master))
	c.JSON(http.StatusOK, response)
}

func UpdateGroupKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var dataMaster Tbl_group_kategoris
	if err := db.Where("kd_group = ? and flag_aktif=? ", c.Param("kdgroup"), 0).First(&dataMaster).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Not Found ..."}
		response := helper.APIResponse("Update Data Failed ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var input GroupKategoriInput
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

	data := Tbl_group_kategoris{
		Kd_jenis:   input.Kd_jenis,
		Nm_group:   input.Nm_group,
		Edited_by:  currentUser.(string),
		Edited_on:  datenowx,
		Flag_aktif: 0,
	}

	err = db.Model(&dataMaster).Updates(data).Error
	if err != nil {
		response := helper.APIResponse("Update Data Failed ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Update Data Successfully ...", http.StatusOK, "success", dataMaster)
	c.JSON(http.StatusOK, response)
}

func DeleteGroupKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var dataMaster Tbl_group_kategoris
	if err := db.Where("kd_group = ? and flag_aktif=? ", c.Param("kdgroup"), 0).First(&dataMaster).Error; err != nil {
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

	data := Tbl_group_kategoris{
		Edited_by:  currentUser.(string),
		Edited_on:  datenowx,
		Flag_aktif: 9,
	}

	err = db.Model(&dataMaster).Updates(data).Error
	if err != nil {
		response := helper.APIResponse("Delete Data Failed ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Delete Data Successfully ...", http.StatusOK, "success", dataMaster)
	c.JSON(http.StatusOK, response)
}
