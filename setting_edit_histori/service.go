package setting_edit_histori

import (
	"errors"
	"net/http"
	"rest_api_bendahara/helper"
	"time"

	"rest_api_bendahara/table_data"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func Get_Akses(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var statusOpen StatusOpen

	if err := c.ShouldBindJSON(&statusOpen); err != nil {
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

	var settingdata table_data.Tbl_open_lock_historis
	err = db.Raw("UPDATE tbl_open_lock_historis SET open = ?,request_by=?,request_on=? ", statusOpen.Open, currentUser.(string), datenowx).Scan(&settingdata).Error
	if err != nil {
		response := helper.APIResponse("Update Data Ke tbl_open_lock_historis Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var pesan string
	if statusOpen.Open == "Y" {
		pesan = "Akses Tombol didapatkan ..."
	} else if statusOpen.Open == "N" {
		pesan = "Akses Tombol dihilangkan ..."
	} else {
		pesan = "Parameter Masukan Salah ..."
	}

	response := helper.APIResponse(pesan, http.StatusOK, "success", statusOpen)
	c.JSON(http.StatusOK, response)

}
