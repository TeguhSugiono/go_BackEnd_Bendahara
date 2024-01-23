package master_sumber_dana

import (
	"fmt"
	"net/http"
	"rest_api_bendahara/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SumberDana(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var paramData ParamData
	if err := c.ShouldBindJSON(&paramData); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var PosDanaMasuk []DanaMasuk
	sqlDanaMasuk := " SELECT SUM(total_bayar) 'TotalBayar',kd_group 'Kd_group',nm_group 'Nm_group' FROM vw_group_uang_masuk   "
	if paramData.Kd_group != "" {
		sqlDanaMasuk = fmt.Sprintf("%s where kd_group = '%s'", sqlDanaMasuk, paramData.Kd_group)
	}
	sqlDanaMasuk = fmt.Sprintf("%s GROUP BY kd_group", sqlDanaMasuk)
	db.Raw(sqlDanaMasuk).Scan(&PosDanaMasuk)

	var PosDanaKeluar []DanaKeluar
	sqlDanaKeluar := " SELECT SUM(total_bayar) 'TotalBayar',kd_group 'Kd_group',nm_group 'Nm_group' FROM vw_group_uang_keluar "
	if paramData.Kd_group != "" {
		sqlDanaKeluar = fmt.Sprintf("%s where kd_group = '%s'", sqlDanaKeluar, paramData.Kd_group)
	}
	sqlDanaKeluar = fmt.Sprintf("%s GROUP BY kd_group", sqlDanaKeluar)
	db.Raw(sqlDanaKeluar).Scan(&PosDanaKeluar)

	SetArrayDana := []DanaMasuk{}

	for _, dataMasuk := range PosDanaMasuk {

		ArrayDana := DanaMasuk{}

		DanaMasuk_Kd_group := dataMasuk.Kd_group
		DanaMasuk_Nm_group := dataMasuk.Nm_group
		DanaMasuk_TotalBayar := dataMasuk.TotalBayar

		for _, dataKeluar := range PosDanaKeluar {
			DanaKeluar_Kd_group := dataKeluar.Kd_group
			DanaKeluar_TotalBayar := dataKeluar.TotalBayar

			if DanaMasuk_Kd_group == DanaKeluar_Kd_group {
				DanaMasuk_TotalBayar = DanaMasuk_TotalBayar - DanaKeluar_TotalBayar
			}

		}

		ArrayDana.Kd_group = DanaMasuk_Kd_group
		ArrayDana.Nm_group = DanaMasuk_Nm_group
		ArrayDana.TotalBayar = DanaMasuk_TotalBayar

		SetArrayDana = append(SetArrayDana, ArrayDana)
	}

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayDana)
	c.JSON(http.StatusOK, response)
}
