package transaksi_uang_masuk_spp

import (
	"net/http"
	"rest_api_bendahara/helper"
	"rest_api_bendahara/master_group_kategori"
	"rest_api_bendahara/master_kategori_uang"

	//"rest_api_bendahara/master_kategori_uang"

	"github.com/gin-gonic/gin"
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

	response := helper.APIResponse("List Data ..."+sql, http.StatusOK, "success", master_group_kategori.FormatGroupKategori(master))
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

func InsertTransSpp(c *gin.Context) {

}

// func LinkKdGroup(input GetLinkKdGroup) (string, error) {
// 	return "", nil
// }

// func LinkKdKategori(input GetLinkKdKategori) (string, error) {
// 	return "", nil
// }

// var result DataTokenInput
// 	db.Raw("SELECT Id_user,Password,Username FROM tbl_users WHERE Username = ?", dataInput.Username).Scan(&result)
// token, err := GenerateToken(result)
// 	if err != nil {
// 		response := helper.APIResponse("Generate Token Gagal ...", http.StatusBadRequest, "error", nil)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}
