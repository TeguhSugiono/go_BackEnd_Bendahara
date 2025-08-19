package report_uangmasuk

import (
	"errors"
	"fmt"
	"net/http"
	"rest_api_bendahara/helper"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ReportSPP(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramReport ParamReportSPP
	if err := c.ShouldBindJSON(&paramReport); err != nil {
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

	var PeriodeBayar1 string
	var PeriodeBayar2 string

	if paramReport.Periode_bayar1 != "" {
		TPeriodeBayar1, err2 := time.Parse("01-2006", paramReport.Periode_bayar1)
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
		PeriodeBayar1 = TPeriodeBayar1.Format("2006-01")
	}

	if paramReport.Periode_bayar2 != "" {
		TPeriodeBayar2, err2 := time.Parse("01-2006", paramReport.Periode_bayar2)
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
		PeriodeBayar2 = TPeriodeBayar2.Format("2006-01")
	}

	var TglBayar1 string
	var TglBayar2 string

	if paramReport.Tgl_bayar1 != "" {
		tTgl_bayar1, err2 := time.Parse("02-01-2006", paramReport.Tgl_bayar1)
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
		TglBayar1 = tTgl_bayar1.Format("2006-01-02")
	}

	if paramReport.Tgl_bayar1 != "" {
		tTglBayar2, err2 := time.Parse("02-01-2006", paramReport.Tgl_bayar2)
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
		TglBayar2 = tTglBayar2.Format("2006-01-02")
	}

	var nm_group string
	var nm_kategori string
	var tahun_akademik string
	var nis_siswa string
	var nm_siswa string
	var total_biaya float64
	var total_bayar float64
	var sisa_biaya float64
	var keterangan string
	var kd_trans_masuk int
	var nm_kelas string

	SetArrayData := []GetDataHeaderSPP{}
	ssql := " SELECT nm_group,nm_kategori,tahun_akademik,nis_siswa,nm_siswa,nm_kelas, " +
		" total_biaya,total_bayar,sisa_biaya,keterangan,kd_trans_masuk " +
		" FROM vw_report_spp where nm_group<>'' "
	if paramReport.Tahun_akademik != "" {
		ssql = fmt.Sprintf("%s and tahun_akademik = '%s'", ssql, paramReport.Tahun_akademik)
	}
	if paramReport.Nis_siswa != "" {
		ssql = fmt.Sprintf("%s and nis_siswa = '%s'", ssql, paramReport.Nis_siswa)
	}
	if paramReport.Periode_bayar1 != "" {
		ssql = fmt.Sprintf("%s and date_periode_bayar >= '%s'", ssql, PeriodeBayar1)
	}
	if paramReport.Periode_bayar2 != "" {
		ssql = fmt.Sprintf("%s and date_periode_bayar <= '%s'", ssql, PeriodeBayar2)
	}
	if paramReport.Tgl_bayar1 != "" {
		ssql = fmt.Sprintf("%s and tgl_bayar >= '%s'", ssql, TglBayar1)
	}
	if paramReport.Tgl_bayar2 != "" {
		ssql = fmt.Sprintf("%s and tgl_bayar <= '%s'", ssql, TglBayar2)
	}
	if paramReport.Kd_pembayaran != "" {
		ssql = fmt.Sprintf("%s and kd_pembayaran = '%s'", ssql, paramReport.Kd_pembayaran)
	}

	ssql = fmt.Sprintf("%s group by kd_trans_masuk ", ssql)

	ssqldetail := ""

	rows, _ := db.Raw(ssql).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&nm_group, &nm_kategori, &tahun_akademik, &nis_siswa, &nm_siswa, &nm_kelas, &total_biaya, &total_bayar, &sisa_biaya, &keterangan, &kd_trans_masuk)
		arraydata := GetDataHeaderSPP{}
		arraydata.Nm_group = nm_group
		arraydata.Nm_kategori = nm_kategori
		arraydata.Tahun_akademik = tahun_akademik
		arraydata.Nis_siswa = nis_siswa
		arraydata.Nm_siswa = nm_siswa
		arraydata.Nm_kelas = nm_kelas
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = keterangan

		ssqldetail = " SELECT periode_bayar,date_format(tgl_bayar,'%d-%m-%Y') 'tgl_bayar',jml_tagihan,jml_bayar,keterangan_detail,tipe_pembayaran FROM vw_report_spp "
		ssqldetail = fmt.Sprintf("%s where kd_trans_masuk = %d", ssqldetail, kd_trans_masuk)
		if paramReport.Tahun_akademik != "" {
			ssqldetail = fmt.Sprintf("%s and tahun_akademik = '%s'", ssqldetail, paramReport.Tahun_akademik)
		}
		if paramReport.Nis_siswa != "" {
			ssqldetail = fmt.Sprintf("%s and nis_siswa = '%s'", ssqldetail, paramReport.Nis_siswa)
		}
		if paramReport.Periode_bayar1 != "" {
			ssqldetail = fmt.Sprintf("%s and date_periode_bayar >= '%s'", ssqldetail, PeriodeBayar1)
		}
		if paramReport.Periode_bayar2 != "" {
			ssqldetail = fmt.Sprintf("%s and date_periode_bayar <= '%s'", ssqldetail, PeriodeBayar2)
		}
		if paramReport.Tgl_bayar1 != "" {
			ssqldetail = fmt.Sprintf("%s and tgl_bayar >= '%s'", ssqldetail, TglBayar1)
		}
		if paramReport.Tgl_bayar2 != "" {
			ssqldetail = fmt.Sprintf("%s and tgl_bayar <= '%s'", ssqldetail, TglBayar2)
		}
		if paramReport.Kd_pembayaran != "" {
			ssqldetail = fmt.Sprintf("%s and kd_pembayaran = '%s'", ssqldetail, paramReport.Kd_pembayaran)
		}

		var getDataDetail []GetDataDetailSPP
		db.Raw(ssqldetail).Scan(&getDataDetail)

		arraydata.Detail = getDataDetail

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

func ReportPPDB(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramReport ParamReportPPDB
	if err := c.ShouldBindJSON(&paramReport); err != nil {
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

	var TglBayar1 string
	var TglBayar2 string

	if paramReport.Tgl_bayar1 != "" {
		tTgl_bayar1, err2 := time.Parse("02-01-2006", paramReport.Tgl_bayar1)
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
		TglBayar1 = tTgl_bayar1.Format("2006-01-02")
	}

	if paramReport.Tgl_bayar1 != "" {
		tTglBayar2, err2 := time.Parse("02-01-2006", paramReport.Tgl_bayar2)
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
		TglBayar2 = tTglBayar2.Format("2006-01-02")
	}

	var TglDaftar1 string
	var TglDaftar2 string

	if paramReport.Tgldaftar1 != "" {
		tTglDaftar1, err2 := time.Parse("02-01-2006", paramReport.Tgldaftar1)
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
		TglDaftar1 = tTglDaftar1.Format("2006-01-02")
	}

	if paramReport.Tgldaftar2 != "" {
		tTgldaftar2, err2 := time.Parse("02-01-2006", paramReport.Tgldaftar2)
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
		TglDaftar2 = tTgldaftar2.Format("2006-01-02")
	}

	var nm_group string
	var nm_kategori string
	var nik string
	var nm_siswa string
	var tgldaftar string
	var tahun_daftar string
	var tahun_akademik string
	var total_biaya float64
	var total_bayar float64
	var sisa_biaya float64
	var keterangan string
	var kd_trans_masuk_ppdb int

	SetArrayData := []GetDataHeaderPPDB{}
	ssql := " SELECT nm_group,nm_kategori,nik,nm_siswa,tgldaftar,tahun_daftar,tahun_akademik, " +
		" total_biaya,total_bayar,sisa_biaya,keterangan,kd_trans_masuk_ppdb " +
		" FROM vw_report_ppdb where nm_group<>'' "
	if paramReport.Tahun_akademik != "" {
		ssql = fmt.Sprintf("%s and tahun_akademik = '%s'", ssql, paramReport.Tahun_akademik)
	}
	if paramReport.Nik != "" {
		ssql = fmt.Sprintf("%s and nik = '%s'", ssql, paramReport.Nik)
	}
	if paramReport.Tgldaftar1 != "" {
		ssql = fmt.Sprintf("%s and tgldaftar >= '%s'", ssql, TglDaftar1)
	}
	if paramReport.Tgldaftar2 != "" {
		ssql = fmt.Sprintf("%s and tgldaftar <= '%s'", ssql, TglDaftar2)
	}
	if paramReport.Tgl_bayar1 != "" {
		ssql = fmt.Sprintf("%s and tgl_bayar >= '%s'", ssql, TglBayar1)
	}
	if paramReport.Tgl_bayar2 != "" {
		ssql = fmt.Sprintf("%s and tgl_bayar <= '%s'", ssql, TglBayar2)
	}
	if paramReport.Kd_pembayaran != "" {
		ssql = fmt.Sprintf("%s and kd_pembayaran = '%s'", ssql, paramReport.Kd_pembayaran)
	}

	ssql = fmt.Sprintf("%s group by kd_trans_masuk_ppdb ", ssql)

	rows, _ := db.Raw(ssql).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&nm_group, &nm_kategori, &nik, &nm_siswa, &tgldaftar, &tahun_daftar, &tahun_akademik, &total_biaya, &total_bayar, &sisa_biaya, &keterangan, &kd_trans_masuk_ppdb)
		arraydata := GetDataHeaderPPDB{}
		arraydata.Nm_group = nm_group
		arraydata.Nm_kategori = nm_kategori
		arraydata.Nik = nik
		arraydata.Nm_siswa = nm_siswa
		arraydata.Tgldaftar = tgldaftar
		arraydata.Tahun_daftar = tahun_daftar
		arraydata.Tahun_akademik = tahun_akademik
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = keterangan

		ssqldetail := " SELECT date_format(tgl_bayar,'%d-%m-%Y') 'tgl_bayar',jml_bayar,keterangan_detail,tipe_pembayaran FROM vw_report_ppdb "
		ssqldetail = fmt.Sprintf("%s where kd_trans_masuk_ppdb = %d", ssqldetail, kd_trans_masuk_ppdb)
		if paramReport.Tahun_akademik != "" {
			ssqldetail = fmt.Sprintf("%s and tahun_akademik = '%s'", ssqldetail, paramReport.Tahun_akademik)
		}
		if paramReport.Nik != "" {
			ssqldetail = fmt.Sprintf("%s and nik = '%s'", ssqldetail, paramReport.Nik)
		}
		if paramReport.Tgldaftar1 != "" {
			ssqldetail = fmt.Sprintf("%s and tgldaftar >= '%s'", ssqldetail, TglDaftar1)
		}
		if paramReport.Tgldaftar2 != "" {
			ssqldetail = fmt.Sprintf("%s and tgldaftar <= '%s'", ssqldetail, TglDaftar2)
		}
		if paramReport.Tgl_bayar1 != "" {
			ssqldetail = fmt.Sprintf("%s and tgl_bayar >= '%s'", ssqldetail, TglBayar1)
		}
		if paramReport.Tgl_bayar2 != "" {
			ssqldetail = fmt.Sprintf("%s and tgl_bayar <= '%s'", ssqldetail, TglBayar2)
		}
		if paramReport.Kd_pembayaran != "" {
			ssql = fmt.Sprintf("%s and kd_pembayaran = '%s'", ssql, paramReport.Kd_pembayaran)
		}

		var getDataDetail []GetDataDetailPPDB
		db.Raw(ssqldetail).Scan(&getDataDetail)

		arraydata.Detail = getDataDetail

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

func ReportUmSiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramReport ParamReportUmSiswa
	if err := c.ShouldBindJSON(&paramReport); err != nil {
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

	var TglBayar1 string
	var TglBayar2 string

	if paramReport.Tgl_bayar1 != "" {
		tTgl_bayar1, err2 := time.Parse("02-01-2006", paramReport.Tgl_bayar1)
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
		TglBayar1 = tTgl_bayar1.Format("2006-01-02")
	}

	if paramReport.Tgl_bayar1 != "" {
		tTglBayar2, err2 := time.Parse("02-01-2006", paramReport.Tgl_bayar2)
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
		TglBayar2 = tTglBayar2.Format("2006-01-02")
	}

	var nm_group string
	var nm_kategori string
	var tahun_akademik string
	var nis_siswa string
	var nm_siswa string
	var nm_kelas string
	var total_biaya float64
	var total_bayar float64
	var sisa_biaya float64
	var keterangan string
	var kd_trans_masuk_siswa int

	SetArrayData := []GetDataHeaderUmSiswa{}
	ssql := " SELECT nm_group,nm_kategori,tahun_akademik,nis_siswa,nm_siswa,nm_kelas, " +
		" total_biaya,total_bayar,sisa_biaya,keterangan,kd_trans_masuk_siswa " +
		" FROM vw_report_umsiswa where nm_group<>'' "
	if paramReport.Tahun_akademik != "" {
		ssql = fmt.Sprintf("%s and tahun_akademik = '%s'", ssql, paramReport.Tahun_akademik)
	}
	if paramReport.Nis_siswa != "" {
		ssql = fmt.Sprintf("%s and nis_siswa = '%s'", ssql, paramReport.Nis_siswa)
	}
	if paramReport.Nm_kelas != "" {
		ssql = fmt.Sprintf("%s and nm_kelas = '%s'", ssql, paramReport.Nm_kelas)
	}
	if paramReport.Tgl_bayar1 != "" {
		ssql = fmt.Sprintf("%s and tgl_bayar >= '%s'", ssql, TglBayar1)
	}
	if paramReport.Tgl_bayar2 != "" {
		ssql = fmt.Sprintf("%s and tgl_bayar <= '%s'", ssql, TglBayar2)
	}
	if paramReport.Kd_pembayaran != "" {
		ssql = fmt.Sprintf("%s and kd_pembayaran = '%s'", ssql, paramReport.Kd_pembayaran)
	}

	ssql = fmt.Sprintf("%s group by kd_trans_masuk_siswa ", ssql)

	rows, _ := db.Raw(ssql).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&nm_group, &nm_kategori, &tahun_akademik, &nis_siswa, &nm_siswa, &nm_kelas, &total_biaya, &total_bayar, &sisa_biaya, &keterangan, &kd_trans_masuk_siswa)
		arraydata := GetDataHeaderUmSiswa{}
		arraydata.Nm_group = nm_group
		arraydata.Nm_kategori = nm_kategori
		arraydata.Tahun_akademik = tahun_akademik
		arraydata.Nis_siswa = nis_siswa
		arraydata.Nm_siswa = nm_siswa
		arraydata.Nm_kelas = nm_kelas
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = keterangan

		ssqldetail := " SELECT date_format(tgl_bayar,'%d-%m-%Y') 'tgl_bayar',jml_bayar,keterangan_detail,tipe_pembayaran FROM vw_report_umsiswa "
		ssqldetail = fmt.Sprintf("%s where kd_trans_masuk_siswa = %d", ssqldetail, kd_trans_masuk_siswa)
		if paramReport.Tahun_akademik != "" {
			ssqldetail = fmt.Sprintf("%s and tahun_akademik = '%s'", ssqldetail, paramReport.Tahun_akademik)
		}
		if paramReport.Nis_siswa != "" {
			ssqldetail = fmt.Sprintf("%s and nis_siswa = '%s'", ssqldetail, paramReport.Nis_siswa)
		}
		if paramReport.Nm_kelas != "" {
			ssqldetail = fmt.Sprintf("%s and nm_kelas = '%s'", ssqldetail, paramReport.Nm_kelas)
		}
		if paramReport.Tgl_bayar1 != "" {
			ssqldetail = fmt.Sprintf("%s and tgl_bayar >= '%s'", ssqldetail, TglBayar1)
		}
		if paramReport.Tgl_bayar2 != "" {
			ssqldetail = fmt.Sprintf("%s and tgl_bayar <= '%s'", ssqldetail, TglBayar2)
		}
		if paramReport.Kd_pembayaran != "" {
			ssqldetail = fmt.Sprintf("%s and kd_pembayaran = '%s'", ssqldetail, paramReport.Kd_pembayaran)
		}

		var getDataDetail []GetDataDetailUmSiswa
		db.Raw(ssqldetail).Scan(&getDataDetail)

		arraydata.Detail = getDataDetail

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

func ReportUmLain(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramReport ParamReportUmLain
	if err := c.ShouldBindJSON(&paramReport); err != nil {
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

	var TglBayar1 string
	var TglBayar2 string

	if paramReport.Tgl_bayar1 != "" {
		tTgl_bayar1, err2 := time.Parse("02-01-2006", paramReport.Tgl_bayar1)
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
		TglBayar1 = tTgl_bayar1.Format("2006-01-02")
	}

	if paramReport.Tgl_bayar1 != "" {
		tTglBayar2, err2 := time.Parse("02-01-2006", paramReport.Tgl_bayar2)
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
		TglBayar2 = tTglBayar2.Format("2006-01-02")
	}

	var TglDocument1 string
	var TglDocument2 string

	if paramReport.Tgl_document1 != "" {
		tTglDocument1, err2 := time.Parse("02-01-2006", paramReport.Tgl_document1)
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
		TglDocument1 = tTglDocument1.Format("2006-01-02")
	}

	if paramReport.Tgl_document2 != "" {
		tTglDocument2, err2 := time.Parse("02-01-2006", paramReport.Tgl_document2)
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
		TglDocument2 = tTglDocument2.Format("2006-01-02")
	}

	var nm_group string
	var nm_kategori string
	var no_document string
	var tgl_document string
	var total_biaya float64
	var total_bayar float64
	var sisa_biaya float64
	var keterangan string
	var kd_trans_masuk_lain int

	SetArrayData := []GetDataHeaderUmLain{}
	ssql := " SELECT nm_group,nm_kategori,no_document,tgl_document, " +
		" total_biaya,total_bayar,sisa_biaya,keterangan,kd_trans_masuk_lain " +
		" FROM vw_report_umlain where nm_group<>'' "
	if paramReport.No_document != "" {
		ssql = fmt.Sprintf("%s and no_document = '%s'", ssql, paramReport.No_document)
	}
	if paramReport.Tgl_bayar1 != "" {
		ssql = fmt.Sprintf("%s and tgl_bayar >= '%s'", ssql, TglBayar1)
	}
	if paramReport.Tgl_bayar2 != "" {
		ssql = fmt.Sprintf("%s and tgl_bayar <= '%s'", ssql, TglBayar2)
	}
	if paramReport.Tgl_document1 != "" {
		ssql = fmt.Sprintf("%s and tgl_bayar >= '%s'", ssql, TglDocument1)
	}
	if paramReport.Tgl_document2 != "" {
		ssql = fmt.Sprintf("%s and tgl_bayar <= '%s'", ssql, TglDocument2)
	}
	if paramReport.Kd_pembayaran != "" {
		ssql = fmt.Sprintf("%s and kd_pembayaran = '%s'", ssql, paramReport.Kd_pembayaran)
	}

	ssql = fmt.Sprintf("%s group by kd_trans_masuk_lain ", ssql)

	rows, _ := db.Raw(ssql).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&nm_group, &nm_kategori, &no_document, &tgl_document, &total_biaya, &total_bayar, &sisa_biaya, &keterangan, &kd_trans_masuk_lain)
		arraydata := GetDataHeaderUmLain{}
		arraydata.Nm_group = nm_group
		arraydata.Nm_kategori = nm_kategori
		arraydata.No_document = no_document
		arraydata.Tgl_document = tgl_document
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = keterangan

		ssqldetail := " SELECT date_format(tgl_bayar,'%d-%m-%Y') 'tgl_bayar',jml_bayar,keterangan_detail,tipe_pembayaran FROM vw_report_umlain "
		ssqldetail = fmt.Sprintf("%s where kd_trans_masuk_lain = %d", ssqldetail, kd_trans_masuk_lain)
		if paramReport.No_document != "" {
			ssqldetail = fmt.Sprintf("%s and no_document = '%s'", ssqldetail, paramReport.No_document)
		}
		if paramReport.Tgl_bayar1 != "" {
			ssqldetail = fmt.Sprintf("%s and tgl_bayar >= '%s'", ssqldetail, TglBayar1)
		}
		if paramReport.Tgl_bayar2 != "" {
			ssqldetail = fmt.Sprintf("%s and tgl_bayar <= '%s'", ssqldetail, TglBayar2)
		}
		if paramReport.Tgl_document1 != "" {
			ssqldetail = fmt.Sprintf("%s and tgl_bayar >= '%s'", ssqldetail, TglDocument1)
		}
		if paramReport.Tgl_document2 != "" {
			ssqldetail = fmt.Sprintf("%s and tgl_bayar <= '%s'", ssqldetail, TglDocument2)
		}
		if paramReport.Kd_pembayaran != "" {
			ssqldetail = fmt.Sprintf("%s and kd_pembayaran = '%s'", ssqldetail, paramReport.Kd_pembayaran)
		}

		var getDataDetail []GetDataDetailUmLain
		db.Raw(ssqldetail).Scan(&getDataDetail)

		arraydata.Detail = getDataDetail

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
