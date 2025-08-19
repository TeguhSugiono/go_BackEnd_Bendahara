package transaksi_uang_keluar_pra_act

import (
	"errors"
	"fmt"
	"net/http"
	"rest_api_bendahara/helper"
	"rest_api_bendahara/master_group_kategori"
	"rest_api_bendahara/table_data"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func PostUangMasuk(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//kd_jenis=2 adalah uang keluar
	var master []master_group_kategori.ListData
	sql := "  SELECT a.*,b.proses_uang FROM tbl_group_kategoris as a " +
		" inner join tbl_jenis_trans as b on a.kd_jenis=b.kd_jenis  " +
		" where a.flag_aktif=0 and b.flag_aktif=0  " +
		" and a.kd_jenis in('1') order by b.proses_uang "

	db.Raw(sql).Scan(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", master_group_kategori.FormatGroupKategori(master))
	c.JSON(http.StatusOK, response)
}

func ListDokument(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var ListDok []ListDokumentPRA
	sql := "  SELECT a.kd_trans_keluar,a.no_document " +
		" from tbl_trans_uang_keluar_pra_headers a " +
		" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar = b.kd_trans_keluar " +
		" where a.flag_aktif=0 and a.flag_act=0 and b.flag_aktif=0  group by a.kd_trans_keluar "

	db.Raw(sql).Scan(&ListDok)

	response := helper.APIResponse("List Dokument ...", http.StatusOK, "success", ListDok)
	c.JSON(http.StatusOK, response)
}

func ListGroupKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//kd_jenis=1 adalah uang masuk
	var master []master_group_kategori.ListData
	sql := "  SELECT a.*,b.proses_uang FROM tbl_group_kategoris as a " +
		" inner join tbl_jenis_trans as b on a.kd_jenis=b.kd_jenis  " +
		" where a.flag_aktif=0 and b.flag_aktif=0  " +
		" and a.kd_jenis in('1') order by b.proses_uang "

	db.Raw(sql).Scan(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", master_group_kategori.FormatGroupKategori(master))
	c.JSON(http.StatusOK, response)
}

func CreateUangKeluar(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramCreateACT ParamCreateACT
	if err := c.ShouldBindJSON(&paramCreateACT); err != nil {
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

	var intJmldata int
	db.Raw(" SELECT count(*) jmldata FROM tbl_trans_uang_keluar_pra_headers a "+
		" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar=b.kd_trans_keluar "+
		" where a.flag_aktif=0 and a.flag_act=0 and b.flag_aktif=0 and a.no_document=? and a.kd_trans_keluar=? limit 1 ", paramCreateACT.No_document, paramCreateACT.Kd_trans_keluar).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Data Perencanaan Tidak Ditemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
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

	var intKd_trans_keluar int
	db.Raw("SELECT ifnull(max(kd_trans_keluar),0) + 1 as 'run_number' FROM tbl_trans_uang_keluar_pra_act_headers ").Scan(&intKd_trans_keluar)

	var kd_group int
	var kd_kategori int
	var no_document string
	var tgl_document string
	var total_biaya float64
	var total_bayar float64
	var sisa_biaya float64
	var keterangan string

	ssql := " select kd_group,kd_kategori,no_document,date_format(tgl_document,'%Y-%m-%d'),total_biaya,total_bayar,sisa_biaya,keterangan from tbl_trans_uang_keluar_pra_headers where flag_aktif=0   "
	ssql = fmt.Sprintf("%s and kd_trans_keluar = %d", ssql, paramCreateACT.Kd_trans_keluar)
	ssql = fmt.Sprintf("%s and no_document = '%s'", ssql, paramCreateACT.No_document)

	rows, _ := db.Raw(ssql).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&kd_group, &kd_kategori, &no_document, &tgl_document, &total_biaya, &total_bayar, &sisa_biaya, &keterangan)
	}

	currentUser := c.MustGet("currentUser")
	data := table_data.Tbl_trans_uang_keluar_pra_act_headers{
		Kd_group:        kd_group,
		Kd_kategori:     kd_kategori,
		Kd_trans_keluar: intKd_trans_keluar,
		No_document:     no_document,
		Tgl_document:    tgl_document,
		Total_biaya:     total_biaya,
		Total_bayar:     total_bayar,
		Sisa_biaya:      sisa_biaya,
		Keterangan:      keterangan,
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

	var jml_bayar float64
	var intKd_trans_keluar_detail int
	//var int_seqno int

	int_seqno := 1

	ssqldetail := " select jml_bayar,keterangan from tbl_trans_uang_keluar_pra_details where flag_aktif=0 "
	ssqldetail = fmt.Sprintf("%s and kd_trans_keluar = %d", ssqldetail, paramCreateACT.Kd_trans_keluar)
	ssqldetail = fmt.Sprintf("%s ORDER BY seqno %s", ssqldetail, "asc")

	rowsdetail, _ := db.Raw(ssqldetail).Rows()
	defer rowsdetail.Close()
	for rowsdetail.Next() {
		rowsdetail.Scan(&jml_bayar, &keterangan)

		db.Raw("SELECT ifnull(max(kd_trans_keluar_detail),0) + 1 as 'run_number' FROM tbl_trans_uang_keluar_pra_act_details ").Scan(&intKd_trans_keluar_detail)

		datadetail := table_data.Tbl_trans_uang_keluar_pra_act_details{
			Kd_trans_keluar:        intKd_trans_keluar,
			Kd_trans_keluar_detail: intKd_trans_keluar_detail,
			Seqno:                  int_seqno,			
			Jml_bayar:              jml_bayar,
			Keterangan:             keterangan,
			Created_by:             currentUser.(string),
			Created_on:             datenowx,
			Flag_aktif:             0,
		}//Tgl_bayar:              tgl_document,

		err = db.Omit("Edited_on", "Edited_by", "Kd_post_uang_masuk","Tgl_bayar").Create(&datadetail).Error
		if err != nil {
			response := helper.APIResponse("Simpan Data Detail Gagal ...", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		int_seqno++
	}

	var dataHeader table_data.Tbl_trans_uang_keluar_pra_headers
	err = db.Raw("UPDATE tbl_trans_uang_keluar_pra_headers SET flag_act=1, "+
		" edited_on = ? , edited_by = ?   "+
		" WHERE kd_trans_keluar = ? "+
		" and flag_aktif=0 ", datenowx, currentUser.(string), paramCreateACT.Kd_trans_keluar).Scan(&dataHeader).Error
	if err != nil {
		response := helper.APIResponse("Update Data Ke Tbl_trans_uang_keluar_pra_headers Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//setting tampilan habis save document
	SetArrayData := []GetBiayaAndSisa{}
	var kd_trans_keluar int
	var nm_group string
	var nm_kategori string
	var ket string

	ssql_show := " SELECT distinct a.kd_trans_keluar,a.kd_group,d.nm_group,a.kd_kategori,e.nm_kategori,a.no_document,a.tgl_document, " +
		" a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan " +
		" FROM tbl_trans_uang_keluar_pra_act_headers a " +
		" INNER JOIN tbl_trans_uang_keluar_pra_act_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
		" INNER JOIN tbl_group_kategoris d on a.kd_group = d.kd_group " +
		" INNER JOIN tbl_kategori_uangs e on a.kd_kategori = e.kd_kategori " +
		" where a.flag_aktif=0 and b.flag_aktif=0  "

	ssql_show = fmt.Sprintf("%s and a.kd_trans_keluar = %d", ssql_show, intKd_trans_keluar)
	// ssql_show = fmt.Sprintf("%s ORDER BY a.tgl_document %s", ssql_show, "asc")

	rows_show, _ := db.Raw(ssql_show).Rows()
	defer rows_show.Close()
	for rows_show.Next() {
		rows_show.Scan(&kd_trans_keluar, &kd_group, &nm_group, &kd_kategori, &nm_kategori, &no_document, &tgl_document, &total_biaya, &total_bayar, &sisa_biaya, &ket)
		arraydata := GetBiayaAndSisa{}
		arraydata.Kd_trans_keluar = kd_trans_keluar
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group
		arraydata.Kd_kategori = kd_kategori
		arraydata.Nm_kategori = nm_kategori
		arraydata.No_document = no_document
		arraydata.Tgl_document = tgl_document
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = ket

		sql := " SELECT b.kd_trans_keluar_detail,b.seqno, " +
			" b.kd_post_uang_masuk,c.nm_group,b.tgl_bayar,b.jml_bayar,b.keterangan,d.kd_pembayaran,d.tipe_pembayaran " +
			" FROM tbl_trans_uang_keluar_pra_act_headers a " +
			" INNER JOIN tbl_trans_uang_keluar_pra_act_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
			" left join tbl_group_kategoris c on b.kd_post_uang_masuk = c.kd_group " +
			" LEFT JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran " +
			" where a.flag_aktif=0 and b.flag_aktif=0  "

		sql = fmt.Sprintf("%s and a.kd_trans_keluar = %d", sql, kd_trans_keluar)

		sql = fmt.Sprintf("%s ORDER BY a.kd_trans_keluar %s,b.seqno %s", sql, "asc", "asc")

		var getDataUmSiswa []GetDataUmSiswa
		db.Raw(sql).Scan(&getDataUmSiswa)

		arraydata.Detail = getDataUmSiswa
		SetArrayData = append(SetArrayData, arraydata)
	}

	if len(SetArrayData) == 0 {
		response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
	c.JSON(http.StatusOK, response)

}

func EditUangKeluar(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	kd_trans_keluar := c.Param("idhead")

	var editBiayaHeader EditBiayaHeader
	if err := c.ShouldBindJSON(&editBiayaHeader); err != nil {
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

	var intJmldata int
	db.Raw(" SELECT count(*) jmldata FROM tbl_trans_uang_keluar_pra_act_headers a "+
		" INNER JOIN tbl_trans_uang_keluar_pra_act_details b on a.kd_trans_keluar=b.kd_trans_keluar "+
		" where a.flag_aktif=0 and b.flag_aktif=0 and a.kd_trans_keluar=? limit 1 ", kd_trans_keluar).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Data Uang Keluar Tidak Ditemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser")

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

	var sumJmlBayar float64
	db.Raw("SELECT sum(jml_bayar) 'jml_bayar' FROM tbl_trans_uang_keluar_pra_act_details "+
		" where flag_aktif=0 and tgl_bayar is not null and kd_trans_keluar=?", kd_trans_keluar).Scan(&sumJmlBayar)

	var sisa_biaya float64 = editBiayaHeader.Total_biaya - sumJmlBayar

	var dataHeader table_data.Tbl_trans_uang_keluar_pra_act_headers
	err = db.Raw("UPDATE tbl_trans_uang_keluar_pra_act_headers SET total_biaya=?,sisa_biaya=?, "+
		" edited_on = ? , edited_by = ?   "+
		" WHERE kd_trans_keluar = ? "+
		" and flag_aktif=0 ", editBiayaHeader.Total_biaya, sisa_biaya, datenowx, currentUser.(string), kd_trans_keluar).Scan(&dataHeader).Error
	if err != nil {
		response := helper.APIResponse("Update Data Ke Tbl_trans_uang_keluar_pra_act_headers Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//setting tampilan habis save document
	SetArrayData := []GetBiayaAndSisa{}
	var kd_group int
	var kd_kategori int
	var total_biaya float64
	var total_bayar float64
	var no_document string
	var tgl_document string
	var nm_group string
	var nm_kategori string
	var intkd_trans_keluar int
	var ket string

	ssql_show := " SELECT distinct a.kd_trans_keluar,a.kd_group,d.nm_group,a.kd_kategori,e.nm_kategori,a.no_document,a.tgl_document, " +
		" a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan " +
		" FROM tbl_trans_uang_keluar_pra_act_headers a " +
		" INNER JOIN tbl_trans_uang_keluar_pra_act_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
		" INNER JOIN tbl_group_kategoris d on a.kd_group = d.kd_group " +
		" INNER JOIN tbl_kategori_uangs e on a.kd_kategori = e.kd_kategori " +
		" where a.flag_aktif=0 and b.flag_aktif=0  "

	ssql_show = fmt.Sprintf("%s and a.kd_trans_keluar = '%s'", ssql_show, kd_trans_keluar)
	// ssql_show = fmt.Sprintf("%s ORDER BY a.tgl_document %s", ssql_show, "asc")

	rows_show, _ := db.Raw(ssql_show).Rows()
	defer rows_show.Close()
	for rows_show.Next() {
		rows_show.Scan(&intkd_trans_keluar, &kd_group, &nm_group, &kd_kategori, &nm_kategori, &no_document, &tgl_document, &total_biaya, &total_bayar, &sisa_biaya, &ket)
		arraydata := GetBiayaAndSisa{}
		arraydata.Kd_trans_keluar = intkd_trans_keluar
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group
		arraydata.Kd_kategori = kd_kategori
		arraydata.Nm_kategori = nm_kategori
		arraydata.No_document = no_document
		arraydata.Tgl_document = tgl_document
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = ket

		sql := " SELECT b.kd_trans_keluar_detail,b.seqno, " +
			" b.kd_post_uang_masuk,c.nm_group,b.tgl_bayar,b.jml_bayar,b.keterangan,d.kd_pembayaran,d.tipe_pembayaran " +
			" FROM tbl_trans_uang_keluar_pra_act_headers a " +
			" INNER JOIN tbl_trans_uang_keluar_pra_act_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
			" left join tbl_group_kategoris c on b.kd_post_uang_masuk = c.kd_group " +
			" left JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran " +
			" where a.flag_aktif=0 and b.flag_aktif=0  "

		sql = fmt.Sprintf("%s and a.kd_trans_keluar = '%s'", sql, kd_trans_keluar)

		sql = fmt.Sprintf("%s ORDER BY a.kd_trans_keluar %s,b.seqno %s", sql, "asc", "asc")

		var getDataUmSiswa []GetDataUmSiswa
		db.Raw(sql).Scan(&getDataUmSiswa)

		arraydata.Detail = getDataUmSiswa
		SetArrayData = append(SetArrayData, arraydata)
	}

	if len(SetArrayData) == 0 {
		response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
	c.JSON(http.StatusOK, response)

}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func UpdateUangKeluarDetail(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	kd_trans_keluar := c.Param("idhead")
	kd_trans_keluar_detail := c.Param("iddetail")

	var paramEditDetail ParamEditDetail
	if err := c.ShouldBindJSON(&paramEditDetail); err != nil {
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

	//VALIDASI POS UANG MASUK BERDASARKAN KD_GROUP
	var jml_inputan_lama float64
	db.Raw(" SELECT jml_bayar FROM tbl_trans_uang_keluar_pra_act_details where kd_trans_keluar_detail=? ", kd_trans_keluar_detail).Scan(&jml_inputan_lama)

	jml_inputan_baru := paramEditDetail.Jml_bayar

	if jml_inputan_baru > jml_inputan_lama {

		var jumlah_total_masuk float64
		db.Raw(" SELECT sum(total_bayar) FROM vw_group_uang_masuk where kd_group=? ", paramEditDetail.Kd_post_uang_masuk).Scan(&jumlah_total_masuk)

		var jumlah_total_keluar float64
		db.Raw(" SELECT sum(total_bayar) FROM vw_group_uang_keluar where kd_group=? ", paramEditDetail.Kd_post_uang_masuk).Scan(&jumlah_total_keluar)

		jumlah_total_keluar = jumlah_total_keluar - jml_inputan_lama

		sisa_total_masuk := jumlah_total_masuk - (jumlah_total_keluar + jml_inputan_baru)

		if sisa_total_masuk < 0 {
			errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
			response := helper.APIResponse("Pos Uang Masuk Kategori ini Tidak Cukup ...", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}
	//END VALIDASI POS UANG MASUK BERDASARKAN KD_GROUP

	var intJmldata int
	db.Raw(" SELECT count(*) jmldata FROM tbl_trans_uang_keluar_pra_act_headers a "+
		" INNER JOIN tbl_trans_uang_keluar_pra_act_details b on a.kd_trans_keluar=b.kd_trans_keluar "+
		" where a.flag_aktif=0 and b.flag_aktif=0 and a.kd_trans_keluar=? limit 1 ", kd_trans_keluar).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Data Uang Keluar Tidak Ditemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	db.Raw(" SELECT count(*) jmldata FROM tbl_trans_uang_keluar_pra_act_headers a "+
		" INNER JOIN tbl_trans_uang_keluar_pra_act_details b on a.kd_trans_keluar=b.kd_trans_keluar "+
		" where a.flag_aktif=0 and b.flag_aktif=0 and a.kd_trans_keluar=? and kd_trans_keluar_detail=? limit 1 ", kd_trans_keluar, kd_trans_keluar_detail).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Data Uang Keluar Tidak Ditemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser")

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

	tTglBayar, err2 := time.Parse("02-01-2006", paramEditDetail.Tgl_bayar)
	if err2 != nil {
		var ve validator.ValidationErrors
		if errors.As(err2, &ve) {
			errors := helper.FormatValidationError(err2)
			errorMessage := gin.H{"errors": errors}
			response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		var error_binding []string
		error_binding = append(error_binding, err2.Error())
		errorMessage := gin.H{"errors": error_binding}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	dateStr := tTglBayar.Format("2006-01-02")

	var dataDetail table_data.Tbl_trans_uang_keluar_pra_act_details
	err = db.Raw("update tbl_trans_uang_keluar_pra_act_details set kd_post_uang_masuk=?, tgl_bayar=?, "+
		" jml_bayar=?,keterangan=?,edited_by=?,edited_on=?,kd_pembayaran=? "+
		" where kd_trans_keluar_detail=? and kd_trans_keluar=? and flag_aktif=0 ", paramEditDetail.Kd_post_uang_masuk, dateStr,
		paramEditDetail.Jml_bayar, paramEditDetail.Keterangan, currentUser.(string), datenowx, paramEditDetail.Kd_pembayaran,
		kd_trans_keluar_detail, kd_trans_keluar).Scan(&dataDetail).Error
	if err != nil {
		response := helper.APIResponse("Update Data Ke Tbl_trans_uang_keluar_pra_act_details Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var sumJmlBayar float64
	db.Raw("SELECT sum(jml_bayar) 'jml_bayar' FROM tbl_trans_uang_keluar_pra_act_details "+
		" where flag_aktif=0 and tgl_bayar is not null and kd_trans_keluar=?", kd_trans_keluar).Scan(&sumJmlBayar)

	var total_biaya float64
	db.Raw("SELECT total_biaya FROM tbl_trans_uang_keluar_pra_act_headers where flag_aktif=0 and kd_trans_keluar=?", kd_trans_keluar).Scan(&total_biaya)
	var sisa_biaya float64 = total_biaya - sumJmlBayar

	var dataHeader table_data.Tbl_trans_uang_keluar_pra_act_headers
	err = db.Raw("UPDATE tbl_trans_uang_keluar_pra_act_headers SET total_bayar = ?, sisa_biaya = ?, "+
		" edited_on = ? , edited_by = ? "+
		" WHERE kd_trans_keluar = ? and flag_aktif=0 ", sumJmlBayar, sisa_biaya, datenowx, currentUser.(string), kd_trans_keluar).Scan(&dataHeader).Error
	if err != nil {
		response := helper.APIResponse("Update Data Ke Tbl_trans_uang_keluar_pra_act_headers Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//setting tampilan habis save siswa

	SetArrayData := []GetBiayaAndSisa{}
	var kd_group int
	var nm_group string
	var kd_kategori int
	var nm_kategori string
	var total_bayar float64
	var no_document string
	var tgl_document string
	var ket string

	ssql := " SELECT distinct b.kd_trans_keluar,a.kd_group,d.nm_group,a.kd_kategori,e.nm_kategori,a.no_document,a.tgl_document, " +
		" a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan " +
		" FROM tbl_trans_uang_keluar_pra_act_headers a " +
		" INNER JOIN tbl_trans_uang_keluar_pra_act_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
		" INNER JOIN tbl_group_kategoris d on a.kd_group = d.kd_group " +
		" INNER JOIN tbl_kategori_uangs e on a.kd_kategori = e.kd_kategori " +
		" where a.flag_aktif=0 and b.flag_aktif=0  "

	ssql = fmt.Sprintf("%s and a.kd_trans_keluar= '%s'", ssql, kd_trans_keluar)
	ssql = fmt.Sprintf("%s ORDER BY a.tgl_document %s", ssql, "asc")

	int_kd_trans_keluar, _ := strconv.Atoi(kd_trans_keluar)

	rows, _ := db.Raw(ssql).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&int_kd_trans_keluar, &kd_group, &nm_group, &kd_kategori, &nm_kategori, &no_document, &tgl_document, &total_biaya, &total_bayar, &sisa_biaya, &ket)
		arraydata := GetBiayaAndSisa{}
		arraydata.Kd_trans_keluar = int_kd_trans_keluar
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group
		arraydata.Kd_kategori = kd_kategori
		arraydata.Nm_kategori = nm_kategori
		arraydata.No_document = no_document
		arraydata.Tgl_document = tgl_document
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = ket

		sql := " SELECT b.kd_trans_keluar_detail,b.seqno,b.kd_post_uang_masuk,c.nm_group, " +
			" b.tgl_bayar,b.jml_bayar,b.keterangan,d.kd_pembayaran,d.tipe_pembayaran " +
			" FROM tbl_trans_uang_keluar_pra_act_headers a " +
			" INNER JOIN tbl_trans_uang_keluar_pra_act_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
			" left join tbl_group_kategoris c on b.kd_post_uang_masuk = c.kd_group " +
			" left JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran " +
			" where a.flag_aktif=0 and b.flag_aktif=0  "

		sql = fmt.Sprintf("%s and a.kd_trans_keluar = '%s'", sql, kd_trans_keluar)

		sql = fmt.Sprintf("%s ORDER BY a.kd_trans_keluar %s,b.seqno %s", sql, "asc", "asc")

		var getDataUmSiswa []GetDataUmSiswa
		db.Raw(sql).Scan(&getDataUmSiswa)

		arraydata.Detail = getDataUmSiswa
		SetArrayData = append(SetArrayData, arraydata)
	}

	if len(SetArrayData) == 0 {
		response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
	c.JSON(http.StatusOK, response)

}

func CreateUangKeluarDetail(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramAddDetail ParamAddDetail
	if err := c.ShouldBindJSON(&paramAddDetail); err != nil {
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

	var intJmldata int
	//cek data transaksi
	db.Raw(" SELECT count(*) jmldata FROM tbl_trans_uang_keluar_pra_act_headers a "+
		" INNER JOIN tbl_trans_uang_keluar_pra_act_details b on a.kd_trans_keluar=b.kd_trans_keluar "+
		" where a.kd_trans_keluar=? and a.flag_aktif=0 and b.flag_aktif=0 ", paramAddDetail.Kd_trans_keluar).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Data Transaksi Document Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser")

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

	var intkd_trans_keluar_detail int
	db.Raw("SELECT ifnull(max(kd_trans_keluar_detail),0) + 1 as 'run_number' FROM tbl_trans_uang_keluar_pra_act_details ").Scan(&intkd_trans_keluar_detail)

	var int_seqno int
	db.Raw("SELECT (seqno + 1) as 'run_number' "+
		" FROM tbl_trans_uang_keluar_pra_act_details where flag_aktif=0 and kd_trans_keluar=?", paramAddDetail.Kd_trans_keluar).Scan(&int_seqno)

	datadetail := table_data.Tbl_trans_uang_keluar_pra_act_details{
		Kd_trans_keluar:        paramAddDetail.Kd_trans_keluar,
		Kd_trans_keluar_detail: intkd_trans_keluar_detail,
		Seqno:                  int_seqno,
		Jml_bayar:              0,
		Keterangan:             "",
		Created_by:             currentUser.(string),
		Created_on:             datenowx,
		Flag_aktif:             0,
	}

	err = db.Omit("Edited_on", "Edited_by", "Tgl_bayar", "Kd_post_uang_masuk").Create(&datadetail).Error
	if err != nil {
		response := helper.APIResponse("Simpan Data Detail Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//setting tampilan habis save siswa
	SetArrayData := []GetBiayaAndSisa{}
	var kd_trans_keluar int
	var kd_group int
	var nm_group string
	var kd_kategori int
	var nm_kategori string
	var total_biaya float64
	var total_bayar float64
	var sisa_biaya float64
	var no_document string
	var tgl_document string
	var ket string

	ssql := " SELECT distinct b.kd_trans_keluar,a.kd_group,d.nm_group,a.kd_kategori,e.nm_kategori,a.no_document,a.tgl_document, " +
		" a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan " +
		" FROM tbl_trans_uang_keluar_pra_act_headers a " +
		" INNER JOIN tbl_trans_uang_keluar_pra_act_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
		" INNER JOIN tbl_group_kategoris d on a.kd_group = d.kd_group " +
		" INNER JOIN tbl_kategori_uangs e on a.kd_kategori = e.kd_kategori " +
		" where a.flag_aktif=0 and b.flag_aktif=0   "

	ssql = fmt.Sprintf("%s and a.kd_trans_keluar= %d", ssql, paramAddDetail.Kd_trans_keluar)
	ssql = fmt.Sprintf("%s ORDER BY a.tgl_document %s", ssql, "asc")

	//int_Kd_trans_masuk_lain, _ := strconv.Atoi(Kd_trans_masuk_lain)

	rows, _ := db.Raw(ssql).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&kd_trans_keluar, &kd_group, &nm_group, &kd_kategori, &nm_kategori, &no_document, &tgl_document, &total_biaya, &total_bayar, &sisa_biaya, &ket)
		arraydata := GetBiayaAndSisa{}
		arraydata.Kd_trans_keluar = kd_trans_keluar
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group
		arraydata.Kd_kategori = kd_kategori
		arraydata.Nm_kategori = nm_kategori
		arraydata.No_document = no_document
		arraydata.Tgl_document = tgl_document
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = ket

		sql := " SELECT b.kd_trans_keluar_detail,b.seqno,b.kd_post_uang_masuk,c.nm_group, " +
			" b.tgl_bayar,b.jml_bayar,b.keterangan,d.kd_pembayaran,d.tipe_pembayaran " +
			" FROM tbl_trans_uang_keluar_pra_act_headers a " +
			" INNER JOIN tbl_trans_uang_keluar_pra_act_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
			" left join tbl_group_kategoris c on b.kd_post_uang_masuk = c.kd_group " +
			" left JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran " +
			" where a.flag_aktif=0 and b.flag_aktif=0  "

		sql = fmt.Sprintf("%s and a.kd_trans_keluar = %d", sql, kd_trans_keluar)

		sql = fmt.Sprintf("%s ORDER BY a.kd_trans_keluar %s,b.seqno %s", sql, "asc", "asc")

		var getDataUmSiswa []GetDataUmSiswa
		db.Raw(sql).Scan(&getDataUmSiswa)

		arraydata.Detail = getDataUmSiswa
		SetArrayData = append(SetArrayData, arraydata)
	}

	if len(SetArrayData) == 0 {
		response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
	c.JSON(http.StatusOK, response)

}

func ListData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramChangeSiswa ParamChangeSiswa
	if err := c.ShouldBindJSON(&paramChangeSiswa); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	tTglTransaksi1, err2 := time.Parse("02-01-2006", paramChangeSiswa.Tgl_document1)
	if err2 != nil {
		var ve validator.ValidationErrors
		if errors.As(err2, &ve) {
			errors := helper.FormatValidationError(err2)
			errorMessage := gin.H{"errors": errors}
			response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		var error_binding []string
		error_binding = append(error_binding, err2.Error())
		errorMessage := gin.H{"errors": error_binding}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	dateStr1 := tTglTransaksi1.Format("2006-01-02")

	tTglTransaksi2, err2 := time.Parse("02-01-2006", paramChangeSiswa.Tgl_document2)
	if err2 != nil {
		var ve validator.ValidationErrors
		if errors.As(err2, &ve) {
			errors := helper.FormatValidationError(err2)
			errorMessage := gin.H{"errors": errors}
			response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		var error_binding []string
		error_binding = append(error_binding, err2.Error())
		errorMessage := gin.H{"errors": error_binding}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	dateStr2 := tTglTransaksi2.Format("2006-01-02")

	SetArrayData := []GetBiayaAndSisa{}
	var kd_trans_keluar int
	var kd_group int
	var nm_group string
	var kd_kategori int
	var nm_kategori string
	var total_biaya float64
	var total_bayar float64
	var sisa_biaya float64
	var no_document string
	var tgl_document string
	var ket string

	ssql := " SELECT distinct b.kd_trans_keluar,a.kd_group,d.nm_group,a.kd_kategori,e.nm_kategori,a.no_document,a.tgl_document, " +
		" a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan " +
		" FROM tbl_trans_uang_keluar_pra_act_headers a " +
		" INNER JOIN tbl_trans_uang_keluar_pra_act_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
		" INNER JOIN tbl_group_kategoris d on a.kd_group = d.kd_group " +
		" INNER JOIN tbl_kategori_uangs e on a.kd_kategori = e.kd_kategori " +
		" where a.flag_aktif=0 and b.flag_aktif=0   "

	ssql = fmt.Sprintf("%s and a.tgl_document >= '%s'", ssql, dateStr1)
	ssql = fmt.Sprintf("%s and a.tgl_document <= '%s'", ssql, dateStr2)

	if paramChangeSiswa.No_document != "" {
		ssql = fmt.Sprintf("%s and a.no_document = '%s'", ssql, paramChangeSiswa.No_document)
	}

	ssql = fmt.Sprintf("%s ORDER BY a.tgl_document %s", ssql, "asc")
	rows, _ := db.Raw(ssql).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&kd_trans_keluar, &kd_group, &nm_group, &kd_kategori, &nm_kategori, &no_document, &tgl_document, &total_biaya, &total_bayar, &sisa_biaya, &ket)
		arraydata := GetBiayaAndSisa{}
		arraydata.Kd_trans_keluar = kd_trans_keluar
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group
		arraydata.Kd_kategori = kd_kategori
		arraydata.Nm_kategori = nm_kategori
		arraydata.No_document = no_document
		arraydata.Tgl_document = tgl_document
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = ket

		sql := " SELECT b.kd_trans_keluar_detail,b.seqno,b.kd_post_uang_masuk,c.nm_group, " +
			" b.tgl_bayar,b.jml_bayar,b.keterangan,d.kd_pembayaran,d.tipe_pembayaran " +
			" FROM tbl_trans_uang_keluar_pra_act_headers a " +
			" INNER JOIN tbl_trans_uang_keluar_pra_act_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
			" left join tbl_group_kategoris c on b.kd_post_uang_masuk = c.kd_group " +
			" left JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran " +
			" where a.flag_aktif=0 and b.flag_aktif=0  "

		sql = fmt.Sprintf("%s and a.kd_trans_keluar = %d", sql, kd_trans_keluar)

		sql = fmt.Sprintf("%s ORDER BY a.kd_trans_keluar %s,b.seqno %s", sql, "asc", "asc")

		var getDataUmSiswa []GetDataUmSiswa
		db.Raw(sql).Scan(&getDataUmSiswa)

		arraydata.Detail = getDataUmSiswa
		SetArrayData = append(SetArrayData, arraydata)
	}

	if len(SetArrayData) == 0 {
		response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
	c.JSON(http.StatusOK, response)
}

func DeleteUangKeluarDetail(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramEditUmSiswaDetail ParamDeleteUmLainDetail

	if err := c.ShouldBindJSON(&paramEditUmSiswaDetail); err != nil {
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

	kd_trans_keluar := paramEditUmSiswaDetail.Kd_trans_keluar
	kd_trans_keluar_detail := paramEditUmSiswaDetail.Kd_trans_keluar_detail

	var dataUtama table_data.Tbl_trans_uang_keluar_pra_act_details
	if err := db.Where("flag_aktif=0 and kd_trans_keluar_detail=? and kd_trans_keluar=?", kd_trans_keluar_detail, kd_trans_keluar).First(&dataUtama).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
		response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser")

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

	var dataDetail table_data.Tbl_trans_uang_keluar_pra_act_details
	err = db.Raw("update tbl_trans_uang_keluar_pra_act_details set flag_aktif=9,edited_by=?,edited_on=? "+
		" where kd_trans_keluar_detail=? and kd_trans_keluar=? and flag_aktif=0 ",
		currentUser.(string), datenowx, kd_trans_keluar_detail, kd_trans_keluar).Scan(&dataDetail).Error
	if err != nil {
		response := helper.APIResponse("Delete Data Ke Tbl_trans_uang_keluar_pra_act_details Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var sumJmlBayar float64
	db.Raw("SELECT sum(jml_bayar) 'jml_bayar' FROM tbl_trans_uang_keluar_pra_act_details "+
		" where flag_aktif=0 and tgl_bayar is not null and kd_trans_keluar=?", kd_trans_keluar).Scan(&sumJmlBayar)

	var total_biaya float64
	db.Raw("SELECT total_biaya FROM tbl_trans_uang_keluar_pra_act_headers where flag_aktif=0 and kd_trans_keluar=?", kd_trans_keluar).Scan(&total_biaya)
	var sisa_biaya float64 = total_biaya - sumJmlBayar

	var dataHeader table_data.Tbl_trans_uang_keluar_pra_act_headers
	err = db.Raw("UPDATE tbl_trans_uang_keluar_pra_act_headers SET total_bayar = ?, sisa_biaya = ?, "+
		" edited_on = ? , edited_by = ? "+
		" WHERE kd_trans_keluar = ? and flag_aktif=0 ", sumJmlBayar, sisa_biaya, datenowx, currentUser.(string), kd_trans_keluar).Scan(&dataHeader).Error
	if err != nil {
		response := helper.APIResponse("Update Data Ke Tbl_trans_uang_keluar_pra_act_headers Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//setting tampilan habis save siswa

	SetArrayData := []GetBiayaAndSisa{}
	var kd_group int
	var nm_group string
	var kd_kategori int
	var nm_kategori string
	var total_bayar float64
	var no_document string
	var tgl_document string
	var ket string

	ssql := " SELECT distinct b.kd_trans_keluar,a.kd_group,d.nm_group,a.kd_kategori,e.nm_kategori,a.no_document,a.tgl_document, " +
		" a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan " +
		" FROM tbl_trans_uang_keluar_pra_act_headers a " +
		" INNER JOIN tbl_trans_uang_keluar_pra_act_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
		" INNER JOIN tbl_group_kategoris d on a.kd_group = d.kd_group " +
		" INNER JOIN tbl_kategori_uangs e on a.kd_kategori = e.kd_kategori " +
		" where a.flag_aktif=0 and b.flag_aktif=0   "

	ssql = fmt.Sprintf("%s and a.kd_trans_keluar= '%s'", ssql, kd_trans_keluar)
	ssql = fmt.Sprintf("%s ORDER BY a.tgl_document %s", ssql, "asc")

	int_kd_trans_keluar, _ := strconv.Atoi(kd_trans_keluar)

	rows, _ := db.Raw(ssql).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&kd_trans_keluar, &kd_group, &nm_group, &kd_kategori, &nm_kategori, &no_document, &tgl_document, &total_biaya, &total_bayar, &sisa_biaya, &ket)
		arraydata := GetBiayaAndSisa{}
		arraydata.Kd_trans_keluar = int_kd_trans_keluar
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group
		arraydata.Kd_kategori = kd_kategori
		arraydata.Nm_kategori = nm_kategori
		arraydata.No_document = no_document
		arraydata.Tgl_document = tgl_document
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = ket

		sql := " SELECT b.kd_trans_keluar_detail,b.seqno,b.kd_post_uang_masuk,c.nm_group, " +
			" b.tgl_bayar,b.jml_bayar,b.keterangan,d.kd_pembayaran,d.tipe_pembayaran " +
			" FROM tbl_trans_uang_keluar_pra_act_headers a " +
			" INNER JOIN tbl_trans_uang_keluar_pra_act_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
			" left join tbl_group_kategoris c on b.kd_post_uang_masuk = c.kd_group " +
			" left JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran " +
			" where a.flag_aktif=0 and b.flag_aktif=0  "

		sql = fmt.Sprintf("%s and a.kd_trans_keluar = %d", sql, int_kd_trans_keluar)

		sql = fmt.Sprintf("%s ORDER BY a.kd_trans_keluar %s,b.seqno %s", sql, "asc", "asc")

		var getDataUmSiswa []GetDataUmSiswa
		db.Raw(sql).Scan(&getDataUmSiswa)

		arraydata.Detail = getDataUmSiswa
		SetArrayData = append(SetArrayData, arraydata)
	}

	if len(SetArrayData) == 0 {
		response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
	c.JSON(http.StatusOK, response)
}

func DeleteAllUangKeluar(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	kd_trans_keluar := c.Param("idhead")

	var dataUtama table_data.Tbl_trans_uang_keluar_pra_act_headers
	if err := db.Where("flag_aktif=0 and kd_trans_keluar=?", kd_trans_keluar).First(&dataUtama).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Header Tidak Ditemukan ..."}
		response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var dataUtamaDet table_data.Tbl_trans_uang_keluar_pra_act_details
	if err := db.Where("flag_aktif=0 and kd_trans_keluar=?", kd_trans_keluar).First(&dataUtamaDet).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Detail Tidak Ditemukan ..."}
		response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser")

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

	var dataDetail table_data.Tbl_trans_uang_keluar_pra_act_details
	err = db.Raw("update tbl_trans_uang_keluar_pra_act_details set flag_aktif=9,edited_by=?,edited_on=? "+
		" where kd_trans_keluar=? and flag_aktif=0 ",
		currentUser.(string), datenowx, kd_trans_keluar).Scan(&dataDetail).Error
	if err != nil {
		response := helper.APIResponse("Delete Data Ke Tbl_trans_uang_keluar_pra_act_details Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var dataHead table_data.Tbl_trans_uang_keluar_pra_act_headers
	err = db.Raw("update tbl_trans_uang_keluar_pra_act_headers set flag_aktif=9,edited_by=?,edited_on=? "+
		" where kd_trans_keluar=? and flag_aktif=0 ",
		currentUser.(string), datenowx, kd_trans_keluar).Scan(&dataHead).Error
	if err != nil {
		response := helper.APIResponse("Delete Data Ke Tbl_trans_uang_keluar_pra_act_headers Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//setting tampilan habis save siswa

	SetArrayData := []GetBiayaAndSisa{}
	var kd_group int
	var nm_group string
	var kd_kategori int
	var nm_kategori string
	var total_bayar float64
	var no_document string
	var tgl_document string
	var total_biaya float64
	var sisa_biaya float64
	var ket string

	ssql := " SELECT distinct b.kd_trans_keluar,a.kd_group,d.nm_group,a.kd_kategori,e.nm_kategori,a.no_document,a.tgl_document, " +
		" a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan " +
		" FROM tbl_trans_uang_keluar_pra_act_headers a " +
		" INNER JOIN tbl_trans_uang_keluar_pra_act_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
		" INNER JOIN tbl_group_kategoris d on a.kd_group = d.kd_group " +
		" INNER JOIN tbl_kategori_uangs e on a.kd_kategori = e.kd_kategori " +
		" where a.flag_aktif=0 and b.flag_aktif=0   "

	ssql = fmt.Sprintf("%s ORDER BY a.tgl_document %s", ssql, "asc")

	//int_kd_trans_keluar, _ := strconv.Atoi(kd_trans_keluar)

	rows, _ := db.Raw(ssql).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&kd_trans_keluar, &kd_group, &nm_group, &kd_kategori, &nm_kategori, &no_document, &tgl_document, &total_biaya, &total_bayar, &sisa_biaya, &ket)
		int_kd_trans_keluarx, _ := strconv.Atoi(kd_trans_keluar)
		arraydata := GetBiayaAndSisa{}
		arraydata.Kd_trans_keluar = int_kd_trans_keluarx
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group
		arraydata.Kd_kategori = kd_kategori
		arraydata.Nm_kategori = nm_kategori
		arraydata.No_document = no_document
		arraydata.Tgl_document = tgl_document
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = ket

		sql := " SELECT b.kd_trans_keluar_detail,b.seqno,b.kd_post_uang_masuk,c.nm_group, " +
			" b.tgl_bayar,b.jml_bayar,b.keterangan,d.kd_pembayaran,d.tipe_pembayaran " +
			" FROM tbl_trans_uang_keluar_pra_act_headers a " +
			" INNER JOIN tbl_trans_uang_keluar_pra_act_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
			" left join tbl_group_kategoris c on b.kd_post_uang_masuk = c.kd_group " +
			" left JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran " +
			" where a.flag_aktif=0 and b.flag_aktif=0  "

		sql = fmt.Sprintf("%s and a.kd_trans_keluar = '%s'", sql, kd_trans_keluar)

		sql = fmt.Sprintf("%s ORDER BY a.kd_trans_keluar %s,b.seqno %s", sql, "asc", "asc")

		var getDataUmSiswa []GetDataUmSiswa
		db.Raw(sql).Scan(&getDataUmSiswa)

		arraydata.Detail = getDataUmSiswa
		SetArrayData = append(SetArrayData, arraydata)
	}

	if len(SetArrayData) == 0 {
		response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
	c.JSON(http.StatusOK, response)
}
