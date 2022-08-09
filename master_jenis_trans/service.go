package master_jenis_trans

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"rest_api_bendahara/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// func ShowJenisTrans(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)

// 	var master []Tbl_jenis_trans
// 	//db.Find(&master)
// 	db.Where("flag_aktif = ?", 0).Find(&master)

// 	response := helper.APIResponse("List Data ...", http.StatusOK, "success", FormatJenisTrans(master))
// 	c.JSON(http.StatusOK, response)
// }

func ShowJenisTrans(c *gin.Context) {

	//pagination terdiri dari
	//1. page (halaman kebarapa)
	//2. perpage (jumlah perhalaman yang ditampilkan)
	//3. search (data yang dicari berdasarkan setiap coulumn)
	//4. ordering (sorting data berdasarkan asc dan desc)

	db := c.MustGet("db").(*gorm.DB)

	var master []Tbl_jenis_trans

	sql := "SELECT * FROM tbl_jenis_trans where flag_aktif=0 "

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
			sql = fmt.Sprintf("%s and proses_uang LIKE '%%%s%%' ", sql, s)
		}

	}

	if sort := c.Query("sort"); sort != "" {
		sql = fmt.Sprintf("%s ORDER BY proses_uang %s", sql, sort)
	}

	page := c.Query("page")
	perPage := c.Query("perpage")

	intpage, _ := strconv.Atoi(page)
	intperPage, _ := strconv.Atoi(perPage)

	//strconv.ParseInt(s, 10, 64)
	//sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, intperPage, (intpage-1)*intperPage)

	// page, _ := strconv.Atoi(c.Query("page"), "1")
	// perPage := 9

	var total int64

	db.Raw(sql).Count(&total)

	sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, intperPage, (intpage-1)*intperPage)
	db.Raw(sql).Scan(&master)

	CompTableData := TableData{
		Total:     total,
		Page:      intpage,
		Last_page: int(math.Ceil(float64(total) / float64(intperPage))),
	}

	response := helper.APIResponseTable("List Data ...", http.StatusOK, "success", sql, CompTableData, FormatJenisTrans(master))
	c.JSON(http.StatusOK, response)

	// return c.JSON(http.StatusOK.Map{
	// 	"data":      products,
	// 	"total":     total,
	// 	"page":      page,
	// 	"last_page": math.Ceil(float64(total / int64(perPage))),
	// })
}

func InsertJenisTrans(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var dataInput JenisTransInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
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
	data := Tbl_jenis_trans{
		Proses_uang: dataInput.Proses_uang,
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

	response := helper.APIResponse("Save Data Successfully ...", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func UpdateJenisTrans(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var dataMaster Tbl_jenis_trans
	if err := db.Where("kd_jenis = ? and flag_aktif=? ", c.Param("kdjenis"), 0).First(&dataMaster).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Not Found ..."}
		response := helper.APIResponse("Update Data Failed ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Validate input
	var input JenisTransInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
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

	data := Tbl_jenis_trans{
		Proses_uang: input.Proses_uang,
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

func DeleteJenisTrans(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var dataMaster Tbl_jenis_trans
	if err := db.Where("kd_jenis = ? and flag_aktif=?", c.Param("kdjenis"), 0).First(&dataMaster).Error; err != nil {
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

	data := Tbl_jenis_trans{
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
