package report_uangkeluar

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

func ReportPRA(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramReport ParamReportPRA
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
	var kd_trans_keluar int

	SetArrayData := []GetDataHeaderPRA{}
	ssql := " SELECT nm_group,nm_kategori,no_document,tgl_document, " +
		" total_biaya,total_bayar,sisa_biaya,keterangan,kd_trans_keluar " +
		" FROM vw_report_pra where nm_group<>'' "
	if paramReport.No_document != "" {
		ssql = fmt.Sprintf("%s and no_document = '%s'", ssql, paramReport.No_document)
	}
	if paramReport.Tgl_document1 != "" {
		ssql = fmt.Sprintf("%s and tgl_document >= '%s'", ssql, TglDocument1)
	}
	if paramReport.Tgl_document2 != "" {
		ssql = fmt.Sprintf("%s and tgl_document <= '%s'", ssql, TglDocument2)
	}

	ssql = fmt.Sprintf("%s group by kd_trans_keluar ", ssql)

	rows, _ := db.Raw(ssql).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&nm_group, &nm_kategori, &no_document, &tgl_document, &total_biaya, &total_bayar, &sisa_biaya, &keterangan, &kd_trans_keluar)
		arraydata := GetDataHeaderPRA{}
		arraydata.Nm_group = nm_group
		arraydata.Nm_kategori = nm_kategori
		arraydata.No_document = no_document
		arraydata.Tgl_document = tgl_document
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = keterangan

		ssqldetail := " SELECT jml_bayar,keterangan_detail FROM vw_report_pra "
		ssqldetail = fmt.Sprintf("%s where kd_trans_keluar = %d", ssqldetail, kd_trans_keluar)
		if paramReport.No_document != "" {
			ssqldetail = fmt.Sprintf("%s and no_document = '%s'", ssqldetail, paramReport.No_document)
		}
		if paramReport.Tgl_document1 != "" {
			ssqldetail = fmt.Sprintf("%s and tgl_document >= '%s'", ssqldetail, TglDocument1)
		}
		if paramReport.Tgl_document2 != "" {
			ssqldetail = fmt.Sprintf("%s and tgl_document <= '%s'", ssqldetail, TglDocument2)
		}

		var getDataDetail []GetDataDetailPRA
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

func ReportPRAACT(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramReport ParamReportPRAACT
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
	var no_document string
	var tgl_document string
	var total_biaya float64
	var total_bayar float64
	var sisa_biaya float64
	var keterangan string
	var kd_trans_keluar int

	SetArrayData := []GetDataHeaderPRAACT{}
	ssql := " SELECT nm_group,nm_kategori,no_document,tgl_document, " +
		" total_biaya,total_bayar,sisa_biaya,keterangan,kd_trans_keluar " +
		" FROM vw_report_pra_act where nm_group<>'' "
	if paramReport.No_document != "" {
		ssql = fmt.Sprintf("%s and no_document = '%s'", ssql, paramReport.No_document)
	}
	if paramReport.Tgl_document1 != "" {
		ssql = fmt.Sprintf("%s and tgl_document >= '%s'", ssql, TglDocument1)
	}
	if paramReport.Tgl_document2 != "" {
		ssql = fmt.Sprintf("%s and tgl_document <= '%s'", ssql, TglDocument2)
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

	ssql = fmt.Sprintf("%s group by kd_trans_keluar ", ssql)

	rows, _ := db.Raw(ssql).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&nm_group, &nm_kategori, &no_document, &tgl_document, &total_biaya, &total_bayar, &sisa_biaya, &keterangan, &kd_trans_keluar)
		arraydata := GetDataHeaderPRAACT{}
		arraydata.Nm_group = nm_group
		arraydata.Nm_kategori = nm_kategori
		arraydata.No_document = no_document
		arraydata.Tgl_document = tgl_document
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = keterangan

		ssqldetail := " SELECT pos_uang_masuk,date_format(tgl_bayar,'%d-%m-%Y') 'tgl_bayar',jml_bayar,keterangan_detail,tipe_pembayaran FROM vw_report_pra_act "
		ssqldetail = fmt.Sprintf("%s where kd_trans_keluar = %d", ssqldetail, kd_trans_keluar)
		if paramReport.No_document != "" {
			ssqldetail = fmt.Sprintf("%s and no_document = '%s'", ssqldetail, paramReport.No_document)
		}
		if paramReport.Tgl_document1 != "" {
			ssqldetail = fmt.Sprintf("%s and tgl_document >= '%s'", ssqldetail, TglDocument1)
		}
		if paramReport.Tgl_document2 != "" {
			ssqldetail = fmt.Sprintf("%s and tgl_document <= '%s'", ssqldetail, TglDocument2)
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

		var getDataDetail []GetDataDetailPRAACT
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

func ReportACT(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramReport ParamReportPRAACT
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
	var no_document string
	var tgl_document string
	var total_biaya float64
	var total_bayar float64
	var sisa_biaya float64
	var keterangan string
	var kd_trans_keluar int

	SetArrayData := []GetDataHeaderPRAACT{}
	ssql := " SELECT nm_group,nm_kategori,no_document,tgl_document, " +
		" total_biaya,total_bayar,sisa_biaya,keterangan,kd_trans_keluar " +
		" FROM vw_report_act where nm_group<>'' "
	if paramReport.No_document != "" {
		ssql = fmt.Sprintf("%s and no_document = '%s'", ssql, paramReport.No_document)
	}
	if paramReport.Tgl_document1 != "" {
		ssql = fmt.Sprintf("%s and tgl_document >= '%s'", ssql, TglDocument1)
	}
	if paramReport.Tgl_document2 != "" {
		ssql = fmt.Sprintf("%s and tgl_document <= '%s'", ssql, TglDocument2)
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

	ssql = fmt.Sprintf("%s group by kd_trans_keluar ORDER BY tgl_document ", ssql)

	rows, _ := db.Raw(ssql).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&nm_group, &nm_kategori, &no_document, &tgl_document, &total_biaya, &total_bayar, &sisa_biaya, &keterangan, &kd_trans_keluar)
		arraydata := GetDataHeaderPRAACT{}
		arraydata.Nm_group = nm_group
		arraydata.Nm_kategori = nm_kategori
		arraydata.No_document = no_document
		arraydata.Tgl_document = tgl_document
		arraydata.Total_biaya = total_biaya
		arraydata.Total_bayar = total_bayar
		arraydata.Sisa_biaya = sisa_biaya
		arraydata.Keterangan = keterangan

		ssqldetail := " SELECT pos_uang_masuk,date_format(tgl_bayar,'%d-%m-%Y') 'tgl_bayar',jml_bayar,keterangan_detail,tipe_pembayaran FROM vw_report_act "
		ssqldetail = fmt.Sprintf("%s where kd_trans_keluar = %d", ssqldetail, kd_trans_keluar)
		if paramReport.No_document != "" {
			ssqldetail = fmt.Sprintf("%s and no_document = '%s'", ssqldetail, paramReport.No_document)
		}
		if paramReport.Tgl_document1 != "" {
			ssqldetail = fmt.Sprintf("%s and tgl_document >= '%s'", ssqldetail, TglDocument1)
		}
		if paramReport.Tgl_document2 != "" {
			ssqldetail = fmt.Sprintf("%s and tgl_document <= '%s'", ssqldetail, TglDocument2)
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

		var getDataDetail []GetDataDetailPRAACT
		db.Raw(ssqldetail).Scan(&getDataDetail)

		arraydata.Detail = getDataDetail

		SetArrayData = append(SetArrayData, arraydata)
	}

	if len(SetArrayData) == 0 {
		response := helper.APIResponse("List Data ..."+ssql, http.StatusOK, "success", SetArrayData)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("List Data ..."+ssql, http.StatusOK, "success", SetArrayData)
	c.JSON(http.StatusOK, response)
}
