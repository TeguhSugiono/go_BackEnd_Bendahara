package master_siswa_akademik

import (
	"fmt"
	"net/http"
	"rest_api_bendahara/connection"
	"rest_api_bendahara/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListSiswaAkademik(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var master []ListData
	//db.Find(&master)
	sql := " SELECT nis,nm_siswa FROM tbl_siswa " +
		" WHERE flag_siswa = 0 AND status_siswa NOT IN ('Tidak Aktif','LULUS') order by nm_siswa "

	db.Raw(sql).Scan(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", FormatShowData(master))
	c.JSON(http.StatusOK, response)
}

func ListSiswaLulus(c *gin.Context) {
	dbSIA := connection.SetupConnectionSIA()

	//dbSIA := c.MustGet("dbSIA").(*gorm.DB)

	var paramSearch SearchSiswaLulus
	if err := c.ShouldBindJSON(&paramSearch); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var master []ListDataSiswa
	sql := " SELECT a.id_siswa,a.nm_siswa,a.tahun_lulus,a.no_peserta,b.nis " +
		"	FROM tbl_lulus a  " +
		"	INNER JOIN tbl_siswa b on a.id_siswa=b.id_siswa " +
		"	where a.flag_lulus=0 and b.flag_siswa=0 and b.status_siswa <> 'Tidak Aktif' "

	// if paramSearch.Id_siswa != "" {
	// 	sql = fmt.Sprintf("%s and id_siswa = '%s'", sql, paramSearch.Id_siswa)
	// }

	// if paramSearch.Tahun_lulus != "" {
	// 	sql = fmt.Sprintf("%s and tahun_lulus = '%s'", sql, paramSearch.Tahun_lulus)
	// }

	if paramSearch.No_peserta != "" {
		sql = fmt.Sprintf("%s and a.no_peserta = '%s'", sql, paramSearch.No_peserta)
	}

	dbSIA.Raw(sql).Scan(&master)

	if len(master) == 0 {
		SetArrayData := []ListDataSiswa{}
		response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", master)
	c.JSON(http.StatusOK, response)
}

func ListDataSiswaAll(c *gin.Context) {
	//Function ini berfungsi mencari data siswa yang statusnya aktif, lulus

	dbSIA := connection.SetupConnectionSIA()

	//dbSIA := c.MustGet("dbSIA").(*gorm.DB)

	var paramSiswaAll ParamSiswaAll
	if err := c.ShouldBindJSON(&paramSiswaAll); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validasi ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var master []ListSiswaAll
	sql := " SELECT id_siswa,nis,nm_siswa from tbl_siswa where status_siswa  in ('LULUS','Aktif') and flag_siswa=0 "

	if paramSiswaAll.Nis != "" {
		sql = fmt.Sprintf("%s and nis = '%s'", sql, paramSiswaAll.Nis)
	}

	dbSIA.Raw(sql).Scan(&master)

	if len(master) == 0 {
		SetArrayData := []ListSiswaAll{}
		response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetArrayData)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", master)
	c.JSON(http.StatusOK, response)
}
