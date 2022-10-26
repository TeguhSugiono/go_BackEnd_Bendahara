package transaksi_uang_masuk_siswa

import (
	"net/http"
	"rest_api_bendahara/helper"
	"rest_api_bendahara/master_group_kategori"
	"rest_api_bendahara/master_kategori_uang"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListGroupKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	nomor := 0
	var result_kd_group string
	var str_kd_group string

	//kd_jenis=1 adalah uang masuk
	rows, _ := db.Raw("SELECT kd_group FROM tbl_group_kategoris where flag_aktif=0 and kd_jenis=1 ").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&result_kd_group)
		if nomor == 0 {
			str_kd_group = result_kd_group
		} else {
			str_kd_group = str_kd_group + "," + result_kd_group
		}
		nomor++
	}

	var master []master_group_kategori.ListData
	sql := "  SELECT a.*,b.proses_uang FROM tbl_group_kategoris as a " +
		" inner join tbl_jenis_trans as b on a.kd_jenis=b.kd_jenis  " +
		" where a.flag_aktif=0 and b.flag_aktif=0  " +
		" and a.kd_group in(" + str_kd_group + ")  order by b.proses_uang "

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

func ListDataAddSiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramGetSiswaAdd ParamGetSiswaAdd
	if err := c.ShouldBindJSON(&paramGetSiswaAdd); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var master []ListData
	sql := " SELECT DISTINCT id_tahun_aktif 'id_tahun',tahun_aktif 'tahun_akademik', " +
		" REPLACE(REPLACE(REPLACE(nm_kelas,'MIA',''),'IIS',''),' ','') as 'id_kelas', " +
		" REPLACE(REPLACE(REPLACE(nm_kelas,'MIA',''),'IIS',''),' ','') as 'nm_kelas' " +
		" FROM tbl_siswa " +
		" WHERE flag_siswa = 0 AND status_siswa NOT IN ('Tidak Aktif','LULUS') and nis='" + paramGetSiswaAdd.Nis + "' "
	db.Raw(sql).Scan(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", master)
	c.JSON(http.StatusOK, response)
}

// func ListSiswa(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)

// 	var paramChangeNmKelas ParamChangeNmKelas
// 	if err := c.ShouldBindJSON(&paramChangeNmKelas); err != nil {
// 		errors := helper.FormatValidationError(err)
// 		errorMessage := gin.H{"errors": errors}
// 		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	var getNisAndNameSiswa []GetNisAndNameSiswa
// 	db.Raw("SELECT DISTINCT a.nis,a.nm_siswa FROM tbl_siswa a  "+
// 		" LEFT JOIN Tbl_trans_uang_masuk_siswa_headers b on a.nis = b.nis_siswa "+
// 		" and b.flag_aktif=0 and b.tahun_akademik=? and b.nm_kelas=? "+
// 		" where a.flag_siswa = 0 AND a.status_siswa NOT IN ('Tidak Aktif') "+
// 		" and (a.tahun_aktif = ? or a.tahun_aktif = REPLACE(?,'-','/')) "+
// 		" and REPLACE(REPLACE(a.nm_kelas,'MIA',''),'IIS','') = ? "+
// 		" ORDER BY a.nm_siswa ", paramChangeNmKelas.Tahun_akademik, paramChangeNmKelas.Nm_kelas, paramChangeNmKelas.Tahun_akademik, paramChangeNmKelas.Tahun_akademik, paramChangeNmKelas.Nm_kelas).Scan(&getNisAndNameSiswa)

// 	response := helper.APIResponse("List Data ...", http.StatusOK, "success", getNisAndNameSiswa)
// 	c.JSON(http.StatusOK, response)
// }

// func ListKelas(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)

// 	var getIdAndNameKelas []GetIdAndNameKelas
// 	db.Raw("SELECT DISTINCT * FROM vw_kelas_trans").Scan(&getIdAndNameKelas)

// 	response := helper.APIResponse("List Data ...", http.StatusOK, "success", getIdAndNameKelas)
// 	c.JSON(http.StatusOK, response)
// }

// func ListData(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)

// 	var paramChangeSiswa ParamChangeSiswa
// 	if err := c.ShouldBindJSON(&paramChangeSiswa); err != nil {
// 		errors := helper.FormatValidationError(err)
// 		errorMessage := gin.H{"errors": errors}
// 		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	var GetDataUmSiswa []GetDataUmSiswa
// 	db.Raw("SELECT b.kd_trans_masuk_detail,b.seqno,b.periode_bayar, "+
// 		" b.tgl_bayar,b.jml_bayar,b.keterangan "+
// 		" FROM Tbl_trans_uang_masuk_siswa_headers a "+
// 		" INNER JOIN Tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
// 		" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis "+
// 		" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "+
// 		" and a.tahun_akademik=? and a.nm_kelas=? and a.nis_siswa = ? "+
// 		" order by b.seqno ", paramChangeSiswa.Tahun_akademik, paramChangeSiswa.Nm_kelas, paramChangeSiswa.Nis_siswa).Scan(&GetDataUmSiswa)

// 	SetArrayData := []GetBiayaAndSisa{}
// 	var kd_trans_masuk int
// 	var total_biaya float64
// 	var total_bayar float64
// 	var sisa_biaya float64
// 	rows, _ := db.Raw("SELECT distinct b.kd_trans_masuk,a.total_biaya,a.total_bayar,a.sisa_biaya "+
// 		" FROM Tbl_trans_uang_masuk_siswa_headers a "+
// 		" INNER JOIN Tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
// 		" where a.flag_aktif=0 and b.flag_aktif=0  "+
// 		" and a.tahun_akademik=? and a.nm_kelas=? and a.nis_siswa = ? ", paramChangeSiswa.Tahun_akademik, paramChangeSiswa.Nm_kelas, paramChangeSiswa.Nis_siswa).Rows()
// 	defer rows.Close()
// 	for rows.Next() {
// 		rows.Scan(&kd_trans_masuk, &total_biaya, &total_bayar, &sisa_biaya)
// 		arraydata := GetBiayaAndSisa{}
// 		arraydata.Kd_trans_masuk = kd_trans_masuk
// 		arraydata.Total_biaya = total_biaya
// 		arraydata.Total_bayar = total_bayar
// 		arraydata.Sisa_biaya = sisa_biaya
// 		arraydata.Detail = GetDataUmSiswa
// 		SetArrayData = append(SetArrayData, arraydata)
// 	}

// 	if len(GetDataUmSiswa) == 0 {
// 		response := helper.APIResponse("List Data ...", http.StatusOK, "success", GetDataUmSiswa)
// 		c.JSON(http.StatusOK, response)
// 		return
// 	}

// 	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
// 	c.JSON(http.StatusOK, response)
// }

// func CreateUangMasukSpp(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)
// 	var ParamInputSiswa ParamInputSiswa
// 	if err := c.ShouldBindJSON(&ParamInputSiswa); err != nil {
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

// 	//cek data siswa
// 	var intJmldata int
// 	db.Raw("SELECT count(*) jmldata FROM tbl_siswa where flag_siswa=0 and status_siswa not in('Tidak Aktif','LULUS') and nis=? "+
// 		" and (tahun_aktif=? or tahun_aktif = REPLACE(?,'-','/')) "+
// 		" and REPLACE(REPLACE(REPLACE(nm_kelas,'MIA',''),'IIS',''),' ','') = ?", ParamInputSiswa.Nis_siswa, ParamInputSiswa.Tahun_akademik, ParamInputSiswa.Tahun_akademik, ParamInputSiswa.Nm_kelas).Scan(&intJmldata)
// 	if intJmldata == 0 {
// 		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
// 		response := helper.APIResponse("Data Siswa Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	db.Raw("SELECT count(*) jmldata FROM tbl_conf_periode_spps "+
// 		" where flag_aktif = 0 and tahun_akademik=?  and nm_kelas=?", ParamInputSiswa.Tahun_akademik, ParamInputSiswa.Nm_kelas).Scan(&intJmldata)
// 	if intJmldata == 0 {
// 		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
// 		response := helper.APIResponse("Data Configurasi SPP Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	db.Raw(" SELECT count(*) jmldata from Tbl_trans_uang_masuk_siswa_headers a "+
// 		" INNER JOIN Tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
// 		" where a.flag_aktif=0 and b.flag_aktif=0 and kd_group=?  "+
// 		" and kd_kategori=? and nis_siswa=? and nm_kelas=? "+
// 		" and tahun_akademik=? ", ParamInputSiswa.Kd_group, ParamInputSiswa.Kd_kategori, ParamInputSiswa.Nis_siswa, ParamInputSiswa.Nm_kelas, ParamInputSiswa.Tahun_akademik).Scan(&intJmldata)
// 	if intJmldata > 0 {
// 		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
// 		response := helper.APIResponse("Data Pembayaran SPP Sudah Ada ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 		//}

// 		// var CekDataUangMasuk table_data.Tbl_trans_uang_masuk_siswa_headers
// 		// checkUser := db.Select("*").Where("flag_aktif = 0 and kd_group= ? and kd_kategori= ? and nis_siswa=? and nm_kelas=? and tahun_akademik = ?", ParamInputSiswa.Kd_group, ParamInputSiswa.Kd_kategori, ParamInputSiswa.Nis_siswa, ParamInputSiswa.Nm_kelas, ParamInputSiswa.Tahun_akademik).Find(&CekDataUangMasuk)
// 		// if checkUser.RowsAffected > 0 {
// 		// errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
// 		// response := helper.APIResponse("Data Pembayaran SPP Sudah Ada ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		// c.JSON(http.StatusUnprocessableEntity, response)
// 		//return
// 	} else {

// 		//Jika data spp belum dibuat
// 		var datenows string = time.Now().UTC().Format("2006-01-02 15:04:05")
// 		date := "2006-01-02 15:04:05"
// 		datenowx, err := time.Parse(date, datenows)
// 		if err != nil {
// 			errors := helper.FormatValidationError(err)
// 			errorMessage := gin.H{"errors": errors, "tgl": datenowx}
// 			response := helper.APIResponse("Format Tanggal Salah ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 			c.JSON(http.StatusUnprocessableEntity, response)
// 			return
// 		}

// 		var intKd_trans_masuk int
// 		db.Raw("SELECT ifnull(max(kd_trans_masuk),0) + 1 as 'run_number' FROM Tbl_trans_uang_masuk_siswa_headers where flag_aktif=0").Scan(&intKd_trans_masuk)

// 		var float_biaya_spp float64
// 		db.Raw("SELECT sum(biaya_spp) 'biaya_spp' FROM tbl_conf_periode_spps where tahun_akademik=? and nm_kelas=? and flag_aktif=0", ParamInputSiswa.Tahun_akademik, ParamInputSiswa.Nm_kelas).Scan(&float_biaya_spp)

// 		currentUser := c.MustGet("currentUser")
// 		data := table_data.Tbl_trans_uang_masuk_siswa_headers{
// 			Kd_group:       ParamInputSiswa.Kd_group,
// 			Kd_kategori:    ParamInputSiswa.Kd_kategori,
// 			Kd_trans_masuk: intKd_trans_masuk,
// 			Nis_siswa:      ParamInputSiswa.Nis_siswa,
// 			Nm_kelas:       ParamInputSiswa.Nm_kelas,
// 			Tahun_akademik: ParamInputSiswa.Tahun_akademik,
// 			Total_biaya:    float_biaya_spp,
// 			Total_bayar:    0,
// 			Sisa_biaya:     float_biaya_spp,
// 			Keterangan:     ParamInputSiswa.Keterangan,
// 			Created_by:     currentUser.(string),
// 			Created_on:     datenowx,
// 			Flag_aktif:     0,
// 		}

// 		err = db.Omit("Edited_on", "Edited_by").Create(&data).Error
// 		if err != nil {
// 			response := helper.APIResponse("Simpan Data Gagal ...", http.StatusBadRequest, "error", err)
// 			c.JSON(http.StatusBadRequest, response)
// 			return
// 		}

// 		var intKd_trans_masuk_detail int
// 		db.Raw("SELECT ifnull(max(kd_trans_masuk_detail),0) + 1 as 'run_number' FROM Tbl_trans_uang_masuk_siswa_details where flag_aktif=0").Scan(&intKd_trans_masuk_detail)

// 		var int_seqno int
// 		var periodebayar string
// 		rows, _ := db.Raw("SELECT seqno,CONCAT(kd_bulan,'-',tahun) 'periodebayar',biaya_spp FROM tbl_conf_periode_spps WHERE flag_aktif=0 and tahun_akademik=? and nm_kelas=? ORDER BY seqno", ParamInputSiswa.Tahun_akademik, ParamInputSiswa.Nm_kelas).Rows()
// 		defer rows.Close()
// 		for rows.Next() {
// 			rows.Scan(&int_seqno, &periodebayar)

// 			datadetail := table_data.Tbl_trans_uang_masuk_siswa_details{
// 				Kd_trans_masuk:        intKd_trans_masuk,
// 				Kd_trans_masuk_detail: intKd_trans_masuk_detail,
// 				Seqno:                 int_seqno,
// 				Periode_bayar:         periodebayar,
// 				Jml_bayar:             0,
// 				Keterangan:            "",
// 				Created_by:            currentUser.(string),
// 				Created_on:            datenowx,
// 				Flag_aktif:            0,
// 			}

// 			err = db.Omit("Edited_on", "Edited_by", "Tgl_bayar").Create(&datadetail).Error
// 			if err != nil {
// 				response := helper.APIResponse("Simpan Data Detail Gagal ...", http.StatusBadRequest, "error", err)
// 				c.JSON(http.StatusBadRequest, response)
// 				return
// 			}

// 			intKd_trans_masuk_detail++
// 		}

// 		//setting tampilan habis save spp

// 		var GetDataUmSiswa []GetDataUmSiswa
// 		db.Raw("SELECT b.kd_trans_masuk_detail,b.seqno,b.periode_bayar, "+
// 			" b.tgl_bayar,b.jml_bayar,b.keterangan "+
// 			" FROM Tbl_trans_uang_masuk_siswa_headers a "+
// 			" INNER JOIN Tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
// 			" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis "+
// 			" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "+
// 			" and a.tahun_akademik=? and a.nm_kelas=? and a.nis_siswa = ? "+
// 			" order by b.seqno ", ParamInputSiswa.Tahun_akademik, ParamInputSiswa.Nm_kelas, ParamInputSiswa.Nis_siswa).Scan(&GetDataUmSiswa)

// 		SetArrayData := []GetBiayaAndSisa{}
// 		var kd_trans_masuk int
// 		var total_biaya float64
// 		var total_bayar float64
// 		var sisa_biaya float64
// 		rowss, _ := db.Raw("SELECT distinct b.kd_trans_masuk,a.total_biaya,a.total_bayar,a.sisa_biaya "+
// 			" FROM Tbl_trans_uang_masuk_siswa_headers a "+
// 			" INNER JOIN Tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
// 			" where a.flag_aktif=0 and b.flag_aktif=0  "+
// 			" and a.tahun_akademik=? and a.nm_kelas=? and a.nis_siswa = ? ", ParamInputSiswa.Tahun_akademik, ParamInputSiswa.Nm_kelas, ParamInputSiswa.Nis_siswa).Rows()
// 		defer rowss.Close()
// 		for rowss.Next() {
// 			rowss.Scan(&kd_trans_masuk, &total_biaya, &total_bayar, &sisa_biaya)
// 			arraydata := GetBiayaAndSisa{}
// 			arraydata.Kd_trans_masuk = kd_trans_masuk
// 			arraydata.Total_biaya = total_biaya
// 			arraydata.Total_bayar = total_bayar
// 			arraydata.Sisa_biaya = sisa_biaya
// 			arraydata.Detail = GetDataUmSiswa
// 			SetArrayData = append(SetArrayData, arraydata)
// 		}

// 		response := helper.APIResponse("Simpan Data Sukses ...", http.StatusOK, "success", SetArrayData)
// 		c.JSON(http.StatusOK, response)
// 	}

// }

// func UpdateUangMasukSpp(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)

// 	var ParamEditSiswaDetail ParamEditSiswaDetail
// 	if err := c.ShouldBindJSON(&ParamEditSiswaDetail); err != nil {
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

// 	var dataUtama table_data.Tbl_trans_uang_masuk_siswa_details
// 	if err := db.Where("flag_aktif=0 and kd_trans_masuk_detail=?", c.Param("iddetail")).First(&dataUtama).Error; err != nil {
// 		errorMessage := gin.H{"errors": "Data Tidak Ditemukan ..."}
// 		response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	currentUser := c.MustGet("currentUser")

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

// 	tTglBayar, err2 := time.Parse("02-01-2006", ParamEditSiswaDetail.Tgl_bayar)
// 	if err2 != nil {
// 		var ve validator.ValidationErrors
// 		if errors.As(err2, &ve) {
// 			errors := helper.FormatValidationError(err2)
// 			errorMessage := gin.H{"errors": errors}
// 			response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 			c.JSON(http.StatusUnprocessableEntity, response)
// 			return
// 		}
// 		var error_binding []string
// 		error_binding = append(error_binding, err2.Error())
// 		errorMessage := gin.H{"errors": error_binding}
// 		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}
// 	dateStr := tTglBayar.Format("2006-01-02")

// 	var dataDetail table_data.Tbl_trans_uang_masuk_siswa_details
// 	err = db.Raw("update Tbl_trans_uang_masuk_siswa_details set tgl_bayar=?,jml_bayar=?,keterangan=?,edited_by=?,edited_on=? "+
// 		" where kd_trans_masuk_detail=? and flag_aktif=0 ", dateStr, ParamEditSiswaDetail.Jml_bayar, ParamEditSiswaDetail.Keterangan, currentUser.(string), datenowx, c.Param("iddetail")).Scan(&dataDetail).Error
// 	if err != nil {
// 		response := helper.APIResponse("Update Data Ke Tbl_trans_uang_masuk_siswa_details Gagal ...", http.StatusBadRequest, "error", err)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	var sumJmlBayar float64
// 	db.Raw("SELECT sum(jml_bayar) 'jml_bayar' FROM Tbl_trans_uang_masuk_siswa_details "+
// 		" where flag_aktif=0 and tgl_bayar is not null and kd_trans_masuk=?", c.Param("idhead")).Scan(&sumJmlBayar)

// 	var total_biaya float64
// 	db.Raw("SELECT total_biaya FROM Tbl_trans_uang_masuk_siswa_headers where flag_aktif=0 and kd_trans_masuk=?", c.Param("idhead")).Scan(&total_biaya)
// 	var sisa_biaya float64 = total_biaya - sumJmlBayar

// 	var dataHeader table_data.Tbl_trans_uang_masuk_siswa_headers
// 	err = db.Raw("UPDATE Tbl_trans_uang_masuk_siswa_headers SET total_bayar = ?, sisa_biaya = ?, "+
// 		" edited_on = ? , edited_by = ? "+
// 		" WHERE kd_trans_masuk = ? and flag_aktif=0 ", sumJmlBayar, sisa_biaya, datenowx, currentUser.(string), c.Param("idhead")).Scan(&dataHeader).Error
// 	if err != nil {
// 		response := helper.APIResponse("Update Data Ke Tbl_trans_uang_masuk_siswa_headers Gagal ...", http.StatusBadRequest, "error", err)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	//setting tampilan habis update spp

// 	var GetDataUmSiswa []GetDataUmSiswa
// 	db.Raw("SELECT b.kd_trans_masuk_detail,b.seqno,b.periode_bayar, "+
// 		" b.tgl_bayar,b.jml_bayar,b.keterangan "+
// 		" FROM Tbl_trans_uang_masuk_siswa_headers a "+
// 		" INNER JOIN Tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
// 		" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis "+
// 		" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "+
// 		" and a.kd_trans_masuk=? "+
// 		" order by b.seqno ", c.Param("idhead")).Scan(&GetDataUmSiswa)

// 	SetArrayData := []GetBiayaAndSisa{}
// 	var kd_trans_masuk int
// 	var total_bayar float64
// 	rowss, _ := db.Raw("SELECT distinct b.kd_trans_masuk,a.total_biaya,a.total_bayar,a.sisa_biaya "+
// 		" FROM Tbl_trans_uang_masuk_siswa_headers a "+
// 		" INNER JOIN Tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
// 		" where a.flag_aktif=0 and b.flag_aktif=0  "+
// 		" and a.kd_trans_masuk=? ", c.Param("idhead")).Rows()
// 	defer rowss.Close()
// 	for rowss.Next() {
// 		rowss.Scan(&kd_trans_masuk, &total_biaya, &total_bayar, &sisa_biaya)
// 		arraydata := GetBiayaAndSisa{}
// 		arraydata.Kd_trans_masuk = kd_trans_masuk
// 		arraydata.Total_biaya = total_biaya
// 		arraydata.Total_bayar = total_bayar
// 		arraydata.Sisa_biaya = sisa_biaya
// 		arraydata.Detail = GetDataUmSiswa
// 		SetArrayData = append(SetArrayData, arraydata)
// 	}

// 	response := helper.APIResponse("Update Data Sukses ...", http.StatusOK, "success", SetArrayData)
// 	c.JSON(http.StatusOK, response)

// }
