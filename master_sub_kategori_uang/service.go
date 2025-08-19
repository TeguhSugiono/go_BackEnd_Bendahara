package master_sub_kategori_uang

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"rest_api_bendahara/helper"
	"rest_api_bendahara/table_data"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ListSubKategoriUang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var master []ListData
	//db.Find(&master)
	//db.Where("flag_aktif = ?", 0).Find(&master)
	sql := " SELECT a.*,b.nm_kategori FROM tbl_sub_kategori_uangs a " +
		" INNER JOIN tbl_kategori_uangs b on a.kd_kategori=b.kd_kategori " +
		" where a.flag_aktif=0 and b.flag_aktif=0 "

	db.Raw(sql).Scan(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", FormatShowData(master))
	c.JSON(http.StatusOK, response)
}

func ShowSubKategoriUang(c *gin.Context) {
	//pagination terdiri dari
	//1. page (halaman kebarapa)
	//2. perpage (jumlah perhalaman yang ditampilkan)
	//3. search (data yang dicari berdasarkan setiap coulumn)
	//4. ordering (sorting data berdasarkan asc dan desc)

	db := c.MustGet("db").(*gorm.DB)

	var master []ListData

	sql := " SELECT a.*,b.nm_kategori FROM tbl_sub_kategori_uangs a " +
		" INNER JOIN tbl_kategori_uangs b on a.kd_kategori=b.kd_kategori " +
		" where a.flag_aktif=0 and b.flag_aktif=0 "

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
			sql = fmt.Sprintf("%s and a.nm_sub_kategori LIKE '%%%s%%' ", sql, s)
		}

	}

	if filter := c.Query("filter"); filter != "" {
		sql = fmt.Sprintf("%s and b.nm_kategori = '%s'", sql, filter)
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

	// return c.JSON(http.StatusOK.Map{
	// 	"data":      products,
	// 	"total":     total,
	// 	"page":      page,
	// 	"last_page": math.Ceil(float64(total / int64(perPage))),
	// })
}

func InsertSubKategoriUang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var dataInput SKategoriUangInput
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

	var dataMaster table_data.Tbl_kategori_uangs
	if err := db.Where("kd_kategori = ? and flag_aktif=? ", dataInput.Kd_kategori, 0).First(&dataMaster).Error; err != nil {
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
	data := table_data.Tbl_sub_kategori_uangs{
		Kd_kategori:     dataInput.Kd_kategori,
		Nm_sub_kategori: dataInput.Nm_sub_kategori,
		Created_by:      currentUser.(string),
		Created_on:      datenowx,
		Flag_aktif:      0,
	}

	err = db.Omit("Edited_on", "Edited_by").Create(&data).Error
	if err != nil {
		response := helper.APIResponse("Simpan Data Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//=========================================

	var nm_sub_kategori string
	var str_nm_sub_kategori string

	str_nm_sub_kategori = ""
	nomor := 0
	rowss, _ := db.Raw("select nm_sub_kategori from tbl_sub_kategori_uangs where flag_aktif = ? and kd_kategori=?", 0, dataInput.Kd_kategori).Rows()
	defer rowss.Close()
	for rowss.Next() {
		rowss.Scan(&nm_sub_kategori)
		if nomor == 0 {
			str_nm_sub_kategori = nm_sub_kategori
		} else {
			str_nm_sub_kategori = str_nm_sub_kategori + "," + nm_sub_kategori
		}
		nomor++
	}

	err = db.Raw("UPDATE tbl_kategori_uangs SET nm_detail = ? WHERE kd_kategori = ? ", str_nm_sub_kategori, dataInput.Kd_kategori).Scan(&dataMaster).Error
	if err != nil {
		response := helper.APIResponse("Update Data Ke Tbl_group_kategoris Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//=========================================

	response := helper.APIResponse("Simpan Data Sukses ...", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func UpdateSubKategoriUang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input SKategoriUangInput
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

	var dataUtama table_data.Tbl_sub_kategori_uangs
	if err := db.Where("kd_sub_kategori = ? and flag_aktif=? ", c.Param("kdsubkategori"), 0).First(&dataUtama).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
		response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var dataMaster table_data.Tbl_kategori_uangs
	if err := db.Where("kd_kategori = ? and flag_aktif=? ", input.Kd_kategori, 0).First(&dataMaster).Error; err != nil {
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

	data := table_data.Tbl_sub_kategori_uangs{
		Kd_kategori:     input.Kd_kategori,
		Nm_sub_kategori: input.Nm_sub_kategori,
		Edited_by:       currentUser.(string),
		Edited_on:       datenowx,
		Flag_aktif:      0,
	}

	err = db.Model(&dataUtama).Updates(data).Error
	if err != nil {
		response := helper.APIResponse("Update Data Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//=========================================

	var nm_sub_kategori string
	var str_nm_sub_kategori string

	str_nm_sub_kategori = ""
	nomor := 0
	rowss, _ := db.Raw("select nm_sub_kategori from tbl_sub_kategori_uangs where flag_aktif = ? and kd_kategori=?", 0, input.Kd_kategori).Rows()
	defer rowss.Close()
	for rowss.Next() {
		rowss.Scan(&nm_sub_kategori)
		if nomor == 0 {
			str_nm_sub_kategori = nm_sub_kategori
		} else {
			str_nm_sub_kategori = str_nm_sub_kategori + "," + nm_sub_kategori
		}
		nomor++
	}

	err = db.Raw("UPDATE tbl_kategori_uangs SET nm_detail = ? WHERE kd_kategori = ? ", str_nm_sub_kategori, input.Kd_kategori).Scan(&dataMaster).Error
	if err != nil {
		response := helper.APIResponse("Update Data Ke Tbl_group_kategoris Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//=========================================

	response := helper.APIResponse("Update Data Sukses ...", http.StatusOK, "success", dataUtama)
	c.JSON(http.StatusOK, response)

}

func DeleteSubKategoriUang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var dataMaster table_data.Tbl_sub_kategori_uangs
	if err := db.Where("kd_sub_kategori = ? and flag_aktif=?", c.Param("kdsubkategori"), 0).First(&dataMaster).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
		response := helper.APIResponse("Delete Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
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

	data := table_data.Tbl_sub_kategori_uangs{
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

	//=========================================

	var nm_sub_kategori string
	var str_nm_sub_kategori string

	str_nm_sub_kategori = ""
	nomor := 0
	rowss, _ := db.Raw("select nm_sub_kategori from tbl_sub_kategori_uangs where flag_aktif = ? and kd_kategori=?", 0, dataMaster.Kd_kategori).Rows()
	defer rowss.Close()
	for rowss.Next() {
		rowss.Scan(&nm_sub_kategori)
		if nomor == 0 {
			str_nm_sub_kategori = nm_sub_kategori
		} else {
			str_nm_sub_kategori = str_nm_sub_kategori + "," + nm_sub_kategori
		}
		nomor++
	}

	var dataMasterx table_data.Tbl_kategori_uangs
	err = db.Raw("UPDATE tbl_kategori_uangs SET nm_detail = ? WHERE kd_kategori = ? ", str_nm_sub_kategori, dataMaster.Kd_kategori).Scan(&dataMasterx).Error
	if err != nil {
		response := helper.APIResponse("Update Data Ke Tbl_group_kategoris Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//=========================================

	response := helper.APIResponse("Delete Data Sukses ...", http.StatusOK, "success", dataMaster)
	c.JSON(http.StatusOK, response)
}
