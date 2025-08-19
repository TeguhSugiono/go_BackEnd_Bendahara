package transaksi_uang_masuk_siswa

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

	// nomor := 0
	// var result_kd_group string
	// var str_kd_group string

	// //kd_jenis=1 adalah uang masuk
	// rows, _ := db.Raw("SELECT kd_group FROM tbl_group_kategoris where flag_aktif=0 and kd_jenis=1 ").Rows()
	// defer rows.Close()
	// for rows.Next() {
	// 	rows.Scan(&result_kd_group)
	// if nomor == 0 {
	// 	str_kd_group = result_kd_group
	// } else {
	// 	str_kd_group = str_kd_group + "," + result_kd_group
	// }
	// 	nomor++
	// }

	nomor := 0
	var result_kd_group string
	var str_kd_group string
	rows, _ := db.Raw("SELECT kd_group FROM tbl_link_kategoris where link_name in('form_biaya_spp','form_biaya_ppdb')").Rows()
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

	// var result_kd_jenis string
	// rows1, _ := db.Raw("SELECT kd_jenis FROM tbl_group_kategoris where flag_aktif=0 and kd_jenis=1 ").Rows()
	// defer rows1.Close()
	// for rows1.Next() {
	// 	rows1.Scan(&result_kd_jenis)
	// }

	//kd_jenis=1 adalah uang masuk
	var master []master_group_kategori.ListData
	sql := "  SELECT a.*,b.proses_uang FROM tbl_group_kategoris as a " +
		" inner join tbl_jenis_trans as b on a.kd_jenis=b.kd_jenis  " +
		" where a.flag_aktif=0 and b.flag_aktif=0  " +
		" and a.kd_group not in(" + str_kd_group + ")  and a.kd_jenis in('1') order by b.proses_uang "

	db.Raw(sql).Scan(&master)

	response := helper.APIResponse("List Data ..."+result_kd_group, http.StatusOK, "success", master_group_kategori.FormatGroupKategori(master))
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

	//p_kd_kategori := paramGetSiswaAdd.Kd_kategori
	//p_kd_group := paramGetSiswaAdd.Kd_group
	//p_nis := paramGetSiswaAdd.Nis

	if err := c.ShouldBindJSON(&paramGetSiswaAdd); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//cari setting tarif
	var float_jml_biaya float64
	db.Raw(" SELECT jml_biaya FROM tbl_biaya_masuk_keluars where kd_kategori=?", paramGetSiswaAdd.Kd_kategori).Scan(&float_jml_biaya)

	var id_tahun int
	var tahun_akademik string
	var id_kelas string
	var nm_kelas string

	SetArrayData := []ListAddSiswa{}
	rows, _ := db.Raw(" SELECT DISTINCT id_tahun_aktif 'id_tahun',tahun_aktif 'tahun_akademik', "+
		" nm_kelas as 'id_kelas', "+
		" nm_kelas as 'nm_kelas' "+
		" FROM tbl_siswa "+
		" WHERE flag_siswa = 0 AND status_siswa NOT IN ('Tidak Aktif','LULUS') and nis=?", paramGetSiswaAdd.Nis).Rows()
	defer rows.Close()
	for rows.Next() {
		arraydata := ListAddSiswa{}
		rows.Scan(&id_tahun, &tahun_akademik, &id_kelas, &nm_kelas)
		arraydata.Id_tahun = id_tahun
		arraydata.Tahun_akademik = tahun_akademik
		arraydata.Id_kelas = id_kelas
		arraydata.Nm_kelas = nm_kelas
		arraydata.Total_biaya = float_jml_biaya
		arraydata.Total_bayar = 0
		arraydata.Sisa_biaya = float_jml_biaya
		SetArrayData = append(SetArrayData, arraydata)
	}

	var intJmldata int

	ssql := " SELECT count(*) 'jdata' FROM tbl_trans_uang_masuk_siswa_headers a " +
		" inner join tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
		" where a.flag_aktif=0 and b.flag_aktif=0 "
	ssql = fmt.Sprintf("%s and a.kd_kategori= %d", ssql, paramGetSiswaAdd.Kd_kategori)
	ssql = fmt.Sprintf("%s and a.kd_group= %d", ssql, paramGetSiswaAdd.Kd_group)
	ssql = fmt.Sprintf("%s and a.nis_siswa= '%s'", ssql, paramGetSiswaAdd.Nis)
	ssql = fmt.Sprintf("%s and a.tahun_akademik= '%s'", ssql, tahun_akademik)
	ssql = fmt.Sprintf("%s and a.nm_kelas= '%s'", ssql, nm_kelas)
	//db.Raw(ssql).Scan(&intJmldata)

	rowsdata, _ := db.Raw(ssql).Rows()
	defer rowsdata.Close()
	for rowsdata.Next() {
		rowsdata.Scan(&intJmldata)
	}

	if intJmldata > 0 {
		//setting tampilan habis save siswa

		SetArrayData := []GetBiayaAndSisa{}
		var kd_trans_masuk_siswa int
		var kd_group int
		var nm_group string
		var kd_kategori int
		var nm_kategori string
		var total_biaya float64
		var total_bayar float64
		var sisa_biaya float64
		//var tahun_akademik string
		var nis_siswa string
		//var nm_kelas string
		var nm_siswa string
		var ket string

		ssqlx := " SELECT distinct b.kd_trans_masuk_siswa,a.kd_group,d.nm_group,a.kd_kategori,e.nm_kategori,a.tahun_akademik,a.nis_siswa, " +
			" c.nm_siswa,a.nm_kelas,a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan " +
			" FROM tbl_trans_uang_masuk_siswa_headers a " +
			" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
			" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
			" INNER JOIN tbl_group_kategoris d on a.kd_group = d.kd_group " +
			" INNER JOIN tbl_kategori_uangs e on a.kd_kategori = e.kd_kategori " +
			" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif')  "

		ssqlx = fmt.Sprintf("%s and a.kd_kategori= %d", ssqlx, paramGetSiswaAdd.Kd_kategori)
		ssqlx = fmt.Sprintf("%s and a.kd_group= %d", ssqlx, paramGetSiswaAdd.Kd_group)
		ssqlx = fmt.Sprintf("%s and a.nis_siswa= '%s'", ssqlx, paramGetSiswaAdd.Nis)
		ssqlx = fmt.Sprintf("%s and a.tahun_akademik= '%s'", ssqlx, tahun_akademik)
		ssqlx = fmt.Sprintf("%s and a.nm_kelas= '%s'", ssqlx, nm_kelas)

		ssqlx = fmt.Sprintf("%s ORDER BY a.tahun_akademik %s,a.nm_kelas %s", ssqlx, "asc", "asc")
		rows, _ := db.Raw(ssqlx).Rows()
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&kd_trans_masuk_siswa, &kd_group, &nm_group, &kd_kategori, &nm_kategori, &tahun_akademik, &nis_siswa, &nm_siswa, &nm_kelas, &total_biaya, &total_bayar, &sisa_biaya, &ket)
			arraydata := GetBiayaAndSisa{}
			arraydata.Kd_trans_masuk_siswa = kd_trans_masuk_siswa
			arraydata.Kd_group = kd_group
			arraydata.Nm_group = nm_group
			arraydata.Kd_kategori = kd_kategori
			arraydata.Nm_kategori = nm_kategori
			arraydata.Tahun_akademik = tahun_akademik
			arraydata.Nis_siswa = nis_siswa
			arraydata.Nm_siswa = nm_siswa
			arraydata.Nm_kelas = nm_kelas
			arraydata.Total_biaya = total_biaya
			arraydata.Total_bayar = total_bayar
			arraydata.Sisa_biaya = sisa_biaya
			arraydata.Keterangan = ket

			sql := " SELECT b.kd_trans_masuk_detail_siswa,b.seqno, " +
				" b.tgl_bayar,b.jml_bayar,b.keterangan,d.kd_pembayaran,d.tipe_pembayaran " +
				" FROM tbl_trans_uang_masuk_siswa_headers a " +
				" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
				" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
				" LEFT JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran " +
				" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "

			sql = fmt.Sprintf("%s and a.kd_trans_masuk_siswa = %d", sql, kd_trans_masuk_siswa)

			sql = fmt.Sprintf("%s ORDER BY a.kd_trans_masuk_siswa %s,b.seqno %s", sql, "asc", "asc")

			var getDataUmSiswa []GetDataUmSiswa
			db.Raw(sql).Scan(&getDataUmSiswa)

			arraydata.Detail = getDataUmSiswa
			SetArrayData = append(SetArrayData, arraydata)
		}

		// if len(SetArrayData) == 0 {
		// 	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
		// 	c.JSON(http.StatusOK, response)
		// 	return
		// }

		response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
		c.JSON(http.StatusOK, response)
	} else {
		response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
		c.JSON(http.StatusOK, response)
	}

}

func CreateUangMasukSiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var paramInputSiswa ParamInputSiswa
	if err := c.ShouldBindJSON(&paramInputSiswa); err != nil {
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
	db.Raw(" SELECT count(*) jmldata FROM tbl_group_kategoris where kd_group=? and flag_aktif=0 ", paramInputSiswa.Kd_group).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Kode Group Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	db.Raw(" SELECT count(*) jmldata FROM tbl_kategori_uangs where kd_kategori=? and flag_aktif=0 ", paramInputSiswa.Kd_kategori).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Kode Kategori Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//cek data siswa

	db.Raw("SELECT count(*) jmldata FROM tbl_siswa where flag_siswa=0 and status_siswa not in('Tidak Aktif','LULUS') and nis=? "+
		" and (tahun_aktif=? or tahun_aktif = REPLACE(?,'-','/')) "+
		" and nm_kelas = ?", paramInputSiswa.Nis_siswa, paramInputSiswa.Tahun_akademik, paramInputSiswa.Tahun_akademik, paramInputSiswa.Nm_kelas).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Data Siswa Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	db.Raw(" SELECT count(*) jmldata from tbl_trans_uang_masuk_siswa_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa "+
		" where a.flag_aktif=0 and b.flag_aktif=0 and kd_group=?  "+
		" and kd_kategori=? and nis_siswa=? and nm_kelas=? "+
		" and tahun_akademik=? ", paramInputSiswa.Kd_group, paramInputSiswa.Kd_kategori, paramInputSiswa.Nis_siswa, paramInputSiswa.Nm_kelas, paramInputSiswa.Tahun_akademik).Scan(&intJmldata)

	if intJmldata > 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Data Pembayaran Siswa Sudah Ada ...", http.StatusUnprocessableEntity, "error", errorMessage)
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
		db.Raw("SELECT ifnull(max(kd_trans_masuk_siswa),0) + 1 as 'run_number' FROM tbl_trans_uang_masuk_siswa_headers ").Scan(&intKd_trans_masuk)

		var float_biaya_spp float64
		db.Raw("SELECT jml_biaya FROM tbl_biaya_masuk_keluars where kd_kategori=?", paramInputSiswa.Kd_kategori).Scan(&float_biaya_spp)

		if paramInputSiswa.Total_biaya > 0 {
			float_biaya_spp = paramInputSiswa.Total_biaya
		}

		currentUser := c.MustGet("currentUser")
		data := table_data.Tbl_trans_uang_masuk_siswa_headers{
			Kd_group:             paramInputSiswa.Kd_group,
			Kd_kategori:          paramInputSiswa.Kd_kategori,
			Kd_trans_masuk_siswa: intKd_trans_masuk,
			Nis_siswa:            paramInputSiswa.Nis_siswa,
			Nm_kelas:             paramInputSiswa.Nm_kelas,
			Tahun_akademik:       paramInputSiswa.Tahun_akademik,
			Total_biaya:          float_biaya_spp,
			Total_bayar:          0,
			Sisa_biaya:           float_biaya_spp,
			Keterangan:           paramInputSiswa.Keterangan,
			Created_by:           currentUser.(string),
			Created_on:           datenowx,
			Flag_aktif:           0,
		}

		err = db.Omit("Edited_on", "Edited_by").Create(&data).Error
		if err != nil {
			response := helper.APIResponse("Simpan Data Gagal ...", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		var intKd_trans_masuk_detail int
		db.Raw("SELECT ifnull(max(kd_trans_masuk_detail_siswa),0) + 1 as 'run_number' FROM tbl_trans_uang_masuk_siswa_details ").Scan(&intKd_trans_masuk_detail)

		datadetail := table_data.Tbl_trans_uang_masuk_siswa_details{
			Kd_trans_masuk_siswa:        intKd_trans_masuk,
			Kd_trans_masuk_detail_siswa: intKd_trans_masuk_detail,
			Seqno:                       1,
			Jml_bayar:                   0,
			Keterangan:                  "",
			Created_by:                  currentUser.(string),
			Created_on:                  datenowx,
			Flag_aktif:                  0,
		}

		err = db.Omit("Edited_on", "Edited_by", "Tgl_bayar", "Kd_pembayaran").Create(&datadetail).Error
		if err != nil {
			response := helper.APIResponse("Simpan Data Detail Gagal ...", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		//setting tampilan habis save siswa

		SetArrayData := []GetBiayaAndSisa{}
		var kd_trans_masuk_siswa int
		var kd_group int
		var nm_group string
		var kd_kategori int
		var nm_kategori string
		var total_biaya float64
		var total_bayar float64
		var sisa_biaya float64
		var tahun_akademik string
		var nis_siswa string
		var nm_kelas string
		var nm_siswa string
		var ket string

		ssql := " SELECT distinct b.kd_trans_masuk_siswa,a.kd_group,d.nm_group,a.kd_kategori,e.nm_kategori,a.tahun_akademik,a.nis_siswa, " +
			" c.nm_siswa,a.nm_kelas,a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan " +
			" FROM tbl_trans_uang_masuk_siswa_headers a " +
			" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
			" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
			" INNER JOIN tbl_group_kategoris d on a.kd_group = d.kd_group " +
			" INNER JOIN tbl_kategori_uangs e on a.kd_kategori = e.kd_kategori " +
			" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif')  "

		ssql = fmt.Sprintf("%s and a.tahun_akademik= '%s'", ssql, paramInputSiswa.Tahun_akademik)

		ssql = fmt.Sprintf("%s ORDER BY a.tahun_akademik %s,a.nm_kelas %s", ssql, "asc", "asc")
		rows, _ := db.Raw(ssql).Rows()
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&kd_trans_masuk_siswa, &kd_group, &nm_group, &kd_kategori, &nm_kategori, &tahun_akademik, &nis_siswa, &nm_siswa, &nm_kelas, &total_biaya, &total_bayar, &sisa_biaya, &ket)
			arraydata := GetBiayaAndSisa{}
			arraydata.Kd_trans_masuk_siswa = kd_trans_masuk_siswa
			arraydata.Kd_group = kd_group
			arraydata.Nm_group = nm_group
			arraydata.Kd_kategori = kd_kategori
			arraydata.Nm_kategori = nm_kategori
			arraydata.Tahun_akademik = tahun_akademik
			arraydata.Nis_siswa = nis_siswa
			arraydata.Nm_siswa = nm_siswa
			arraydata.Nm_kelas = nm_kelas
			arraydata.Total_biaya = total_biaya
			arraydata.Total_bayar = total_bayar
			arraydata.Sisa_biaya = sisa_biaya
			arraydata.Keterangan = ket

			sql := " SELECT b.kd_trans_masuk_detail_siswa,b.seqno, " +
				" b.tgl_bayar,b.jml_bayar,b.keterangan,d.kd_pembayaran,d.tipe_pembayaran " +
				" FROM tbl_trans_uang_masuk_siswa_headers a " +
				" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
				" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
				" LEFT JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran " +
				" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "

			sql = fmt.Sprintf("%s and a.kd_trans_masuk_siswa = %d", sql, intKd_trans_masuk)

			sql = fmt.Sprintf("%s ORDER BY a.kd_trans_masuk_siswa %s,b.seqno %s", sql, "asc", "asc")

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

func EditUangMasukSiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	kd_trans_masuk_siswa := c.Param("idhead")

	var paramInputSiswa ParamInputSiswaEdit
	if err := c.ShouldBindJSON(&paramInputSiswa); err != nil {
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

	db.Raw(" SELECT count(*) jmldata FROM tbl_group_kategoris where kd_group=? and flag_aktif=0 ", paramInputSiswa.Kd_group).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Kode Group Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	db.Raw(" SELECT count(*) jmldata FROM tbl_kategori_uangs where kd_kategori=? and flag_aktif=0 ", paramInputSiswa.Kd_kategori).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Kode Kategori Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//cek data siswa

	db.Raw("SELECT count(*) jmldata FROM tbl_siswa where flag_siswa=0 and status_siswa not in('Tidak Aktif','LULUS') and nis=? "+
		" and (tahun_aktif=? or tahun_aktif = REPLACE(?,'-','/')) "+
		" and nm_kelas = ?", paramInputSiswa.Nis_siswa, paramInputSiswa.Tahun_akademik, paramInputSiswa.Tahun_akademik, paramInputSiswa.Nm_kelas).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Data Siswa Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//cek data transaksi
	db.Raw(" SELECT count(*) jmldata FROM tbl_trans_uang_masuk_siswa_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa "+
		" where a.kd_trans_masuk_siswa=? and a.flag_aktif=0 and b.flag_aktif=0 ", kd_trans_masuk_siswa).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Data Transaksi Siswa Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//Validasi data sama
	var kd_group_old int
	var kd_kategori_old int
	var tahun_akademik_old string
	var nis_siswa_old string
	var nm_kelas_old string

	rowA, _ := db.Raw("SELECT a.kd_group,a.kd_kategori,a.tahun_akademik,a.nis_siswa,a.nm_kelas FROM tbl_trans_uang_masuk_siswa_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa "+
		" where a.kd_trans_masuk_siswa=?  and a.flag_aktif=0 and b.flag_aktif=0  limit 1", kd_trans_masuk_siswa).Rows()

	defer rowA.Close()
	for rowA.Next() {
		rowA.Scan(&kd_group_old, &kd_kategori_old, &tahun_akademik_old, &nis_siswa_old, &nm_kelas_old)
	}

	//===============================================================================================================================
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
	db.Raw("SELECT ifnull(sum(jml_bayar),0.00) 'jml_bayar' FROM tbl_trans_uang_masuk_siswa_details "+
		" where flag_aktif=0 and tgl_bayar is not null and kd_trans_masuk_siswa=?", kd_trans_masuk_siswa).Scan(&sumJmlBayar)

	var sisa_biaya float64 = paramInputSiswa.Total_biaya - sumJmlBayar

	// response := helper.APIResponse("Simpan Data Sukses ..."+strconv.Itoa(kd_group_old)+"---"+strconv.Itoa(kd_kategori_old)+"---"+tahun_akademik_old+"---"+nis_siswa_old+"---"+nm_kelas_old, http.StatusOK, "success"+strconv.Itoa(paramInputSiswa.Kd_group)+"---"+strconv.Itoa(paramInputSiswa.Kd_kategori)+paramInputSiswa.Tahun_akademik+"---"+paramInputSiswa.Nis_siswa+"---"+paramInputSiswa.Nm_kelas, paramInputSiswa)
	// c.JSON(http.StatusOK, response)

	//===============================================================================================================================

	if (kd_group_old != paramInputSiswa.Kd_group) || (kd_kategori_old != paramInputSiswa.Kd_kategori) || (tahun_akademik_old != paramInputSiswa.Tahun_akademik) || (nis_siswa_old != paramInputSiswa.Nis_siswa) || (nm_kelas_old != paramInputSiswa.Nm_kelas) {
		db.Raw(" SELECT count(*) jmldata from tbl_trans_uang_masuk_siswa_headers a "+
			" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa "+
			" where a.flag_aktif=0 and b.flag_aktif=0 and kd_group=?  "+
			" and kd_kategori=? and nis_siswa=? and nm_kelas=? "+
			" and tahun_akademik=? ", paramInputSiswa.Kd_group, paramInputSiswa.Kd_kategori, paramInputSiswa.Nis_siswa,
			paramInputSiswa.Nm_kelas, paramInputSiswa.Tahun_akademik).Scan(&intJmldata)

		if intJmldata > 0 {
			errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
			response := helper.APIResponse("Data Pembayaran Siswa Sudah Ada ...", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		} else {

			var dataHeader table_data.Tbl_trans_uang_masuk_siswa_headers
			err = db.Raw("UPDATE tbl_trans_uang_masuk_siswa_headers SET total_biaya=? ,total_bayar = ?, sisa_biaya = ?, "+
				" edited_on = ? , edited_by = ? ,kd_group=?,kd_kategori=?,nis_siswa=?,nm_kelas=?,tahun_akademik=?,kd_group=?,kd_kategori=?  "+
				" WHERE kd_trans_masuk_siswa = ? "+
				" and flag_aktif=0 ", paramInputSiswa.Total_biaya, sumJmlBayar, sisa_biaya, datenowx, currentUser.(string),
				paramInputSiswa.Kd_group, paramInputSiswa.Kd_kategori, paramInputSiswa.Nis_siswa,
				paramInputSiswa.Nm_kelas, paramInputSiswa.Tahun_akademik, paramInputSiswa.Kd_group, paramInputSiswa.Kd_kategori, kd_trans_masuk_siswa).Scan(&dataHeader).Error
			if err != nil {
				response := helper.APIResponse("Update Data Ke tbl_trans_uang_masuk_siswa_headers Gagal ...", http.StatusBadRequest, "error", err)
				c.JSON(http.StatusBadRequest, response)
				return
			}

			// var dataDetail table_data.Tbl_trans_uang_masuk_siswa_details
			// err = db.Raw("update Tbl_trans_uang_masuk_siswa_details set tgl_bayar=?,jml_bayar=?,keterangan=?,edited_by=?,edited_on=? "+
			// 	" where kd_trans_masuk_detail=? and flag_aktif=0 ", dateStr, ParamEditSiswaDetail.Jml_bayar, ParamEditSiswaDetail.Keterangan, currentUser.(string), datenowx, c.Param("iddetail")).Scan(&dataDetail).Error
			// if err != nil {
			// 	response := helper.APIResponse("Update Data Ke Tbl_trans_uang_masuk_siswa_details Gagal ...", http.StatusBadRequest, "error", err)
			// 	c.JSON(http.StatusBadRequest, response)
			// 	return
			// }

		}

	} else {

		var dataHeader table_data.Tbl_trans_uang_masuk_siswa_headers
		err = db.Raw("UPDATE tbl_trans_uang_masuk_siswa_headers SET total_biaya=? ,total_bayar = ?, sisa_biaya = ?, "+
			" edited_on = ? , edited_by = ?   "+
			" WHERE kd_trans_masuk_siswa = ? "+
			" and flag_aktif=0 ", paramInputSiswa.Total_biaya, sumJmlBayar, sisa_biaya, datenowx, currentUser.(string), kd_trans_masuk_siswa).Scan(&dataHeader).Error
		if err != nil {
			response := helper.APIResponse("Update Data Ke tbl_trans_uang_masuk_siswa_headers Gagal ...", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

	}

	//End Validasi data sama

	//setting tampilan habis save siswa

	SetArrayData := []GetBiayaAndSisa{}
	//var kd_trans_masuk_siswa int
	var kd_group int
	var nm_group string
	var kd_kategori int
	var nm_kategori string
	var total_biaya float64
	var total_bayar float64
	//var sisa_biaya float64
	var tahun_akademik string
	var nis_siswa string
	var nm_kelas string
	var nm_siswa string
	var ket string

	ssql := " SELECT distinct b.kd_trans_masuk_siswa,a.kd_group,d.nm_group,a.kd_kategori,e.nm_kategori,a.tahun_akademik,a.nis_siswa, " +
		" c.nm_siswa,a.nm_kelas,a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan " +
		" FROM tbl_trans_uang_masuk_siswa_headers a " +
		" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
		" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
		" INNER JOIN tbl_group_kategoris d on a.kd_group = d.kd_group " +
		" INNER JOIN tbl_kategori_uangs e on a.kd_kategori = e.kd_kategori " +
		" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif')  "

	ssql = fmt.Sprintf("%s and a.kd_trans_masuk_siswa= '%s'", ssql, kd_trans_masuk_siswa)
	ssql = fmt.Sprintf("%s ORDER BY a.tahun_akademik %s,a.nm_kelas %s", ssql, "asc", "asc")

	int_kd_trans_masuk_siswa, _ := strconv.Atoi(kd_trans_masuk_siswa)

	rows, _ := db.Raw(ssql).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&kd_trans_masuk_siswa, &kd_group, &nm_group, &kd_kategori, &nm_kategori, &tahun_akademik, &nis_siswa, &nm_siswa, &nm_kelas, &total_biaya, &total_bayar, &sisa_biaya, &ket)
		arraydata := GetBiayaAndSisa{}
		arraydata.Kd_trans_masuk_siswa = int_kd_trans_masuk_siswa
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group
		arraydata.Kd_kategori = kd_kategori
		arraydata.Nm_kategori = nm_kategori
		arraydata.Tahun_akademik = tahun_akademik
		arraydata.Nis_siswa = nis_siswa
		arraydata.Nm_siswa = nm_siswa
		arraydata.Nm_kelas = nm_kelas
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = ket

		sql := " SELECT b.kd_trans_masuk_detail_siswa,b.seqno, " +
			" b.tgl_bayar,b.jml_bayar,b.keterangan,d.kd_pembayaran,d.tipe_pembayaran " +
			" FROM tbl_trans_uang_masuk_siswa_headers a " +
			" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
			" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
			" LEFT JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran " +
			" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "

		sql = fmt.Sprintf("%s and a.kd_trans_masuk_siswa = '%s'", sql, kd_trans_masuk_siswa)

		sql = fmt.Sprintf("%s ORDER BY a.kd_trans_masuk_siswa %s,b.seqno %s", sql, "asc", "asc")

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

	// SetArrayData := []GetBiayaAndSisa{}
	// var kd_trans_masuk_siswaX int
	// var total_biaya float64
	// var total_bayar float64
	// //var sisa_biaya float64
	// var tahun_akademik string
	// var nis_siswa string
	// var nm_kelas string
	// var nm_siswa string

	// ssql := " SELECT distinct b.kd_trans_masuk_siswa,a.tahun_akademik,a.nis_siswa,c.nm_siswa,a.nm_kelas,a.total_biaya,a.total_bayar,a.sisa_biaya " +
	// 	" FROM tbl_trans_uang_masuk_siswa_headers a " +
	// 	" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
	// 	" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
	// 	" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif')  "

	// ssql = fmt.Sprintf("%s and a.kd_trans_masuk_siswa= '%s'", ssql, kd_trans_masuk_siswa)

	// ssql = fmt.Sprintf("%s ORDER BY a.tahun_akademik %s,a.nm_kelas %s", ssql, "asc", "asc")
	// rows, _ := db.Raw(ssql).Rows()
	// defer rows.Close()
	// for rows.Next() {
	// 	rows.Scan(&kd_trans_masuk_siswaX, &tahun_akademik, &nis_siswa, &nm_siswa, &nm_kelas, &total_biaya, &total_bayar, &sisa_biaya)
	// 	arraydata := GetBiayaAndSisa{}
	// 	arraydata.Kd_trans_masuk_siswa = kd_trans_masuk_siswaX
	// 	arraydata.Tahun_akademik = tahun_akademik
	// 	arraydata.Nis_siswa = nis_siswa
	// 	arraydata.Nm_siswa = nm_siswa
	// 	arraydata.Nm_kelas = nm_kelas
	// 	arraydata.Total_biaya = total_biaya
	// 	arraydata.Total_bayar = total_bayar
	// 	arraydata.Sisa_biaya = sisa_biaya

	// 	sql := " SELECT b.kd_trans_masuk_detail_siswa,b.seqno, " +
	// 		" b.tgl_bayar,b.jml_bayar,b.keterangan " +
	// 		" FROM tbl_trans_uang_masuk_siswa_headers a " +
	// 		" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
	// 		" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
	// 		" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "

	// 	sql = fmt.Sprintf("%s and a.kd_trans_masuk_siswa = '%s'", sql, kd_trans_masuk_siswa)

	// 	sql = fmt.Sprintf("%s ORDER BY a.kd_trans_masuk_siswa %s,b.seqno %s", sql, "asc", "asc")

	// 	var getDataUmSiswa []GetDataUmSiswa
	// 	db.Raw(sql).Scan(&getDataUmSiswa)

	// 	arraydata.Detail = getDataUmSiswa
	// 	SetArrayData = append(SetArrayData, arraydata)
	// }

	// response := helper.APIResponse("Simpan Data Sukses ...", http.StatusOK, "success", SetArrayData)
	// c.JSON(http.StatusOK, response)

	// response := helper.APIResponse("Simpan Data Sukses ..."+strconv.Itoa(kd_group_old)+"---"+strconv.Itoa(kd_kategori_old)+"---"+tahun_akademik_old+"---"+nis_siswa_old+"---"+nm_kelas_old, http.StatusOK, "success"+strconv.Itoa(paramInputSiswa.Kd_group)+"---"+strconv.Itoa(paramInputSiswa.Kd_kategori)+paramInputSiswa.Tahun_akademik+"---"+paramInputSiswa.Nis_siswa+"---"+paramInputSiswa.Nm_kelas, paramInputSiswa)
	// c.JSON(http.StatusOK, response)
}

func DeleteUangMasukSiswaDetail(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramEditUmSiswaDetail ParamDeleteUmSiswaDetail
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

	kd_trans_masuk_siswa := paramEditUmSiswaDetail.Kd_trans_masuk_siswa
	kd_trans_masuk_detail_siswa := paramEditUmSiswaDetail.Kd_trans_masuk_detail_siswa

	var dataUtama table_data.Tbl_trans_uang_masuk_siswa_details
	if err := db.Where("flag_aktif=0 and kd_trans_masuk_detail_siswa=? and kd_trans_masuk_siswa=?", kd_trans_masuk_detail_siswa, kd_trans_masuk_siswa).First(&dataUtama).Error; err != nil {
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

	var dataDetail table_data.Tbl_trans_uang_masuk_siswa_details
	err = db.Raw("update tbl_trans_uang_masuk_siswa_details set flag_aktif=9,edited_by=?,edited_on=? "+
		" where kd_trans_masuk_detail_siswa=? and kd_trans_masuk_siswa=? and flag_aktif=0 ",
		currentUser.(string), datenowx, kd_trans_masuk_detail_siswa, kd_trans_masuk_siswa).Scan(&dataDetail).Error
	if err != nil {
		response := helper.APIResponse("Delete Data Ke Tbl_trans_uang_masuk_siswa_details Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var sumJmlBayar float64
	db.Raw("SELECT ifnull(sum(jml_bayar),0.00) 'jml_bayar' FROM tbl_trans_uang_masuk_siswa_details "+
		" where flag_aktif=0 and tgl_bayar is not null and kd_trans_masuk_siswa=?", kd_trans_masuk_siswa).Scan(&sumJmlBayar)

	var total_biaya float64
	db.Raw("SELECT total_biaya FROM tbl_trans_uang_masuk_siswa_headers where flag_aktif=0 and kd_trans_masuk_siswa=?", kd_trans_masuk_siswa).Scan(&total_biaya)
	var sisa_biaya float64 = total_biaya - sumJmlBayar

	var dataHeader table_data.Tbl_trans_uang_masuk_siswa_headers
	err = db.Raw("UPDATE tbl_trans_uang_masuk_siswa_headers SET total_bayar = ?, sisa_biaya = ?, "+
		" edited_on = ? , edited_by = ? "+
		" WHERE kd_trans_masuk_siswa = ? and flag_aktif=0 ", sumJmlBayar, sisa_biaya, datenowx, currentUser.(string), kd_trans_masuk_siswa).Scan(&dataHeader).Error
	if err != nil {
		response := helper.APIResponse("Update Data Ke Tbl_trans_uang_masuk_siswa_headers Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//setting tampilan habis save siswa

	SetArrayData := []GetBiayaAndSisa{}
	//var kd_trans_masuk_siswa int
	var kd_group int
	var nm_group string
	var kd_kategori int
	var nm_kategori string
	//var total_biaya float64
	var total_bayar float64
	//var sisa_biaya float64
	var tahun_akademik string
	var nis_siswa string
	var nm_kelas string
	var nm_siswa string
	var ket string

	ssql := " SELECT distinct b.kd_trans_masuk_siswa,a.kd_group,d.nm_group,a.kd_kategori,e.nm_kategori,a.tahun_akademik,a.nis_siswa, " +
		" c.nm_siswa,a.nm_kelas,a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan " +
		" FROM tbl_trans_uang_masuk_siswa_headers a " +
		" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
		" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
		" INNER JOIN tbl_group_kategoris d on a.kd_group = d.kd_group " +
		" INNER JOIN tbl_kategori_uangs e on a.kd_kategori = e.kd_kategori " +
		" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif')  "

	ssql = fmt.Sprintf("%s and a.kd_trans_masuk_siswa= '%s'", ssql, kd_trans_masuk_siswa)
	ssql = fmt.Sprintf("%s ORDER BY a.tahun_akademik %s,a.nm_kelas %s", ssql, "asc", "asc")

	int_kd_trans_masuk_siswa, _ := strconv.Atoi(kd_trans_masuk_siswa)

	rows, _ := db.Raw(ssql).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&kd_trans_masuk_siswa, &kd_group, &nm_group, &kd_kategori, &nm_kategori, &tahun_akademik, &nis_siswa, &nm_siswa, &nm_kelas, &total_biaya, &total_bayar, &sisa_biaya, &ket)
		arraydata := GetBiayaAndSisa{}
		arraydata.Kd_trans_masuk_siswa = int_kd_trans_masuk_siswa
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group
		arraydata.Kd_kategori = kd_kategori
		arraydata.Nm_kategori = nm_kategori
		arraydata.Tahun_akademik = tahun_akademik
		arraydata.Nis_siswa = nis_siswa
		arraydata.Nm_siswa = nm_siswa
		arraydata.Nm_kelas = nm_kelas
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = ket

		sql := " SELECT b.kd_trans_masuk_detail_siswa,b.seqno, " +
			" b.tgl_bayar,b.jml_bayar,b.keterangan,d.kd_pembayaran,d.tipe_pembayaran " +
			" FROM tbl_trans_uang_masuk_siswa_headers a " +
			" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
			" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
			" LEFT JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran " +
			" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "

		sql = fmt.Sprintf("%s and a.kd_trans_masuk_siswa = '%s'", sql, kd_trans_masuk_siswa)

		sql = fmt.Sprintf("%s ORDER BY a.kd_trans_masuk_siswa %s,b.seqno %s", sql, "asc", "asc")

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

func UpdateUangMasukSiswaDetail(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	kd_trans_masuk_siswa := c.Param("idhead")
	kd_trans_masuk_detail_siswa := c.Param("iddetail")

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

	var dataUtama table_data.Tbl_trans_uang_masuk_siswa_details
	if err := db.Where("flag_aktif=0 and kd_trans_masuk_detail_siswa=? and kd_trans_masuk_siswa=?", kd_trans_masuk_detail_siswa, kd_trans_masuk_siswa).First(&dataUtama).Error; err != nil {
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

	tTglBayar, err2 := time.Parse("02-01-2006", paramEditUmSiswaDetail.Tgl_bayar)
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

	var dataDetail table_data.Tbl_trans_uang_masuk_siswa_details
	err = db.Raw("update tbl_trans_uang_masuk_siswa_details set tgl_bayar=?,jml_bayar=?,keterangan=?,"+
		" edited_by=?,edited_on=?,kd_pembayaran=? "+
		" where kd_trans_masuk_detail_siswa=? and kd_trans_masuk_siswa=? and flag_aktif=0 ", dateStr,
		paramEditUmSiswaDetail.Jml_bayar, paramEditUmSiswaDetail.Keterangan, currentUser.(string), datenowx,
		paramEditUmSiswaDetail.Kd_pembayaran, kd_trans_masuk_detail_siswa, kd_trans_masuk_siswa).Scan(&dataDetail).Error
	if err != nil {
		response := helper.APIResponse("Update Data Ke Tbl_trans_uang_masuk_siswa_details Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var sumJmlBayar float64
	db.Raw("SELECT ifnull(sum(jml_bayar),0.00) 'jml_bayar' FROM tbl_trans_uang_masuk_siswa_details "+
		" where flag_aktif=0 and tgl_bayar is not null and kd_trans_masuk_siswa=?", kd_trans_masuk_siswa).Scan(&sumJmlBayar)

	var total_biaya float64
	db.Raw("SELECT total_biaya FROM tbl_trans_uang_masuk_siswa_headers where flag_aktif=0 and kd_trans_masuk_siswa=?", kd_trans_masuk_siswa).Scan(&total_biaya)
	var sisa_biaya float64 = total_biaya - sumJmlBayar

	var dataHeader table_data.Tbl_trans_uang_masuk_siswa_headers
	err = db.Raw("UPDATE tbl_trans_uang_masuk_siswa_headers SET total_bayar = ?, sisa_biaya = ?, "+
		" edited_on = ? , edited_by = ? "+
		" WHERE kd_trans_masuk_siswa = ? and flag_aktif=0 ", sumJmlBayar, sisa_biaya, datenowx, currentUser.(string), kd_trans_masuk_siswa).Scan(&dataHeader).Error
	if err != nil {
		response := helper.APIResponse("Update Data Ke Tbl_trans_uang_masuk_siswa_headers Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//setting tampilan habis save siswa

	SetArrayData := []GetBiayaAndSisa{}
	//var kd_trans_masuk_siswa int
	var kd_group int
	var nm_group string
	var kd_kategori int
	var nm_kategori string
	//var total_biaya float64
	var total_bayar float64
	//var sisa_biaya float64
	var tahun_akademik string
	var nis_siswa string
	var nm_kelas string
	var nm_siswa string
	var ket string

	ssql := " SELECT distinct b.kd_trans_masuk_siswa,a.kd_group,d.nm_group,a.kd_kategori,e.nm_kategori,a.tahun_akademik,a.nis_siswa, " +
		" c.nm_siswa,a.nm_kelas,a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan " +
		" FROM tbl_trans_uang_masuk_siswa_headers a " +
		" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
		" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
		" INNER JOIN tbl_group_kategoris d on a.kd_group = d.kd_group " +
		" INNER JOIN tbl_kategori_uangs e on a.kd_kategori = e.kd_kategori " +
		" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif')  "

	ssql = fmt.Sprintf("%s and a.kd_trans_masuk_siswa= '%s'", ssql, kd_trans_masuk_siswa)
	ssql = fmt.Sprintf("%s ORDER BY a.tahun_akademik %s,a.nm_kelas %s", ssql, "asc", "asc")

	int_kd_trans_masuk_siswa, _ := strconv.Atoi(kd_trans_masuk_siswa)

	rows, _ := db.Raw(ssql).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&kd_trans_masuk_siswa, &kd_group, &nm_group, &kd_kategori, &nm_kategori, &tahun_akademik, &nis_siswa, &nm_siswa, &nm_kelas, &total_biaya, &total_bayar, &sisa_biaya, &ket)
		arraydata := GetBiayaAndSisa{}
		arraydata.Kd_trans_masuk_siswa = int_kd_trans_masuk_siswa
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group
		arraydata.Kd_kategori = kd_kategori
		arraydata.Nm_kategori = nm_kategori
		arraydata.Tahun_akademik = tahun_akademik
		arraydata.Nis_siswa = nis_siswa
		arraydata.Nm_siswa = nm_siswa
		arraydata.Nm_kelas = nm_kelas
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = ket

		sql := " SELECT b.kd_trans_masuk_detail_siswa,b.seqno, " +
			" b.tgl_bayar,b.jml_bayar,b.keterangan,d.kd_pembayaran,d.tipe_pembayaran " +
			" FROM tbl_trans_uang_masuk_siswa_headers a " +
			" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
			" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
			" LEFT JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran " +
			" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "

		sql = fmt.Sprintf("%s and a.kd_trans_masuk_siswa = '%s'", sql, kd_trans_masuk_siswa)

		sql = fmt.Sprintf("%s ORDER BY a.kd_trans_masuk_siswa %s,b.seqno %s", sql, "asc", "asc")

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

	// SetArrayData := []GetBiayaAndSisa{}
	// var kd_trans_masuk_siswaX int
	// //var total_biaya float64
	// var total_bayar float64
	// //var sisa_biaya float64
	// var tahun_akademik string
	// var nis_siswa string
	// var nm_kelas string
	// var nm_siswa string

	// ssql := " SELECT distinct b.kd_trans_masuk_siswa,a.tahun_akademik,a.nis_siswa,c.nm_siswa,a.nm_kelas,a.total_biaya,a.total_bayar,a.sisa_biaya " +
	// 	" FROM tbl_trans_uang_masuk_siswa_headers a " +
	// 	" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
	// 	" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
	// 	" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif')  "

	// ssql = fmt.Sprintf("%s and a.kd_trans_masuk_siswa= '%s'", ssql, kd_trans_masuk_siswa)

	// ssql = fmt.Sprintf("%s ORDER BY a.tahun_akademik %s,a.nm_kelas %s", ssql, "asc", "asc")
	// rows, _ := db.Raw(ssql).Rows()
	// defer rows.Close()
	// for rows.Next() {
	// 	rows.Scan(&kd_trans_masuk_siswaX, &tahun_akademik, &nis_siswa, &nm_siswa, &nm_kelas, &total_biaya, &total_bayar, &sisa_biaya)
	// 	arraydata := GetBiayaAndSisa{}
	// 	arraydata.Kd_trans_masuk_siswa = kd_trans_masuk_siswaX
	// 	arraydata.Tahun_akademik = tahun_akademik
	// 	arraydata.Nis_siswa = nis_siswa
	// 	arraydata.Nm_siswa = nm_siswa
	// 	arraydata.Nm_kelas = nm_kelas
	// 	arraydata.Total_biaya = total_biaya
	// 	arraydata.Total_bayar = total_bayar
	// 	arraydata.Sisa_biaya = sisa_biaya

	// 	sql := " SELECT b.kd_trans_masuk_detail_siswa,b.seqno, " +
	// 		" b.tgl_bayar,b.jml_bayar,b.keterangan " +
	// 		" FROM tbl_trans_uang_masuk_siswa_headers a " +
	// 		" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
	// 		" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
	// 		" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "

	// 	sql = fmt.Sprintf("%s and a.kd_trans_masuk_siswa = '%s'", sql, kd_trans_masuk_siswa)

	// 	sql = fmt.Sprintf("%s ORDER BY a.kd_trans_masuk_siswa %s,b.seqno %s", sql, "asc", "asc")

	// 	var getDataUmSiswa []GetDataUmSiswa
	// 	db.Raw(sql).Scan(&getDataUmSiswa)

	// 	arraydata.Detail = getDataUmSiswa
	// 	SetArrayData = append(SetArrayData, arraydata)
	// }

	// response := helper.APIResponse("Simpan Data Sukses ...", http.StatusOK, "success", SetArrayData)
	// c.JSON(http.StatusOK, response)
}

func CreateUangMasukSiswaDetail(c *gin.Context) {
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
	db.Raw(" SELECT count(*) jmldata FROM tbl_trans_uang_masuk_siswa_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa "+
		" where a.kd_trans_masuk_siswa=? and a.flag_aktif=0 and b.flag_aktif=0 ", paramAddDetail.Kd_trans_masuk_siswa).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Data Transaksi Siswa Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
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

	var intKd_trans_masuk_detail int
	db.Raw("SELECT ifnull(max(kd_trans_masuk_detail_siswa),0) + 1 as 'run_number' FROM tbl_trans_uang_masuk_siswa_details ").Scan(&intKd_trans_masuk_detail)

	var int_seqno int
	db.Raw("SELECT (seqno + 1) as 'run_number' "+
		" FROM tbl_trans_uang_masuk_siswa_details where flag_aktif=0 and kd_trans_masuk_siswa=?", paramAddDetail.Kd_trans_masuk_siswa).Scan(&int_seqno)

	datadetail := table_data.Tbl_trans_uang_masuk_siswa_details{
		Kd_trans_masuk_siswa:        paramAddDetail.Kd_trans_masuk_siswa,
		Kd_trans_masuk_detail_siswa: intKd_trans_masuk_detail,
		Seqno:                       int_seqno,
		Jml_bayar:                   0,
		Keterangan:                  "",
		Created_by:                  currentUser.(string),
		Created_on:                  datenowx,
		Flag_aktif:                  0,
	}

	err = db.Omit("Edited_on", "Edited_by", "Tgl_bayar", "Kd_pembayaran").Create(&datadetail).Error
	if err != nil {
		response := helper.APIResponse("Simpan Data Detail Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//setting tampilan habis save siswa
	SetArrayData := []GetBiayaAndSisa{}
	var kd_trans_masuk_siswa int
	var kd_group int
	var nm_group string
	var kd_kategori int
	var nm_kategori string
	var total_biaya float64
	var total_bayar float64
	var sisa_biaya float64
	var tahun_akademik string
	var nis_siswa string
	var nm_kelas string
	var nm_siswa string
	var ket string

	ssql := " SELECT distinct b.kd_trans_masuk_siswa,a.kd_group,d.nm_group,a.kd_kategori,e.nm_kategori,a.tahun_akademik,a.nis_siswa, " +
		" c.nm_siswa,a.nm_kelas,a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan " +
		" FROM tbl_trans_uang_masuk_siswa_headers a " +
		" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
		" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
		" INNER JOIN tbl_group_kategoris d on a.kd_group = d.kd_group " +
		" INNER JOIN tbl_kategori_uangs e on a.kd_kategori = e.kd_kategori " +
		" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif')  "

	ssql = fmt.Sprintf("%s and a.kd_trans_masuk_siswa= %d", ssql, paramAddDetail.Kd_trans_masuk_siswa)
	ssql = fmt.Sprintf("%s ORDER BY a.tahun_akademik %s,a.nm_kelas %s", ssql, "asc", "asc")

	//int_kd_trans_masuk_siswa, _ := strconv.Atoi(kd_trans_masuk_siswa)

	rows, _ := db.Raw(ssql).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&kd_trans_masuk_siswa, &kd_group, &nm_group, &kd_kategori, &nm_kategori, &tahun_akademik, &nis_siswa, &nm_siswa, &nm_kelas, &total_biaya, &total_bayar, &sisa_biaya, &ket)
		arraydata := GetBiayaAndSisa{}
		arraydata.Kd_trans_masuk_siswa = paramAddDetail.Kd_trans_masuk_siswa
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group
		arraydata.Kd_kategori = kd_kategori
		arraydata.Nm_kategori = nm_kategori
		arraydata.Tahun_akademik = tahun_akademik
		arraydata.Nis_siswa = nis_siswa
		arraydata.Nm_siswa = nm_siswa
		arraydata.Nm_kelas = nm_kelas
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = ket

		sql := " SELECT b.kd_trans_masuk_detail_siswa,b.seqno, " +
			" b.tgl_bayar,b.jml_bayar,b.keterangan,d.kd_pembayaran,d.tipe_pembayaran " +
			" FROM tbl_trans_uang_masuk_siswa_headers a " +
			" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
			" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
			" LEFT JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran " +
			" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "

		sql = fmt.Sprintf("%s and a.kd_trans_masuk_siswa = %d", sql, kd_trans_masuk_siswa)

		sql = fmt.Sprintf("%s ORDER BY a.kd_trans_masuk_siswa %s,b.seqno %s", sql, "asc", "asc")

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

	// SetArrayData := []GetBiayaAndSisa{}
	// var kd_trans_masuk_siswa int
	// var total_biaya float64
	// var total_bayar float64
	// var sisa_biaya float64
	// var tahun_akademik string
	// var nis_siswa string
	// var nm_kelas string
	// var nm_siswa string

	// ssql := " SELECT distinct b.kd_trans_masuk_siswa,a.tahun_akademik,a.nis_siswa,c.nm_siswa,a.nm_kelas,a.total_biaya,a.total_bayar,a.sisa_biaya " +
	// 	" FROM tbl_trans_uang_masuk_siswa_headers a " +
	// 	" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
	// 	" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
	// 	" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif')  "

	// ssql = fmt.Sprintf("%s and a.kd_trans_masuk_siswa= %d", ssql, paramAddDetail.Kd_trans_masuk_siswa)

	// ssql = fmt.Sprintf("%s ORDER BY a.tahun_akademik %s,a.nm_kelas %s", ssql, "asc", "asc")
	// rows, _ := db.Raw(ssql).Rows()
	// defer rows.Close()
	// for rows.Next() {
	// 	rows.Scan(&kd_trans_masuk_siswa, &tahun_akademik, &nis_siswa, &nm_siswa, &nm_kelas, &total_biaya, &total_bayar, &sisa_biaya)
	// 	arraydata := GetBiayaAndSisa{}
	// 	arraydata.Kd_trans_masuk_siswa = kd_trans_masuk_siswa
	// 	arraydata.Tahun_akademik = tahun_akademik
	// 	arraydata.Nis_siswa = nis_siswa
	// 	arraydata.Nm_siswa = nm_siswa
	// 	arraydata.Nm_kelas = nm_kelas
	// 	arraydata.Total_biaya = total_biaya
	// 	arraydata.Total_bayar = total_bayar
	// 	arraydata.Sisa_biaya = sisa_biaya

	// 	sql := " SELECT b.kd_trans_masuk_detail_siswa,b.seqno, " +
	// 		" b.tgl_bayar,b.jml_bayar,b.keterangan " +
	// 		" FROM tbl_trans_uang_masuk_siswa_headers a " +
	// 		" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
	// 		" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
	// 		" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "

	// 	sql = fmt.Sprintf("%s and a.kd_trans_masuk_siswa = %d", sql, paramAddDetail.Kd_trans_masuk_siswa)

	// 	sql = fmt.Sprintf("%s ORDER BY a.kd_trans_masuk_siswa %s,b.seqno %s", sql, "asc", "asc")

	// 	var getDataUmSiswa []GetDataUmSiswa
	// 	db.Raw(sql).Scan(&getDataUmSiswa)

	// 	arraydata.Detail = getDataUmSiswa
	// 	SetArrayData = append(SetArrayData, arraydata)
	// }

	// response := helper.APIResponse("Simpan Data Sukses ...", http.StatusOK, "success", SetArrayData)
	// c.JSON(http.StatusOK, response)

}

func ListData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramChangeSiswa ParamChangeSiswa
	// if err := c.ShouldBindJSON(&paramChangeSiswa); err != nil {
	// 	errors := helper.FormatValidationError(err)
	// 	errorMessage := gin.H{"errors": errors}
	// 	response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
	// 	c.JSON(http.StatusUnprocessableEntity, response)
	// 	return
	// }

	if err := c.ShouldBindJSON(&paramChangeSiswa); err != nil {
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

	SetArrayData := []GetBiayaAndSisa{}
	var kd_trans_masuk_siswa int
	var kd_group int
	var nm_group string
	var kd_kategori int
	var nm_kategori string
	var total_biaya float64
	var total_bayar float64
	var sisa_biaya float64
	var tahun_akademik string
	var nis_siswa string
	var nm_kelas string
	var nm_siswa string
	var ket string

	ssql := " SELECT distinct b.kd_trans_masuk_siswa,a.kd_group,d.nm_group,a.kd_kategori,e.nm_kategori,a.tahun_akademik,a.nis_siswa, " +
		" c.nm_siswa,a.nm_kelas,a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan " +
		" FROM tbl_trans_uang_masuk_siswa_headers a " +
		" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
		" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
		" INNER JOIN tbl_group_kategoris d on a.kd_group = d.kd_group " +
		" INNER JOIN tbl_kategori_uangs e on a.kd_kategori = e.kd_kategori " +
		" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif')  "

	ssql = fmt.Sprintf("%s and a.tahun_akademik= '%s'", ssql, paramChangeSiswa.Tahun_akademik)

	if paramChangeSiswa.Nm_kelas != "" {
		ssql = fmt.Sprintf("%s and a.nm_kelas= '%s'", ssql, paramChangeSiswa.Nm_kelas)
	}
	if paramChangeSiswa.Nis_siswa != "" {
		ssql = fmt.Sprintf("%s and a.nis_siswa= '%s'", ssql, paramChangeSiswa.Nis_siswa)
	}

	ssql = fmt.Sprintf("%s ORDER BY a.tahun_akademik %s,a.nm_kelas %s", ssql, "asc", "asc")
	rows, _ := db.Raw(ssql).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&kd_trans_masuk_siswa, &kd_group, &nm_group, &kd_kategori, &nm_kategori, &tahun_akademik, &nis_siswa, &nm_siswa, &nm_kelas, &total_biaya, &total_bayar, &sisa_biaya, &ket)
		arraydata := GetBiayaAndSisa{}
		arraydata.Kd_trans_masuk_siswa = kd_trans_masuk_siswa
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group
		arraydata.Kd_kategori = kd_kategori
		arraydata.Nm_kategori = nm_kategori
		arraydata.Tahun_akademik = tahun_akademik
		arraydata.Nis_siswa = nis_siswa
		arraydata.Nm_siswa = nm_siswa
		arraydata.Nm_kelas = nm_kelas
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = ket

		sql := " SELECT b.kd_trans_masuk_detail_siswa,b.seqno, " +
			" b.tgl_bayar,b.jml_bayar,b.keterangan,d.kd_pembayaran,d.tipe_pembayaran " +
			" FROM tbl_trans_uang_masuk_siswa_headers a " +
			" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
			" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
			" LEFT JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran " +
			" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "

		sql = fmt.Sprintf("%s and a.kd_trans_masuk_siswa = %d", sql, kd_trans_masuk_siswa)

		sql = fmt.Sprintf("%s ORDER BY a.kd_trans_masuk_siswa %s,b.seqno %s", sql, "asc", "asc")

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

func DeleteAllUangMasuk(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	idhead := c.Param("idhead")

	var dataUtama table_data.Tbl_trans_uang_masuk_siswa_headers
	if err := db.Where("flag_aktif=0 and kd_trans_masuk_siswa=?", idhead).First(&dataUtama).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Header Tidak Ditemukan ..."}
		response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var dataUtamaDet table_data.Tbl_trans_uang_masuk_siswa_details
	if err := db.Where("flag_aktif=0 and kd_trans_masuk_siswa=?", idhead).First(&dataUtamaDet).Error; err != nil {
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

	var dataDetail table_data.Tbl_trans_uang_masuk_siswa_details
	err = db.Raw("update tbl_trans_uang_masuk_siswa_details set flag_aktif=9,edited_by=?,edited_on=? "+
		" where kd_trans_masuk_siswa=? and flag_aktif=0 ",
		currentUser.(string), datenowx, idhead).Scan(&dataDetail).Error
	if err != nil {
		response := helper.APIResponse("Delete Data Ke Tbl_trans_uang_masuk_siswa_details Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var dataHead table_data.Tbl_trans_uang_masuk_siswa_headers
	err = db.Raw("update tbl_trans_uang_masuk_siswa_headers set flag_aktif=9,edited_by=?,edited_on=? "+
		" where kd_trans_masuk_siswa=? and flag_aktif=0 ",
		currentUser.(string), datenowx, idhead).Scan(&dataHead).Error
	if err != nil {
		response := helper.APIResponse("Delete Data Ke tbl_trans_uang_masuk_siswa_headers Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//setting tampilan habis save siswa
	SetArrayData := []GetBiayaAndSisa{}
	var kd_trans_masuk_siswa int
	var kd_group int
	var nm_group string
	var kd_kategori int
	var nm_kategori string
	var total_biaya float64
	var total_bayar float64
	var sisa_biaya float64
	var tahun_akademik string
	var nis_siswa string
	var nm_kelas string
	var nm_siswa string
	var ket string

	ssql := " SELECT distinct b.kd_trans_masuk_siswa,a.kd_group,d.nm_group,a.kd_kategori,e.nm_kategori,a.tahun_akademik,a.nis_siswa, " +
		" c.nm_siswa,a.nm_kelas,a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan " +
		" FROM tbl_trans_uang_masuk_siswa_headers a " +
		" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
		" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
		" INNER JOIN tbl_group_kategoris d on a.kd_group = d.kd_group " +
		" INNER JOIN tbl_kategori_uangs e on a.kd_kategori = e.kd_kategori " +
		" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif')  "

	//ssql = fmt.Sprintf("%s and a.kd_trans_masuk_siswa= '%s'", ssql, idhead)
	ssql = fmt.Sprintf("%s ORDER BY a.tahun_akademik %s,a.nm_kelas %s", ssql, "asc", "asc")

	//int_kd_trans_masuk_siswa, _ := strconv.Atoi(kd_trans_masuk_siswa)

	rows, _ := db.Raw(ssql).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&kd_trans_masuk_siswa, &kd_group, &nm_group, &kd_kategori, &nm_kategori, &tahun_akademik, &nis_siswa, &nm_siswa, &nm_kelas, &total_biaya, &total_bayar, &sisa_biaya, &ket)
		arraydata := GetBiayaAndSisa{}
		arraydata.Kd_trans_masuk_siswa = kd_trans_masuk_siswa
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group
		arraydata.Kd_kategori = kd_kategori
		arraydata.Nm_kategori = nm_kategori
		arraydata.Tahun_akademik = tahun_akademik
		arraydata.Nis_siswa = nis_siswa
		arraydata.Nm_siswa = nm_siswa
		arraydata.Nm_kelas = nm_kelas
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = ket

		sql := " SELECT b.kd_trans_masuk_detail_siswa,b.seqno, " +
			" b.tgl_bayar,b.jml_bayar,b.keterangan,d.kd_pembayaran,d.tipe_pembayaran " +
			" FROM tbl_trans_uang_masuk_siswa_headers a " +
			" INNER JOIN tbl_trans_uang_masuk_siswa_details b on a.kd_trans_masuk_siswa=b.kd_trans_masuk_siswa " +
			" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis " +
			" LEFT JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran " +
			" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "

		sql = fmt.Sprintf("%s and a.kd_trans_masuk_siswa = %d", sql, kd_trans_masuk_siswa)

		sql = fmt.Sprintf("%s ORDER BY a.kd_trans_masuk_siswa %s,b.seqno %s", sql, "asc", "asc")

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
