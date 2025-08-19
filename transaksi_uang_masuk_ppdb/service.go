package transaksi_uang_masuk_ppdb

import (
	"errors"
	"fmt"
	"math"
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

	var paramPPdb ParamPPdb
	if err := c.ShouldBindJSON(&paramPPdb); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	ssql := " SELECT nik,concat('Double Data ==> ','',nm_siswa) 'nm_siswa' FROM tbl_user_ppdb " +
	" WHERE (nik <> '' or nik is not null) " +
	" and status = 'sudah diverifikasi' and flag_verifikasidata='1' and flag_wawancara='1' " +
	" and flag_pembayaran='1' and flag=0 and status_berkas <> 'DiCabut' and flag_import<>'9' "

	if paramPPdb.Tahun_daftar != "" {
		ssql = fmt.Sprintf("%s and tahun_daftar = '%s'", ssql, paramPPdb.Tahun_daftar)
	}

	ssql = fmt.Sprintf("%s GROUP BY nik HAVING count(*) > 1  ", ssql)

	SetArrayData := []GetNikAndNameSiswa{}
	var nik string
	var nm_siswa string
	rows, _ := db.Raw(ssql).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&nik, &nm_siswa)
		arraydata := GetNikAndNameSiswa{}
		arraydata.Nik = nik
		arraydata.Nm_siswa = nm_siswa
		SetArrayData = append(SetArrayData, arraydata)
	}

	if len(SetArrayData) > 0 {
		response := helper.APIResponse("Terdapat Data Siswa Yang Double ...", http.StatusUnprocessableEntity, "error", SetArrayData)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	ssqls := " SELECT nik,nm_siswa FROM tbl_user_ppdb " +
	" WHERE (nik <> '' or nik is not null) " +
	" and status = 'sudah diverifikasi' and flag_verifikasidata='1' and flag_wawancara='1' " +
	" and flag_pembayaran='1' and flag=0 and status_berkas <> 'DiCabut' and flag_import<>'9' "

	if paramPPdb.Tahun_daftar != "" {
		ssqls = fmt.Sprintf("%s and tahun_daftar = '%s'", ssqls, paramPPdb.Tahun_daftar)
	}
	ssqls = fmt.Sprintf("%s order by nm_siswa  ", ssqls)
	SetArrayDatax := []GetNikAndNameSiswa{}
	rowsd, _ := db.Raw(ssqls).Rows()
	defer rowsd.Close()
	for rowsd.Next() {
		rowsd.Scan(&nik, &nm_siswa)
		arraydata := GetNikAndNameSiswa{}
		arraydata.Nik = nik
		arraydata.Nm_siswa = nm_siswa
		SetArrayDatax = append(SetArrayDatax, arraydata)
	}

	// var getNikAndNameSiswa []GetNikAndNameSiswa
	// db.Raw(" SELECT nik,nm_siswa FROM tbl_user_ppdb WHERE (nik <> '' or nik is not null)  " +
	// 	" and status = 'sudah diverifikasi' and flag_verifikasidata='1' and flag_wawancara='1' " +
	// 	" and flag_pembayaran='1' and flag=0 and status_berkas <> 'DiCabut' and flag_import<>'9' order by nm_siswa").Scan(&getNikAndNameSiswa)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayDatax)
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
		" b.tgl_bayar,b.jml_bayar,d.kd_pembayaran,d.tipe_pembayaran "+
		" FROM tbl_trans_uang_masuk_ppdb_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
		" INNER JOIN tbl_user_ppdb c on a.nik = c.nik "+
		" LEFT JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran "+
		" where a.flag_aktif=0 and b.flag_aktif=0 and c.status = 'sudah diverifikasi' and c.flag_verifikasidata='1' and c.flag_wawancara='1' "+
		" and c.flag_pembayaran='1' and c.flag=0 and flag_import<>'9' and status_berkas <> 'DiCabut' and a.nik=? "+
		" order by b.seqno ", paramChangeNik.Nik).Scan(&getDataPPDB)

		SetArrayData := []GetBiayaAndSisa{}
		var kd_trans_masuk_ppdb int
		var tgldaftar string
		var tahun_daftar string
		var tahun_akademik string
		var total_biaya float64
		var total_bayar float64
		var sisa_biaya float64
		var ket string
		rows, _ := db.Raw("SELECT distinct b.kd_trans_masuk_ppdb,convert(a.tgldaftar,CHAR) 'tgldaftar', "+
		" a.tahun_daftar,a.tahun_akademik,a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan "+
		" FROM tbl_trans_uang_masuk_ppdb_headers a "+
		" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
		" where a.flag_aktif=0 and b.flag_aktif=0  "+
		" and a.nik=? ", paramChangeNik.Nik).Rows()
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&kd_trans_masuk_ppdb, &tgldaftar, &tahun_daftar, &tahun_akademik, &total_biaya, &total_bayar, &sisa_biaya, &ket)
			arraydata := GetBiayaAndSisa{}
			arraydata.Kd_trans_masuk_ppdb = kd_trans_masuk_ppdb
			tTglDaftar, _ := time.Parse("2006-01-02", tgldaftar)
			dateStrTglDaftar := tTglDaftar.Format("02-01-2006")
			arraydata.Tgldaftar = dateStrTglDaftar
			arraydata.Tahun_daftar = tahun_daftar
			arraydata.Tahun_akademik = tahun_akademik
			arraydata.Total_biaya = total_biaya
			arraydata.Total_bayar = total_bayar
			arraydata.Sisa_biaya = sisa_biaya
			arraydata.Keterangan = ket
			arraydata.Detail = getDataPPDB
			SetArrayData = append(SetArrayData, arraydata)
		}

		// if kd_trans_masuk_ppdb == 0 {
			// 	SetArrayKosong := []GetBiayaAndSisa{}
			// 	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayKosong)
			// 	c.JSON(http.StatusOK, response)
			// 	return
			// }

			if len(getDataPPDB) == 0 {

				rows1, _ := db.Raw(" SELECT CONVERT(tgldaftar,CHAR) 'tgldaftar',tahun_daftar, "+
				" (SELECT CAST(REPLACE(REPLACE(TRIM(jumlah_pembayaran), 'Rp.', ''), '.', '') AS UNSIGNED) from tbl_biayadaftar where id=id_biaya) 'total_biaya' "+
				" FROM tbl_user_ppdb WHERE (nik <> '' or nik is not null)  "+
				" and status = 'sudah diverifikasi' and flag_verifikasidata='1' and flag_wawancara='1' "+
				" and flag_pembayaran='1' and flag=0 and status_berkas <> 'DiCabut' and flag_import<>'9' and nik=?", paramChangeNik.Nik).Rows()
				defer rows1.Close()
				for rows1.Next() {
					rows1.Scan(&tgldaftar, &tahun_daftar, &total_biaya)
					arraydata := GetBiayaAndSisa{}
					arraydata.Kd_trans_masuk_ppdb = 0
					tTglDaftar, _ := time.Parse("2006-01-02", tgldaftar)
					dateStrTglDaftar := tTglDaftar.Format("02-01-2006")
					arraydata.Tgldaftar = dateStrTglDaftar
					arraydata.Tahun_daftar = tahun_daftar

					sql := "SELECT tahun_akademik FROM tbl_tahun_akademik where flag_tahun=0 "
					sql = fmt.Sprintf("%s and tahun_akademik LIKE '%s%%%%' ", sql, tahun_daftar)
					db.Raw(sql).Scan(&tahun_akademik)

					arraydata.Tahun_akademik = tahun_akademik

					arraydata.Total_biaya = total_biaya
					arraydata.Total_bayar = 0
					arraydata.Sisa_biaya = total_biaya

					//arraydata.Detail = getDataPPDB

					//var getDataPPDB_Null []GetDataPPDB
					SetNullArray := []GetDataPPDB{}
					arraydata.Detail = SetNullArray

					SetArrayData = append(SetArrayData, arraydata)
				}

				response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
				c.JSON(http.StatusOK, response)
				return
			} else {
				response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
				c.JSON(http.StatusOK, response)
			}
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

			var intJmldata int
			db.Raw(" SELECT count(*) jmldata FROM tbl_group_kategoris where kd_group=? and flag_aktif=0 ", paramInputPPdb.Kd_group).Scan(&intJmldata)
			if intJmldata == 0 {
				errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
				response := helper.APIResponse("Kode Group Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
				c.JSON(http.StatusUnprocessableEntity, response)
				return
			}
			db.Raw(" SELECT count(*) jmldata FROM tbl_kategori_uangs where kd_kategori=? and flag_aktif=0 ", paramInputPPdb.Kd_kategori).Scan(&intJmldata)
			if intJmldata == 0 {
				errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
				response := helper.APIResponse("Kode Kategori Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
				c.JSON(http.StatusUnprocessableEntity, response)
				return
			}

			//cek data siswa ppdb

			db.Raw("SELECT count(*) jmldata "+
			" FROM tbl_user_ppdb WHERE nik = ? and tgldaftar=? and tahun_daftar=? "+
			" and status = 'sudah diverifikasi' and flag_verifikasidata='1' and flag_wawancara='1' "+
			" and flag_pembayaran='1' and flag=0 and status_berkas <> 'DiCabut' and flag_import<>'9'", paramInputPPdb.Nik, dateStrTglDaftar, paramInputPPdb.Tahun_daftar).Scan(&intJmldata)
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
				db.Raw("SELECT ifnull(max(kd_trans_masuk_ppdb),0) + 1 as 'run_number' FROM tbl_trans_uang_masuk_ppdb_headers ").Scan(&intKd_trans_masuk)

				var tgldaftar string
				var tahun_daftar string
				var total_biaya float64
				var total_bayar float64
				rows, _ := db.Raw(" SELECT CONVERT(tgldaftar,CHAR) 'tgldaftar',tahun_daftar, "+
				" (SELECT CAST(REPLACE(REPLACE(TRIM(jumlah_pembayaran), 'Rp.', ''), '.', '') AS UNSIGNED) from tbl_biayadaftar where id=id_biaya) 'total_biaya',0 as 'total_bayar' "+
				" FROM tbl_user_ppdb WHERE (nik <> '' or nik is not null)  "+
				" and status = 'sudah diverifikasi' and flag_verifikasidata='1' and flag_wawancara='1' "+
				" and flag_pembayaran='1' and flag=0 and status_berkas <> 'DiCabut' and flag_import<>'9' and nik=?", paramInputPPdb.Nik).Rows()
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
					Total_bayar:         total_biaya,
					Sisa_biaya:          0,
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
				db.Raw("SELECT ifnull(max(kd_trans_masuk_detail_ppdb),0) + 1 as 'run_number' FROM tbl_trans_uang_masuk_ppdb_details ").Scan(&intKd_trans_masuk_detail)

				var int_seqno int = 1

				datadetail := table_data.Tbl_trans_uang_masuk_ppdb_details{
					Kd_trans_masuk_ppdb:        intKd_trans_masuk,
					Kd_trans_masuk_detail_ppdb: intKd_trans_masuk_detail,
					Seqno:                      int_seqno,
					Kategori_biaya_ppdb:        "Biaya Import dari PPDB",
					Tgl_bayar:                  dateStrTglDaftar, //tgl byar belum tau dapat dri mana untuk sementara saya ambil dari tgl daftar
					Jml_bayar:                  total_biaya,      //biaya cicilan belum tau dapat dri mana untuk sementara saya ambil dari header nya
					Created_by:                 currentUser.(string),
					Created_on:                 datenowx,
					Flag_aktif:                 0,
				}

				err = db.Omit("Edited_on", "Edited_by", "Keterangan", "Kd_pembayaran").Create(&datadetail).Error
				if err != nil {
					response := helper.APIResponse("Simpan Data Detail Gagal ...", http.StatusBadRequest, "error", err)
					c.JSON(http.StatusBadRequest, response)
					return
				}

				intKd_trans_masuk_detail++
				int_seqno++

			}

			//setting tampilan habis save ppdb

			var getDataPPDB []GetDataPPDB
			db.Raw(" SELECT DISTINCT b.kd_trans_masuk_detail_ppdb,b.seqno,b.kategori_biaya_ppdb, "+
				" b.tgl_bayar,b.jml_bayar,d.kd_pembayaran,d.tipe_pembayaran "+
				" FROM tbl_trans_uang_masuk_ppdb_headers a "+
				" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
				" INNER JOIN tbl_user_ppdb c on a.nik=c.nik and a.tgldaftar=c.tgldaftar and a.tahun_daftar=c.tahun_daftar "+
				" LEFT JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran "+
				" where a.flag_aktif=0 and b.flag_aktif=0 and flag_import<>'9' and a.nik=? and a.tahun_daftar=? and a.tgldaftar=? "+
				" ORDER BY seqno ", paramInputPPdb.Nik, paramInputPPdb.Tahun_daftar, dateStrTglDaftar).Scan(&getDataPPDB)

				SetArrayData := []GetBiayaAndSisa{}
				var kd_trans_masuk_ppdb int
				var total_biaya float64
				var total_bayar float64
				var sisa_biaya float64
				var tgldaftar string
				var tahun_daftar string
				var tahun_akademik string
				var ket string
				rowss, _ := db.Raw("SELECT distinct b.kd_trans_masuk_ppdb,CONVERT(a.tgldaftar,CHAR) 'tgldaftar',a.tahun_daftar,a.tahun_akademik,a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan "+
				" FROM tbl_trans_uang_masuk_ppdb_headers a "+
				" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
				" where a.flag_aktif=0 and b.flag_aktif=0  "+
				" and a.nik=? and a.tahun_daftar=? and a.tgldaftar=? ", paramInputPPdb.Nik, paramInputPPdb.Tahun_daftar, dateStrTglDaftar).Rows()
				defer rowss.Close()
				for rowss.Next() {
					rowss.Scan(&kd_trans_masuk_ppdb, &tgldaftar, &tahun_daftar, &tahun_akademik, &total_biaya, &total_bayar, &sisa_biaya, &ket)
					arraydata := GetBiayaAndSisa{}
					arraydata.Kd_trans_masuk_ppdb = kd_trans_masuk_ppdb
					tTglDaftar, _ := time.Parse("2006-01-02", tgldaftar)
					dateStrTglDaftar := tTglDaftar.Format("02-01-2006")
					arraydata.Tgldaftar = dateStrTglDaftar
					arraydata.Tahun_daftar = tahun_daftar
					arraydata.Tahun_akademik = tahun_akademik
					arraydata.Total_biaya = total_biaya
					arraydata.Total_bayar = total_bayar
					arraydata.Sisa_biaya = sisa_biaya
					arraydata.Keterangan = ket
					arraydata.Detail = getDataPPDB
					SetArrayData = append(SetArrayData, arraydata)
				}

				response := helper.APIResponse("Simpan Data Sukses ...", http.StatusOK, "success", SetArrayData)
				c.JSON(http.StatusOK, response)

			}

			func ImportAll(c *gin.Context) {
				db := c.MustGet("db").(*gorm.DB)

				var paramDataPPDB ParamDataPPDB
				if err := c.ShouldBindJSON(&paramDataPPDB); err != nil {
					errors := helper.FormatValidationError(err)
					errorMessage := gin.H{"errors": errors}
					response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
					c.JSON(http.StatusUnprocessableEntity, response)
					return
				}

				var Strnik string = ""
				var StatusImport string
				var Nik string

				rowsData, _ := db.Raw(" SELECT c.nik, "+
					" (SELECT if(count(*) > 0,'Sukses','') FROM tbl_trans_uang_masuk_ppdb_headers a "+
					" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
					" where a.flag_aktif=0 and b.flag_aktif=0 and a.nik=c.nik  GROUP BY a.nik) 'StatusImport' "+
					" FROM tbl_user_ppdb c WHERE (c.nik <> '' or c.nik is not null)   "+
					" and c.status = 'sudah diverifikasi' and c.flag_verifikasidata='1' and c.flag_wawancara='1'  "+
					" and c.flag_pembayaran='1' and c.flag=0 and c.status_berkas <> 'DiCabut' and c.flag_import<>'9' and c.tahun_daftar=? "+
					" ORDER BY c.nm_siswa ", paramDataPPDB.Tahun_daftar).Rows()

					defer rowsData.Close()
					for rowsData.Next() {
						StatusImport = ""
						Nik = ""

						rowsData.Scan(&Nik, &StatusImport)

						if StatusImport == "" && Strnik == "" {
							Strnik = "'" + Nik + "'"
						} else if StatusImport == "" {
							Strnik = Strnik + ",'" + Nik + "'"
						}
					}

					if Strnik == "" {
						response := helper.APIResponse("Data Sudah Sinkron Semua...", http.StatusUnprocessableEntity, "error", nil)
						c.JSON(http.StatusUnprocessableEntity, response)
						return
					}

					var tgldaftar string
					var tahun_daftar string
					var total_biaya float64
					var total_bayar float64
					var nik string
					var tahun_akademik string

					rows, _ := db.Raw(" SELECT nik,CONVERT(tgldaftar,CHAR) 'tgldaftar',tahun_daftar, " +
					" (SELECT CAST(REPLACE(REPLACE(TRIM(jumlah_pembayaran), 'Rp.', ''), '.', '') AS UNSIGNED) from tbl_biayadaftar where id=id_biaya) 'total_biaya',0 as 'total_bayar' " +
					" FROM tbl_user_ppdb WHERE (nik <> '' or nik is not null)  " +
					" and status = 'sudah diverifikasi' and flag_verifikasidata='1' and flag_wawancara='1' " +
					" and flag_pembayaran='1' and flag=0 and status_berkas <> 'DiCabut' and flag_import<>'9' and nik in(" + Strnik + ")  ").Rows()
					defer rows.Close()
					for rows.Next() {
						rows.Scan(&nik, &tgldaftar, &tahun_daftar, &total_biaya, &total_bayar)

						var intKd_trans_masuk int
						db.Raw("SELECT ifnull(max(kd_trans_masuk_ppdb),0) + 1 as 'run_number' FROM tbl_trans_uang_masuk_ppdb_headers ").Scan(&intKd_trans_masuk)

						sql := "SELECT tahun_akademik FROM tbl_tahun_akademik where flag_tahun=0 "
						sql = fmt.Sprintf("%s and tahun_akademik LIKE '%s%%%%' ", sql, tahun_daftar)
						db.Raw(sql).Scan(&tahun_akademik)

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
						data := table_data.Tbl_trans_uang_masuk_ppdb_headers{
							Kd_group:            2,
							Kd_kategori:         2,
							Kd_trans_masuk_ppdb: intKd_trans_masuk,
							Nik:                 nik,
							Tgldaftar:           tgldaftar,
							Tahun_daftar:        tahun_daftar,
							Tahun_akademik:      tahun_akademik,
							Total_biaya:         total_biaya,
							Total_bayar:         total_biaya,
							Sisa_biaya:          0,
							Keterangan:          "",
							Created_by:          currentUser.(string),
							Created_on:          datenowx,
							Flag_aktif:          0,
						}
						err = db.Omit("Edited_on", "Edited_by").Create(&data).Error
						if err != nil {
							response := helper.APIResponse("Simpan Data Header Gagal ...", http.StatusBadRequest, "error", err)
							c.JSON(http.StatusBadRequest, response)
							return
						}

						var intKd_trans_masuk_detail int
						db.Raw("SELECT ifnull(max(kd_trans_masuk_detail_ppdb),0) + 1 as 'run_number' FROM tbl_trans_uang_masuk_ppdb_details ").Scan(&intKd_trans_masuk_detail)

						var int_seqno int = 1

						datadetail := table_data.Tbl_trans_uang_masuk_ppdb_details{
							Kd_trans_masuk_ppdb:        intKd_trans_masuk,
							Kd_trans_masuk_detail_ppdb: intKd_trans_masuk_detail,
							Seqno:                      int_seqno,
							Kategori_biaya_ppdb:        "Biaya Import dari PPDB",
							Tgl_bayar:                  tgldaftar,   //tgl byar belum tau dapat dri mana untuk sementara saya ambil dari tgl daftar
							Jml_bayar:                  total_biaya, //biaya cicilan belum tau dapat dri mana untuk sementara saya ambil dari header nya
							Created_by:                 currentUser.(string),
							Created_on:                 datenowx,
							Flag_aktif:                 0,
							Kd_pembayaran:              1,
						}

						err = db.Omit("Edited_on", "Edited_by", "Keterangan", "Kd_pembayaran").Create(&datadetail).Error
						if err != nil {
							response := helper.APIResponse("Simpan Data Detail Gagal ...", http.StatusBadRequest, "error", err)
							c.JSON(http.StatusBadRequest, response)
							return
						}

					}

					pesan := gin.H{"success": "Simpan Data Detail Berhasil ..."}
					response := helper.APIResponse("Import Data Sukses ...", http.StatusOK, "success", pesan)
					c.JSON(http.StatusOK, response)

					// var totalDataPendaftar int
					// var jmlDataSudahImport int
					// var jmlDataBelumImport int

					// db.Raw("SELECT count(*) jmldata "+
					// 	" FROM tbl_user_ppdb WHERE tahun_daftar=? "+
					// 	" and status = 'sudah diverifikasi' and flag_verifikasidata='1' and flag_wawancara='1' "+
					// 	" and flag_pembayaran='1' and flag=0 and status_berkas <> 'DiCabut' and flag_import<>'9'", paramDataPPDB.Tahun_daftar).Scan(&totalDataPendaftar)
					// // if totalDataPendaftar == 0 {
						// 	// errorMessage := gin.H{"errors": "Pencarian Data Pendaftar PPDB Error ..."}
						// 	// response := helper.APIResponse("Data Siswa Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
						// 	// c.JSON(http.StatusUnprocessableEntity, response)
						// // 	return
						// // }

						// db.Raw(" SELECT COUNT(jmldata) as 'jmldata' from (SELECT count(*) jmldata FROM tbl_trans_uang_masuk_ppdb_headers a "+
						// 	" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
						// 	" where a.tahun_daftar=?  and a.flag_aktif=0 and b.flag_aktif=0 GROUP BY a.nik) as jmldata ", paramDataPPDB.Tahun_daftar).Scan(&jmlDataSudahImport)

						// jmlDataBelumImport = totalDataPendaftar - jmlDataSudahImport

						// SetArrayData := []ListDataPPDB{}
						// arraydata := ListDataPPDB{}
						// arraydata.TotalDataPendaftar = totalDataPendaftar
						// arraydata.JmlDataSudahImport = jmlDataSudahImport
						// arraydata.JmlDataBelumImport = jmlDataBelumImport

						// var listDataPPDBDetail []ListDataPPDBDetail
						// db.Raw(" SELECT c.nik,c.nm_siswa,CONVERT(DATE_FORMAT(c.tgldaftar,'%d-%m-%Y'),CHAR) 'tgldaftar', "+
						// 	" (SELECT if(count(*) > 0,'Sukses','') FROM tbl_trans_uang_masuk_ppdb_headers a "+
						// 	" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
						// 	" where a.flag_aktif=0 and b.flag_aktif=0 and a.nik=c.nik  GROUP BY a.nik) 'StatusImport' "+
						// 	" FROM tbl_user_ppdb c WHERE (c.nik <> '' or c.nik is not null)   "+
						// 	" and c.status = 'sudah diverifikasi' and c.flag_verifikasidata='1' and c.flag_wawancara='1'  "+
						// 	" and c.flag_pembayaran='1' and c.flag=0 and c.status_berkas <> 'DiCabut' and c.flag_import<>'9' and c.tahun_daftar=? "+
						// 	" ORDER BY (SELECT if(count(*) > 0,'Sukses','') FROM tbl_trans_uang_masuk_ppdb_headers a  "+
						// 	" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
						// 	" where a.flag_aktif=0 and b.flag_aktif=0 and a.nik=c.nik  GROUP BY a.nik) desc, c.nm_siswa ", paramDataPPDB.Tahun_daftar).Scan(&listDataPPDBDetail)

						// arraydata.Detail = listDataPPDBDetail
						// SetArrayData = append(SetArrayData, arraydata)

						// response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
						// c.JSON(http.StatusOK, response)
					}

					func DataListPPDB(c *gin.Context) {
						db := c.MustGet("db").(*gorm.DB)

						var paramDataPPDB ParamDataPPDB
						if err := c.ShouldBindJSON(&paramDataPPDB); err != nil {
							errors := helper.FormatValidationError(err)
							errorMessage := gin.H{"errors": errors}
							response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
							c.JSON(http.StatusUnprocessableEntity, response)
							return
						}

						var totalDataPendaftar int
						var jmlDataSudahImport int
						var jmlDataBelumImport int

						db.Raw("SELECT count(*) jmldata "+
						" FROM tbl_user_ppdb WHERE (REPLACE(nik,' ','') <> '' or REPLACE(nik,' ','') is not null) and tahun_daftar=? "+
						" and status = 'sudah diverifikasi' and flag_verifikasidata='1' and flag_wawancara='1' "+
						" and flag_pembayaran='1' and flag=0 and status_berkas <> 'DiCabut' and flag_import<>'9'", paramDataPPDB.Tahun_daftar).Scan(&totalDataPendaftar)
						if totalDataPendaftar == 0 {
							errorMessage := gin.H{"errors": "Pencarian Data Pendaftar PPDB Error ..."}
							response := helper.APIResponse("Data Siswa Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
							c.JSON(http.StatusUnprocessableEntity, response)
							return
						}

						db.Raw(" SELECT COUNT(jmldata) as 'jmldata' from (SELECT count(*) jmldata FROM tbl_trans_uang_masuk_ppdb_headers a "+
						" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
						" where a.tahun_daftar=?  and a.flag_aktif=0 and b.flag_aktif=0 GROUP BY a.nik) as jmldata ", paramDataPPDB.Tahun_daftar).Scan(&jmlDataSudahImport)

						jmlDataBelumImport = totalDataPendaftar - jmlDataSudahImport

						var listDataPPDBDetail []ListDataPPDBDetail
						sql := " SELECT c.nik,c.nm_siswa,CONVERT(DATE_FORMAT(c.tgldaftar,'%d-%m-%Y'),CHAR) 'tgldaftar', " +
						" (SELECT if(count(*) > 0,'Sukses','') FROM tbl_trans_uang_masuk_ppdb_headers a " +
						" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb " +
						" where a.flag_aktif=0 and b.flag_aktif=0 and a.nik=c.nik  GROUP BY a.nik) 'StatusImport' " +
						" FROM tbl_user_ppdb c WHERE (c.nik <> '' or c.nik is not null)   " +
						" and c.status = 'sudah diverifikasi' and c.flag_verifikasidata='1' and c.flag_wawancara='1'  " +
						" and c.flag_pembayaran='1' and c.flag=0 and c.status_berkas <> 'DiCabut' and c.flag_import<>'9'  "

						sql = fmt.Sprintf("%s and c.tahun_daftar='%s'", sql, paramDataPPDB.Tahun_daftar)

						if s := c.Query("search"); s != "" {
							if len(c.Query("search")) >= 3 {
								sql = fmt.Sprintf("%s and (c.nm_siswa LIKE '%%%s%%' ", sql, s)
								sql = fmt.Sprintf("%s or c.nik LIKE '%%%s%%') ", sql, s)
							}
						}

						if sort := c.Query("sort"); sort != "" {
							sql = fmt.Sprintf("%s ORDER BY (SELECT if(count(*) > 0,'Sukses','') FROM tbl_trans_uang_masuk_ppdb_headers a  "+
							" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
							" where a.flag_aktif=0 and b.flag_aktif=0 and a.nik=c.nik  GROUP BY a.nik) %s,c.nm_siswa asc", sql, "desc")
						} else {
							sql = fmt.Sprintf("%s ORDER BY (SELECT if(count(*) > 0,'Sukses','') FROM tbl_trans_uang_masuk_ppdb_headers a  "+
							" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
							" where a.flag_aktif=0 and b.flag_aktif=0 and a.nik=c.nik  GROUP BY a.nik) %s,c.nm_siswa asc", sql, "asc")
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
						db.Raw(sql).Scan(&listDataPPDBDetail)

						CompTableData := TableDataList{
							TotalDataPendaftar: totalDataPendaftar,
							JmlDataSudahImport: jmlDataSudahImport,
							JmlDataBelumImport: jmlDataBelumImport,
							Total:              total,
							Page:               intpage,
							Last_page:          int(math.Ceil(float64(total) / float64(intperPage))),
						}

						response := helper.APIResponseTable("List Data ...", http.StatusOK, "success", "", CompTableData, listDataPPDBDetail)
						c.JSON(http.StatusOK, response)
					}

					func UpdateUangMasukPPdb(c *gin.Context) {
						db := c.MustGet("db").(*gorm.DB)

						var paramEditPPdbDetail ParamEditPPdbDetail
						if err := c.ShouldBindJSON(&paramEditPPdbDetail); err != nil {
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

						var dataUtama table_data.Tbl_trans_uang_masuk_ppdb_details
						if err := db.Where("flag_aktif=0 and kd_trans_masuk_detail_ppdb=?", c.Param("iddetail")).First(&dataUtama).Error; err != nil {
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

						tTglBayar, err2 := time.Parse("02-01-2006", paramEditPPdbDetail.Tgl_bayar)
						if err2 != nil {
							errors := helper.FormatValidationError(err2)
							errorMessage := gin.H{"errors": errors, "date": tTglBayar}
							response := helper.APIResponse("Tanggal Format Salah ...", http.StatusUnprocessableEntity, "error", errorMessage)
							c.JSON(http.StatusUnprocessableEntity, response)
							return
						}
						dateStr := tTglBayar.Format("2006-01-02")

						var dataDetail table_data.Tbl_trans_uang_masuk_ppdb_details
						err = db.Raw("update tbl_trans_uang_masuk_ppdb_details set tgl_bayar=?,jml_bayar=?,kategori_biaya_ppdb=?,edited_by=?,edited_on=?,kd_pembayaran=? "+
							" where kd_trans_masuk_detail_ppdb=? and flag_aktif=0 ", dateStr, paramEditPPdbDetail.Jml_bayar,
							paramEditPPdbDetail.Kategori_biaya_ppdb, currentUser.(string), datenowx, paramEditPPdbDetail.Kd_pembayaran,
							c.Param("iddetail")).Scan(&dataDetail).Error
							if err != nil {
								response := helper.APIResponse("Update Data Ke Tbl_trans_uang_masuk_ppdb_details Gagal ...", http.StatusBadRequest, "error", err)
								c.JSON(http.StatusBadRequest, response)
								return
							}

							var sumJmlBayar float64
							db.Raw("SELECT sum(jml_bayar) 'jml_bayar' FROM tbl_trans_uang_masuk_ppdb_details "+
							" where flag_aktif=0 and tgl_bayar is not null and kd_trans_masuk_ppdb=?", c.Param("idhead")).Scan(&sumJmlBayar)

							var total_biaya float64
							// db.Raw("SELECT total_biaya FROM tbl_trans_uang_masuk_ppdb_headers where flag_aktif=0 and kd_trans_masuk_ppdb=?", c.Param("idhead")).Scan(&total_biaya)
							// total_biaya = total_biaya + sumJmlBayar

							var total_bayar float64
							// db.Raw("SELECT total_bayar FROM tbl_trans_uang_masuk_ppdb_headers where flag_aktif=0 and kd_trans_masuk_ppdb=?", c.Param("idhead")).Scan(&total_bayar)
							// total_bayar = total_bayar + sumJmlBayar

							var sisa_biaya float64

							var dataHeader table_data.Tbl_trans_uang_masuk_ppdb_headers
							err = db.Raw("UPDATE tbl_trans_uang_masuk_ppdb_headers SET total_bayar = ?,total_biaya = ?, sisa_biaya = ?, "+
								" edited_on = ? , edited_by = ? "+
								" WHERE kd_trans_masuk_ppdb = ? and flag_aktif=0 ", sumJmlBayar, sumJmlBayar, 0, datenowx, currentUser.(string), c.Param("idhead")).Scan(&dataHeader).Error
								if err != nil {
									response := helper.APIResponse("Update Data Ke Tbl_trans_uang_masuk_ppdb_headers Gagal ...", http.StatusBadRequest, "error", err)
									c.JSON(http.StatusBadRequest, response)
									return
								}

								//setting tampilan habis update ppdb
								var getDataPPDB []GetDataPPDB
								db.Raw(" SELECT DISTINCT b.kd_trans_masuk_detail_ppdb,b.seqno,b.kategori_biaya_ppdb, "+
									" b.tgl_bayar,b.jml_bayar,d.kd_pembayaran,d.tipe_pembayaran "+
									" FROM tbl_trans_uang_masuk_ppdb_headers a "+
									" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
									" INNER JOIN tbl_user_ppdb c on a.nik=c.nik and a.tgldaftar=c.tgldaftar and a.tahun_daftar=c.tahun_daftar "+
									" LEFT JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran "+
									" where a.flag_aktif=0 and b.flag_aktif=0 and flag_import<>'9' and a.kd_trans_masuk_ppdb=? "+
									" ORDER BY seqno ", c.Param("idhead")).Scan(&getDataPPDB)

									SetArrayData := []GetBiayaAndSisa{}
									var kd_trans_masuk_ppdb int
									//var total_bayar float64
									var tgldaftar string
									var tahun_daftar string
									var tahun_akademik string
									var ket string
									rowss, _ := db.Raw("SELECT distinct b.kd_trans_masuk_ppdb,CONVERT(a.tgldaftar,CHAR) 'tgldaftar',a.tahun_daftar,a.tahun_akademik,a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan "+
									" FROM tbl_trans_uang_masuk_ppdb_headers a "+
									" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
									" where a.flag_aktif=0 and b.flag_aktif=0  "+
									" and a.kd_trans_masuk_ppdb=?  ", c.Param("idhead")).Rows()
									defer rowss.Close()
									for rowss.Next() {
										rowss.Scan(&kd_trans_masuk_ppdb, &tgldaftar, &tahun_daftar, &tahun_akademik, &total_biaya, &total_bayar, &sisa_biaya, &ket)
										arraydata := GetBiayaAndSisa{}
										arraydata.Kd_trans_masuk_ppdb = kd_trans_masuk_ppdb
										tTglDaftar, _ := time.Parse("2006-01-02", tgldaftar)
										dateStrTglDaftar := tTglDaftar.Format("02-01-2006")
										arraydata.Tgldaftar = dateStrTglDaftar
										arraydata.Tahun_daftar = tahun_daftar
										arraydata.Tahun_akademik = tahun_akademik
										arraydata.Total_biaya = total_biaya
										arraydata.Total_bayar = total_bayar
										arraydata.Sisa_biaya = sisa_biaya
										arraydata.Keterangan = ket
										arraydata.Detail = getDataPPDB
										SetArrayData = append(SetArrayData, arraydata)
									}

									response := helper.APIResponse("Update Data Sukses ...", http.StatusOK, "success", SetArrayData)
									c.JSON(http.StatusOK, response)

								}

								func DeleteAllUangMasuk(c *gin.Context) {
									db := c.MustGet("db").(*gorm.DB)

									idhead := c.Param("idhead")

									var dataUtama table_data.Tbl_trans_uang_masuk_ppdb_headers
									if err := db.Where("flag_aktif=0 and kd_trans_masuk_ppdb=?", idhead).First(&dataUtama).Error; err != nil {
										errorMessage := gin.H{"errors": "Data Header Tidak Ditemukan ..."}
										response := helper.APIResponse("Update Data Gagal ...", http.StatusUnprocessableEntity, "error", errorMessage)
										c.JSON(http.StatusUnprocessableEntity, response)
										return
									}

									var dataUtamaDet table_data.Tbl_trans_uang_masuk_ppdb_details
									if err := db.Where("flag_aktif=0 and kd_trans_masuk_ppdb=?", idhead).First(&dataUtamaDet).Error; err != nil {
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

									var dataDetail table_data.Tbl_trans_uang_masuk_ppdb_details
									err = db.Raw("update tbl_trans_uang_masuk_ppdb_details set flag_aktif=9,edited_by=?,edited_on=? "+
										" where kd_trans_masuk_ppdb=? and flag_aktif=0 ",
										currentUser.(string), datenowx, idhead).Scan(&dataDetail).Error
										if err != nil {
											response := helper.APIResponse("Delete Data Ke Tbl_trans_uang_masuk_ppdb_details Gagal ...", http.StatusBadRequest, "error", err)
											c.JSON(http.StatusBadRequest, response)
											return
										}

										var dataHead table_data.Tbl_trans_uang_masuk_ppdb_headers
										err = db.Raw("update tbl_trans_uang_masuk_ppdb_headers set flag_aktif=9,edited_by=?,edited_on=? "+
											" where kd_trans_masuk_ppdb=? and flag_aktif=0 ",
											currentUser.(string), datenowx, idhead).Scan(&dataHead).Error
											if err != nil {
												response := helper.APIResponse("Delete Data Ke Tbl_trans_uang_masuk_ppdb_headers Gagal ...", http.StatusBadRequest, "error", err)
												c.JSON(http.StatusBadRequest, response)
												return
											}

											//setting tampilan habis update ppdb
											var getDataPPDB []GetDataPPDB
											db.Raw(" SELECT DISTINCT b.kd_trans_masuk_detail_ppdb,b.seqno,b.kategori_biaya_ppdb, "+
												" b.tgl_bayar,b.jml_bayar,d.kd_pembayaran,d.tipe_pembayaran "+
												" FROM tbl_trans_uang_masuk_ppdb_headers a "+
												" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
												" INNER JOIN tbl_user_ppdb c on a.nik=c.nik and a.tgldaftar=c.tgldaftar and a.tahun_daftar=c.tahun_daftar "+
												" LEFT JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran "+
												" where a.flag_aktif=0 and b.flag_aktif=0 and flag_import<>'9' and a.kd_trans_masuk_ppdb=? "+
												" ORDER BY seqno ", idhead).Scan(&getDataPPDB)

												SetArrayData := []GetBiayaAndSisa{}
												var kd_trans_masuk_ppdb int
												var total_bayar float64
												var tgldaftar string
												var tahun_daftar string
												var tahun_akademik string
												var total_biaya float64
												var sisa_biaya float64
												var ket string
												rowss, _ := db.Raw("SELECT distinct b.kd_trans_masuk_ppdb,CONVERT(a.tgldaftar,CHAR) 'tgldaftar',a.tahun_daftar,a.tahun_akademik,a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan "+
												" FROM tbl_trans_uang_masuk_ppdb_headers a "+
												" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
												" where a.flag_aktif=0 and b.flag_aktif=0  "+
												" and a.kd_trans_masuk_ppdb=?  ", idhead).Rows()
												defer rowss.Close()
												for rowss.Next() {
													rowss.Scan(&kd_trans_masuk_ppdb, &tgldaftar, &tahun_daftar, &tahun_akademik, &total_biaya, &total_bayar, &sisa_biaya, &ket)
													arraydata := GetBiayaAndSisa{}
													arraydata.Kd_trans_masuk_ppdb = kd_trans_masuk_ppdb
													tTglDaftar, _ := time.Parse("2006-01-02", tgldaftar)
													dateStrTglDaftar := tTglDaftar.Format("02-01-2006")
													arraydata.Tgldaftar = dateStrTglDaftar
													arraydata.Tahun_daftar = tahun_daftar
													arraydata.Tahun_akademik = tahun_akademik
													arraydata.Total_biaya = total_biaya
													arraydata.Total_bayar = total_bayar
													arraydata.Sisa_biaya = sisa_biaya
													arraydata.Keterangan = ket
													arraydata.Detail = getDataPPDB
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

											func CreateUangMasukDetail(c *gin.Context) {
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
												db.Raw(" SELECT count(*) jmldata FROM tbl_trans_uang_masuk_ppdb_headers a "+
												" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
												" where a.kd_trans_masuk_ppdb=? and a.flag_aktif=0 and b.flag_aktif=0 ", paramAddDetail.Kd_trans_masuk_ppdb).Scan(&intJmldata)
												if intJmldata == 0 {
													errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
													response := helper.APIResponse("Data Transaksi PPDB Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
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
												db.Raw("SELECT ifnull(max(kd_trans_masuk_detail_ppdb),0) + 1 as 'run_number' FROM tbl_trans_uang_masuk_ppdb_details ").Scan(&intKd_trans_masuk_detail)

												var int_seqno int
												db.Raw("SELECT (seqno + 1) as 'run_number' "+
												" FROM tbl_trans_uang_masuk_ppdb_details where flag_aktif=0 and kd_trans_masuk_ppdb=?", paramAddDetail.Kd_trans_masuk_ppdb).Scan(&int_seqno)

												datadetail := table_data.Tbl_trans_uang_masuk_ppdb_details{
													Kd_trans_masuk_ppdb:        paramAddDetail.Kd_trans_masuk_ppdb,
													Kd_trans_masuk_detail_ppdb: intKd_trans_masuk_detail,
													Seqno:                      int_seqno,
													Jml_bayar:                  0,
													Kategori_biaya_ppdb:        "",
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

												//setting tampilan habis save siswa
												var getDataPPDB []GetDataPPDB
												db.Raw(" SELECT DISTINCT b.kd_trans_masuk_detail_ppdb,b.seqno,b.kategori_biaya_ppdb, "+
													" b.tgl_bayar,b.jml_bayar,d.kd_pembayaran,d.tipe_pembayaran "+
													" FROM tbl_trans_uang_masuk_ppdb_headers a "+
													" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
													" INNER JOIN tbl_user_ppdb c on a.nik=c.nik and a.tgldaftar=c.tgldaftar and a.tahun_daftar=c.tahun_daftar "+
													" LEFT JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran "+
													" where a.flag_aktif=0 and b.flag_aktif=0 and flag_import<>'9' and a.kd_trans_masuk_ppdb=? "+
													" ORDER BY seqno ", paramAddDetail.Kd_trans_masuk_ppdb).Scan(&getDataPPDB)

													SetArrayData := []GetBiayaAndSisa{}
													var kd_trans_masuk_ppdb int
													var total_bayar float64
													var tgldaftar string
													var tahun_daftar string
													var tahun_akademik string
													var total_biaya float64
													var sisa_biaya float64
													var ket string
													rowss, _ := db.Raw("SELECT distinct b.kd_trans_masuk_ppdb,CONVERT(a.tgldaftar,CHAR) 'tgldaftar',a.tahun_daftar,a.tahun_akademik,a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan "+
													" FROM tbl_trans_uang_masuk_ppdb_headers a "+
													" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
													" where a.flag_aktif=0 and b.flag_aktif=0  "+
													" and a.kd_trans_masuk_ppdb=?  ", paramAddDetail.Kd_trans_masuk_ppdb).Rows()
													defer rowss.Close()
													for rowss.Next() {
														rowss.Scan(&kd_trans_masuk_ppdb, &tgldaftar, &tahun_daftar, &tahun_akademik, &total_biaya, &total_bayar, &sisa_biaya, &ket)
														arraydata := GetBiayaAndSisa{}
														arraydata.Kd_trans_masuk_ppdb = kd_trans_masuk_ppdb
														tTglDaftar, _ := time.Parse("2006-01-02", tgldaftar)
														dateStrTglDaftar := tTglDaftar.Format("02-01-2006")
														arraydata.Tgldaftar = dateStrTglDaftar
														arraydata.Tahun_daftar = tahun_daftar
														arraydata.Tahun_akademik = tahun_akademik
														arraydata.Total_biaya = total_biaya
														arraydata.Total_bayar = total_bayar
														arraydata.Sisa_biaya = sisa_biaya
														arraydata.Keterangan = ket
														arraydata.Detail = getDataPPDB
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

												func DeleteUangMasukDetail(c *gin.Context) {
													db := c.MustGet("db").(*gorm.DB)

													var paramDeleteDetail ParamDeleteDetail
													if err := c.ShouldBindJSON(&paramDeleteDetail); err != nil {
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

													kd_trans_masuk_ppdb := paramDeleteDetail.Kd_trans_masuk_ppdb
													kd_trans_masuk_detail_ppdb := paramDeleteDetail.Kd_trans_masuk_detail_ppdb

													var dataUtama table_data.Tbl_trans_uang_masuk_ppdb_details
													if err := db.Where("flag_aktif=0 and kd_trans_masuk_detail_ppdb=? and kd_trans_masuk_ppdb=?", kd_trans_masuk_detail_ppdb, kd_trans_masuk_ppdb).First(&dataUtama).Error; err != nil {
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

													var dataDetail table_data.Tbl_trans_uang_masuk_ppdb_details
													err = db.Raw("update tbl_trans_uang_masuk_ppdb_details set flag_aktif=9,edited_by=?,edited_on=? "+
														" where kd_trans_masuk_detail_ppdb=? and kd_trans_masuk_ppdb=? and flag_aktif=0 ",
														currentUser.(string), datenowx, kd_trans_masuk_detail_ppdb, kd_trans_masuk_ppdb).Scan(&dataDetail).Error
														if err != nil {
															response := helper.APIResponse("Delete Data Ke Tbl_trans_uang_masuk_ppdb_details Gagal ...", http.StatusBadRequest, "error", err)
															c.JSON(http.StatusBadRequest, response)
															return
														}

														var sumJmlBayar float64
														db.Raw("SELECT sum(jml_bayar) 'jml_bayar' FROM tbl_trans_uang_masuk_ppdb_details "+
														" where flag_aktif=0 and tgl_bayar is not null and kd_trans_masuk_ppdb=?", kd_trans_masuk_ppdb).Scan(&sumJmlBayar)

														var total_biaya float64
														db.Raw("SELECT total_biaya FROM tbl_trans_uang_masuk_ppdb_headers where flag_aktif=0 and kd_trans_masuk_ppdb=?", kd_trans_masuk_ppdb).Scan(&total_biaya)
														var sisa_biaya float64 = total_biaya - sumJmlBayar

														var dataHeader table_data.Tbl_trans_uang_masuk_ppdb_headers
														err = db.Raw("UPDATE tbl_trans_uang_masuk_ppdb_headers SET total_bayar = ?, sisa_biaya = ?, "+
															" edited_on = ? , edited_by = ? "+
															" WHERE kd_trans_masuk_ppdb = ? and flag_aktif=0 ", sumJmlBayar, sisa_biaya, datenowx, currentUser.(string), kd_trans_masuk_ppdb).Scan(&dataHeader).Error
															if err != nil {
																response := helper.APIResponse("Update Data Ke tbl_trans_uang_masuk_ppdb_headers Gagal ...", http.StatusBadRequest, "error", err)
																c.JSON(http.StatusBadRequest, response)
																return
															}

															//setting tampilan habis save siswa
															var getDataPPDB []GetDataPPDB
															db.Raw(" SELECT DISTINCT b.kd_trans_masuk_detail_ppdb,b.seqno,b.kategori_biaya_ppdb, "+
																" b.tgl_bayar,b.jml_bayar,d.kd_pembayaran,d.tipe_pembayaran "+
																" FROM tbl_trans_uang_masuk_ppdb_headers a "+
																" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
																" INNER JOIN tbl_user_ppdb c on a.nik=c.nik and a.tgldaftar=c.tgldaftar and a.tahun_daftar=c.tahun_daftar "+
																" LEFT JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran "+
																" where a.flag_aktif=0 and b.flag_aktif=0 and flag_import<>'9' and a.kd_trans_masuk_ppdb=? "+
																" ORDER BY seqno ", paramDeleteDetail.Kd_trans_masuk_ppdb).Scan(&getDataPPDB)

																SetArrayData := []GetBiayaAndSisa{}
																var total_bayar float64
																var tgldaftar string
																var tahun_daftar string
																var tahun_akademik string
																var ket string
																rowss, _ := db.Raw("SELECT distinct b.kd_trans_masuk_ppdb,CONVERT(a.tgldaftar,CHAR) 'tgldaftar',a.tahun_daftar,a.tahun_akademik,a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan "+
																" FROM tbl_trans_uang_masuk_ppdb_headers a "+
																" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
																" where a.flag_aktif=0 and b.flag_aktif=0  "+
																" and a.kd_trans_masuk_ppdb=?  ", paramDeleteDetail.Kd_trans_masuk_ppdb).Rows()
																defer rowss.Close()
																for rowss.Next() {
																	rowss.Scan(&kd_trans_masuk_ppdb, &tgldaftar, &tahun_daftar, &tahun_akademik, &total_biaya, &total_bayar, &sisa_biaya, &ket)
																	arraydata := GetBiayaAndSisa{}
																	arraydata.Kd_trans_masuk_ppdb = kd_trans_masuk_ppdb
																	tTglDaftar, _ := time.Parse("2006-01-02", tgldaftar)
																	dateStrTglDaftar := tTglDaftar.Format("02-01-2006")
																	arraydata.Tgldaftar = dateStrTglDaftar
																	arraydata.Tahun_daftar = tahun_daftar
																	arraydata.Tahun_akademik = tahun_akademik
																	arraydata.Total_biaya = total_biaya
																	arraydata.Total_bayar = total_bayar
																	arraydata.Sisa_biaya = sisa_biaya
																	arraydata.Keterangan = ket
																	arraydata.Detail = getDataPPDB
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

															func EditUangMasuk(c *gin.Context) {
																db := c.MustGet("db").(*gorm.DB)

																kd_trans_masuk_ppdb := c.Param("idhead")

																var paramEditPPDB ParamEditPPDB
																if err := c.ShouldBindJSON(&paramEditPPDB); err != nil {
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

																//cek data pembayaran mengada-ngada
																db.Raw(" SELECT count(*) jmldata FROM tbl_trans_uang_masuk_ppdb_headers a "+
																" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
																" where a.kd_trans_masuk_ppdb=? and a.flag_aktif=0 and b.flag_aktif=0 ", kd_trans_masuk_ppdb).Scan(&intJmldata)
																if intJmldata == 0 {
																	errorMessage := gin.H{"errors": "Simpan Data Gagal ..."}
																	response := helper.APIResponse("Data Pembayaran PPDB Tidak DiTemukan ...", http.StatusUnprocessableEntity, "error", errorMessage)
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
																db.Raw("SELECT sum(jml_bayar) 'jml_bayar' FROM tbl_trans_uang_masuk_ppdb_details "+
																" where flag_aktif=0 and tgl_bayar is not null and kd_trans_masuk_ppdb=?", kd_trans_masuk_ppdb).Scan(&sumJmlBayar)

																var sisa_biaya float64 = paramEditPPDB.Total_biaya - sumJmlBayar

																var dataHeader table_data.Tbl_trans_uang_masuk_ppdb_headers

																err = db.Raw("UPDATE tbl_trans_uang_masuk_ppdb_headers SET total_biaya=? ,total_bayar = ?, sisa_biaya = ?, "+
																	" edited_on = ? , edited_by = ? , keterangan=?   "+
																	" WHERE kd_trans_masuk_ppdb = ? "+
																	" and flag_aktif=0 ", paramEditPPDB.Total_biaya, sumJmlBayar, sisa_biaya, datenowx, currentUser.(string),
																	paramEditPPDB.Keterangan, kd_trans_masuk_ppdb).Scan(&dataHeader).Error
																	if err != nil {
																		response := helper.APIResponse("Update Data Ke tbl_trans_uang_masuk_ppdb_headers Gagal ...", http.StatusBadRequest, "error", err)
																		c.JSON(http.StatusBadRequest, response)
																		return
																	}

																	//setting tampilan habis save siswa

																	var getDataPPDB []GetDataPPDB
																	db.Raw(" SELECT DISTINCT b.kd_trans_masuk_detail_ppdb,b.seqno,b.kategori_biaya_ppdb, "+
																		" b.tgl_bayar,b.jml_bayar,d.kd_pembayaran,d.tipe_pembayaran "+
																		" FROM tbl_trans_uang_masuk_ppdb_headers a "+
																		" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
																		" INNER JOIN tbl_user_ppdb c on a.nik=c.nik and a.tgldaftar=c.tgldaftar and a.tahun_daftar=c.tahun_daftar "+
																		" LEFT JOIN tbl_tipe_pembayarans d on b.kd_pembayaran=d.kd_pembayaran "+
																		" where a.flag_aktif=0 and b.flag_aktif=0 and flag_import<>'9' and a.kd_trans_masuk_ppdb=? "+
																		" ORDER BY seqno ", kd_trans_masuk_ppdb).Scan(&getDataPPDB)

																		SetArrayData := []GetBiayaAndSisa{}
																		var total_bayar float64
																		var tgldaftar string
																		var tahun_daftar string
																		var tahun_akademik string
																		var ket string
																		var total_biaya float64

																		int_kd_trans_masuk_ppdb, _ := strconv.Atoi(kd_trans_masuk_ppdb)

																		rowss, _ := db.Raw("SELECT distinct b.kd_trans_masuk_ppdb,CONVERT(a.tgldaftar,CHAR) 'tgldaftar',a.tahun_daftar,a.tahun_akademik, "+
																		" a.total_biaya,a.total_bayar,a.sisa_biaya,a.keterangan "+
																		" FROM tbl_trans_uang_masuk_ppdb_headers a "+
																		" INNER JOIN tbl_trans_uang_masuk_ppdb_details b on a.kd_trans_masuk_ppdb=b.kd_trans_masuk_ppdb "+
																		" where a.flag_aktif=0 and b.flag_aktif=0  "+
																		" and a.kd_trans_masuk_ppdb=?  ", kd_trans_masuk_ppdb).Rows()
																		defer rowss.Close()
																		for rowss.Next() {
																			rowss.Scan(&kd_trans_masuk_ppdb, &tgldaftar, &tahun_daftar, &tahun_akademik, &total_biaya, &total_bayar, &sisa_biaya, &ket)
																			arraydata := GetBiayaAndSisa{}
																			arraydata.Kd_trans_masuk_ppdb = int_kd_trans_masuk_ppdb
																			tTglDaftar, _ := time.Parse("2006-01-02", tgldaftar)
																			dateStrTglDaftar := tTglDaftar.Format("02-01-2006")
																			arraydata.Tgldaftar = dateStrTglDaftar
																			arraydata.Tahun_daftar = tahun_daftar
																			arraydata.Tahun_akademik = tahun_akademik
																			arraydata.Total_biaya = total_biaya
																			arraydata.Total_bayar = total_bayar
																			arraydata.Sisa_biaya = sisa_biaya
																			arraydata.Keterangan = ket
																			arraydata.Detail = getDataPPDB
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
