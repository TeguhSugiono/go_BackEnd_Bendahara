package master_sett_periode

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"rest_api_bendahara/connection"
	"rest_api_bendahara/helper"
	"rest_api_bendahara/table_data"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ListConfPeriode(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var master []ListData

	sql := " SELECT tahun_akademik,nm_kelas,biaya_spp,nm_sett from tbl_conf_periode_spps where flag_aktif = 0 GROUP BY tahun_akademik,nm_kelas order by tahun_akademik,nm_kelas "

	db.Raw(sql).Scan(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", FormatShowData(master))
	c.JSON(http.StatusOK, response)
}

func ShowConfPeriode(c *gin.Context) {

	//pagination terdiri dari
	//1. page (halaman kebarapa)
	//2. perpage (jumlah perhalaman yang ditampilkan)
	//3. search (data yang dicari berdasarkan setiap coulumn)
	//4. ordering (sorting data berdasarkan asc dan desc)

	db := c.MustGet("db").(*gorm.DB)

	var master []ListData

	sql := " SELECT tahun_akademik,nm_kelas,biaya_spp,nm_sett from tbl_conf_periode_spps  " +
		" where flag_aktif=0 "

	if s := c.Query("search"); s != "" {
		if len(c.Query("search")) >= 3 {
			//sql = fmt.Sprintf("%s and nm_sett LIKE '%%%s%%' ", sql, s)
			sql = fmt.Sprintf("%s and (kd_bulan LIKE '%%%s%%' or tahun LIKE '%%%s%%' "+
				" or nm_sett LIKE '%%%s%%' or tahun_akademik LIKE '%%%s%%' or nm_kelas LIKE '%%%s%%' "+
				" or biaya_spp LIKE '%%%s%%') ", sql, s, s, s, s, s, s)
		}
	}

	sql = fmt.Sprintf("%s GROUP BY tahun_akademik,nm_kelas ", sql)

	if sort := c.Query("sort"); sort != "" {
		sql = fmt.Sprintf("%s ORDER BY tahun_akademik %s,nm_kelas %s, seqno %s", sql, "asc", "asc", "asc")
	} else {
		sql = fmt.Sprintf("%s ORDER BY tahun_akademik %s,nm_kelas %s,seqno %s", sql, "desc", "desc", "desc")
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

func InsertConfPeriode(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	dbSIA := connection.SetupConnectionSIA()

	var dataInput InputTahunAkademik
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

	var conf_periode_spps table_data.Tbl_conf_periode_spps
	checkUser := db.Select("*").Where("flag_aktif = ? and tahun_akademik = ? and nm_kelas=?", 0, dataInput.Tahun_akademik, dataInput.Nm_kelas).Find(&conf_periode_spps)
	if checkUser.RowsAffected > 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Data Setting Tahun Periode atau Tahun Akademik Sudah Ada ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var tahun_akademik_s string
	row := dbSIA.Table("tbl_tahun_akademik").Where("flag_tahun = ? and tahun_akademik=?", 0, dataInput.Tahun_akademik).Select("tahun_akademik").Row()
	row.Scan(&tahun_akademik_s)
	if tahun_akademik_s == "" {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Tahun Akademik Tidak Ada ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var nm_kelas_s string = ""

	sqlkelas := " SELECT nm_kelas FROM tbl_kelas "
	sqlkelas = fmt.Sprintf("%s where flag_kelas = 0  and nm_kelas = '%s'", sqlkelas, dataInput.Nm_kelas)
	//sqlkelas = fmt.Sprintf("%s GROUP BY REPLACE(REPLACE(nm_kelas,'MIA',''),'IIS','') ", sqlkelas)

	rowskelas, _ := dbSIA.Raw(sqlkelas).Rows()
	defer rowskelas.Close()
	for rowskelas.Next() {
		rowskelas.Scan(&nm_kelas_s)
	}

	if nm_kelas_s == "" {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Tingkat Kelas (Nama Kelas) Tidak Ada ...", http.StatusUnprocessableEntity, "error", errorMessage)
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

	tahunakademik := dataInput.Tahun_akademik
	ReplaceExpressionSearch := regexp.MustCompile(`[-/]`)
	StringSlice := ReplaceExpressionSearch.Split(tahunakademik, -1)
	var Settname string = "Tahun Akademik " + StringSlice[0] + " s/d " + StringSlice[1] + " "

	var IntTahunFirst int
	var IntTahunSecond int
	IntTahunFirst, _ = strconv.Atoi(StringSlice[0])
	IntTahunSecond, _ = strconv.Atoi(StringSlice[1])

	var namabulan = [12]string{"07", "08", "09", "10", "11", "12", "01", "02", "03", "04", "05", "06"}

	//SetArrayData := []table_data.Tbl_conf_periode_spps{}
	for seqno := 0; seqno < 12; seqno++ {
		arraydata := table_data.Tbl_conf_periode_spps{}
		// arraydata.Kd_periode_spp = Int_kd_periode_spp
		arraydata.Seqno = seqno + 1
		arraydata.Kd_bulan = namabulan[seqno]

		if namabulan[seqno] == "01" || namabulan[seqno] == "02" || namabulan[seqno] == "03" || namabulan[seqno] == "04" || namabulan[seqno] == "05" || namabulan[seqno] == "06" {
			arraydata.Tahun = IntTahunSecond
		} else {
			arraydata.Tahun = IntTahunFirst
		}

		arraydata.Nm_sett = Settname
		arraydata.Tahun_akademik = dataInput.Tahun_akademik
		arraydata.Created_by = currentUser.(string)
		arraydata.Created_on = datenowx
		arraydata.Flag_aktif = 0
		arraydata.Nm_kelas = dataInput.Nm_kelas
		arraydata.Biaya_spp = dataInput.Biaya_spp
		//SetArrayData = append(SetArrayData, arraydata)

		err = db.Omit("Edited_on", "Edited_by").Create(&arraydata).Error
		if err != nil {
			response := helper.APIResponse("Simpan Data Gagal ...", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	var nm_sett string
	var tahun_akademik string

	SetArrayDataA := []Return_Up{}
	rows, _ := db.Raw("select nm_sett,tahun_akademik from tbl_conf_periode_spps where flag_aktif=0 and tahun_akademik=? GROUP BY tahun_akademik", dataInput.Tahun_akademik).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&nm_sett, &tahun_akademik)
		arraydata := Return_Up{}
		arraydata.Nm_Sett = nm_sett
		arraydata.Tahun_akademik = tahun_akademik

		var id_conf int
		var seqno int
		var kd_bulan string
		var tahun int
		var nm_kelas string
		var biaya_spp float64

		SetArrayDataDetail := []Return_Down{}
		rowsdet, _ := db.Raw("SELECT id_conf,seqno,kd_bulan,tahun,nm_kelas,biaya_spp "+
			" from tbl_conf_periode_spps where flag_aktif=0 and tahun_akademik=? order by seqno", dataInput.Tahun_akademik).Rows()
		defer rowsdet.Close()
		for rowsdet.Next() {
			arraydatadetail := Return_Down{}
			rowsdet.Scan(&id_conf, &seqno, &kd_bulan, &tahun, &nm_kelas, &biaya_spp)
			arraydatadetail.Id_conf = id_conf
			arraydatadetail.Seqno = seqno
			arraydatadetail.Kd_bulan = kd_bulan
			arraydatadetail.Tahun = tahun
			arraydatadetail.Nm_kelas = nm_kelas
			arraydatadetail.Biaya_spp = biaya_spp
			SetArrayDataDetail = append(SetArrayDataDetail, arraydatadetail)
		}

		arraydata.Detail = SetArrayDataDetail

		SetArrayDataA = append(SetArrayDataA, arraydata)
	}

	response := helper.APIResponse("Simpan Data Sukses ...", http.StatusOK, "success", SetArrayDataA)
	c.JSON(http.StatusOK, response)
}

type Return_Up struct {
	Nm_Sett        string      `json:"nm_sett"`
	Tahun_akademik string      `json:"tahun_akademik"`
	Detail         interface{} `json:"detail"`
}

type Return_Down struct {
	Id_conf   int     `json:"id_conf"`
	Seqno     int     `json:"seqno"`
	Kd_bulan  string  `json:"kd_bulan"`
	Tahun     int     `json:"tahun"`
	Nm_kelas  string  `json:"nm_kelas"`
	Biaya_spp float64 `json:"biaya_spp"`
}

// func UpdateConfPeriode(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)

// 	var dataMaster table_data.Tbl_conf_periode_spps
// 	if err := db.Where("id_conf = ? and flag_aktif=? ", c.Param("idconf"), 0).First(&dataMaster).Error; err != nil {
// 		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
// 		response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	var dataInput EditTahunAkademik
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

// 	data := table_data.Tbl_conf_periode_spps{
// 		Biaya_spp:  dataInput.Biaya_spp,
// 		Edited_by:  currentUser.(string),
// 		Edited_on:  datenowx,
// 		Flag_aktif: 0,
// 	}

// 	err = db.Model(&dataMaster).Updates(data).Error
// 	if err != nil {
// 		response := helper.APIResponse("Update Data Gagal ...", http.StatusBadRequest, "error", err)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := helper.APIResponse("Update Data Sukses ...", http.StatusOK, "success", dataMaster)
// 	c.JSON(http.StatusOK, response)

// }

func UpdateConfPeriodeAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	dbSIA := connection.SetupConnectionSIA()

	var dataInput table_data.Tbl_conf_periode_spps
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

	var dataMaster table_data.Tbl_conf_periode_spps
	if err := db.Where("tahun_akademik = ? and flag_aktif=? and nm_kelas=?", dataInput.Tahun_akademik, 0, dataInput.Nm_kelas).First(&dataMaster).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
		response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var tahun_akademik_s string
	row := dbSIA.Table("tbl_tahun_akademik").Where("flag_tahun = ? and tahun_akademik=?", 0, dataInput.Tahun_akademik).Select("tahun_akademik").Row()
	row.Scan(&tahun_akademik_s)
	if tahun_akademik_s == "" {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Tahun Akademik Tidak Ada ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var nm_kelas_s string = ""

	sqlkelas := " SELECT nm_kelas FROM tbl_kelas "
	sqlkelas = fmt.Sprintf("%s where flag_kelas = 0  and nm_kelas = '%s'", sqlkelas, dataInput.Nm_kelas)

	rowskelas, _ := dbSIA.Raw(sqlkelas).Rows()
	defer rowskelas.Close()
	for rowskelas.Next() {
		rowskelas.Scan(&nm_kelas_s)
	}

	if nm_kelas_s == "" {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Tingkat Kelas (Nama Kelas) Tidak Ada ...", http.StatusUnprocessableEntity, "error", errorMessage)
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

	var tblupdate table_data.Tbl_conf_periode_spps
	err = db.Raw("update tbl_conf_periode_spps SET edited_by  = ? , edited_on = ?  , biaya_spp = ? "+
		" WHERE tahun_akademik = ? and nm_kelas=?", currentUser.(string), datenowx, dataInput.Biaya_spp, dataInput.Tahun_akademik, dataInput.Nm_kelas).Scan(&tblupdate).Error
	if err != nil {
		response := helper.APIResponse("Delete Data Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var nm_sett string
	var tahun_akademik string

	SetArrayDataA := []Return_Up{}
	rows, _ := db.Raw("select nm_sett,tahun_akademik from tbl_conf_periode_spps where flag_aktif=0 and tahun_akademik=? GROUP BY tahun_akademik", dataInput.Tahun_akademik).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&nm_sett, &tahun_akademik)
		arraydata := Return_Up{}
		arraydata.Nm_Sett = nm_sett
		arraydata.Tahun_akademik = tahun_akademik

		var id_conf int
		var seqno int
		var kd_bulan string
		var tahun int
		var nm_kelas string
		var biaya_spp float64

		SetArrayDataDetail := []Return_Down{}
		rowsdet, _ := db.Raw("SELECT id_conf,seqno,kd_bulan,tahun,nm_kelas,biaya_spp "+
			" from tbl_conf_periode_spps where flag_aktif=0 and tahun_akademik=? order by seqno", dataInput.Tahun_akademik).Rows()
		defer rowsdet.Close()
		for rowsdet.Next() {
			arraydatadetail := Return_Down{}
			rowsdet.Scan(&id_conf, &seqno, &kd_bulan, &tahun, &nm_kelas, &biaya_spp)
			arraydatadetail.Id_conf = id_conf
			arraydatadetail.Seqno = seqno
			arraydatadetail.Kd_bulan = kd_bulan
			arraydatadetail.Tahun = tahun
			arraydatadetail.Nm_kelas = nm_kelas
			arraydatadetail.Biaya_spp = biaya_spp
			SetArrayDataDetail = append(SetArrayDataDetail, arraydatadetail)
		}

		arraydata.Detail = SetArrayDataDetail

		SetArrayDataA = append(SetArrayDataA, arraydata)
	}

	response := helper.APIResponse("Update Data Sukses ...", http.StatusOK, "success", SetArrayDataA)
	c.JSON(http.StatusOK, response)

	// response := helper.APIResponse("Update Data Sukses ...", http.StatusOK, "success", tblupdate)
	// c.JSON(http.StatusOK, response)

}

func DeleteConfPeriode(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var dataInput DeleteTahunAkademik
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var conf_periode_spps table_data.Tbl_conf_periode_spps
	checkUser := db.Select("*").Where("flag_aktif = ? and tahun_akademik = ? ", 0, dataInput.Tahun_akademik).Find(&conf_periode_spps)
	if checkUser.RowsAffected == 0 {
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
		errorMessage := gin.H{"errors": errors, "date": datenowx}
		response := helper.APIResponse("Tanggal Format Salah ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser")

	var tblupdate table_data.Tbl_conf_periode_spps
	err = db.Raw("update tbl_conf_periode_spps SET Flag_aktif = ?, edited_by  = ? , edited_on = ?  WHERE tahun_akademik = ? and nm_kelas=?", 9, currentUser.(string), datenowx, dataInput.Tahun_akademik, dataInput.Nm_kelas).Scan(&tblupdate).Error
	if err != nil {
		response := helper.APIResponse("Delete Data Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	SetArrayData := []ResultDelete{}
	var seqno int
	var kd_bulan string
	rows, _ := db.Raw("SELECT seqno,kd_bulan FROM tbl_conf_periode_spps where flag_aktif=? and edited_on=? and edited_by=? order by seqno", 9, datenowx, currentUser.(string)).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&seqno, &kd_bulan)
		arraydata := ResultDelete{}
		arraydata.Seqno = seqno
		arraydata.Kd_bulan = kd_bulan

		SetArrayData = append(SetArrayData, arraydata)
	}

	response := APIResponseDelete("Delete Data Sukses ...", http.StatusOK, "success", conf_periode_spps.Tahun, conf_periode_spps.Tahun_akademik, SetArrayData)
	c.JSON(http.StatusOK, response)

}

type ResultDelete struct {
	Seqno    int
	Kd_bulan string
}

type ResponseDelete struct {
	Meta MetaDelete `json:"meta"`
	Data DataDelete `json:"data"`
}

type MetaDelete struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type DataDelete struct {
	//Kd_periode_spp int         `json:"kd_periode_spp"`
	Tahun          int         `json:"tahun"`
	Tahun_akademik string      `json:"tahun_akademik"`
	Detail         interface{} `json:"detail"`
}

func APIResponseDelete(message string, code int, status string, tahun int, tahun_akademik string, datax interface{}) ResponseDelete {
	meta := MetaDelete{
		Message: message,
		Code:    code,
		Status:  status,
	}

	data := DataDelete{
		// Kd_periode_spp: kd_periode_spp,
		Tahun:          tahun,
		Tahun_akademik: tahun_akademik,
		Detail:         datax,
	}

	jsonResponse := ResponseDelete{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}
