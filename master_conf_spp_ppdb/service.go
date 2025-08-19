package master_conf_spp_ppdb

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

func ShowConfSppPPDB(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var master []ListData

	sql := " SELECT a.id_link,a.link_name,a.kd_group,b.nm_group,a.kd_kategori,c.nm_kategori " +
		" FROM tbl_link_kategoris a " +
		" LEFT JOIN tbl_group_kategoris b on a.kd_group=b.kd_group " +
		" LEFT JOIN tbl_kategori_uangs c on a.kd_kategori=c.kd_kategori "

	if s := c.Query("search"); s != "" {
		if len(c.Query("search")) >= 3 {
			sql = fmt.Sprintf("%s and b.nm_group LIKE '%%%s%%' ", sql, s)
			sql = fmt.Sprintf("%s and c.nm_kategori LIKE '%%%s%%' ", sql, s)
		}
	}

	if sort := c.Query("sort"); sort != "" {
		sql = fmt.Sprintf("%s ORDER BY a.link_name %s", sql, sort)
	} else {
		sql = fmt.Sprintf("%s ORDER BY a.link_name %s", sql, "desc")
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

func UpdateConfSppPPDB(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var dataMaster table_data.Tbl_link_kategoris
	if err := db.Where("id_link = ? ", c.Param("idlink")).First(&dataMaster).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
		response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Validate input
	var input ParamInput
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

	//Chek data Tbl_group_kategoris
	var tbl_group_kategoris table_data.Tbl_group_kategoris
	if err := db.Where("flag_aktif=0 and kd_group=?", input.Kd_group).First(&tbl_group_kategoris).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Group Tidak Ditemukan ..."}
		response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//Chek data Tbl_kategori_uangs
	var tbl_kategori_uangs table_data.Tbl_kategori_uangs
	if err := db.Where("flag_aktif=0 and kd_kategori=?", input.Kd_kategori).First(&tbl_kategori_uangs).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Kategori Tidak Ditemukan ..."}
		response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := table_data.Tbl_link_kategoris{
		Kd_group:    input.Kd_group,
		Kd_kategori: input.Kd_kategori,
	}

	err := db.Model(&dataMaster).Updates(data).Error
	if err != nil {
		response := helper.APIResponse("Update Data Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Update Data Sukses ...", http.StatusOK, "success", dataMaster)
	c.JSON(http.StatusOK, response)
}
