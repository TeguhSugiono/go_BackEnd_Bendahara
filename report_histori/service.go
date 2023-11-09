package report_histori

import (
	"net/http"
	"rest_api_bendahara/helper"
	"rest_api_bendahara/connection"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListNikNis(c *gin.Context) {
	//db := c.MustGet("db").(*gorm.DB)
	dbSIA := connection.SetupConnectionSIA()

	var listdata []ListData
	sql := "  SELECT nis,nik,nm_siswa from tbl_siswa " +
		" where status_siswa NOT IN ('Tidak Aktif') and flag_siswa=0 "

	dbSIA.Raw(sql).Scan(&listdata)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", listdata)
	c.JSON(http.StatusOK, response)
}

func ReportHistori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramSearch ParamSearch
	if err := c.ShouldBindJSON(&paramSearch); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	SetArrayData := []DetailHistori{}
	arraydata := DetailHistori{}

	arraydata.Nis = paramSearch.Nis
	arraydata.Nik = paramSearch.Nik
	arraydata.Nm_siswa = paramSearch.Nm_siswa

	// AMBIL DATA PPDB
	var detailPPDB []DetailPPDB
	db.Raw(" SELECT  tgl_bayar,jml_bayar,kategori_biaya_ppdb as 'keterangan_detail',kd_pembayaran,tipe_pembayaran,kd_trans_masuk_ppdb,kd_trans_masuk_detail_ppdb "+
		" FROM vw_report_ppdb where nik=? ", paramSearch.Nik).Scan(&detailPPDB)

	SetArrayHeaderPPDB := []HeaderPPDB{}
	var nm_group string
	var nm_kategori string
	var tgldaftar string
	var tahun_daftar string
	var tahun_akademik string
	var total_biaya float64
	var total_bayar float64
	var sisa_biaya float64
	var keterangan string
	var kd_trans_masuk_siswa int
	get_data_ppdb, _ := db.Raw(" SELECT nm_group,nm_kategori,tgldaftar,tahun_daftar,tahun_akademik,total_biaya,total_bayar,sisa_biaya,keterangan "+
		" FROM vw_report_ppdb where nik=? group by kd_trans_masuk_ppdb ", paramSearch.Nik).Rows()
	defer get_data_ppdb.Close()
	for get_data_ppdb.Next() {
		get_data_ppdb.Scan(&nm_group, &nm_kategori, &tgldaftar, &tahun_daftar, &tahun_akademik, &total_biaya, &total_bayar, &sisa_biaya, &keterangan)
		arrayPPDB := HeaderPPDB{}
		arrayPPDB.Nm_group = nm_group
		arrayPPDB.Nm_kategori = nm_kategori
		arrayPPDB.Tgldaftar = tgldaftar
		arrayPPDB.Tahun_daftar = tahun_daftar
		arrayPPDB.Tahun_akademik = tahun_akademik
		arrayPPDB.Total_biaya = total_biaya
		arrayPPDB.Total_bayar = total_bayar
		arrayPPDB.Sisa_biaya = sisa_biaya
		arrayPPDB.Keterangan = keterangan
		arrayPPDB.Detail = detailPPDB
		SetArrayHeaderPPDB = append(SetArrayHeaderPPDB, arrayPPDB)
	}
	arraydata.DataPPDB = SetArrayHeaderPPDB
	// END AMBIL DATA PPDB

	// AMBIL DATA SPP
	SetArrayGroupTahunKelas := []GroupTahunKelas{}
	var nm_kelas string
	get_data_spp, _ := db.Raw(" SELECT distinct tahun_akademik,nm_kelas FROM vw_report_spp where nis_siswa=?  ORDER BY nm_kelas desc ", paramSearch.Nis).Rows()
	defer get_data_spp.Close()
	for get_data_spp.Next() {
		get_data_spp.Scan(&tahun_akademik, &nm_kelas)
		arraySPP := GroupTahunKelas{}
		arraySPP.Tahun_akademik = tahun_akademik
		arraySPP.Nm_kelas = nm_kelas

		SetArrayHeaderSPP := []HeaderSPP{}
		get_data_spp_header, _ := db.Raw(" SELECT nm_group,nm_kategori, "+
			" total_biaya,total_bayar,sisa_biaya,keterangan "+
			" FROM vw_report_spp where nis_siswa=? and tahun_akademik=?  "+
			" and nm_kelas=? GROUP BY kd_trans_masuk  ", paramSearch.Nis, tahun_akademik, nm_kelas).Rows()
		defer get_data_spp_header.Close()
		for get_data_spp_header.Next() {
			get_data_spp_header.Scan(&nm_group, &nm_kategori, &total_biaya, &total_bayar, &sisa_biaya, &keterangan)
			arraySPPHeader := HeaderSPP{}
			arraySPPHeader.Nm_group = nm_group
			arraySPPHeader.Nm_kategori = nm_kategori
			arraySPPHeader.Total_biaya = total_biaya
			arraySPPHeader.Total_bayar = total_bayar
			arraySPPHeader.Sisa_biaya = sisa_biaya
			arraySPPHeader.Keterangan = keterangan

			var detailSPP []DetailSPP
			db.Raw(" SELECT  periode_bayar,tgl_bayar,jml_tagihan,jml_bayar,keterangan_detail,kd_pembayaran,tipe_pembayaran,kd_trans_masuk,kd_trans_masuk_detail "+
				" FROM vw_report_spp where nis_siswa=? and tahun_akademik=? and nm_kelas=? ", paramSearch.Nis, tahun_akademik, nm_kelas).Scan(&detailSPP)
			arraySPPHeader.DetailData = detailSPP

			SetArrayHeaderSPP = append(SetArrayHeaderSPP, arraySPPHeader)
		}

		arraySPP.HeaderData = SetArrayHeaderSPP
		SetArrayGroupTahunKelas = append(SetArrayGroupTahunKelas, arraySPP)
	}
	arraydata.DataSPP = SetArrayGroupTahunKelas
	//END AMBIL DATA SPP

	// AMBIL DATA UM SISWA
	SetArrayGroupTahunKelasSiswa := []GroupTahunKelas{}
	get_data_siswa, _ := db.Raw(" SELECT distinct tahun_akademik,nm_kelas FROM vw_report_umsiswa where nis_siswa=?  ORDER BY nm_kelas desc ", paramSearch.Nis).Rows()
	defer get_data_siswa.Close()
	for get_data_siswa.Next() {
		get_data_siswa.Scan(&tahun_akademik, &nm_kelas)
		arraySPP := GroupTahunKelas{}
		arraySPP.Tahun_akademik = tahun_akademik
		arraySPP.Nm_kelas = nm_kelas

		SetArrayHeaderSPP := []HeaderSPP{}
		get_data_siswa_header, _ := db.Raw(" SELECT nm_group,nm_kategori, "+
			" total_biaya,total_bayar,sisa_biaya,keterangan,kd_trans_masuk_siswa "+
			" FROM vw_report_umsiswa where nis_siswa=? and tahun_akademik=?  "+
			" and nm_kelas=? GROUP BY kd_trans_masuk_siswa  ", paramSearch.Nis, tahun_akademik, nm_kelas).Rows()
		defer get_data_siswa_header.Close()
		for get_data_siswa_header.Next() {
			get_data_siswa_header.Scan(&nm_group, &nm_kategori, &total_biaya, &total_bayar, &sisa_biaya, &keterangan, &kd_trans_masuk_siswa)
			arraySPPHeader := HeaderSPP{}
			arraySPPHeader.Nm_group = nm_group
			arraySPPHeader.Nm_kategori = nm_kategori
			arraySPPHeader.Total_biaya = total_biaya
			arraySPPHeader.Total_bayar = total_bayar
			arraySPPHeader.Sisa_biaya = sisa_biaya
			arraySPPHeader.Keterangan = keterangan

			var detailUmSiswa []DetailUmSiswa
			db.Raw(" SELECT  tgl_bayar,jml_bayar,keterangan_detail,kd_pembayaran,tipe_pembayaran,kd_trans_masuk_siswa,kd_trans_masuk_detail_siswa "+
				" FROM vw_report_umsiswa where nis_siswa=? and tahun_akademik=? and nm_kelas=? "+
				" and kd_trans_masuk_siswa=? ", paramSearch.Nis, tahun_akademik, nm_kelas, kd_trans_masuk_siswa).Scan(&detailUmSiswa)
			arraySPPHeader.DetailData = detailUmSiswa

			SetArrayHeaderSPP = append(SetArrayHeaderSPP, arraySPPHeader)
		}

		arraySPP.HeaderData = SetArrayHeaderSPP
		SetArrayGroupTahunKelasSiswa = append(SetArrayGroupTahunKelasSiswa, arraySPP)
	}
	arraydata.DataUmSiswa = SetArrayGroupTahunKelasSiswa
	//END AMBIL DATA UM SISWA

	SetArrayData = append(SetArrayData, arraydata)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
	c.JSON(http.StatusOK, response)

}


