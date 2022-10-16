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
	db := c.MustGet("db").(*gorm.DB)

	var getIdAndNameKelas []GetIdAndNameKelas
	db.Raw("SELECT DISTINCT * FROM vw_kelas_trans").Scan(&getIdAndNameKelas)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", getIdAndNameKelas)
	c.JSON(http.StatusOK, response)
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

	var paramChangeSiswa ParamChangeSiswa
	if err := c.ShouldBindJSON(&paramChangeSiswa); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//cek data transaksi ppdb
	var intJmldata int
	db.Raw("SELECT count(*) jmldata from tbl_trans_uang_masuk_ppdb_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
		" where a.flag_aktif=0 and b.flag_aktif=0 where nik=?", paramChangeSiswa.Nik).Scan(&intJmldata)
	if intJmldata == 0 {
		//Jika Belum Ada Didalam Transaksi
		SetArrayData := []GetDataPPDB{}
		var tgldaftar string
		var tahun_daftar string
		var total_biaya float64
		var total_bayar float64
		var sisa_biaya float64
		rows, _ := db.Raw(" SELECT tgldaftar,tahun_daftar, "+
			" (SELECT jumlah_pembayaran from tbl_biayadaftar where id=id_biaya) 'total_biaya',0 as 'total_bayar', 0 as 'sisa_biaya' "+
			" FROM tbl_user_ppdb WHERE (nik <> '' or nik is not null)  "+
			" and status = 'sudah diverifikasi' and flag_verifikasidata='1' and flag_wawancara='1' "+
			" and flag_pembayaran='1' and flag=0 and status_berkas <> 'DiCabut' and nik=?", paramChangeSiswa.Nik).Rows()
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&tgldaftar, &tahun_daftar, &total_biaya, &total_bayar, &sisa_biaya)
			arraydata := GetDataPPDB{}

			//tTgldaftar, _ := time.Parse("02-01-2006", tgldaftar)
			// if err != nil {
			// 	errors := helper.FormatValidationError(err)
			// 	errorMessage := gin.H{"errors": errors, "date": tTgldaftar}
			// 	response := helper.APIResponse("Tanggal Format Salah ...", http.StatusUnprocessableEntity, "error", errorMessage)
			// 	c.JSON(http.StatusUnprocessableEntity, response)
			// 	return
			// }
			//dateStr := tTgldaftar.Format("2006-01-02")

			arraydata.Tgldaftar = tgldaftar
			arraydata.Tahun_daftar = tahun_daftar
			arraydata.Total_biaya = total_biaya
			arraydata.Total_bayar = total_bayar
			arraydata.Sisa_biaya = sisa_biaya
			SetArrayData = append(SetArrayData, arraydata)
		}

		response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
		c.JSON(http.StatusOK, response)
	} else {
		//Jika Sudah Ada Didalam Transaksi

	}

	// var getDataUmSpp []GetDataUmSpp
	// db.Raw("SELECT b.kd_trans_masuk_detail,b.seqno,b.periode_bayar, "+
	// 	" b.tgl_bayar,b.jml_bayar,b.keterangan "+
	// 	" FROM tbl_trans_uang_masuk_spp_headers a "+
	// 	" INNER JOIN tbl_trans_uang_masuk_spp_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
	// 	" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis "+
	// 	" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "+
	// 	" and a.tahun_akademik=? and a.nm_kelas=? and a.nis_siswa = ? "+
	// 	" order by b.seqno ", paramChangeSiswa.Tahun_akademik, paramChangeSiswa.Nm_kelas, paramChangeSiswa.Nis_siswa).Scan(&getDataUmSpp)

	// SetArrayData := []GetBiayaAndSisa{}
	// var kd_trans_masuk int
	// var total_biaya float64
	// var total_bayar float64
	// var sisa_biaya float64
	// rows, _ := db.Raw("SELECT distinct b.kd_trans_masuk,a.total_biaya,a.total_bayar,a.sisa_biaya "+
	// 	" FROM tbl_trans_uang_masuk_spp_headers a "+
	// 	" INNER JOIN tbl_trans_uang_masuk_spp_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
	// 	" where a.flag_aktif=0 and b.flag_aktif=0  "+
	// 	" and a.tahun_akademik=? and a.nm_kelas=? and a.nis_siswa = ? ", paramChangeSiswa.Tahun_akademik, paramChangeSiswa.Nm_kelas, paramChangeSiswa.Nis_siswa).Rows()
	// defer rows.Close()
	// for rows.Next() {
	// 	rows.Scan(&kd_trans_masuk, &total_biaya, &total_bayar, &sisa_biaya)
	// 	arraydata := GetBiayaAndSisa{}
	// arraydata.Kd_trans_masuk = kd_trans_masuk
	// arraydata.Total_biaya = total_biaya
	// arraydata.Total_bayar = total_bayar
	// arraydata.Sisa_biaya = sisa_biaya
	// arraydata.Detail = getDataUmSpp
	// SetArrayData = append(SetArrayData, arraydata)
	// }

	// if len(getDataUmSpp) == 0 {
	// 	response := helper.APIResponse("List Data ...", http.StatusOK, "success", getDataUmSpp)
	// 	c.JSON(http.StatusOK, response)
	// 	return
	// }

}

func CreateUangMasukSpp(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var paramInputSPP ParamInputSPP
	if err := c.ShouldBindJSON(&paramInputSPP); err != nil {
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

	//cek data siswa
	var intJmldata int
	db.Raw("SELECT count(*) jmldata FROM tbl_siswa where flag_siswa=0 and status_siswa not in('Tidak Aktif','LULUS') and nis=? "+
		" and (tahun_aktif=? or tahun_aktif = REPLACE(?,'-','/')) "+
		" and REPLACE(REPLACE(REPLACE(nm_kelas,'MIA',''),'IIS',''),' ','') = ?", paramInputSPP.Nis_siswa, paramInputSPP.Tahun_akademik, paramInputSPP.Tahun_akademik, paramInputSPP.Nm_kelas).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Data Siswa Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var CekDataUangMasuk table_data.Tbl_trans_uang_masuk_spp_headers
	checkUser := db.Select("*").Where("flag_aktif = 0 and kd_group= ? and kd_kategori= ? and nis_siswa=? and nm_kelas=? and tahun_akademik = ?", paramInputSPP.Kd_group, paramInputSPP.Kd_kategori, paramInputSPP.Nis_siswa, paramInputSPP.Nm_kelas, paramInputSPP.Tahun_akademik).Find(&CekDataUangMasuk)
	if checkUser.RowsAffected > 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Data Pembayaran SPP Sudah Ada ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	} else {

		//Jika data spp belum dibuat
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
		db.Raw("SELECT ifnull(max(kd_trans_masuk),0) + 1 as 'run_number' FROM tbl_trans_uang_masuk_spp_headers where flag_aktif=0").Scan(&intKd_trans_masuk)

		var float_biaya_spp float64
		db.Raw("SELECT sum(biaya_spp) 'biaya_spp' FROM tbl_conf_periode_spps where tahun_akademik=? and nm_kelas=? and flag_aktif=0", paramInputSPP.Tahun_akademik, paramInputSPP.Nm_kelas).Scan(&float_biaya_spp)

		currentUser := c.MustGet("currentUser")
		data := table_data.Tbl_trans_uang_masuk_spp_headers{
			Kd_group:       paramInputSPP.Kd_group,
			Kd_kategori:    paramInputSPP.Kd_kategori,
			Kd_trans_masuk: intKd_trans_masuk,
			Nis_siswa:      paramInputSPP.Nis_siswa,
			Nm_kelas:       paramInputSPP.Nm_kelas,
			Tahun_akademik: paramInputSPP.Tahun_akademik,
			Total_biaya:    float_biaya_spp,
			Total_bayar:    0,
			Sisa_biaya:     float_biaya_spp,
			Keterangan:     paramInputSPP.Keterangan,
			Created_by:     currentUser.(string),
			Created_on:     datenowx,
			Flag_aktif:     0,
		}

		err = db.Omit("Edited_on", "Edited_by").Create(&data).Error
		if err != nil {
			response := helper.APIResponse("Simpan Data Gagal ...", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		var intKd_trans_masuk_detail int
		db.Raw("SELECT ifnull(max(kd_trans_masuk_detail),0) + 1 as 'run_number' FROM tbl_trans_uang_masuk_spp_details where flag_aktif=0").Scan(&intKd_trans_masuk_detail)

		var int_seqno int
		var periodebayar string
		var jmlbayar float64
		rows, _ := db.Raw("SELECT seqno,CONCAT(kd_bulan,'-',tahun) 'periodebayar',biaya_spp FROM tbl_conf_periode_spps WHERE flag_aktif=0 and tahun_akademik=? and nm_kelas=? ORDER BY seqno", paramInputSPP.Tahun_akademik, paramInputSPP.Nm_kelas).Rows()
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&int_seqno, &periodebayar, &jmlbayar)

			datadetail := table_data.Tbl_trans_uang_masuk_spp_details{
				Kd_trans_masuk:        intKd_trans_masuk,
				Kd_trans_masuk_detail: intKd_trans_masuk_detail,
				Seqno:                 int_seqno,
				Periode_bayar:         periodebayar,
				Jml_bayar:             jmlbayar,
				Keterangan:            "",
				Created_by:            currentUser.(string),
				Created_on:            datenowx,
				Flag_aktif:            0,
			}

			err = db.Omit("Edited_on", "Edited_by", "Tgl_bayar").Create(&datadetail).Error
			if err != nil {
				response := helper.APIResponse("Simpan Data Detail Gagal ...", http.StatusBadRequest, "error", err)
				c.JSON(http.StatusBadRequest, response)
				return
			}

			intKd_trans_masuk_detail++
		}

		//setting tampilan habis save spp

		var getDataUmSpp []GetDataUmSpp
		db.Raw("SELECT b.kd_trans_masuk_detail,b.seqno,b.periode_bayar, "+
			" b.tgl_bayar,b.jml_bayar,b.keterangan "+
			" FROM tbl_trans_uang_masuk_spp_headers a "+
			" INNER JOIN tbl_trans_uang_masuk_spp_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
			" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis "+
			" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "+
			" and a.tahun_akademik=? and a.nm_kelas=? and a.nis_siswa = ? "+
			" order by b.seqno ", paramInputSPP.Tahun_akademik, paramInputSPP.Nm_kelas, paramInputSPP.Nis_siswa).Scan(&getDataUmSpp)

		SetArrayData := []GetBiayaAndSisa{}
		var kd_trans_masuk int
		var total_biaya float64
		var total_bayar float64
		var sisa_biaya float64
		rowss, _ := db.Raw("SELECT distinct b.kd_trans_masuk,a.total_biaya,a.total_bayar,a.sisa_biaya "+
			" FROM tbl_trans_uang_masuk_spp_headers a "+
			" INNER JOIN tbl_trans_uang_masuk_spp_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
			" where a.flag_aktif=0 and b.flag_aktif=0  "+
			" and a.tahun_akademik=? and a.nm_kelas=? and a.nis_siswa = ? ", paramInputSPP.Tahun_akademik, paramInputSPP.Nm_kelas, paramInputSPP.Nis_siswa).Rows()
		defer rowss.Close()
		for rowss.Next() {
			rowss.Scan(&kd_trans_masuk, &total_biaya, &total_bayar, &sisa_biaya)
			arraydata := GetBiayaAndSisa{}
			arraydata.Kd_trans_masuk = kd_trans_masuk
			arraydata.Total_biaya = total_biaya
			arraydata.Total_bayar = total_bayar
			arraydata.Sisa_biaya = sisa_biaya
			arraydata.Detail = getDataUmSpp
			SetArrayData = append(SetArrayData, arraydata)
		}

		response := helper.APIResponse("Simpan Data Sukses ...", http.StatusOK, "success", SetArrayData)
		c.JSON(http.StatusOK, response)
	}

}

func UpdateUangMasukSpp(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramEditSPPDetail ParamEditSPPDetail
	if err := c.ShouldBindJSON(&paramEditSPPDetail); err != nil {
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

	var dataUtama table_data.Tbl_trans_uang_masuk_spp_details
	if err := db.Where("flag_aktif=0 and kd_trans_masuk_detail=?", c.Param("iddetail")).First(&dataUtama).Error; err != nil {
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

	tTglBayar, err2 := time.Parse("02-01-2006", paramEditSPPDetail.Tgl_bayar)
	if err2 != nil {
		errors := helper.FormatValidationError(err2)
		errorMessage := gin.H{"errors": errors, "date": tTglBayar}
		response := helper.APIResponse("Tanggal Format Salah ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	dateStr := tTglBayar.Format("2006-01-02")

	var dataDetail table_data.Tbl_trans_uang_masuk_spp_details
	err = db.Raw("update tbl_trans_uang_masuk_spp_details set tgl_bayar=?,jml_bayar=?,keterangan=?,edited_by=?,edited_on=? "+
		" where kd_trans_masuk_detail=? and flag_aktif=0 ", dateStr, paramEditSPPDetail.Jml_bayar, paramEditSPPDetail.Keterangan, currentUser.(string), datenowx, c.Param("iddetail")).Scan(&dataDetail).Error
	if err != nil {
		response := helper.APIResponse("Update Data Ke Tbl_trans_uang_masuk_spp_headers Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var sumJmlBayar float64
	db.Raw("SELECT sum(jml_bayar) 'jml_bayar' FROM tbl_trans_uang_masuk_spp_details "+
		" where flag_aktif=0 and tgl_bayar is not null and kd_trans_masuk=?", c.Param("idhead")).Scan(&sumJmlBayar)

	var total_biaya float64
	db.Raw("SELECT total_biaya FROM tbl_trans_uang_masuk_spp_headers where flag_aktif=0 and kd_trans_masuk=?", c.Param("idhead")).Scan(&total_biaya)
	var sisa_biaya float64 = total_biaya - sumJmlBayar

	var dataHeader table_data.Tbl_trans_uang_masuk_spp_headers
	err = db.Raw("UPDATE tbl_trans_uang_masuk_spp_headers SET total_bayar = ?, sisa_biaya = ?, "+
		" edited_on = ? , edited_by = ? "+
		" WHERE kd_trans_masuk = ? and flag_aktif=0 ", sumJmlBayar, sisa_biaya, datenowx, currentUser.(string), c.Param("idhead")).Scan(&dataHeader).Error
	if err != nil {
		response := helper.APIResponse("Update Data Ke Tbl_trans_uang_masuk_spp_headers Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//setting tampilan habis update spp

	var getDataUmSpp []GetDataUmSpp
	db.Raw("SELECT b.kd_trans_masuk_detail,b.seqno,b.periode_bayar, "+
		" b.tgl_bayar,b.jml_bayar,b.keterangan "+
		" FROM tbl_trans_uang_masuk_spp_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_spp_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
		" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis "+
		" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "+
		" and a.kd_trans_masuk=? "+
		" order by b.seqno ", c.Param("idhead")).Scan(&getDataUmSpp)

	SetArrayData := []GetBiayaAndSisa{}
	var kd_trans_masuk int
	//var total_biaya float64
	var total_bayar float64
	//var sisa_biaya float64
	rowss, _ := db.Raw("SELECT distinct b.kd_trans_masuk,a.total_biaya,a.total_bayar,a.sisa_biaya "+
		" FROM tbl_trans_uang_masuk_spp_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_spp_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
		" where a.flag_aktif=0 and b.flag_aktif=0  "+
		" and a.kd_trans_masuk=? ", c.Param("idhead")).Rows()
	defer rowss.Close()
	for rowss.Next() {
		rowss.Scan(&kd_trans_masuk, &total_biaya, &total_bayar, &sisa_biaya)
		arraydata := GetBiayaAndSisa{}
		arraydata.Kd_trans_masuk = kd_trans_masuk
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Detail = getDataUmSpp
		SetArrayData = append(SetArrayData, arraydata)
	}

	response := helper.APIResponse("Update Data Sukses ...", http.StatusOK, "success", SetArrayData)
	c.JSON(http.StatusOK, response)

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
