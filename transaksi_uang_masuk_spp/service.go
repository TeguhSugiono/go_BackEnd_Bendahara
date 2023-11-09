package transaksi_uang_masuk_spp

import (
	"errors"
	"net/http"
	"rest_api_bendahara/connection"
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
	rows, _ := db.Raw("SELECT kd_group FROM tbl_link_kategoris where link_name=? ", "form_biaya_spp").Rows()
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
	rows, _ := db.Raw("SELECT kd_kategori FROM tbl_link_kategoris where link_name=? ", "form_biaya_spp").Rows()
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

// func ListKelas(c *gin.Context) {
// 	dbSIA := connection.SetupConnectionSIA()

// 	var getIdAndNameKelas []GetIdAndNameKelas
// 	dbSIA.Raw("SELECT DISTINCT * FROM vw_kelas_trans").Scan(&getIdAndNameKelas)

// 	response := helper.APIResponse("List Data ...", http.StatusOK, "success", getIdAndNameKelas)
// 	c.JSON(http.StatusOK, response)
// }

func ListKelas(c *gin.Context) {
	dbSIA := connection.SetupConnectionSIA()
	db := c.MustGet("db").(*gorm.DB)

	SetArrayMerge := []GetIdAndNameKelas{}

	var str_id_kelas string
	var str_nm_kelas string
	rows, _ := dbSIA.Raw(" select tbl_kelas.nm_kelas AS id_kelas,tbl_kelas.nm_kelas AS nm_kelas " +
		" from tbl_kelas where (tbl_kelas.flag_kelas = 0) order by id_kelas").Rows()

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&str_id_kelas, &str_nm_kelas)
		arraydata := GetIdAndNameKelas{}
		arraydata.Id_kelas = str_id_kelas
		arraydata.Nm_kelas = str_nm_kelas
		SetArrayMerge = append(SetArrayMerge, arraydata)
	}

	rowBendahara, _ := db.Raw(" select distinct tbl_trans_uang_masuk_spp_headers.nm_kelas AS id_kelas, " +
		" tbl_trans_uang_masuk_spp_headers.nm_kelas AS nm_kelas " +
		" from tbl_trans_uang_masuk_spp_headers ").Rows()

	defer rowBendahara.Close()
	for rowBendahara.Next() {
		rowBendahara.Scan(&str_id_kelas, &str_nm_kelas)
		arraydataBendahara := GetIdAndNameKelas{}
		arraydataBendahara.Id_kelas = str_id_kelas
		arraydataBendahara.Nm_kelas = str_nm_kelas
		SetArrayMerge = append(SetArrayMerge, arraydataBendahara)
	}

	//SetArrayMerge = removeDuplicate(SetArrayMerge)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayMerge)
	c.JSON(http.StatusOK, response)
}

// func removeDuplicate[T string | int](sliceList []T) []T {
//     allKeys := make(map[T]bool)
//     list := []T{}
//     for _, item := range sliceList {
//         if _, value := allKeys[item]; !value {
//             allKeys[item] = true
//             list = append(list, item)
//         }
//     }
//     return list
// }

func ListSiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramChangeNmKelas ParamChangeNmKelas
	if err := c.ShouldBindJSON(&paramChangeNmKelas); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var getNisAndNameSiswa []GetNisAndNameSiswa
	db.Raw(" SELECT DISTINCT a.nis,a.nm_siswa,a.nm_kelas FROM tbl_siswa a  "+
		" LEFT JOIN tbl_trans_uang_masuk_spp_headers b on a.nis = b.nis_siswa and b.flag_aktif=0 and b.tahun_akademik=? and b.nm_kelas=? "+
		" and a.nm_kelas = b.nm_kelas "+
		" INNER JOIN tbl_kelas c on a.nm_kelas = c.nm_kelas  and c.flag_kelas = 0 "+
		" where a.flag_siswa = 0 AND a.status_siswa NOT IN ('Tidak Aktif')  "+
		" and (a.tahun_aktif = ? or a.tahun_aktif = REPLACE(?,'-','/'))  "+
		" and a.nm_kelas = ?  "+
		" ORDER BY a.nm_siswa ", paramChangeNmKelas.Tahun_akademik, paramChangeNmKelas.Nm_kelas, paramChangeNmKelas.Tahun_akademik, paramChangeNmKelas.Tahun_akademik, paramChangeNmKelas.Nm_kelas).Scan(&getNisAndNameSiswa)

	if len(getNisAndNameSiswa) == 0 {
		SetArrayData := []GetBiayaAndSisa{}
		response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", getNisAndNameSiswa)
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

	var getDataUmSpp []GetDataUmSpp
	db.Raw("SELECT b.kd_trans_masuk_detail,b.seqno,b.periode_bayar, "+
		" b.tgl_bayar,b.jml_tagihan,b.jml_bayar,b.keterangan,d.kd_pembayaran,d.tipe_pembayaran "+
		" FROM tbl_trans_uang_masuk_spp_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_spp_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
		" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis "+
		" LEFT JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran "+
		" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "+
		" and a.tahun_akademik=? and a.nm_kelas=? and a.nis_siswa = ? "+
		" order by b.seqno ", paramChangeSiswa.Tahun_akademik, paramChangeSiswa.Nm_kelas, paramChangeSiswa.Nis_siswa).Scan(&getDataUmSpp)

	SetArrayData := []GetBiayaAndSisa{}
	var kd_trans_masuk int
	var total_biaya float64
	var total_bayar float64
	var sisa_biaya float64
	var ket string
	var nmkelas string

	rows, _ := db.Raw("SELECT distinct b.kd_trans_masuk,a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan,c.nm_kelas "+
		" FROM tbl_trans_uang_masuk_spp_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_spp_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
		" INNER JOIN tbl_kelas c on a.nm_kelas = c.nm_kelas "+
		" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_kelas=0 "+
		" and a.tahun_akademik=? and a.nm_kelas=? and a.nis_siswa = ? ", paramChangeSiswa.Tahun_akademik, paramChangeSiswa.Nm_kelas, paramChangeSiswa.Nis_siswa).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&kd_trans_masuk, &total_biaya, &total_bayar, &sisa_biaya, &ket, &nmkelas)
		arraydata := GetBiayaAndSisa{}
		arraydata.Kd_trans_masuk = kd_trans_masuk
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = ket
		arraydata.Nm_kelas = nmkelas
		arraydata.Detail = getDataUmSpp
		SetArrayData = append(SetArrayData, arraydata)
	}

	if len(SetArrayData) == 0 {
		SetArrayData := []GetBiayaAndSisa{}
		response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
	c.JSON(http.StatusOK, response)
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
	var intJmldata int
	db.Raw(" SELECT count(*) jmldata FROM tbl_group_kategoris where kd_group=? and flag_aktif=0 ", paramInputSPP.Kd_group).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Kode Group Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	db.Raw(" SELECT count(*) jmldata FROM tbl_kategori_uangs where kd_kategori=? and flag_aktif=0 ", paramInputSPP.Kd_kategori).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Kode Kategori Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//cek data siswa

	db.Raw("SELECT count(*) jmldata FROM tbl_siswa where flag_siswa=0 and status_siswa not in('Tidak Aktif','LULUS') and nis=? "+
		" and (tahun_aktif=? or tahun_aktif = REPLACE(?,'-','/')) "+
		" and nm_kelas = ?", paramInputSPP.Nis_siswa, paramInputSPP.Tahun_akademik, paramInputSPP.Tahun_akademik, paramInputSPP.Nm_kelas).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Data Siswa Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	db.Raw("SELECT count(*) jmldata FROM tbl_conf_periode_spps "+
		" where flag_aktif = 0 and tahun_akademik=?  and nm_kelas=?", paramInputSPP.Tahun_akademik, paramInputSPP.Nm_kelas).Scan(&intJmldata)
	if intJmldata == 0 {
		errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
		response := helper.APIResponse("Data Configurasi SPP Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	db.Raw(" SELECT count(*) jmldata from tbl_trans_uang_masuk_spp_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_spp_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
		" where a.flag_aktif=0 and b.flag_aktif=0 and kd_group=?  "+
		" and kd_kategori=? and nis_siswa=? and nm_kelas=? "+
		" and tahun_akademik=? ", paramInputSPP.Kd_group, paramInputSPP.Kd_kategori, paramInputSPP.Nis_siswa, paramInputSPP.Nm_kelas, paramInputSPP.Tahun_akademik).Scan(&intJmldata)
	if intJmldata > 0 {
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
		db.Raw("SELECT ifnull(max(kd_trans_masuk),0) + 1 as 'run_number' FROM tbl_trans_uang_masuk_spp_headers ").Scan(&intKd_trans_masuk)

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
		db.Raw("SELECT ifnull(max(kd_trans_masuk_detail),0) + 1 as 'run_number' FROM tbl_trans_uang_masuk_spp_details ").Scan(&intKd_trans_masuk_detail)

		var int_seqno int
		var periodebayar string
		var jmltagihan float64
		rows, _ := db.Raw("SELECT seqno,CONCAT(kd_bulan,'-',tahun) 'periodebayar',biaya_spp FROM tbl_conf_periode_spps WHERE flag_aktif=0 and tahun_akademik=? and nm_kelas=? ORDER BY seqno", paramInputSPP.Tahun_akademik, paramInputSPP.Nm_kelas).Rows()
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&int_seqno, &periodebayar, &jmltagihan)

			datadetail := table_data.Tbl_trans_uang_masuk_spp_details{
				Kd_trans_masuk:        intKd_trans_masuk,
				Kd_trans_masuk_detail: intKd_trans_masuk_detail,
				Seqno:                 int_seqno,
				Periode_bayar:         periodebayar,
				Jml_tagihan:           jmltagihan,
				Jml_bayar:             0,
				Keterangan:            "",
				Created_by:            currentUser.(string),
				Created_on:            datenowx,
				Flag_aktif:            0,
			}

			err = db.Omit("Edited_on", "Edited_by", "Tgl_bayar", "Kd_pembayaran").Create(&datadetail).Error
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
			" b.tgl_bayar,b.jml_tagihan,b.jml_bayar,b.keterangan,d.kd_pembayaran,d.tipe_pembayaran "+
			" FROM tbl_trans_uang_masuk_spp_headers a "+
			" INNER JOIN tbl_trans_uang_masuk_spp_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
			" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis "+
			" LEFT JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran "+
			" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "+
			" and a.tahun_akademik=? and a.nm_kelas=? and a.nis_siswa = ? "+
			" order by b.seqno ", paramInputSPP.Tahun_akademik, paramInputSPP.Nm_kelas, paramInputSPP.Nis_siswa).Scan(&getDataUmSpp)

		SetArrayData := []GetBiayaAndSisa{}
		var kd_trans_masuk int
		var total_biaya float64
		var total_bayar float64
		var sisa_biaya float64
		var ket string
		var nmkelas string

		rowss, _ := db.Raw("SELECT distinct b.kd_trans_masuk,a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan,c.nm_kelas "+
			" FROM tbl_trans_uang_masuk_spp_headers a "+
			" INNER JOIN tbl_trans_uang_masuk_spp_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
			" INNER JOIN tbl_kelas c on a.nm_kelas = c.nm_kelas "+
			" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_kelas=0  "+
			" and a.tahun_akademik=? and a.nm_kelas=? and a.nis_siswa = ? ", paramInputSPP.Tahun_akademik, paramInputSPP.Nm_kelas, paramInputSPP.Nis_siswa).Rows()
		defer rowss.Close()
		for rowss.Next() {
			rowss.Scan(&kd_trans_masuk, &total_biaya, &total_bayar, &sisa_biaya, &ket, &nmkelas)
			arraydata := GetBiayaAndSisa{}
			arraydata.Kd_trans_masuk = kd_trans_masuk
			arraydata.Total_biaya = total_biaya
			arraydata.Total_bayar = total_bayar
			arraydata.Sisa_biaya = sisa_biaya
			arraydata.Keterangan = ket
			arraydata.Nm_kelas = nmkelas
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

	var dataDetail table_data.Tbl_trans_uang_masuk_spp_details
	err = db.Raw("update tbl_trans_uang_masuk_spp_details set tgl_bayar=?,jml_bayar=?,keterangan=?,edited_by=?,edited_on=?,kd_pembayaran=? "+
		" where kd_trans_masuk_detail=? and flag_aktif=0 ", dateStr, paramEditSPPDetail.Jml_bayar,
		paramEditSPPDetail.Keterangan, currentUser.(string), datenowx, paramEditSPPDetail.Kd_pembayaran, c.Param("iddetail")).Scan(&dataDetail).Error
	if err != nil {
		response := helper.APIResponse("Update Data Ke Tbl_trans_uang_masuk_spp_details Gagal ...", http.StatusBadRequest, "error", err)
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
		" b.tgl_bayar,b.jml_tagihan,b.jml_bayar,b.keterangan,d.kd_pembayaran,d.tipe_pembayaran "+
		" FROM tbl_trans_uang_masuk_spp_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_spp_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
		" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis "+
		" LEFT JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran "+
		" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "+
		" and a.kd_trans_masuk=? "+
		" order by b.seqno ", c.Param("idhead")).Scan(&getDataUmSpp)

	SetArrayData := []GetBiayaAndSisa{}
	var kd_trans_masuk int
	//var total_biaya float64
	var total_bayar float64
	//var sisa_biaya float64
	var ket string
	var nmkelas string

	rowss, _ := db.Raw("SELECT distinct b.kd_trans_masuk,a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan,c.nm_kelas "+
		" FROM tbl_trans_uang_masuk_spp_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_spp_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
		" INNER JOIN tbl_kelas c on a.nm_kelas = c.nm_kelas "+
		" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_kelas=0  "+
		" and a.kd_trans_masuk=? ", c.Param("idhead")).Rows()
	defer rowss.Close()
	for rowss.Next() {
		rowss.Scan(&kd_trans_masuk, &total_biaya, &total_bayar, &sisa_biaya, &ket, &nmkelas)
		arraydata := GetBiayaAndSisa{}
		arraydata.Kd_trans_masuk = kd_trans_masuk
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = ket
		arraydata.Nm_kelas = nmkelas
		arraydata.Detail = getDataUmSpp
		SetArrayData = append(SetArrayData, arraydata)
	}

	response := helper.APIResponse("Update Data Sukses ...", http.StatusOK, "success", SetArrayData)
	c.JSON(http.StatusOK, response)

}

func DeleteAllUangMasuk(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	idhead := c.Param("idhead")

	var dataUtama table_data.Tbl_trans_uang_masuk_spp_headers
	if err := db.Where("flag_aktif=0 and kd_trans_masuk=?", idhead).First(&dataUtama).Error; err != nil {
		errorMessage := gin.H{"errors": "Data Header Tidak Ditemukan ..."}
		response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var dataUtamaDet table_data.Tbl_trans_uang_masuk_spp_details
	if err := db.Where("flag_aktif=0 and kd_trans_masuk=?", idhead).First(&dataUtamaDet).Error; err != nil {
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

	var dataDetail table_data.Tbl_trans_uang_masuk_spp_details
	err = db.Raw("update tbl_trans_uang_masuk_spp_details set flag_aktif=9,edited_by=?,edited_on=? "+
		" where kd_trans_masuk=? and flag_aktif=0 ",
		currentUser.(string), datenowx, idhead).Scan(&dataDetail).Error
	if err != nil {
		response := helper.APIResponse("Delete Data Ke Tbl_trans_uang_masuk_spp_details Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var dataHead table_data.Tbl_trans_uang_masuk_spp_headers
	err = db.Raw("update tbl_trans_uang_masuk_spp_headers set flag_aktif=9,edited_by=?,edited_on=? "+
		" where kd_trans_masuk=? and flag_aktif=0 ",
		currentUser.(string), datenowx, idhead).Scan(&dataHead).Error
	if err != nil {
		response := helper.APIResponse("Delete Data Ke tbl_trans_uang_masuk_spp_headers Gagal ...", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//setting tampilan habis update spp

	var getDataUmSpp []GetDataUmSpp
	db.Raw("SELECT b.kd_trans_masuk_detail,b.seqno,b.periode_bayar, "+
		" b.tgl_bayar,b.jml_tagihan,b.jml_bayar,b.keterangan,d.kd_pembayaran,d.tipe_pembayaran "+
		" FROM tbl_trans_uang_masuk_spp_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_spp_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
		" INNER JOIN tbl_siswa c on a.nis_siswa = c.nis "+
		" LEFT JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran "+
		" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_siswa = 0 and status_siswa not in('Tidak Aktif') "+
		" and a.kd_trans_masuk=? "+
		" order by b.seqno ", idhead).Scan(&getDataUmSpp)

	SetArrayData := []GetBiayaAndSisa{}
	var kd_trans_masuk int
	var total_biaya float64
	var total_bayar float64
	var sisa_biaya float64
	var ket string
	var nmkelas string

	rowss, _ := db.Raw("SELECT distinct b.kd_trans_masuk,a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan,c.nm_kelas "+
		" FROM tbl_trans_uang_masuk_spp_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_spp_details b on a.kd_trans_masuk=b.kd_trans_masuk "+
		" INNER JOIN tbl_kelas c on a.nm_kelas = c.nm_kelas "+
		" where a.flag_aktif=0 and b.flag_aktif=0 and c.flag_kelas=0 "+
		" and a.kd_trans_masuk=? ", idhead).Rows()
	defer rowss.Close()
	for rowss.Next() {
		rowss.Scan(&kd_trans_masuk, &total_biaya, &total_bayar, &sisa_biaya, &ket, &nmkelas)
		arraydata := GetBiayaAndSisa{}
		arraydata.Kd_trans_masuk = kd_trans_masuk
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = ket
		arraydata.Nm_kelas = nmkelas
		arraydata.Detail = getDataUmSpp
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
