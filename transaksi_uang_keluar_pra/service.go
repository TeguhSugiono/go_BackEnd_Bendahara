package transaksi_uang_keluar_pra

import (
	"errors"
	"fmt"
	"net/http"
	"rest_api_bendahara/helper"
	"rest_api_bendahara/master_group_kategori"
	"rest_api_bendahara/master_kategori_uang"
	"rest_api_bendahara/table_data"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)



func ListGroupKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//kd_jenis=2 adalah uang keluar
	var master []master_group_kategori.ListData
	sql := "  SELECT a.*,b.proses_uang FROM tbl_group_kategoris as a " +
		" inner join tbl_jenis_trans as b on a.kd_jenis=b.kd_jenis  " +
		" where a.flag_aktif=0 and b.flag_aktif=0  " +
		" and a.kd_jenis in('2') order by b.proses_uang "

	db.Raw(sql).Scan(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", master_group_kategori.FormatGroupKategori(master))
	c.JSON(http.StatusOK, response)
}

func ListKategoriUang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramChangeKategori ParamChangeKategori
	if err := c.ShouldBindJSON(&paramChangeKategori); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var master []master_kategori_uang.ListData
	sql := " SELECT a.*,b.nm_group FROM tbl_kategori_uangs as a  " +
		" INNER JOIN tbl_group_kategoris b on a.kd_group=b.kd_group " +
		" where a.flag_aktif=0 and b.flag_aktif=0 " +
		" and a.kd_group in (" + paramChangeKategori.Kd_group + ") order by b.nm_group "

	db.Raw(sql).Scan(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", master_kategori_uang.FormatShowData(master))
	c.JSON(http.StatusOK, response)
}

func CreateUangKeluar(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var paramInputTransaksi ParamInputTransaksi
	if err := c.ShouldBindJSON(&paramInputTransaksi); err != nil {
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

	tTglTransaksi, err2 := time.Parse("02-01-2006", paramInputTransaksi.Tgl_document)
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
	dateStr := tTglTransaksi.Format("2006-01-02")

	var intJmldata int
	db.Raw(" SELECT count(*) jmldata FROM tbl_group_kategoris where kd_group=? and flag_aktif=0 ", paramInputTransaksi.Kd_group).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Kode Group Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	db.Raw(" SELECT count(*) jmldata FROM tbl_kategori_uangs where kd_kategori=? and flag_aktif=0 ", paramInputTransaksi.Kd_kategori).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Kode Kategori Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	db.Raw(" SELECT count(*) jmldata from tbl_trans_uang_keluar_pra_headers a "+
		" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar=b.kd_trans_keluar "+
		" where a.flag_aktif=0 and b.flag_aktif=0 and kd_group=?  "+
		" and kd_kategori=? and no_document=? and tgl_document=?  ",
		paramInputTransaksi.Kd_group, paramInputTransaksi.Kd_kategori, paramInputTransaksi.No_document, dateStr).Scan(&intJmldata)

	if intJmldata > 0 {

		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Data Document Perencanaan Pengeluaran Sudah Ada ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	} else {

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
		db.Raw("SELECT ifnull(max(kd_trans_keluar),0) + 1 as 'run_number' FROM tbl_trans_uang_keluar_pra_headers ").Scan(&intKd_trans_keluar)

		var float_biaya_spp float64
		db.Raw("SELECT jml_biaya FROM tbl_biaya_masuk_keluars where kd_kategori=?", paramInputTransaksi.Kd_kategori).Scan(&float_biaya_spp)

		if paramInputTransaksi.Total_biaya > 0 {
			float_biaya_spp = paramInputTransaksi.Total_biaya
		}

		currentUser := c.MustGet("currentUser")
		data := table_data.Tbl_trans_uang_keluar_pra_headers{
			Kd_group:        paramInputTransaksi.Kd_group,
			Kd_kategori:     paramInputTransaksi.Kd_kategori,
			Kd_trans_keluar: intKd_trans_keluar,
			No_document:     paramInputTransaksi.No_document,
			Tgl_document:    dateStr,
			Total_biaya:     float_biaya_spp,
			Total_bayar:     0,
			Sisa_biaya:      float_biaya_spp,
			Keterangan:      paramInputTransaksi.Keterangan,
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

		var intKd_trans_keluar_detail int
		db.Raw("SELECT ifnull(max(kd_trans_keluar_detail),0) + 1 as 'run_number' FROM tbl_trans_uang_keluar_pra_details ").Scan(&intKd_trans_keluar_detail)

		datadetail := table_data.Tbl_trans_uang_keluar_pra_details{
			Kd_trans_keluar:        intKd_trans_keluar,
			Kd_trans_keluar_detail: intKd_trans_keluar_detail,
			Seqno:                  1,
			Jml_bayar:              0,
			Keterangan:             "",
			Created_by:             currentUser.(string),
			Created_on:             datenowx,
			Flag_aktif:             0,
		}

		err = db.Omit("Edited_on", "Edited_by").Create(&datadetail).Error
		if err != nil {
			response := helper.APIResponse("Simpan Data Detail Gagal ...", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		//setting tampilan habis save document
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

		ssql := " SELECT distinct a.kd_trans_keluar,a.kd_group,d.nm_group,a.kd_kategori,e.nm_kategori,a.no_document,a.tgl_document, " +
			" a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan " +
			" FROM tbl_trans_uang_keluar_pra_headers a " +
			" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
			" INNER JOIN tbl_group_kategoris d on a.kd_group = d.kd_group " +
			" INNER JOIN tbl_kategori_uangs e on a.kd_kategori = e.kd_kategori " +
			" where a.flag_aktif=0 and b.flag_aktif=0  "

		ssql = fmt.Sprintf("%s and a.kd_trans_keluar = %d", ssql, intKd_trans_keluar)
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

			sql := " SELECT b.kd_trans_keluar_detail,b.seqno, " +
				" b.jml_bayar,b.keterangan " +
				" FROM tbl_trans_uang_keluar_pra_headers a " +
				" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
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
}

func EditUangKeluar(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	Kd_trans_keluar := c.Param("idhead")

	var paramInputTransaksiEdit ParamInputTransaksiEdit
	if err := c.ShouldBindJSON(&paramInputTransaksiEdit); err != nil {
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
	db.Raw(" SELECT count(*) jmldata FROM tbl_group_kategoris where kd_group=? and flag_aktif=0 ", paramInputTransaksiEdit.Kd_group).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Kode Group Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	db.Raw(" SELECT count(*) jmldata FROM tbl_kategori_uangs where kd_kategori=? and flag_aktif=0 ", paramInputTransaksiEdit.Kd_kategori).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Kode Kategori Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	tTglTransaksi, err2 := time.Parse("02-01-2006", paramInputTransaksiEdit.Tgl_document)
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
	dateStr := tTglTransaksi.Format("2006-01-02")

	//cek data Document pembayaran mengada-ngada
	db.Raw(" SELECT count(*) jmldata FROM tbl_trans_uang_keluar_pra_headers a "+
		" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar=b.kd_trans_keluar "+
		" where a.kd_trans_keluar=? and a.flag_aktif=0 and b.flag_aktif=0  ", Kd_trans_keluar).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Data Document Pembayaran Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//Validasi data sama
	var kd_group_old int
	var kd_kategori_old int
	var no_document_old string
	var tgl_document_old string

	rowA, _ := db.Raw("SELECT a.kd_group,a.kd_kategori,a.no_document,date_format(a.tgl_document,'%Y-%m-%d') FROM tbl_trans_uang_keluar_pra_headers a "+
		" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar=b.kd_trans_keluar "+
		" where a.kd_trans_keluar=?  and a.flag_aktif=0 and b.flag_aktif=0 limit 1", Kd_trans_keluar).Rows()

	defer rowA.Close()
	for rowA.Next() {
		rowA.Scan(&kd_group_old, &kd_kategori_old, &no_document_old, &tgl_document_old)
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
	db.Raw("SELECT sum(jml_bayar) 'jml_bayar' FROM tbl_trans_uang_keluar_pra_details "+
		" where flag_aktif=0 and kd_trans_keluar=?", Kd_trans_keluar).Scan(&sumJmlBayar)

	var sisa_biaya float64 = paramInputTransaksiEdit.Total_biaya - sumJmlBayar

	var dataHeader table_data.Tbl_trans_uang_keluar_pra_headers

	//cek data jika ada perubahan di nodocument akan tetapi perubahan tersebut sama dengan data yang sudah ada
	if (paramInputTransaksiEdit.Kd_group != kd_group_old) || (paramInputTransaksiEdit.Kd_kategori != kd_kategori_old) || (paramInputTransaksiEdit.No_document != no_document_old) || (dateStr != tgl_document_old) {
		db.Raw(" SELECT count(*) jmldata FROM tbl_trans_uang_keluar_pra_headers a "+
			" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar=b.kd_trans_keluar "+
			" where a.flag_aktif=0 and b.flag_aktif=0  and a.kd_group=? and a.kd_kategori=? "+
			" and a.no_document=? and a.tgl_document=? ", paramInputTransaksiEdit.Kd_group,
			paramInputTransaksiEdit.Kd_kategori, paramInputTransaksiEdit.No_document, dateStr).Scan(&intJmldata)
		if intJmldata > 0 {
			errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
			response := helper.APIResponse("Data Document Pembayaran Sudah Ada ...", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		err = db.Raw("UPDATE tbl_trans_uang_keluar_pra_headers SET total_biaya=? ,total_bayar = ?, sisa_biaya = ?,keterangan=?,  "+
			" edited_on = ? , edited_by = ? ,kd_group=?,kd_kategori=?,tgl_document=?,no_document=?  "+
			" WHERE kd_trans_keluar = ? "+
			" and flag_aktif=0 ", paramInputTransaksiEdit.Total_biaya, sumJmlBayar, sisa_biaya, paramInputTransaksiEdit.Keterangan,datenowx, currentUser.(string),
			paramInputTransaksiEdit.Kd_group, paramInputTransaksiEdit.Kd_kategori, dateStr, paramInputTransaksiEdit.No_document,
			Kd_trans_keluar).Scan(&dataHeader).Error
		if err != nil {
			response := helper.APIResponse("Update Data Ke Tbl_trans_uang_keluar_pra_headers Gagal ...", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

	} else {

		err = db.Raw("UPDATE tbl_trans_uang_keluar_pra_headers SET total_biaya=? ,total_bayar = ?, sisa_biaya = ?,keterangan=?, "+
			" edited_on = ? , edited_by = ?   "+
			" WHERE kd_trans_keluar = ? "+
			" and flag_aktif=0 ", paramInputTransaksiEdit.Total_biaya, sumJmlBayar, sisa_biaya,paramInputTransaksiEdit.Keterangan, datenowx, currentUser.(string),
			Kd_trans_keluar).Scan(&dataHeader).Error
		if err != nil {
			response := helper.APIResponse("Update Data Ke Tbl_trans_uang_keluar_pra_headers Gagal ...", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

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
	var no_document string
	var tgl_document string
	var ket string

	ssql := " SELECT distinct b.kd_trans_keluar,a.kd_group,d.nm_group,a.kd_kategori,e.nm_kategori,a.no_document,a.tgl_document, " +
		" a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan " +
		" FROM tbl_trans_uang_keluar_pra_headers a " +
		" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
		" INNER JOIN tbl_group_kategoris d on a.kd_group = d.kd_group " +
		" INNER JOIN tbl_kategori_uangs e on a.kd_kategori = e.kd_kategori " +
		" where a.flag_aktif=0 and b.flag_aktif=0  "

	ssql = fmt.Sprintf("%s and a.kd_trans_keluar= '%s'", ssql, Kd_trans_keluar)
	ssql = fmt.Sprintf("%s ORDER BY a.tgl_document %s", ssql, "asc")

	//int_Kd_trans_keluar, _ := strconv.Atoi(Kd_trans_keluar)

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

		sql := " SELECT b.kd_trans_keluar_detail,b.seqno, " +
			" b.jml_bayar,b.keterangan " +
			" FROM tbl_trans_uang_keluar_pra_headers a " +
			" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
			" where a.flag_aktif=0 and b.flag_aktif=0   "

		sql = fmt.Sprintf("%s and a.kd_trans_keluar = '%s'", sql, Kd_trans_keluar)

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

func UpdateUangKeluarDetail(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	kd_trans_keluar := c.Param("idhead")
	kd_trans_keluar_detail := c.Param("iddetail")

	var paramEditUmSiswaDetail ParamEditUmSiswaDetail
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

	var dataUtama table_data.Tbl_trans_uang_keluar_pra_details
	if err := db.Where("flag_aktif=0 and kd_trans_keluar_detail=? and Kd_trans_keluar=?", kd_trans_keluar_detail, kd_trans_keluar).First(&dataUtama).Error; err != nil {
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

	var dataDetail table_data.Tbl_trans_uang_keluar_pra_details
	err = db.Raw("update tbl_trans_uang_keluar_pra_details set jml_bayar=?,keterangan=?,edited_by=?,edited_on=? "+
		" where kd_trans_keluar_detail=? and kd_trans_keluar=? and flag_aktif=0 ",
		paramEditUmSiswaDetail.Jml_bayar, paramEditUmSiswaDetail.Keterangan, currentUser.(string), datenowx, kd_trans_keluar_detail, kd_trans_keluar).Scan(&dataDetail).Error
	if err != nil {
		response := helper.APIResponse("Update Data Ke Tbl_trans_uang_keluar_pra_details Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var sumJmlBayar float64
	db.Raw("SELECT sum(jml_bayar) 'jml_bayar' FROM tbl_trans_uang_keluar_pra_details "+
		" where flag_aktif=0 and kd_trans_keluar=?", kd_trans_keluar).Scan(&sumJmlBayar)

	var total_biaya float64
	db.Raw("SELECT total_biaya FROM tbl_trans_uang_keluar_pra_headers where flag_aktif=0 and kd_trans_keluar=?", kd_trans_keluar).Scan(&total_biaya)
	var sisa_biaya float64 = total_biaya - sumJmlBayar

	var dataHeader table_data.Tbl_trans_uang_keluar_pra_headers
	err = db.Raw("UPDATE tbl_trans_uang_keluar_pra_headers SET total_bayar = ?, sisa_biaya = ?, "+
		" edited_on = ? , edited_by = ? "+
		" WHERE kd_trans_keluar = ? and flag_aktif=0 ", sumJmlBayar, sisa_biaya, datenowx, currentUser.(string), kd_trans_keluar).Scan(&dataHeader).Error
	if err != nil {
		response := helper.APIResponse("Update Data Ke Tbl_trans_uang_keluar_pra_headers Gagal ...", http.StatusBadRequest, "error", err)
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
		" FROM tbl_trans_uang_keluar_pra_headers a " +
		" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
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

		sql := " SELECT b.kd_trans_keluar_detail,b.seqno, " +
			" b.jml_bayar,b.keterangan " +
			" FROM tbl_trans_uang_keluar_pra_headers a " +
			" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
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
	db.Raw(" SELECT count(*) jmldata FROM tbl_trans_uang_keluar_pra_headers a "+
		" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar=b.kd_trans_keluar "+
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
	db.Raw("SELECT ifnull(max(kd_trans_keluar_detail),0) + 1 as 'run_number' FROM tbl_trans_uang_keluar_pra_details ").Scan(&intkd_trans_keluar_detail)

	var int_seqno int
	db.Raw("SELECT (seqno + 1) as 'run_number' "+
		" FROM tbl_trans_uang_keluar_pra_details where flag_aktif=0 and kd_trans_keluar=?", paramAddDetail.Kd_trans_keluar).Scan(&int_seqno)

	datadetail := table_data.Tbl_trans_uang_keluar_pra_details{
		Kd_trans_keluar:        paramAddDetail.Kd_trans_keluar,
		Kd_trans_keluar_detail: intkd_trans_keluar_detail,
		Seqno:                  int_seqno,
		Jml_bayar:              0,
		Keterangan:             "",
		Created_by:             currentUser.(string),
		Created_on:             datenowx,
		Flag_aktif:             0,
	}

	err = db.Omit("Edited_on", "Edited_by").Create(&datadetail).Error
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
		" FROM tbl_trans_uang_keluar_pra_headers a " +
		" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
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

		sql := " SELECT b.kd_trans_keluar_detail,b.seqno, " +
			" b.jml_bayar,b.keterangan " +
			" FROM tbl_trans_uang_keluar_pra_headers a " +
			" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
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
		" FROM tbl_trans_uang_keluar_pra_headers a " +
		" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
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

		sql := " SELECT b.kd_trans_keluar_detail,b.seqno, " +
			" b.jml_bayar,b.keterangan " +
			" FROM tbl_trans_uang_keluar_pra_headers a " +
			" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
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

	var dataUtama table_data.Tbl_trans_uang_keluar_pra_details
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

	var dataDetail table_data.Tbl_trans_uang_keluar_pra_details
	err = db.Raw("update tbl_trans_uang_keluar_pra_details set flag_aktif=9,edited_by=?,edited_on=? "+
		" where kd_trans_keluar_detail=? and kd_trans_keluar=? and flag_aktif=0 ",
		currentUser.(string), datenowx, kd_trans_keluar_detail, kd_trans_keluar).Scan(&dataDetail).Error
	if err != nil {
		response := helper.APIResponse("Delete Data Ke Tbl_trans_uang_keluar_pra_details Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var sumJmlBayar float64
	db.Raw("SELECT sum(jml_bayar) 'jml_bayar' FROM tbl_trans_uang_keluar_pra_details "+
		" where flag_aktif=0 and kd_trans_keluar=?", kd_trans_keluar).Scan(&sumJmlBayar)

	var total_biaya float64
	db.Raw("SELECT total_biaya FROM tbl_trans_uang_keluar_pra_headers where flag_aktif=0 and kd_trans_keluar=?", kd_trans_keluar).Scan(&total_biaya)
	var sisa_biaya float64 = total_biaya - sumJmlBayar

	var dataHeader table_data.Tbl_trans_uang_keluar_pra_headers
	err = db.Raw("UPDATE tbl_trans_uang_keluar_pra_headers SET total_bayar = ?, sisa_biaya = ?, "+
		" edited_on = ? , edited_by = ? "+
		" WHERE kd_trans_keluar = ? and flag_aktif=0 ", sumJmlBayar, sisa_biaya, datenowx, currentUser.(string), kd_trans_keluar).Scan(&dataHeader).Error
	if err != nil {
		response := helper.APIResponse("Update Data Ke Tbl_trans_uang_keluar_pra_headers Gagal ...", http.StatusBadRequest, "error", err)
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
		" FROM tbl_trans_uang_keluar_pra_headers a " +
		" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
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

		sql := " SELECT b.kd_trans_keluar_detail,b.seqno, " +
			" b.jml_bayar,b.keterangan " +
			" FROM tbl_trans_uang_keluar_pra_headers a " +
			" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
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

	var dataUtama table_data.Tbl_trans_uang_keluar_pra_headers
	if err := db.Where("flag_aktif=0 and kd_trans_keluar=?", kd_trans_keluar).First(&dataUtama).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Header Tidak Ditemukan ..."}
		response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var dataUtamaDet table_data.Tbl_trans_uang_keluar_pra_details
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

	var dataDetail table_data.Tbl_trans_uang_keluar_pra_details
	err = db.Raw("update tbl_trans_uang_keluar_pra_details set flag_aktif=9,edited_by=?,edited_on=? "+
		" where kd_trans_keluar=? and flag_aktif=0 ",
		currentUser.(string), datenowx, kd_trans_keluar).Scan(&dataDetail).Error
	if err != nil {
		response := helper.APIResponse("Delete Data Ke tbl_trans_uang_keluar_pra_details Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var dataHead table_data.Tbl_trans_uang_keluar_pra_headers
	err = db.Raw("update tbl_trans_uang_keluar_pra_headers set flag_aktif=9,edited_by=?,edited_on=? "+
		" where kd_trans_keluar=? and flag_aktif=0 ",
		currentUser.(string), datenowx, kd_trans_keluar).Scan(&dataHead).Error
	if err != nil {
		response := helper.APIResponse("Delete Data Ke tbl_trans_uang_keluar_pra_headers Gagal ...", http.StatusBadRequest, "error", err)
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
		" FROM tbl_trans_uang_keluar_pra_headers a " +
		" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
		" INNER JOIN tbl_group_kategoris d on a.kd_group = d.kd_group " +
		" INNER JOIN tbl_kategori_uangs e on a.kd_kategori = e.kd_kategori " +
		" where a.flag_aktif=0 and b.flag_aktif=0   "

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

		sql := " SELECT b.kd_trans_keluar_detail,b.seqno, " +
			" b.jml_bayar,b.keterangan " +
			" FROM tbl_trans_uang_keluar_pra_headers a " +
			" INNER JOIN tbl_trans_uang_keluar_pra_details b on a.kd_trans_keluar=b.kd_trans_keluar " +
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
