package transaksi_uang_masuk_ppdb

import (
	"errors"
	"net/http"
	"rest_api_bendahara/helper"
	"rest_api_bendahara/master_group_kategori"
	"rest_api_bendahara/master_kategori_uang"
	"rest_api_bendahara/table_data"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ListGroupKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var result_kd_group string
	rows, _ := db.Raw("SELECT kd_group FROM tbl_link_kategoris where link_name=? ", "form_biaya_ppdb").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&result_kd_group)
	}

	var master []master_group_kategori.ListData
	sql := "  SELECT a.*,b.proses_uang FROM tbl_group_kategoris as a " +
		" inner join tbl_jenis_trans as b on a.kd_jenis=b.kd_jenis  " +
		" where a.flag_aktif=0 and b.flag_aktif=0  " +
		" and a.kd_group in(" + result_kd_group + ") "

	db.Raw(sql).Scan(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", master_group_kategori.FormatGroupKategori(master))
	c.JSON(http.StatusOK, response)
}

func ListKategoriUang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var result_kd_kategori string
	rows, _ := db.Raw("SELECT kd_kategori FROM tbl_link_kategoris where link_name=? ", "form_biaya_ppdb").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&result_kd_kategori)
	}

	var master []master_kategori_uang.ListData
	sql := " SELECT a.*,b.nm_group FROM tbl_kategori_uangs as a  " +
		" INNER JOIN tbl_group_kategoris b on a.kd_group=b.kd_group " +
		" where a.flag_aktif=0 and b.flag_aktif=0 " +
		" and a.kd_kategori in (" + result_kd_kategori + ") "

	db.Raw(sql).Scan(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", master_kategori_uang.FormatShowData(master))
	c.JSON(http.StatusOK, response)
}

func ListKelas(c *gin.Context) {
	// db := c.MustGet("db").(*gorm.DB)

	// var getIdAndNameKelas []GetIdAndNameKelas
	// db.Raw("SELECT DISTINCT * FROM vw_kelas_trans").Scan(&getIdAndNameKelas)

	// response := helper.APIResponse("List Data ...", http.StatusOK, "success", getIdAndNameKelas)
	// c.JSON(http.StatusOK, response)
}

func ListSiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var getNikAndName []GetNikAndNameSiswa
	db.Raw("SELECT nik,concat('Double Data ==> ','',nm_siswa) 'nm_siswa' FROM tbl_user_ppdb " +
		" WHERE (nik <> '' or nik is not null) " +
		" and status = 'sudah diverifikasi' and flag_verifikasidata='1' and flag_wawancara='1' " +
		" and flag_pembayaran='1' and flag=0 and status_berkas <> 'DiCabut' " +
		" GROUP BY nik " +
		" HAVING count(*) > 1 ").Scan(&getNikAndName)

	if len(getNikAndName) > 0 {
		response := helper.APIResponse("Terdapat Data Siswa Yang Double ...", http.StatusUnprocessableEntity, "error", getNikAndName)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var getNikAndNameSiswa []GetNikAndNameSiswa
	db.Raw(" SELECT nik,nm_siswa FROM tbl_user_ppdb WHERE (nik <> '' or nik is not null)  " +
		" and status = 'sudah diverifikasi' and flag_verifikasidata='1' and flag_wawancara='1' " +
		" and flag_pembayaran='1' and flag=0 and status_berkas <> 'DiCabut' ").Scan(&getNikAndNameSiswa)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", getNikAndNameSiswa)
	c.JSON(http.StatusOK, response)
}

func ListData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramChangeNik ParamChangeNik
	if err := c.ShouldBindJSON(&paramChangeNik); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var getDataPPDB []GetDataPPDB
	db.Raw("SELECT distinct b.kd_trans_masuk_detail_ppdb,b.seqno,b.kategori_biaya_ppdb, "+
		" b.tgl_bayar,b.jml_bayar,b.keterangan "+
		" FROM tbl_trans_uang_masuk_ppdb_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
		" INNER JOIN tbl_user_ppdb c on a.nik = c.nik "+
		" where a.flag_aktif=0 and b.flag_aktif=0 and c.status = 'sudah diverifikasi' and c.flag_verifikasidata='1' and c.flag_wawancara='1' "+
		" and c.flag_pembayaran='1' and c.flag=0 and status_berkas <> 'DiCabut' and a.nik=? "+
		" order by b.seqno ", paramChangeNik.Nik).Scan(&getDataPPDB)

	SetArrayData := []GetBiayaAndSisa{}
	var kd_trans_masuk_ppdb int
	var total_biaya float64
	var total_bayar float64
	var sisa_biaya float64
	rows, _ := db.Raw("SELECT distinct b.kd_trans_masuk_ppdb,a.total_biaya,a.total_bayar,a.sisa_biaya "+
		" FROM tbl_trans_uang_masuk_ppdb_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
		" where a.flag_aktif=0 and b.flag_aktif=0  "+
		" and a.nik=? ", paramChangeNik.Nik).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&kd_trans_masuk_ppdb, &total_biaya, &total_bayar, &sisa_biaya)
		arraydata := GetBiayaAndSisa{}
		arraydata.Kd_trans_masuk_ppdb = kd_trans_masuk_ppdb
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Detail = getDataPPDB
		SetArrayData = append(SetArrayData, arraydata)
	}

	if len(getDataPPDB) == 0 {
		response := helper.APIResponse("List Data ...", http.StatusOK, "success", getDataPPDB)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
	c.JSON(http.StatusOK, response)
}

func CreateUangMasukPPdb(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var paramInputPPdb ParamInputPPdb
	if err := c.ShouldBindJSON(&paramInputPPdb); err != nil {
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

	tTglDaftar, err := time.Parse("02-01-2006", paramInputPPdb.Tgldaftar)
	if err != nil {
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
	dateStrTglDaftar := tTglDaftar.Format("2006-01-02")

	//cek data siswa ppdb
	var intJmldata int
	db.Raw("SELECT count(*) jmldata "+
		" FROM tbl_user_ppdb WHERE nik = ? and tgldaftar=? and tahun_daftar=? "+
		" and status = 'sudah diverifikasi' and flag_verifikasidata='1' and flag_wawancara='1' "+
		" and flag_pembayaran='1' and flag=0 and status_berkas <> 'DiCabut'", paramInputPPdb.Nik, dateStrTglDaftar, paramInputPPdb.Tahun_daftar).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Data Siswa Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//cek data uang masuk ppdb
	db.Raw("SELECT count(*) jmldata from tbl_trans_uang_masuk_ppdb_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
		" where a.flag_aktif=0 and b.flag_aktif=0 and nik=? and kd_group=? "+
		" and kd_kategori=? and tahun_daftar=? and tgldaftar=? ", paramInputPPdb.Nik, paramInputPPdb.Kd_group, paramInputPPdb.Kd_kategori, paramInputPPdb.Tahun_daftar, dateStrTglDaftar).Scan(&intJmldata)

	if intJmldata > 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Data Pembayaran PPdb Sudah Ada ...", http.StatusUnprocessableEntity, "error", errorMessage)
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

		var intKd_trans_masuk int
		db.Raw("SELECT ifnull(max(kd_trans_masuk_ppdb),0) + 1 as 'run_number' FROM tbl_trans_uang_masuk_ppdb_headers where flag_aktif=0").Scan(&intKd_trans_masuk)

		var tgldaftar string
		var tahun_daftar string
		var total_biaya float64
		var total_bayar float64
		rows, _ := db.Raw(" SELECT CONVERT(tgldaftar,CHAR) 'tgldaftar',tahun_daftar, "+
			" (SELECT REPLACE(jumlah_pembayaran,'.','') from tbl_biayadaftar where id=id_biaya) 'total_biaya',0 as 'total_bayar' "+
			" FROM tbl_user_ppdb WHERE (nik <> '' or nik is not null)  "+
			" and status = 'sudah diverifikasi' and flag_verifikasidata='1' and flag_wawancara='1' "+
			" and flag_pembayaran='1' and flag=0 and status_berkas <> 'DiCabut' and nik=?", paramInputPPdb.Nik).Rows()
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&tgldaftar, &tahun_daftar, &total_biaya, &total_bayar)
		}

		currentUser := c.MustGet("currentUser")
		data := table_data.Tbl_trans_uang_masuk_ppdb_headers{
			Kd_group:            paramInputPPdb.Kd_group,
			Kd_kategori:         paramInputPPdb.Kd_kategori,
			Kd_trans_masuk_ppdb: intKd_trans_masuk,
			Nik:                 paramInputPPdb.Nik,
			Tgldaftar:           dateStrTglDaftar,
			Tahun_daftar:        paramInputPPdb.Tahun_daftar,
			Tahun_akademik:      paramInputPPdb.Tahun_akademik,
			Total_biaya:         total_biaya,
			Total_bayar:         0,
			Sisa_biaya:          total_biaya,
			Keterangan:          paramInputPPdb.Keterangan,
			Created_by:          currentUser.(string),
			Created_on:          datenowx,
			Flag_aktif:          0,
		}

		err = db.Omit("Edited_on", "Edited_by").Create(&data).Error
		if err != nil {
			response := helper.APIResponse("Simpan Data Gagal ...", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		var intKd_trans_masuk_detail int
		db.Raw("SELECT ifnull(max(kd_trans_masuk_detail_ppdb),0) + 1 as 'run_number' FROM tbl_trans_uang_masuk_ppdb_details where flag_aktif=0").Scan(&intKd_trans_masuk_detail)

		var int_seqno int = 1
		var nm_sub_kategori string
		var jmlbayar float64
		rows1, _ := db.Raw("SELECT nm_sub_kategori FROM tbl_sub_kategori_uangs where flag_aktif=0 and kd_kategori=? order by kd_sub_kategori", paramInputPPdb.Kd_kategori).Rows()
		defer rows1.Close()
		for rows1.Next() {
			rows1.Scan(&nm_sub_kategori)

			datadetail := table_data.Tbl_trans_uang_masuk_ppdb_details{
				Kd_trans_masuk_ppdb:        intKd_trans_masuk,
				Kd_trans_masuk_detail_ppdb: intKd_trans_masuk_detail,
				Seqno:                      int_seqno,
				Kategori_biaya_ppdb:        nm_sub_kategori,
				Jml_bayar:                  jmlbayar,
				Keterangan:                 "",
				Created_by:                 currentUser.(string),
				Created_on:                 datenowx,
				Flag_aktif:                 0,
			}

			err = db.Omit("Edited_on", "Edited_by", "Tgl_bayar").Create(&datadetail).Error
			if err != nil {
				response := helper.APIResponse("Simpan Data Detail Gagal ...", http.StatusBadRequest, "error", err)
				c.JSON(http.StatusBadRequest, response)
				return
			}

			intKd_trans_masuk_detail++
			int_seqno++
		}
	}

	//setting tampilan habis save spp

	var getDataPPDB []GetDataPPDB
	db.Raw(" SELECT DISTINCT b.kd_trans_masuk_detail_ppdb,b.seqno,b.kategori_biaya_ppdb, "+
		" b.tgl_bayar,b.jml_bayar,b.keterangan "+
		" FROM tbl_trans_uang_masuk_ppdb_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
		" INNER JOIN tbl_user_ppdb c on a.nik=c.nik and a.tgldaftar=c.tgldaftar and a.tahun_daftar=c.tahun_daftar "+
		" where a.flag_aktif=0 and b.flag_aktif=0 and a.nik=? and a.tahun_daftar=? and a.tgldaftar=? "+
		" ORDER BY seqno ", paramInputPPdb.Nik, paramInputPPdb.Tahun_daftar, dateStrTglDaftar).Scan(&getDataPPDB)

	SetArrayData := []GetBiayaAndSisa{}
	var kd_trans_masuk_ppdb int
	var total_biaya float64
	var total_bayar float64
	var sisa_biaya float64
	rowss, _ := db.Raw("SELECT distinct b.kd_trans_masuk_ppdb,a.total_biaya,a.total_bayar,a.sisa_biaya "+
		" FROM tbl_trans_uang_masuk_ppdb_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
		" where a.flag_aktif=0 and b.flag_aktif=0  "+
		" and a.nik=? and a.tahun_daftar=? and a.tgldaftar=? ", paramInputPPdb.Nik, paramInputPPdb.Tahun_daftar, dateStrTglDaftar).Rows()
	defer rowss.Close()
	for rowss.Next() {
		rowss.Scan(&kd_trans_masuk_ppdb, &total_biaya, &total_bayar, &sisa_biaya)
		arraydata := GetBiayaAndSisa{}
		arraydata.Kd_trans_masuk_ppdb = kd_trans_masuk_ppdb
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Detail = getDataPPDB
		SetArrayData = append(SetArrayData, arraydata)
	}

	response := helper.APIResponse("Simpan Data Sukses ...", http.StatusOK, "success", SetArrayData)
	c.JSON(http.StatusOK, response)

}

func UpdateUangMasukSpp(c *gin.Context) {
	// db := c.MustGet("db").(*gorm.DB)

	// var paramEditSPPDetail ParamEditSPPDetail
	// if err := c.ShouldBindJSON(&paramEditSPPDetail); err != nil {
	// 	var ve validator.ValidationErrors
	// 	if errors.As(err, &ve) {
	// 		errors := helper.FormatValidationError(err)
	// 		errorMessage := gin.H{"errors": errors}
	// 		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
	// 		c.JSON(http.StatusUnprocessableEntity, response)
	// 		return
	// 	}
	// 	var error_binding []string
	// 	error_binding = append(error_binding, err.Error())
	// 	errorMessage := gin.H{"errors": error_binding}
	// 	response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
	// 	c.JSON(http.StatusUnprocessableEntity, response)
	// 	return
	// }

	// var dataUtama table_data.Tbl_trans_uang_masuk_spp_details
	// if err := db.Where("flag_aktif=0 and kd_trans_masuk_detail=?", c.Param("iddetail")).First(&dataUtama).Error; err != nil {
	// 	errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
	// 	response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
	// 	c.JSON(http.StatusUnprocessableEntity, response)
	// 	return
	// }

	// currentUser := c.MustGet("currentUser")

	// var datenows string = time.Now().UTC().Format("2006-01-02 15:04:05")
	// date := "2006-01-02 15:04:05"
	// datenowx, err := time.Parse(date, datenows)
	// if err != nil {
	// 	errors := helper.FormatValidationError(err)
	// 	errorMessage := gin.H{"errors": errors, "date": datenowx}
	// 	response := helper.APIResponse("Tanggal Format Salah ...", http.StatusUnprocessableEntity, "error", errorMessage)
	// 	c.JSON(http.StatusUnprocessableEntity, response)
	// 	return
	// }

	// tTglBayar, err2 := time.Parse("02-01-2006", paramEditSPPDetail.Tgl_bayar)
	// if err2 != nil {
	// 	errors := helper.FormatValidationError(err2)
	// 	errorMessage := gin.H{"errors": errors, "date": tTglBayar}
	// 	response := helper.APIResponse("Tanggal Format Salah ...", http.StatusUnprocessableEntity, "error", errorMessage)
	// 	c.JSON(http.StatusUnprocessableEntity, response)
	// 	return
	// }
	// dateStr := tTglBayar.Format("2006-01-02")

	// var dataDetail table_data.Tbl_trans_uang_masuk_spp_details
	// err = db.Raw("update tbl_trans_uang_masuk_spp_details set tgl_bayar=?,jml_bayar=?,keterangan=?,edited_by=?,edited_on=? "+
	// 	" where kd_trans_masuk_detail=? and flag_aktif=0 ", dateStr, paramEditSPPDetail.Jml_bayar, paramEditSPPDetail.Keterangan, currentUser.(string), datenowx, c.Param("iddetail")).Scan(&dataDetail).Error
	// if err != nil {
	// 	response := helper.APIResponse("Update Data Ke Tbl_trans_uang_masuk_spp_headers Gagal ...", http.StatusBadRequest, "error", err)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// var sumJmlBayar float64
	// db.Raw("SELECT sum(jml_bayar) 'jml_bayar' FROM tbl_trans_uang_masuk_spp_details "+
	// 	" where flag_aktif=0 and tgl_bayar is not null and kd_trans_masuk=?", c.Param("idhead")).Scan(&sumJmlBayar)

	// var total_biaya float64
	// db.Raw("SELECT total_biaya FROM tbl_trans_uang_masuk_spp_headers where flag_aktif=0 and kd_trans_masuk=?", c.Param("idhead")).Scan(&total_biaya)
	// var sisa_biaya float64 = total_biaya - sumJmlBayar

	// var dataHeader table_data.Tbl_trans_uang_masuk_spp_headers
	// err = db.Raw("UPDATE tbl_trans_uang_masuk_spp_headers SET total_bayar = ?, sisa_biaya = ?, "+
	// 	" edited_on = ? , edited_by = ? "+
	// 	" WHERE kd_trans_masuk = ? and flag_aktif=0 ", sumJmlBayar, sisa_biaya, datenowx, currentUser.(string), c.Param("idhead")).Scan(&dataHeader).Error
	// if err != nil {
	// 	response := helper.APIResponse("Update Data Ke Tbl_trans_uang_masuk_spp_headers Gagal ...", http.StatusBadRequest, "error", err)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// //setting tampilan habis update spp

	// var getDataUmSpp []GetDataUmSpp
	// db.Raw("SELECT b.kd_trans_masuk_detail,b.seqno,b.periode_bayar, "+
	// 	" b.tgl_bayar,b.jml_bayar,b.keterangan "+
	// 	" FROM tbl_trans_uang_masuk_spp_headers a "+
	// 	" INNER JOIN tbl_trans_uang_masuk_spp_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
	// 	" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis "+
	// 	" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "+
	// 	" and a.kd_trans_masuk=? "+
	// 	" order by b.seqno ", c.Param("idhead")).Scan(&getDataUmSpp)

	// SetArrayData := []GetBiayaAndSisa{}
	// var kd_trans_masuk int
	// //var total_biaya float64
	// var total_bayar float64
	// //var sisa_biaya float64
	// rowss, _ := db.Raw("SELECT distinct b.kd_trans_masuk,a.total_biaya,a.total_bayar,a.sisa_biaya "+
	// 	" FROM tbl_trans_uang_masuk_spp_headers a "+
	// 	" INNER JOIN tbl_trans_uang_masuk_spp_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
	// 	" where a.flag_aktif=0 and b.flag_aktif=0  "+
	// 	" and a.kd_trans_masuk=? ", c.Param("idhead")).Rows()
	// defer rowss.Close()
	// for rowss.Next() {
	// 	rowss.Scan(&kd_trans_masuk, &total_biaya, &total_bayar, &sisa_biaya)
	// 	arraydata := GetBiayaAndSisa{}
	// 	arraydata.Kd_trans_masuk = kd_trans_masuk
	// 	arraydata.Total_biaya = total_biaya
	// 	arraydata.Total_bayar = total_bayar
	// 	arraydata.Sisa_biaya = sisa_biaya
	// 	arraydata.Detail = getDataUmSpp
	// 	SetArrayData = append(SetArrayData, arraydata)
	// }

	// response := helper.APIResponse("Update Data Sukses ...", http.StatusOK, "success", SetArrayData)
	// c.JSON(http.StatusOK, response)

}

// func LinkKdGroup(input GetLinkKdGroup) (string, error) {
// 	return "", nil
// }

// func LinkKdKategori(input GetLinkKdKategori) (string, error) {
// 	return "", nil
// }

// var result DataTokenInput
// 	db.Raw("SELECT Id_user,Password,Username FROM tbl_users WHERE Username = ?", dataInput.Username).Scan(&result)
// token, err := GenerateToken(result)
// 	if err != nil {
// 		response := helper.APIResponse("Generate Token Gagal ...", http.StatusBadRequest, "error", nil)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}
