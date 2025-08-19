package report_group

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

func Report_Group_Masuk(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramData ParamData
	if err := c.ShouldBindJSON(&paramData); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var TglBayar1 string
	var TglBayar2 string

	if paramData.Tgl_bayar1 != "" {
		tTgl_bayar1, err2 := time.Parse("02-01-2006", paramData.Tgl_bayar1)
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
	} else {
		TglBayar1 = ""
	}

	if paramData.Tgl_bayar1 != "" {
		tTglBayar2, err2 := time.Parse("02-01-2006", paramData.Tgl_bayar2)
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
	} else {
		TglBayar2 = ""
	}

	// Query untuk mendapatkan data kategori
	//=============================================
	var NewLoopOne []NewLoop1
	sqlOne := " SELECT a.kd_kategori,a.nm_kategori,a.kd_group FROM tbl_kategori_uangs as a " +
		" left join vw_s_masuk as b on a.nm_kategori=b.nm_kategori  " +
		" and b.tgl_bayar>='" + TglBayar1 + "' and b.tgl_bayar<='" + TglBayar2 + "' " +
		" where a.flag_aktif=0 GROUP BY a.kd_kategori  order by a.kd_kategori  "
	db.Raw(sqlOne).Scan(&NewLoopOne)

	var NewLoopTwo []NewLoop2
	sqlTwo := " SELECT ifnull(sum(jml_bayar),0.00) 'total_bayar',kd_kategori FROM vw_s_masuk  where total_bayar <> 0  " +
		" and tgl_bayar>='" + TglBayar1 + "' and tgl_bayar<='" + TglBayar2 + "' " +
		" GROUP BY kd_kategori "
	db.Raw(sqlTwo).Scan(&NewLoopTwo)

	var NewLoopThree []NewLoop3
	sqlThree := " SELECT tgl_bayar,jml_bayar,tipe_pembayaran,kd_kategori,CAST(nik as CHAR(50)) 'Data_ID',nm_siswa 'Data_Name'  FROM vw_s_masuk  where total_bayar <> 0  " +
		" and tgl_bayar>='" + TglBayar1 + "' and tgl_bayar<='" + TglBayar2 + "' "
	db.Raw(sqlThree).Scan(&NewLoopThree)

	SetJenisUang := []JenisGroupUang{}
	arrayJenisUang := JenisGroupUang{}
	arrayJenisUang.JenisUang = "Uang Masuk"

	SetArrayData := []GroupUangMasuk{}
	var kd_group int
	var nm_group string
	sql_group := " SELECT a.kd_group,a.nm_group FROM tbl_group_kategoris as a  " +
		" inner join tbl_kategori_uangs as c on a.kd_group = c.kd_group " +
		" Left join vw_s_masuk as b  on a.nm_group=b.nm_group and c.nm_kategori=b.nm_kategori "

	if paramData.Tgl_bayar1 != "" {
		sql_group = fmt.Sprintf("%s and b.tgl_bayar >= '%s'", sql_group, TglBayar1)
	}
	if paramData.Tgl_bayar2 != "" {
		sql_group = fmt.Sprintf("%s and b.tgl_bayar <= '%s'", sql_group, TglBayar2)
	}

	sql_group = fmt.Sprintf("%s where a.flag_aktif=0 and a.kd_jenis=1 and a.nm_header <> '' and c.flag_aktif=0  ", sql_group)

	sql_group = fmt.Sprintf("%s  GROUP BY a.kd_group order by a.kd_group ", sql_group)

	execute_sql_group, _ := db.Raw(sql_group).Rows()
	defer execute_sql_group.Close()
	for execute_sql_group.Next() {
		execute_sql_group.Scan(&kd_group, &nm_group)
		arraydata := GroupUangMasuk{}
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group

		SetArrayDetail := []GroupUangMasukDetail{}
		targetKdGroup := kd_group
		var resultOne NewLoop1
		for _, data := range NewLoopOne {
			if data.Kd_group == targetKdGroup {
				resultOne = data

				arraydetail := GroupUangMasukDetail{}
				arraydetail.Kd_kategori = resultOne.Kd_kategori
				arraydetail.Nm_kategori = resultOne.Nm_kategori

				targetKdKategori := resultOne.Kd_kategori
				var resultTwo NewLoop2
				for _, data := range NewLoopTwo {
					if data.Kd_kategori == targetKdKategori {
						resultTwo = data
						break
					}
				}

				arraydetail.Total_bayar = resultTwo.Total_bayar

				arrayDetailBayar := []DetailBayar{}
				targetKdKategoriX := resultOne.Kd_kategori
				var resultThree NewLoop3
				for _, data := range NewLoopThree {
					if data.Kd_kategori == targetKdKategoriX {
						resultThree = data

						arrayDetailBayarTemp := DetailBayar{}
						arrayDetailBayarTemp.Tgl_bayar = resultThree.Tgl_bayar
						arrayDetailBayarTemp.Jml_bayar = resultThree.Jml_bayar
						arrayDetailBayarTemp.Tipe_pembayaran = resultThree.Tipe_pembayaran
						arrayDetailBayarTemp.Data_ID = resultThree.Data_ID
						arrayDetailBayarTemp.Data_Name = resultThree.Data_Name

						arrayDetailBayar = append(arrayDetailBayar, arrayDetailBayarTemp)

					}
				}
				arraydetail.DetailBayar = arrayDetailBayar

				SetArrayDetail = append(SetArrayDetail, arraydetail)

			}

		}

		arraydata.Kategori = SetArrayDetail
		SetArrayData = append(SetArrayData, arraydata)

	}

	arrayJenisUang.DetailData = SetArrayData
	SetJenisUang = append(SetJenisUang, arrayJenisUang)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetJenisUang)
	c.JSON(http.StatusOK, response)

}

func Report_Group_Keluar(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramData ParamData
	if err := c.ShouldBindJSON(&paramData); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var TglBayar1 string
	var TglBayar2 string

	if paramData.Tgl_bayar1 != "" {
		tTgl_bayar1, err2 := time.Parse("02-01-2006", paramData.Tgl_bayar1)
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
	} else {
		TglBayar1 = ""
	}

	if paramData.Tgl_bayar1 != "" {
		tTglBayar2, err2 := time.Parse("02-01-2006", paramData.Tgl_bayar2)
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
	} else {
		TglBayar2 = ""
	}

	var NewLoopOne []NewLoop1
	sqlOne := " SELECT a.kd_kategori,a.nm_kategori,a.kd_group FROM tbl_kategori_uangs as a " +
		" left join vw_s_keluar as b on a.nm_kategori=b.nm_kategori  " +
		" and b.tgl_bayar>='" + TglBayar1 + "' and b.tgl_bayar<='" + TglBayar2 + "' "
	if paramData.Kd_post_uang_masuk != "" {
		sqlOne = fmt.Sprintf("%s and b.kd_post_uang_masuk = '%s'", sqlOne, paramData.Kd_post_uang_masuk)
	}
	sqlOne = fmt.Sprintf("%s  where a.flag_aktif=0 ", sqlOne)
	if paramData.Kd_post_uang_masuk != "" {
		sqlOne = fmt.Sprintf("%s and b.kd_post_uang_masuk = '%s'", sqlOne, paramData.Kd_post_uang_masuk)
	}
	sqlOne = fmt.Sprintf("%s  GROUP BY a.kd_kategori  order by a.kd_kategori ", sqlOne)
	db.Raw(sqlOne).Scan(&NewLoopOne)

	var NewLoopTwo []NewLoop2
	sqlTwo := " SELECT ifnull(sum(jml_bayar),0.00) 'total_bayar',kd_kategori FROM vw_s_keluar  where total_bayar <> 0  " +
		" and tgl_bayar>='" + TglBayar1 + "' and tgl_bayar<='" + TglBayar2 + "' "
	if paramData.Kd_post_uang_masuk != "" {
		sqlTwo = fmt.Sprintf("%s and kd_post_uang_masuk = '%s'", sqlTwo, paramData.Kd_post_uang_masuk)
	}
	sqlTwo = fmt.Sprintf("%s  GROUP BY kd_kategori ", sqlTwo)
	db.Raw(sqlTwo).Scan(&NewLoopTwo)

	var NewLoopThreeA []NewLoop3a
	sqlThree := " SELECT tgl_bayar,jml_bayar,tipe_pembayaran,pos_uang_masuk,kd_kategori,no_document 'Data_no_document',tgl_document 'Data_tgl_document'  FROM vw_s_keluar  where total_bayar <> 0  " +
		" and tgl_bayar>='" + TglBayar1 + "' and tgl_bayar<='" + TglBayar2 + "' "
	if paramData.Kd_post_uang_masuk != "" {
		sqlThree = fmt.Sprintf("%s and kd_post_uang_masuk = '%s'", sqlThree, paramData.Kd_post_uang_masuk)
	}
	db.Raw(sqlThree).Scan(&NewLoopThreeA)

	SetJenisUang := []JenisGroupUang{}
	arrayJenisUang := JenisGroupUang{}
	arrayJenisUang.JenisUang = "Uang Keluar"

	SetArrayData := []GroupUangMasuk{}
	var kd_group int
	var nm_group string

	sql_group := " SELECT a.kd_group,a.nm_group FROM tbl_group_kategoris as a  " +
		" inner join tbl_kategori_uangs as c on a.kd_group = c.kd_group " +
		" Left join vw_s_keluar as b  on a.nm_group=b.nm_group and c.nm_kategori=b.nm_kategori "

	if paramData.Tgl_bayar1 != "" {
		sql_group = fmt.Sprintf("%s and b.tgl_bayar >= '%s'", sql_group, TglBayar1)
	}
	if paramData.Tgl_bayar2 != "" {
		sql_group = fmt.Sprintf("%s and b.tgl_bayar <= '%s'", sql_group, TglBayar2)
	}
	if paramData.Kd_post_uang_masuk != "" {
		sql_group = fmt.Sprintf("%s and b.kd_post_uang_masuk = '%s'", sql_group, paramData.Kd_post_uang_masuk)
	}
	sql_group = fmt.Sprintf("%s where a.flag_aktif=0 and a.kd_jenis=2 and a.nm_header <> '' and c.flag_aktif=0  ", sql_group)
	if paramData.Kd_post_uang_masuk != "" {
		sql_group = fmt.Sprintf("%s and b.kd_post_uang_masuk = '%s'", sql_group, paramData.Kd_post_uang_masuk)
	}
	sql_group = fmt.Sprintf("%s  GROUP BY a.kd_group order by a.kd_group ", sql_group)

	execute_sql_group, _ := db.Raw(sql_group).Rows()
	defer execute_sql_group.Close()
	for execute_sql_group.Next() {

		execute_sql_group.Scan(&kd_group, &nm_group)
		arraydata := GroupUangMasuk{}
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group

		SetArrayDetail := []GroupUangMasukDetail{}
		targetKdGroup := kd_group
		var resultOne NewLoop1
		for _, data := range NewLoopOne {
			if data.Kd_group == targetKdGroup {
				resultOne = data

				arraydetail := GroupUangMasukDetail{}
				arraydetail.Kd_kategori = resultOne.Kd_kategori
				arraydetail.Nm_kategori = resultOne.Nm_kategori

				targetKdKategori := resultOne.Kd_kategori
				var resultTwo NewLoop2
				for _, data := range NewLoopTwo {
					if data.Kd_kategori == targetKdKategori {
						resultTwo = data
						break
					}
				}

				arraydetail.Total_bayar = resultTwo.Total_bayar

				arrayDetailBayar := []DetailBayarOut{}
				targetKdKategoriX := resultOne.Kd_kategori
				var resultThreeA NewLoop3a
				for _, data := range NewLoopThreeA {
					if data.Kd_kategori == targetKdKategoriX {
						resultThreeA = data

						arrayDetailBayarTemp := DetailBayarOut{}
						arrayDetailBayarTemp.Tgl_bayar = resultThreeA.Tgl_bayar
						arrayDetailBayarTemp.Jml_bayar = resultThreeA.Jml_bayar
						arrayDetailBayarTemp.Tipe_pembayaran = resultThreeA.Tipe_pembayaran
						arrayDetailBayarTemp.Pos_uang_masuk = resultThreeA.Pos_uang_masuk
						arrayDetailBayarTemp.Data_no_document = resultThreeA.Data_no_document
						arrayDetailBayarTemp.Data_tgl_document = resultThreeA.Data_tgl_document

						arrayDetailBayar = append(arrayDetailBayar, arrayDetailBayarTemp)

					}
				}
				arraydetail.DetailBayar = arrayDetailBayar

				SetArrayDetail = append(SetArrayDetail, arraydetail)

			}

		}

		arraydata.Kategori = SetArrayDetail
		SetArrayData = append(SetArrayData, arraydata)

	}

	arrayJenisUang.DetailData = SetArrayData
	SetJenisUang = append(SetJenisUang, arrayJenisUang)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetJenisUang)
	c.JSON(http.StatusOK, response)
}

func Report_Group_Masuk_Keluar(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramData ParamData
	if err := c.ShouldBindJSON(&paramData); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var TglBayar1 string
	var TglBayar2 string

	if paramData.Tgl_bayar1 != "" {
		tTgl_bayar1, err2 := time.Parse("02-01-2006", paramData.Tgl_bayar1)
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
	} else {
		TglBayar1 = ""
	}

	if paramData.Tgl_bayar1 != "" {
		tTglBayar2, err2 := time.Parse("02-01-2006", paramData.Tgl_bayar2)
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
	} else {
		TglBayar2 = ""
	}

	SetGroupUang := []GroupUang{}
	arrayGroupUang := GroupUang{}

	//================= Uang Masuk ====================================//
	var NewLoopOne []NewLoop1
	sqlOne := " SELECT a.kd_kategori,a.nm_kategori,a.kd_group FROM tbl_kategori_uangs as a " +
		" left join vw_s_masuk as b on a.nm_kategori=b.nm_kategori  " +
		" and b.tgl_bayar>='" + TglBayar1 + "' and b.tgl_bayar<='" + TglBayar2 + "' " +
		" where a.flag_aktif=0 GROUP BY a.kd_kategori  order by a.kd_kategori  "
	db.Raw(sqlOne).Scan(&NewLoopOne)

	var NewLoopTwo []NewLoop2
	sqlTwo := " SELECT ifnull(sum(jml_bayar),0.00) 'total_bayar',kd_kategori FROM vw_s_masuk  where total_bayar <> 0  " +
		" and tgl_bayar>='" + TglBayar1 + "' and tgl_bayar<='" + TglBayar2 + "' " +
		" GROUP BY kd_kategori "
	db.Raw(sqlTwo).Scan(&NewLoopTwo)

	var NewLoopThree []NewLoop3
	sqlThree := " SELECT tgl_bayar,jml_bayar,tipe_pembayaran,kd_kategori,CAST(nik as CHAR(50)) 'Data_ID',nm_siswa 'Data_Name'  FROM vw_s_masuk  where total_bayar <> 0  " +
		" and tgl_bayar>='" + TglBayar1 + "' and tgl_bayar<='" + TglBayar2 + "' "
	db.Raw(sqlThree).Scan(&NewLoopThree)
	//================= End Uang Masuk ====================================//

	SetJenisGroupUang := []JenisGroupUang{}
	arrayJenisGroupUang := JenisGroupUang{}
	arrayJenisGroupUang.JenisUang = "Uang Masuk"

	//uang masuk
	SetArrayData := []GroupUangMasuk{}
	var kd_group int
	var nm_group string
	//sql_group := " SELECT kd_group,nm_group FROM tbl_group_kategoris where flag_aktif=0 and kd_jenis=1 and nm_header <> '' order by kd_group "

	sql_group := " SELECT a.kd_group,a.nm_group FROM tbl_group_kategoris as a  " +
		" inner join tbl_kategori_uangs as c on a.kd_group = c.kd_group " +
		" Left join vw_s_masuk as b  on a.nm_group=b.nm_group and c.nm_kategori=b.nm_kategori "

	//groupnya biar pertanggal
	if paramData.Tgl_bayar1 != "" {
		sql_group = fmt.Sprintf("%s and b.tgl_bayar >= '%s'", sql_group, TglBayar1)
	}
	if paramData.Tgl_bayar2 != "" {
		sql_group = fmt.Sprintf("%s and b.tgl_bayar <= '%s'", sql_group, TglBayar2)
	}
	//end groupnya biar pertanggal

	sql_group = fmt.Sprintf("%s where a.flag_aktif=0 and a.kd_jenis=1 and a.nm_header <> '' and c.flag_aktif=0 ", sql_group)

	sql_group = fmt.Sprintf("%s  GROUP BY a.kd_group order by a.kd_group ", sql_group)

	execute_sql_group, _ := db.Raw(sql_group).Rows()
	defer execute_sql_group.Close()
	for execute_sql_group.Next() {
		execute_sql_group.Scan(&kd_group, &nm_group)
		arraydata := GroupUangMasuk{}
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group

		SetArrayDetail := []GroupUangMasukDetail{}
		targetKdGroup := kd_group
		var resultOne NewLoop1
		for _, data := range NewLoopOne {
			if data.Kd_group == targetKdGroup {
				resultOne = data

				arraydetail := GroupUangMasukDetail{}
				arraydetail.Kd_kategori = resultOne.Kd_kategori
				arraydetail.Nm_kategori = resultOne.Nm_kategori

				targetKdKategori := resultOne.Kd_kategori
				var resultTwo NewLoop2
				for _, data := range NewLoopTwo {
					if data.Kd_kategori == targetKdKategori {
						resultTwo = data
						break
					}
				}

				arraydetail.Total_bayar = resultTwo.Total_bayar

				arrayDetailBayar := []DetailBayar{}
				targetKdKategoriX := resultOne.Kd_kategori
				var resultThree NewLoop3
				for _, data := range NewLoopThree {
					if data.Kd_kategori == targetKdKategoriX {
						resultThree = data

						arrayDetailBayarTemp := DetailBayar{}
						arrayDetailBayarTemp.Tgl_bayar = resultThree.Tgl_bayar
						arrayDetailBayarTemp.Jml_bayar = resultThree.Jml_bayar
						arrayDetailBayarTemp.Tipe_pembayaran = resultThree.Tipe_pembayaran
						arrayDetailBayarTemp.Data_ID = resultThree.Data_ID
						arrayDetailBayarTemp.Data_Name = resultThree.Data_Name

						arrayDetailBayar = append(arrayDetailBayar, arrayDetailBayarTemp)

					}
				}
				arraydetail.DetailBayar = arrayDetailBayar

				SetArrayDetail = append(SetArrayDetail, arraydetail)

			}

		}

		arraydata.Kategori = SetArrayDetail
		SetArrayData = append(SetArrayData, arraydata)

	}
	arrayJenisGroupUang.DetailData = SetArrayData
	SetJenisGroupUang = append(SetJenisGroupUang, arrayJenisGroupUang)
	//end uang masuk

	//=========================================================================================================
	//======================= Uang Keluar ========================//
	//var NewLoopOne []NewLoop1
	sqlOne = " SELECT a.kd_kategori,a.nm_kategori,a.kd_group FROM tbl_kategori_uangs as a " +
		" left join vw_s_keluar as b on a.nm_kategori=b.nm_kategori  " +
		" and b.tgl_bayar>='" + TglBayar1 + "' and b.tgl_bayar<='" + TglBayar2 + "' " +
		" where a.flag_aktif=0 GROUP BY a.kd_kategori  order by a.kd_kategori  "
	db.Raw(sqlOne).Scan(&NewLoopOne)

	//var NewLoopTwo []NewLoop2
	sqlTwo = " SELECT ifnull(sum(jml_bayar),0.00) 'total_bayar',kd_kategori FROM vw_s_keluar  where total_bayar <> 0  " +
		" and tgl_bayar>='" + TglBayar1 + "' and tgl_bayar<='" + TglBayar2 + "' " +
		" GROUP BY kd_kategori "
	db.Raw(sqlTwo).Scan(&NewLoopTwo)

	var NewLoopThreeA []NewLoop3a
	sqlThree = " SELECT tgl_bayar,jml_bayar,tipe_pembayaran,pos_uang_masuk,kd_kategori,no_document 'Data_no_document',tgl_document 'Data_tgl_document'  FROM vw_s_keluar  where total_bayar <> 0  " +
		" and tgl_bayar>='" + TglBayar1 + "' and tgl_bayar<='" + TglBayar2 + "' "
	db.Raw(sqlThree).Scan(&NewLoopThreeA)
	//======================= End Uang Keluar ========================//

	arrayJenisGroupUangKeluar := JenisGroupUang{}
	arrayJenisGroupUangKeluar.JenisUang = "Uang Keluar"

	SetArrayData = []GroupUangMasuk{}
	sql_group = " SELECT a.kd_group,a.nm_group FROM tbl_group_kategoris as a  " +
		" inner join tbl_kategori_uangs as c on a.kd_group = c.kd_group " +
		" Left join vw_s_keluar as b  on a.nm_group=b.nm_group and c.nm_kategori=b.nm_kategori "

	//groupnya biar pertanggal
	if paramData.Tgl_bayar1 != "" {
		sql_group = fmt.Sprintf("%s and b.tgl_bayar >= '%s'", sql_group, TglBayar1)
	}
	if paramData.Tgl_bayar2 != "" {
		sql_group = fmt.Sprintf("%s and b.tgl_bayar <= '%s'", sql_group, TglBayar2)
	}
	//end groupnya biar pertanggal

	sql_group = fmt.Sprintf("%s where a.flag_aktif=0 and a.kd_jenis=2 and a.nm_header <> '' and c.flag_aktif=0 ", sql_group)

	sql_group = fmt.Sprintf("%s  GROUP BY a.kd_group order by a.kd_group ", sql_group)

	execute_sql_group, _ = db.Raw(sql_group).Rows()
	defer execute_sql_group.Close()
	for execute_sql_group.Next() {
		execute_sql_group.Scan(&kd_group, &nm_group)
		arraydata := GroupUangMasuk{}
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group

		SetArrayDetail := []GroupUangMasukDetail{}
		targetKdGroup := kd_group
		var resultOne NewLoop1
		for _, data := range NewLoopOne {
			if data.Kd_group == targetKdGroup {
				resultOne = data

				arraydetail := GroupUangMasukDetail{}
				arraydetail.Kd_kategori = resultOne.Kd_kategori
				arraydetail.Nm_kategori = resultOne.Nm_kategori

				targetKdKategori := resultOne.Kd_kategori
				var resultTwo NewLoop2
				for _, data := range NewLoopTwo {
					if data.Kd_kategori == targetKdKategori {
						resultTwo = data
						break
					}
				}

				arraydetail.Total_bayar = resultTwo.Total_bayar

				arrayDetailBayar := []DetailBayarOut{}
				targetKdKategoriX := resultOne.Kd_kategori
				var resultThreeA NewLoop3a
				for _, data := range NewLoopThreeA {
					if data.Kd_kategori == targetKdKategoriX {
						resultThreeA = data

						arrayDetailBayarTemp := DetailBayarOut{}
						arrayDetailBayarTemp.Tgl_bayar = resultThreeA.Tgl_bayar
						arrayDetailBayarTemp.Jml_bayar = resultThreeA.Jml_bayar
						arrayDetailBayarTemp.Tipe_pembayaran = resultThreeA.Tipe_pembayaran
						arrayDetailBayarTemp.Pos_uang_masuk = resultThreeA.Pos_uang_masuk
						arrayDetailBayarTemp.Data_no_document = resultThreeA.Data_no_document
						arrayDetailBayarTemp.Data_tgl_document = resultThreeA.Data_tgl_document

						arrayDetailBayar = append(arrayDetailBayar, arrayDetailBayarTemp)

					}
				}
				arraydetail.DetailBayar = arrayDetailBayar

				SetArrayDetail = append(SetArrayDetail, arraydetail)

			}

		}

		arraydata.Kategori = SetArrayDetail
		SetArrayData = append(SetArrayData, arraydata)

	}
	arrayJenisGroupUangKeluar.DetailData = SetArrayData
	//=========================================================================================================

	SetJenisGroupUang = append(SetJenisGroupUang, arrayJenisGroupUangKeluar)

	arrayGroupUang.DataUang = SetJenisGroupUang
	SetGroupUang = append(SetGroupUang, arrayGroupUang)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetGroupUang)
	c.JSON(http.StatusOK, response)
}

func Report_Group_Masuk_Keluar_Old(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramData ParamData
	if err := c.ShouldBindJSON(&paramData); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var TglBayar1 string
	var TglBayar2 string

	if paramData.Tgl_bayar1 != "" {
		tTgl_bayar1, err2 := time.Parse("02-01-2006", paramData.Tgl_bayar1)
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
	} else {
		TglBayar1 = ""
	}

	if paramData.Tgl_bayar1 != "" {
		tTglBayar2, err2 := time.Parse("02-01-2006", paramData.Tgl_bayar2)
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
	} else {
		TglBayar2 = ""
	}

	SetGroupUang := []GroupUang{}
	arrayGroupUang := GroupUang{}

	SetJenisGroupUang := []JenisGroupUang{}
	arrayJenisGroupUang := JenisGroupUang{}
	arrayJenisGroupUang.JenisUang = "Uang Masuk"

	//uang masuk
	SetArrayData := []GroupUangMasuk{}
	var kd_group int
	var nm_group string
	//sql_group := " SELECT kd_group,nm_group FROM tbl_group_kategoris where flag_aktif=0 and kd_jenis=1 and nm_header <> '' order by kd_group "

	sql_group := " SELECT a.kd_group,a.nm_group FROM tbl_group_kategoris as a  " +
		" inner join tbl_kategori_uangs as c on a.kd_group = c.kd_group " +
		" Left join vw_s_masuk as b  on a.nm_group=b.nm_group and c.nm_kategori=b.nm_kategori "

	//groupnya biar pertanggal
	if paramData.Tgl_bayar1 != "" {
		sql_group = fmt.Sprintf("%s and b.tgl_bayar >= '%s'", sql_group, TglBayar1)
	}
	if paramData.Tgl_bayar2 != "" {
		sql_group = fmt.Sprintf("%s and b.tgl_bayar <= '%s'", sql_group, TglBayar2)
	}
	//end groupnya biar pertanggal

	sql_group = fmt.Sprintf("%s where a.flag_aktif=0 and a.kd_jenis=1 and a.nm_header <> '' and c.flag_aktif=0 ", sql_group)

	sql_group = fmt.Sprintf("%s  GROUP BY a.kd_group order by a.kd_group ", sql_group)

	execute_sql_group, _ := db.Raw(sql_group).Rows()

	for execute_sql_group.Next() {
		execute_sql_group.Scan(&kd_group, &nm_group)
		arraydata := GroupUangMasuk{}
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group

		SetArrayDetail := []GroupUangMasukDetail{}
		var kd_kategori int
		var nm_kategori string
		// sql_kategori := " SELECT kd_kategori,nm_kategori FROM tbl_kategori_uangs where flag_aktif=0 "
		// sql_kategori = fmt.Sprintf("%s and kd_group = %d", sql_kategori, kd_group)
		// //groupnya biar pertanggal
		// if paramData.Tgl_bayar1 != "" {
		// 	sql_kategori = fmt.Sprintf("%s and tgl_bayar >= '%s'", sql_kategori, TglBayar1)
		// }
		// if paramData.Tgl_bayar2 != "" {
		// 	sql_kategori = fmt.Sprintf("%s and tgl_bayar <= '%s'", sql_kategori, TglBayar2)
		// }
		// //end groupnya biar pertanggal
		// sql_kategori = fmt.Sprintf("%s  order by kd_kategori ", sql_kategori)
		sql_kategori := " SELECT a.kd_kategori,a.nm_kategori FROM tbl_kategori_uangs as a   " +
			" left join vw_s_masuk as b on a.nm_kategori=b.nm_kategori  "

		//groupnya biar pertanggal
		if paramData.Tgl_bayar1 != "" {
			sql_kategori = fmt.Sprintf("%s and b.tgl_bayar >= '%s'", sql_kategori, TglBayar1)
		}
		if paramData.Tgl_bayar2 != "" {
			sql_kategori = fmt.Sprintf("%s and b.tgl_bayar <= '%s'", sql_kategori, TglBayar2)
		}
		//end groupnya biar pertanggal

		sql_kategori = fmt.Sprintf("%s where a.flag_aktif=0 and a.kd_group = %d", sql_kategori, kd_group)

		sql_kategori = fmt.Sprintf("%s GROUP BY a.kd_kategori  order by a.kd_kategori ", sql_kategori)

		execute_sql_kategori, _ := db.Raw(sql_kategori).Rows()
		for execute_sql_kategori.Next() {
			execute_sql_kategori.Scan(&kd_kategori, &nm_kategori)
			arraydetail := GroupUangMasukDetail{}
			arraydetail.Kd_kategori = kd_kategori
			arraydetail.Nm_kategori = nm_kategori

			//cari uang masuk
			var total_bayar float64
			sql_nominal := " SELECT ifnull(sum(jml_bayar),0.00) 'total_bayar' FROM vw_s_masuk  where total_bayar <> 0 "
			sql_nominal = fmt.Sprintf("%s and kd_kategori = %d", sql_nominal, kd_kategori)
			if paramData.Tgl_bayar1 != "" {
				sql_nominal = fmt.Sprintf("%s and tgl_bayar >= '%s'", sql_nominal, TglBayar1)
			}
			if paramData.Tgl_bayar2 != "" {
				sql_nominal = fmt.Sprintf("%s and tgl_bayar <= '%s'", sql_nominal, TglBayar2)
			}
			sql_nominal = fmt.Sprintf("%s  GROUP BY kd_kategori ", sql_nominal)
			execute_sql_nominal, _ := db.Raw(sql_nominal).Rows()
			//defer execute_sql_nominal.Close()
			for execute_sql_nominal.Next() {
				execute_sql_nominal.Scan(&total_bayar)
			}

			arraydetail.Total_bayar = total_bayar
			//end cari uang masuk

			//cari detail pembayaran
			var tgl_bayar string
			var jml_bayar float64
			var tipe_pembayaran string

			arrayDetailBayar := []DetailBayar{}
			sql_nominal_detail := " SELECT tgl_bayar,jml_bayar,tipe_pembayaran  FROM vw_s_masuk  where total_bayar <> 0 "
			sql_nominal_detail = fmt.Sprintf("%s and kd_kategori = %d", sql_nominal_detail, kd_kategori)
			if paramData.Tgl_bayar1 != "" {
				sql_nominal_detail = fmt.Sprintf("%s and tgl_bayar >= '%s'", sql_nominal_detail, TglBayar1)
			}
			if paramData.Tgl_bayar2 != "" {
				sql_nominal_detail = fmt.Sprintf("%s and tgl_bayar <= '%s'", sql_nominal_detail, TglBayar2)
			}
			execute_sql_nominal_detail, _ := db.Raw(sql_nominal_detail).Rows()
			//defer execute_sql_nominal_detail.Close()
			for execute_sql_nominal_detail.Next() {
				execute_sql_nominal_detail.Scan(&tgl_bayar, &jml_bayar, &tipe_pembayaran)

				arrayDetailBayarTemp := DetailBayar{}
				arrayDetailBayarTemp.Tgl_bayar = tgl_bayar
				arrayDetailBayarTemp.Jml_bayar = jml_bayar
				arrayDetailBayarTemp.Tipe_pembayaran = tipe_pembayaran

				arrayDetailBayar = append(arrayDetailBayar, arrayDetailBayarTemp)
			}

			arraydetail.DetailBayar = arrayDetailBayar
			//end cari detail pembayaran

			SetArrayDetail = append(SetArrayDetail, arraydetail)

			defer execute_sql_nominal.Close()
			defer execute_sql_nominal_detail.Close()
		}
		defer execute_sql_kategori.Close()

		arraydata.Kategori = SetArrayDetail
		SetArrayData = append(SetArrayData, arraydata)
	}
	defer execute_sql_group.Close()
	arrayJenisGroupUang.DetailData = SetArrayData
	SetJenisGroupUang = append(SetJenisGroupUang, arrayJenisGroupUang)
	//end uang masuk

	//=========================================================================================================
	arrayJenisGroupUangKeluar := JenisGroupUang{}
	arrayJenisGroupUangKeluar.JenisUang = "Uang Keluar"

	SetArrayData = []GroupUangMasuk{}
	//sql_group = " SELECT kd_group,nm_group FROM tbl_group_kategoris where flag_aktif=0 and kd_jenis=2 and nm_header <> '' order by kd_group "
	sql_group = " SELECT a.kd_group,a.nm_group FROM tbl_group_kategoris as a  " +
		" inner join tbl_kategori_uangs as c on a.kd_group = c.kd_group " +
		" Left join vw_s_keluar as b  on a.nm_group=b.nm_group and c.nm_kategori=b.nm_kategori "

	//groupnya biar pertanggal
	if paramData.Tgl_bayar1 != "" {
		sql_group = fmt.Sprintf("%s and b.tgl_bayar >= '%s'", sql_group, TglBayar1)
	}
	if paramData.Tgl_bayar2 != "" {
		sql_group = fmt.Sprintf("%s and b.tgl_bayar <= '%s'", sql_group, TglBayar2)
	}
	//end groupnya biar pertanggal

	sql_group = fmt.Sprintf("%s where a.flag_aktif=0 and a.kd_jenis=2 and a.nm_header <> '' and c.flag_aktif=0 ", sql_group)

	sql_group = fmt.Sprintf("%s  GROUP BY a.kd_group order by a.kd_group ", sql_group)

	execute_sql_group, _ = db.Raw(sql_group).Rows()
	for execute_sql_group.Next() {
		execute_sql_group.Scan(&kd_group, &nm_group)
		arraydata := GroupUangMasuk{}
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group

		SetArrayDetail := []GroupUangMasukDetail{}
		var kd_kategori int
		var nm_kategori string
		// sql_kategori := " SELECT kd_kategori,nm_kategori FROM tbl_kategori_uangs where flag_aktif=0 "
		// sql_kategori = fmt.Sprintf("%s and kd_group = %d", sql_kategori, kd_group)
		// //groupnya biar pertanggal
		// if paramData.Tgl_bayar1 != "" {
		// 	sql_kategori = fmt.Sprintf("%s and tgl_bayar >= '%s'", sql_kategori, TglBayar1)
		// }
		// if paramData.Tgl_bayar2 != "" {
		// 	sql_kategori = fmt.Sprintf("%s and tgl_bayar <= '%s'", sql_kategori, TglBayar2)
		// }
		// //end groupnya biar pertanggal
		// sql_kategori = fmt.Sprintf("%s  order by kd_kategori ", sql_kategori)

		sql_kategori := " SELECT a.kd_kategori,a.nm_kategori FROM tbl_kategori_uangs as a   " +
			" left join vw_s_keluar as b on a.nm_kategori=b.nm_kategori  "

		//groupnya biar pertanggal
		if paramData.Tgl_bayar1 != "" {
			sql_kategori = fmt.Sprintf("%s and b.tgl_bayar >= '%s'", sql_kategori, TglBayar1)
		}
		if paramData.Tgl_bayar2 != "" {
			sql_kategori = fmt.Sprintf("%s and b.tgl_bayar <= '%s'", sql_kategori, TglBayar2)
		}
		//end groupnya biar pertanggal

		sql_kategori = fmt.Sprintf("%s where a.flag_aktif=0 and a.kd_group = %d", sql_kategori, kd_group)

		sql_kategori = fmt.Sprintf("%s GROUP BY a.kd_kategori  order by a.kd_kategori ", sql_kategori)

		execute_sql_kategori, _ := db.Raw(sql_kategori).Rows()
		for execute_sql_kategori.Next() {
			execute_sql_kategori.Scan(&kd_kategori, &nm_kategori)
			arraydetail := GroupUangMasukDetail{}
			arraydetail.Kd_kategori = kd_kategori
			arraydetail.Nm_kategori = nm_kategori

			//cari uang masuk
			var total_bayar float64
			sql_nominal := " SELECT ifnull(sum(jml_bayar),0.00) 'total_bayar' FROM vw_s_keluar  where total_bayar <> 0 "
			sql_nominal = fmt.Sprintf("%s and kd_kategori = %d", sql_nominal, kd_kategori)
			if paramData.Tgl_bayar1 != "" {
				sql_nominal = fmt.Sprintf("%s and tgl_bayar >= '%s'", sql_nominal, TglBayar1)
			}
			if paramData.Tgl_bayar2 != "" {
				sql_nominal = fmt.Sprintf("%s and tgl_bayar <= '%s'", sql_nominal, TglBayar2)
			}
			sql_nominal = fmt.Sprintf("%s  GROUP BY kd_kategori ", sql_nominal)
			execute_sql_nominal, _ := db.Raw(sql_nominal).Rows()
			//defer execute_sql_nominal.Close()
			for execute_sql_nominal.Next() {
				execute_sql_nominal.Scan(&total_bayar)
			}

			arraydetail.Total_bayar = total_bayar
			//end cari uang masuk

			//cari detail pembayaran
			var tgl_bayar string
			var jml_bayar float64
			var tipe_pembayaran string
			var pos_uang_masuk string

			arrayDetailBayar := []DetailBayarOut{}
			sql_nominal_detail := " SELECT tgl_bayar,jml_bayar,tipe_pembayaran,pos_uang_masuk  FROM vw_s_keluar  where total_bayar <> 0 "
			sql_nominal_detail = fmt.Sprintf("%s and kd_kategori = %d", sql_nominal_detail, kd_kategori)
			if paramData.Tgl_bayar1 != "" {
				sql_nominal_detail = fmt.Sprintf("%s and tgl_bayar >= '%s'", sql_nominal_detail, TglBayar1)
			}
			if paramData.Tgl_bayar2 != "" {
				sql_nominal_detail = fmt.Sprintf("%s and tgl_bayar <= '%s'", sql_nominal_detail, TglBayar2)
			}
			execute_sql_nominal_detail, _ := db.Raw(sql_nominal_detail).Rows()
			// defer execute_sql_nominal_detail.Close()
			for execute_sql_nominal_detail.Next() {
				execute_sql_nominal_detail.Scan(&tgl_bayar, &jml_bayar, &tipe_pembayaran, &pos_uang_masuk)

				arrayDetailBayarTemp := DetailBayarOut{}
				arrayDetailBayarTemp.Tgl_bayar = tgl_bayar
				arrayDetailBayarTemp.Jml_bayar = jml_bayar
				arrayDetailBayarTemp.Tipe_pembayaran = tipe_pembayaran
				arrayDetailBayarTemp.Pos_uang_masuk = pos_uang_masuk

				arrayDetailBayar = append(arrayDetailBayar, arrayDetailBayarTemp)
			}

			arraydetail.DetailBayar = arrayDetailBayar
			//end cari detail pembayaran

			SetArrayDetail = append(SetArrayDetail, arraydetail)

			defer execute_sql_nominal.Close()
			defer execute_sql_nominal_detail.Close()

		}
		defer execute_sql_kategori.Close()

		arraydata.Kategori = SetArrayDetail
		SetArrayData = append(SetArrayData, arraydata)
	}
	arrayJenisGroupUangKeluar.DetailData = SetArrayData
	//=========================================================================================================

	SetJenisGroupUang = append(SetJenisGroupUang, arrayJenisGroupUangKeluar)

	arrayGroupUang.DataUang = SetJenisGroupUang
	SetGroupUang = append(SetGroupUang, arrayGroupUang)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetGroupUang)
	c.JSON(http.StatusOK, response)
}

func Report_Group_Masuk_Old(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramData ParamData
	if err := c.ShouldBindJSON(&paramData); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var TglBayar1 string
	var TglBayar2 string

	if paramData.Tgl_bayar1 != "" {
		tTgl_bayar1, err2 := time.Parse("02-01-2006", paramData.Tgl_bayar1)
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
	} else {
		TglBayar1 = ""
	}

	if paramData.Tgl_bayar1 != "" {
		tTglBayar2, err2 := time.Parse("02-01-2006", paramData.Tgl_bayar2)
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
	} else {
		TglBayar2 = ""
	}

	SetJenisUang := []JenisGroupUang{}
	arrayJenisUang := JenisGroupUang{}
	arrayJenisUang.JenisUang = "Uang Masuk"

	//var iseng string

	SetArrayData := []GroupUangMasuk{}
	var kd_group int
	var nm_group string
	sql_group := " SELECT a.kd_group,a.nm_group FROM tbl_group_kategoris as a  " +
		" inner join tbl_kategori_uangs as c on a.kd_group = c.kd_group " +
		" Left join vw_s_masuk as b  on a.nm_group=b.nm_group and c.nm_kategori=b.nm_kategori "

	//groupnya biar pertanggal
	if paramData.Tgl_bayar1 != "" {
		sql_group = fmt.Sprintf("%s and b.tgl_bayar >= '%s'", sql_group, TglBayar1)
	}
	if paramData.Tgl_bayar2 != "" {
		sql_group = fmt.Sprintf("%s and b.tgl_bayar <= '%s'", sql_group, TglBayar2)
	}
	//end groupnya biar pertanggal

	sql_group = fmt.Sprintf("%s where a.flag_aktif=0 and a.kd_jenis=1 and a.nm_header <> '' and c.flag_aktif=0 ", sql_group)

	sql_group = fmt.Sprintf("%s  GROUP BY a.kd_group order by a.kd_group ", sql_group)

	// iseng = sql_group

	execute_sql_group, _ := db.Raw(sql_group).Rows()
	// defer execute_sql_group.Close()
	for execute_sql_group.Next() {
		execute_sql_group.Scan(&kd_group, &nm_group)
		arraydata := GroupUangMasuk{}
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group

		SetArrayDetail := []GroupUangMasukDetail{}
		var kd_kategori int
		var nm_kategori string
		sql_kategori := " SELECT a.kd_kategori,a.nm_kategori FROM tbl_kategori_uangs as a   " +
			" left join vw_s_masuk as b on a.nm_kategori=b.nm_kategori  "

		//groupnya biar pertanggal
		if paramData.Tgl_bayar1 != "" {
			sql_kategori = fmt.Sprintf("%s and b.tgl_bayar >= '%s'", sql_kategori, TglBayar1)
		}
		if paramData.Tgl_bayar2 != "" {
			sql_kategori = fmt.Sprintf("%s and b.tgl_bayar <= '%s'", sql_kategori, TglBayar2)
		}
		//end groupnya biar pertanggal

		sql_kategori = fmt.Sprintf("%s where a.flag_aktif=0 and a.kd_group = %d", sql_kategori, kd_group)

		sql_kategori = fmt.Sprintf("%s GROUP BY a.kd_kategori  order by a.kd_kategori ", sql_kategori)

		//iseng = sql_kategori

		execute_sql_kategori, _ := db.Raw(sql_kategori).Rows()
		//defer execute_sql_kategori.Close()
		for execute_sql_kategori.Next() {
			execute_sql_kategori.Scan(&kd_kategori, &nm_kategori)
			arraydetail := GroupUangMasukDetail{}
			arraydetail.Kd_kategori = kd_kategori
			arraydetail.Nm_kategori = nm_kategori

			//cari uang masuk
			var total_bayar float64
			sql_nominal := " SELECT ifnull(sum(jml_bayar),0.00) 'total_bayar' FROM vw_s_masuk  where total_bayar <> 0 "
			sql_nominal = fmt.Sprintf("%s and kd_kategori = %d", sql_nominal, kd_kategori)
			if paramData.Tgl_bayar1 != "" {
				sql_nominal = fmt.Sprintf("%s and tgl_bayar >= '%s'", sql_nominal, TglBayar1)
			}
			if paramData.Tgl_bayar2 != "" {
				sql_nominal = fmt.Sprintf("%s and tgl_bayar <= '%s'", sql_nominal, TglBayar2)
			}
			sql_nominal = fmt.Sprintf("%s  GROUP BY kd_kategori ", sql_nominal)
			execute_sql_nominal, _ := db.Raw(sql_nominal).Rows()
			//defer execute_sql_nominal.Close()
			for execute_sql_nominal.Next() {
				execute_sql_nominal.Scan(&total_bayar)
			}
			defer execute_sql_nominal.Close()
			arraydetail.Total_bayar = total_bayar
			//end cari uang masuk

			//cari detail pembayaran
			var tgl_bayar string
			var jml_bayar float64
			var tipe_pembayaran string

			arrayDetailBayar := []DetailBayar{}
			sql_nominal_detail := " SELECT tgl_bayar,jml_bayar,tipe_pembayaran  FROM vw_s_masuk  where total_bayar <> 0 "
			sql_nominal_detail = fmt.Sprintf("%s and kd_kategori = %d", sql_nominal_detail, kd_kategori)
			if paramData.Tgl_bayar1 != "" {
				sql_nominal_detail = fmt.Sprintf("%s and tgl_bayar >= '%s'", sql_nominal_detail, TglBayar1)
			}
			if paramData.Tgl_bayar2 != "" {
				sql_nominal_detail = fmt.Sprintf("%s and tgl_bayar <= '%s'", sql_nominal_detail, TglBayar2)
			}
			execute_sql_nominal_detail, _ := db.Raw(sql_nominal_detail).Rows()
			//defer execute_sql_nominal_detail.Close()
			for execute_sql_nominal_detail.Next() {
				execute_sql_nominal_detail.Scan(&tgl_bayar, &jml_bayar, &tipe_pembayaran)

				arrayDetailBayarTemp := DetailBayar{}
				arrayDetailBayarTemp.Tgl_bayar = tgl_bayar
				arrayDetailBayarTemp.Jml_bayar = jml_bayar
				arrayDetailBayarTemp.Tipe_pembayaran = tipe_pembayaran

				arrayDetailBayar = append(arrayDetailBayar, arrayDetailBayarTemp)
			}
			defer execute_sql_nominal_detail.Close()
			arraydetail.DetailBayar = arrayDetailBayar
			//end cari detail pembayaran

			SetArrayDetail = append(SetArrayDetail, arraydetail)
		}
		defer execute_sql_kategori.Close()

		arraydata.Kategori = SetArrayDetail
		SetArrayData = append(SetArrayData, arraydata)
	}
	defer execute_sql_group.Close()

	arrayJenisUang.DetailData = SetArrayData
	SetJenisUang = append(SetJenisUang, arrayJenisUang)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetJenisUang)
	c.JSON(http.StatusOK, response)

}

func Report_Group_Keluar_Old(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramData ParamData
	if err := c.ShouldBindJSON(&paramData); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var TglBayar1 string
	var TglBayar2 string

	if paramData.Tgl_bayar1 != "" {
		tTgl_bayar1, err2 := time.Parse("02-01-2006", paramData.Tgl_bayar1)
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
	} else {
		TglBayar1 = ""
	}

	if paramData.Tgl_bayar1 != "" {
		tTglBayar2, err2 := time.Parse("02-01-2006", paramData.Tgl_bayar2)
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
	} else {
		TglBayar2 = ""
	}

	SetJenisUang := []JenisGroupUang{}
	arrayJenisUang := JenisGroupUang{}
	arrayJenisUang.JenisUang = "Uang Keluar"

	SetArrayData := []GroupUangMasuk{}
	var kd_group int
	var nm_group string
	//sql_group := " SELECT kd_group,nm_group FROM tbl_group_kategoris where flag_aktif=0 and kd_jenis=2 and nm_header <> '' order by kd_group "

	//var iseng string
	sql_group := " SELECT a.kd_group,a.nm_group FROM tbl_group_kategoris as a  " +
		" inner join tbl_kategori_uangs as c on a.kd_group = c.kd_group " +
		" Left join vw_s_keluar as b  on a.nm_group=b.nm_group and c.nm_kategori=b.nm_kategori "

	//groupnya biar pertanggal
	if paramData.Tgl_bayar1 != "" {
		sql_group = fmt.Sprintf("%s and b.tgl_bayar >= '%s'", sql_group, TglBayar1)
	}
	if paramData.Tgl_bayar2 != "" {
		sql_group = fmt.Sprintf("%s and b.tgl_bayar <= '%s'", sql_group, TglBayar2)
	}
	//end groupnya biar pertanggal

	sql_group = fmt.Sprintf("%s where a.flag_aktif=0 and a.kd_jenis=2 and a.nm_header <> '' and c.flag_aktif=0 ", sql_group)

	sql_group = fmt.Sprintf("%s  GROUP BY a.kd_group order by a.kd_group ", sql_group)

	//iseng = sql_group

	execute_sql_group, _ := db.Raw(sql_group).Rows()
	for execute_sql_group.Next() {
		execute_sql_group.Scan(&kd_group, &nm_group)
		arraydata := GroupUangMasuk{}
		arraydata.Kd_group = kd_group
		arraydata.Nm_group = nm_group

		SetArrayDetail := []GroupUangMasukDetail{}
		var kd_kategori int
		var nm_kategori string
		// sql_kategori := " SELECT kd_kategori,nm_kategori FROM tbl_kategori_uangs where flag_aktif=0 "
		// sql_kategori = fmt.Sprintf("%s and kd_group = %d", sql_kategori, kd_group)
		// //groupnya biar pertanggal
		// if paramData.Tgl_bayar1 != "" {
		// 	sql_kategori = fmt.Sprintf("%s and tgl_bayar >= '%s'", sql_kategori, TglBayar1)
		// }
		// if paramData.Tgl_bayar2 != "" {
		// 	sql_kategori = fmt.Sprintf("%s and tgl_bayar <= '%s'", sql_kategori, TglBayar2)
		// }
		// //end groupnya biar pertanggal
		// sql_kategori = fmt.Sprintf("%s  order by kd_kategori ", sql_kategori)
		sql_kategori := " SELECT a.kd_kategori,a.nm_kategori FROM tbl_kategori_uangs as a   " +
			" left join vw_s_keluar as b on a.nm_kategori=b.nm_kategori  "

		//groupnya biar pertanggal
		if paramData.Tgl_bayar1 != "" {
			sql_kategori = fmt.Sprintf("%s and b.tgl_bayar >= '%s'", sql_kategori, TglBayar1)
		}
		if paramData.Tgl_bayar2 != "" {
			sql_kategori = fmt.Sprintf("%s and b.tgl_bayar <= '%s'", sql_kategori, TglBayar2)
		}
		//end groupnya biar pertanggal

		sql_kategori = fmt.Sprintf("%s where a.flag_aktif=0 and a.kd_group = %d", sql_kategori, kd_group)

		sql_kategori = fmt.Sprintf("%s GROUP BY a.kd_kategori  order by a.kd_kategori ", sql_kategori)

		//iseng = sql_kategori

		execute_sql_kategori, _ := db.Raw(sql_kategori).Rows()
		// defer execute_sql_kategori.Close()
		for execute_sql_kategori.Next() {
			execute_sql_kategori.Scan(&kd_kategori, &nm_kategori)
			arraydetail := GroupUangMasukDetail{}
			arraydetail.Kd_kategori = kd_kategori
			arraydetail.Nm_kategori = nm_kategori

			//cari uang masuk
			var total_bayar float64
			sql_nominal := " SELECT ifnull(sum(jml_bayar),0.00) 'total_bayar' FROM vw_s_keluar  where total_bayar <> 0 "
			sql_nominal = fmt.Sprintf("%s and kd_kategori = %d", sql_nominal, kd_kategori)
			if paramData.Tgl_bayar1 != "" {
				sql_nominal = fmt.Sprintf("%s and tgl_bayar >= '%s'", sql_nominal, TglBayar1)
			}
			if paramData.Tgl_bayar2 != "" {
				sql_nominal = fmt.Sprintf("%s and tgl_bayar <= '%s'", sql_nominal, TglBayar2)
			}
			sql_nominal = fmt.Sprintf("%s  GROUP BY kd_kategori ", sql_nominal)
			execute_sql_nominal, _ := db.Raw(sql_nominal).Rows()
			// defer execute_sql_nominal.Close()
			for execute_sql_nominal.Next() {
				execute_sql_nominal.Scan(&total_bayar)
			}

			arraydetail.Total_bayar = total_bayar
			//end cari uang masuk

			//cari detail pembayaran
			var tgl_bayar string
			var jml_bayar float64
			var tipe_pembayaran string
			var pos_uang_masuk string

			arrayDetailBayar := []DetailBayarOut{}
			sql_nominal_detail := " SELECT tgl_bayar,jml_bayar,tipe_pembayaran,pos_uang_masuk  FROM vw_s_keluar  where total_bayar <> 0 "
			sql_nominal_detail = fmt.Sprintf("%s and kd_kategori = %d", sql_nominal_detail, kd_kategori)
			if paramData.Tgl_bayar1 != "" {
				sql_nominal_detail = fmt.Sprintf("%s and tgl_bayar >= '%s'", sql_nominal_detail, TglBayar1)
			}
			if paramData.Tgl_bayar2 != "" {
				sql_nominal_detail = fmt.Sprintf("%s and tgl_bayar <= '%s'", sql_nominal_detail, TglBayar2)
			}
			execute_sql_nominal_detail, _ := db.Raw(sql_nominal_detail).Rows()
			// defer execute_sql_nominal_detail.Close()
			for execute_sql_nominal_detail.Next() {
				execute_sql_nominal_detail.Scan(&tgl_bayar, &jml_bayar, &tipe_pembayaran, &pos_uang_masuk)

				arrayDetailBayarTemp := DetailBayarOut{}
				arrayDetailBayarTemp.Tgl_bayar = tgl_bayar
				arrayDetailBayarTemp.Jml_bayar = jml_bayar
				arrayDetailBayarTemp.Tipe_pembayaran = tipe_pembayaran
				arrayDetailBayarTemp.Pos_uang_masuk = pos_uang_masuk

				arrayDetailBayar = append(arrayDetailBayar, arrayDetailBayarTemp)
			}

			arraydetail.DetailBayar = arrayDetailBayar
			//end cari detail pembayaran

			SetArrayDetail = append(SetArrayDetail, arraydetail)

			defer execute_sql_nominal.Close()
			defer execute_sql_nominal_detail.Close()

		}
		defer execute_sql_kategori.Close()

		arraydata.Kategori = SetArrayDetail
		SetArrayData = append(SetArrayData, arraydata)
	}
	defer execute_sql_group.Close()

	arrayJenisUang.DetailData = SetArrayData
	SetJenisUang = append(SetJenisUang, arrayJenisUang)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetJenisUang)
	c.JSON(http.StatusOK, response)
}

// //cari grand total uang masuk
// ssql_total := " SELECT sum(jml_bayar) 'grand_total' FROM vw_s_masuk  where total_bayar <> 0 "
// if paramData.Tgl_bayar1 != "" {
// 	ssql_total = fmt.Sprintf("%s and tgl_bayar >= '%s'", ssql_total, TglBayar1)
// }
// if paramData.Tgl_bayar2 != "" {
// 	ssql_total = fmt.Sprintf("%s and tgl_bayar <= '%s'", ssql_total, TglBayar2)
// }

// var grand_total float64
// rowstotal, _ := db.Raw(ssql_total).Rows()
// defer rowstotal.Close()
// for rowstotal.Next() {
// 	rowstotal.Scan(&grand_total)
// }

// SetArrayDataMasuk := []GrandTotal{}
// arraydataMasuk := GrandTotal{}
// arraydataMasuk.Grand_total = grand_total

// //end cari grand total uang masuk

// var nm_group string
// SetArrayData := []GroupUangMasuk{}
// ssql := " SELECT nm_group FROM vw_s_masuk where nm_group <> ''  "
// if paramData.Tgl_bayar1 != "" {
// 	ssql = fmt.Sprintf("%s and tgl_bayar >= '%s'", ssql, TglBayar1)
// }
// if paramData.Tgl_bayar2 != "" {
// 	ssql = fmt.Sprintf("%s and tgl_bayar <= '%s'", ssql, TglBayar2)
// }

// ssql = fmt.Sprintf("%s GROUP BY nm_group ", ssql)
// ssql = fmt.Sprintf("%s order by nm_group asc ", ssql)

// rows, _ := db.Raw(ssql).Rows()
// defer rows.Close()
// for rows.Next() {
// 	rows.Scan(&nm_group)
// 	arraydata := GroupUangMasuk{}
// 	arraydata.Nm_group = nm_group

// 	ssqldetail := " SELECT nm_kategori,sum(jml_bayar) 'total_bayar' FROM vw_s_masuk where nm_group='" + nm_group + "'  "
// if paramData.Tgl_bayar1 != "" {
// 	ssqldetail = fmt.Sprintf("%s and tgl_bayar >= '%s'", ssqldetail, TglBayar1)
// }
// if paramData.Tgl_bayar2 != "" {
// 	ssqldetail = fmt.Sprintf("%s and tgl_bayar <= '%s'", ssqldetail, TglBayar2)
// }

// 	ssqldetail = fmt.Sprintf("%s GROUP BY nm_kategori ", ssqldetail)
// 	ssqldetail = fmt.Sprintf("%s order by nm_kategori asc ", ssqldetail)

// 	SetArrayDetail := []GroupUangMasukDetail{}
// 	var nm_kategori string
// 	var total_bayar float64
// 	rowsdetail, _ := db.Raw(ssqldetail).Rows()
// 	defer rowsdetail.Close()
// 	for rowsdetail.Next() {
// 		rowsdetail.Scan(&nm_kategori, &total_bayar)
// 		arraydetail := GroupUangMasukDetail{}
// 		arraydetail.Nm_kategori = nm_kategori
// 		arraydetail.Total_bayar = total_bayar
// 		SetArrayDetail = append(SetArrayDetail, arraydetail)
// 	}
// 	arraydata.Kategori = SetArrayDetail

// 	SetArrayData = append(SetArrayData, arraydata)
// }

// arraydataMasuk.DetailData = SetArrayData
// SetArrayDataMasuk = append(SetArrayDataMasuk, arraydataMasuk)

// if len(SetArrayDataMasuk) == 0 {
// 	SetArrayDataMasuk := []GrandTotal{}
// 	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayDataMasuk)
// 	c.JSON(http.StatusOK, response)
// 	return
// }
