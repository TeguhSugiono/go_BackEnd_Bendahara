package master_sett_spp

// import (
// 	"errors"
// 	"fmt"
// 	"math"
// 	"net/http"
// 	"rest_api_bendahara/helper"
// 	"rest_api_bendahara/table_data"
// 	"strconv"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/validator/v10"
// 	"gorm.io/gorm"
// )

// func ListSettPeriode(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)

// 	var master []ListData

// 	sql := " SELECT a.*,b.kd_bulan,b.tahun,b.tahun_akademik FROM tbl_sett_periode_spps a " +
// 		" INNER JOIN tbl_conf_periode_spps b on a.kd_periode_spp=b.kd_periode_spp " +
// 		" where a.flag_aktif=0 and b.flag_aktif=0 ORDER BY a.nm_kelas,b.seqno "

// 	db.Raw(sql).Scan(&master)

// 	response := helper.APIResponse("List Data ...", http.StatusOK, "success", FormatShowData(master))
// 	c.JSON(http.StatusOK, response)
// }

// func ShowSettPeriode(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)

// 	var master []ListData

// 	sql := " SELECT a.*,b.kd_bulan,b.tahun,b.tahun_akademik FROM tbl_sett_periode_spps a " +
// 		" INNER JOIN tbl_conf_periode_spps b on a.kd_periode_spp=b.kd_periode_spp " +
// 		" where a.flag_aktif=0 and b.flag_aktif=0 "

// 	if s := c.Query("search"); s != "" {
// 		if len(c.Query("search")) >= 3 {
// 			sql = fmt.Sprintf("%s and a.keterangan LIKE '%%%s%%' ", sql, s)
// 		}
// 	}

// 	if sort := c.Query("sort"); sort != "" {
// 		sql = fmt.Sprintf("%s ORDER BY a.created_on %s", sql, sort)
// 	} else {
// 		sql = fmt.Sprintf("%s ORDER BY a.created_on %s", sql, "desc")
// 	}

// 	page := c.Query("page")
// 	perPage := c.Query("perpage")

// 	intpage, err := strconv.Atoi(page)
// 	if err != nil {
// 		response := helper.APIResponse("Format Page Salah ...", http.StatusUnprocessableEntity, "error", err.Error())
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	intperPage, err := strconv.Atoi(perPage)
// 	if err != nil {
// 		response := helper.APIResponse("Format Perpage Salah ...", http.StatusUnprocessableEntity, "error", err.Error())
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	var total int64

// 	db.Raw(sql).Count(&total)

// 	sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, intperPage, (intpage-1)*intperPage)
// 	db.Raw(sql).Scan(&master)

// 	CompTableData := table_data.TableData{
// 		Total:     total,
// 		Page:      intpage,
// 		Last_page: int(math.Ceil(float64(total) / float64(intperPage))),
// 	}

// 	response := helper.APIResponseTable("List Data ...", http.StatusOK, "success", "", CompTableData, FormatShowData(master))
// 	c.JSON(http.StatusOK, response)

// }

// func InsertSettPeriode(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)

// 	var dataInput SettPeriodeInput
// 	if err := c.ShouldBindJSON(&dataInput); err != nil {
// 		var ve validator.ValidationErrors
// 		if errors.As(err, &ve) {
// 			errors := helper.FormatValidationError(err)
// 			errorMessage := gin.H{"errors": errors}
// 			response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 			c.JSON(http.StatusUnprocessableEntity, response)
// 			return
// 		}
// 		var error_binding []string
// 		error_binding = append(error_binding, err.Error())
// 		errorMessage := gin.H{"errors": error_binding}
// 		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	var CekData table_data.Tbl_conf_periode_spps
// 	if err := db.Where("kd_periode_spp = ? and flag_aktif=? ", dataInput.Kd_periode_spp, 0).First(&CekData).Error; err != nil {
// 		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
// 		response := helper.APIResponse("Simpan Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	var datenows string = time.Now().UTC().Format("2006-01-02 15:04:05")
// 	date := "2006-01-02 15:04:05"
// 	datenowx, err := time.Parse(date, datenows)
// 	if err != nil {
// 		errors := helper.FormatValidationError(err)
// 		errorMessage := gin.H{"errors": errors, "tgl": datenowx}
// 		response := helper.APIResponse("Format Tanggal Salah ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	currentUser := c.MustGet("currentUser")
// 	data := table_data.Tbl_sett_periode_spps{
// 		Kd_periode_spp: dataInput.Kd_periode_spp,
// 		Nm_kelas:       dataInput.Nm_kelas,
// 		Biaya_spp:      dataInput.Biaya_spp,
// 		Keterangan:     dataInput.Keterangan,
// 		Created_by:     currentUser.(string),
// 		Created_on:     datenowx,
// 		Flag_aktif:     0,
// 	}

// 	err = db.Omit("Edited_on", "Edited_by").Create(&data).Error
// 	if err != nil {
// 		response := helper.APIResponse("Simpan Data Gagal ...", http.StatusBadRequest, "error", err)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := helper.APIResponse("Simpan Data Sukses ...", http.StatusOK, "success", data)
// 	c.JSON(http.StatusOK, response)

// }

// func UpdateSettPeriode(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)

// 	var input SettPeriodeEdit
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		var ve validator.ValidationErrors
// 		if errors.As(err, &ve) {
// 			errors := helper.FormatValidationError(err)
// 			errorMessage := gin.H{"errors": errors}
// 			response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 			c.JSON(http.StatusUnprocessableEntity, response)
// 			return
// 		}
// 		var error_binding []string
// 		error_binding = append(error_binding, err.Error())
// 		errorMessage := gin.H{"errors": error_binding}
// 		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	var CekDataUtama table_data.Tbl_sett_periode_spps
// 	if err := db.Where("kd_sett_spp = ? and flag_aktif=? ", c.Param("kdsettspp"), 0).First(&CekDataUtama).Error; err != nil {
// 		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
// 		response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	var CekData table_data.Tbl_conf_periode_spps
// 	if err := db.Where("kd_periode_spp = ? and flag_aktif=? ", input.Kd_periode_spp, 0).First(&CekData).Error; err != nil {
// 		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
// 		response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	var datenows string = time.Now().UTC().Format("2006-01-02 15:04:05")
// 	date := "2006-01-02 15:04:05"
// 	datenowx, err := time.Parse(date, datenows)
// 	if err != nil {
// 		errors := helper.FormatValidationError(err)
// 		errorMessage := gin.H{"errors": errors, "date": datenowx}
// 		response := helper.APIResponse("Tanggal Format Salah ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	currentUser := c.MustGet("currentUser")

// 	data := table_data.Tbl_sett_periode_spps{
// 		//Kd_periode_spp: input.Kd_periode_spp,
// 		Nm_kelas:   input.Nm_kelas,
// 		Biaya_spp:  input.Biaya_spp,
// 		Keterangan: input.Keterangan,
// 		Edited_by:  currentUser.(string),
// 		Edited_on:  datenowx,
// 		Flag_aktif: 0,
// 	}

// 	err = db.Model(&CekDataUtama).Updates(data).Error
// 	if err != nil {
// 		response := helper.APIResponse("Update Data Gagal ...", http.StatusBadRequest, "error", err)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := helper.APIResponse("Update Data Sukses ...", http.StatusOK, "success", CekDataUtama)
// 	c.JSON(http.StatusOK, response)
// }

// func DeleteSettPeriode(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)

// 	var CekDataUtama table_data.Tbl_sett_periode_spps
// 	if err := db.Where("kd_sett_spp = ? and flag_aktif=? ", c.Param("kdsettspp"), 0).First(&CekDataUtama).Error; err != nil {
// 		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
// 		response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	var datenows string = time.Now().UTC().Format("2006-01-02 15:04:05")
// 	date := "2006-01-02 15:04:05"
// 	datenowx, err := time.Parse(date, datenows)
// 	if err != nil {
// 		errors := helper.FormatValidationError(err)
// 		errorMessage := gin.H{"errors": errors, "tgl": datenowx}
// 		response := helper.APIResponse("Format Tanggal Salah ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	currentUser := c.MustGet("currentUser")

// 	data := table_data.Tbl_sett_periode_spps{
// 		Edited_by:  currentUser.(string),
// 		Edited_on:  datenowx,
// 		Flag_aktif: 9,
// 	}

// 	err = db.Model(&CekDataUtama).Updates(data).Error
// 	if err != nil {
// 		response := helper.APIResponse("Delete Data Gagal ...", http.StatusBadRequest, "error", err)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := helper.APIResponse("Delete Data Sukses ...", http.StatusOK, "success", CekDataUtama)
// 	c.JSON(http.StatusOK, response)
// }
