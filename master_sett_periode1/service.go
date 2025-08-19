package master_sett_periode

// import (
// 	"fmt"
// 	"math"
// 	"net/http"
// 	"regexp"
// 	"rest_api_bendahara/helper"
// 	"rest_api_bendahara/table_data"
// 	"strconv"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// func ListConfPeriode(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)

// 	var master []ListData

// 	sql := " SELECT * from tbl_conf_periode_spps where flag_aktif = 0 group by kd_periode_spp order by tahun_akademik "

// 	db.Raw(sql).Scan(&master)

// 	response := helper.APIResponse("List Data ...", http.StatusOK, "success", FormatShowData(master))
// 	c.JSON(http.StatusOK, response)
// }

// func FormatShowData(table []ListData) []ReturnData {
// 	arraydata := []ReturnData{}
// 	for _, resultdata := range table {
// 		arraytemporary := FormatTampungData(resultdata)
// 		arraydata = append(arraydata, arraytemporary)
// 	}

// 	return arraydata
// }

// type ReturnData struct {
// 	Kd_periode_spp int         `json:"kd_periode_spp"`
// 	Tahun          int         `json:"tahun"`
// 	Tahun_akademik string      `json:"tahun_akademik"`
// 	Detail         interface{} `json:"detail"`
// }

// type ReturnDataSub struct {
// 	Id_conf  int    `json:"id_conf"`
// 	Seqno    int    `json:"seqno"`
// 	Kd_bulan string `json:"kd_bulan"`
// }

// type ListData struct {
// 	Kd_periode_spp int
// 	Tahun          int
// 	Tahun_akademik string
// }

// func FormatTampungData(table ListData) ReturnData {
// 	arraydata := ReturnData{}
// 	arraydata.Kd_periode_spp = table.Kd_periode_spp
// 	arraydata.Tahun = table.Tahun
// 	arraydata.Tahun_akademik = table.Tahun_akademik
// 	arraydata.Detail = nil

// 	return arraydata
// }

// func ShowConfPeriode(c *gin.Context) {

// 	//pagination terdiri dari
// 	//1. page (halaman kebarapa)
// 	//2. perpage (jumlah perhalaman yang ditampilkan)
// 	//3. search (data yang dicari berdasarkan setiap coulumn)
// 	//4. ordering (sorting data berdasarkan asc dan desc)

// 	db := c.MustGet("db").(*gorm.DB)

// 	var master []ListData

// 	sql := " SELECT * from tbl_conf_periode_spps  " +
// 		" where flag_aktif=0 "

// 	if s := c.Query("search"); s != "" {
// 		if len(c.Query("search")) >= 3 {
// 			//sql = fmt.Sprintf("%s and nm_sett LIKE '%%%s%%' ", sql, s)
// 			sql = fmt.Sprintf("%s and (kd_bulan LIKE '%%%s%%' or tahun LIKE '%%%s%%' or nm_sett LIKE '%%%s%%'  or tahun_akademik LIKE '%%%s%%') ", sql, s, s, s, s)
// 		}
// 	}

// 	if sort := c.Query("sort"); sort != "" {
// 		sql = fmt.Sprintf("%s ORDER BY tahun_akademik %s,seqno %s", sql, "asc", "asc")
// 	} else {
// 		sql = fmt.Sprintf("%s ORDER BY tahun_akademik %s,seqno %s", sql, "desc", "desc")
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

// 	response := APIResponseTableN("List Data ...", http.StatusOK, "success", sql, CompTableData, FormatShowData(master))
// 	c.JSON(http.StatusOK, response)

// }

// type DataFirst struct {
// 	Kd_periode_spp int
// 	Nm_sett        string
// 	Tahun_akademik string
// }

// func InsertConfPeriode(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)

// 	var dataInput InputTahunAkademik
// 	if err := c.ShouldBindJSON(&dataInput); err != nil {
// 		errors := helper.FormatValidationError(err)
// 		errorMessage := gin.H{"errors": errors}
// 		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	var conf_periode_spps table_data.Tbl_conf_periode_spps
// 	checkUser := db.Select("*").Where("flag_aktif = ? and tahun_akademik = ? ", 0, dataInput.Tahun_akademik).Find(&conf_periode_spps)
// 	if checkUser.RowsAffected > 0 {
// 		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
// 		response := helper.APIResponse("Data Setting Tahun Periode atau Tahun Akademik Sudah Ada ...", http.StatusUnprocessableEntity, "error", errorMessage)
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

// 	var kd_periode_spp string
// 	var Int_kd_periode_spp int

// 	rows, _ := db.Raw("SELECT ifnull(MAX(kd_periode_spp),0) + 1 as 'kd_periode_spp' FROM tbl_conf_periode_spps where flag_aktif = ?", 0).Rows()
// 	defer rows.Close()
// 	for rows.Next() {
// 		rows.Scan(&kd_periode_spp)
// 		Int_kd_periode_spp, _ = strconv.Atoi(kd_periode_spp)
// 	}

// 	currentUser := c.MustGet("currentUser")

// 	tahunakademik := dataInput.Tahun_akademik
// 	ReplaceExpressionSearch := regexp.MustCompile(`[-/]`)
// 	StringSlice := ReplaceExpressionSearch.Split(tahunakademik, -1)
// 	var Settname string = "Tahun Akademik " + StringSlice[0] + " s/d " + StringSlice[1] + " "

// 	var IntTahunFirst int
// 	var IntTahunSecond int
// 	IntTahunFirst, _ = strconv.Atoi(StringSlice[0])
// 	IntTahunSecond, _ = strconv.Atoi(StringSlice[1])

// 	var namabulan = [12]string{"07", "08", "09", "10", "11", "12", "01", "02", "03", "04", "05", "06"}

// 	SetArrayData := []table_data.Tbl_conf_periode_spps{}
// 	for seqno := 0; seqno < 12; seqno++ {
// 		arraydata := table_data.Tbl_conf_periode_spps{}
// 		arraydata.Kd_periode_spp = Int_kd_periode_spp
// 		arraydata.Seqno = seqno + 1
// 		arraydata.Kd_bulan = namabulan[seqno]

// 		if namabulan[seqno] == "01" || namabulan[seqno] == "02" || namabulan[seqno] == "03" || namabulan[seqno] == "04" || namabulan[seqno] == "05" || namabulan[seqno] == "06" {
// 			arraydata.Tahun = IntTahunSecond
// 		} else {
// 			arraydata.Tahun = IntTahunFirst
// 		}

// 		arraydata.Nm_sett = Settname
// 		arraydata.Tahun_akademik = dataInput.Tahun_akademik
// 		arraydata.Created_by = currentUser.(string)
// 		arraydata.Created_on = datenowx
// 		arraydata.Flag_aktif = 0
// 		SetArrayData = append(SetArrayData, arraydata)

// 		err = db.Omit("Edited_on", "Edited_by").Create(&arraydata).Error
// 		if err != nil {
// 			response := helper.APIResponse("Simpan Data Gagal ...", http.StatusBadRequest, "error", err)
// 			c.JSON(http.StatusBadRequest, response)
// 			return
// 		}
// 	}

// 	response := helper.APIResponse("Simpan Data Sukses ...", http.StatusOK, "success", SetArrayData)
// 	c.JSON(http.StatusOK, response)
// }

// func DeleteConfPeriode(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)

// 	var dataInput InputTahunAkademik
// 	if err := c.ShouldBindJSON(&dataInput); err != nil {
// 		errors := helper.FormatValidationError(err)
// 		errorMessage := gin.H{"errors": errors}
// 		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	var conf_periode_spps table_data.Tbl_conf_periode_spps
// 	checkUser := db.Select("*").Where("flag_aktif = ? and tahun_akademik = ? ", 0, dataInput.Tahun_akademik).Find(&conf_periode_spps)
// 	if checkUser.RowsAffected == 0 {
// 		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
// 		response := helper.APIResponse("Delete Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	var result CekDataSettPeriode
// 	db.Raw("SELECT count(*) as 'jumlah' FROM tbl_conf_periode_spps a "+
// 		" INNER JOIN tbl_sett_periode_spps b on a.kd_periode_spp = b.kd_periode_spp "+
// 		" where a.flag_aktif=0 and b.flag_aktif=0 "+
// 		" and tahun_akademik = ? ", dataInput.Tahun_akademik).Scan(&result)

// 	if result.Jumlah > 0 {
// 		errorMessage := gin.H{"errors": "Data Configurasi Periode Sudah Terpakai Di Master Setting SPP ..."}
// 		response := helper.APIResponse("Delete Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
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

// 	var tblupdate table_data.Tbl_conf_periode_spps
// 	err = db.Raw("update tbl_conf_periode_spps SET Flag_aktif = ?, edited_by  = ? , edited_on = ?  WHERE tahun_akademik = ? ", 9, currentUser.(string), datenowx, dataInput.Tahun_akademik).Scan(&tblupdate).Error
// 	if err != nil {
// 		response := helper.APIResponse("Delete Data Gagal ...", http.StatusBadRequest, "error", err)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	SetArrayData := []ResultDelete{}
// 	var seqno int
// 	var kd_bulan string
// 	rows, _ := db.Raw("SELECT seqno,kd_bulan FROM tbl_conf_periode_spps where flag_aktif=? and edited_on=? and edited_by=? order by seqno", 9, datenowx, currentUser.(string)).Rows()
// 	defer rows.Close()
// 	for rows.Next() {
// 		rows.Scan(&seqno, &kd_bulan)
// 		arraydata := ResultDelete{}
// 		arraydata.Seqno = seqno
// 		arraydata.Kd_bulan = kd_bulan

// 		SetArrayData = append(SetArrayData, arraydata)
// 	}

// 	response := APIResponseN("Delete Data Sukses ...", http.StatusOK, "success", conf_periode_spps.Kd_periode_spp, conf_periode_spps.Tahun, conf_periode_spps.Tahun_akademik, SetArrayData)
// 	c.JSON(http.StatusOK, response)

// }

// type ResultDelete struct {
// 	Seqno    int
// 	Kd_bulan string
// }

// type ResponseN struct {
// 	Meta MetaN `json:"meta"`
// 	Data DataN `json:"data"`
// }

// type MetaN struct {
// 	Message string `json:"message"`
// 	Code    int    `json:"code"`
// 	Status  string `json:"status"`
// }

// type DataN struct {
// 	Kd_periode_spp int         `json:"kd_periode_spp"`
// 	Tahun          int         `json:"tahun"`
// 	Tahun_akademik string      `json:"tahun_akademik"`
// 	Detail         interface{} `json:"detail"`
// }

// func APIResponseN(message string, code int, status string, kd_periode_spp int, tahun int, tahun_akademik string, datax interface{}) ResponseN {
// 	meta := MetaN{
// 		Message: message,
// 		Code:    code,
// 		Status:  status,
// 	}

// 	data := DataN{
// 		Kd_periode_spp: kd_periode_spp,
// 		Tahun:          tahun,
// 		Tahun_akademik: tahun_akademik,
// 		Detail:         datax,
// 	}

// 	jsonResponse := ResponseN{
// 		Meta: meta,
// 		Data: data,
// 	}

// 	return jsonResponse
// }

// type ResponseTableN struct {
// 	Meta      MetaTableN  `json:"meta"`
// 	DataTable interface{} `json:"datatable"`
// 	Data      interface{} `json:"data"`
// }

// type MetaTableN struct {
// 	Message string `json:"message"`
// 	Code    int    `json:"code"`
// 	Status  string `json:"status"`
// 	Query   string `json:"query"`
// }

// func APIResponseTableN(message string, code int, status string, query string, datatable interface{}, data interface{}) ResponseTableN {
// 	meta := MetaTableN{
// 		Message: message,
// 		Code:    code,
// 		Status:  status,
// 		Query:   query,
// 	}

// 	jsonResponse := ResponseTableN{
// 		Meta:      meta,
// 		DataTable: datatable,
// 		Data:      data,
// 	}

// 	return jsonResponse
// }
