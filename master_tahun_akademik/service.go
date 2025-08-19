package master_tahun_akademik

import (
	"net/http"
	"rest_api_bendahara/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListTahunAkademik(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var returnListData []ReturnListData
	//db.Find(&master)
	sql := " SELECT tahun_akademik as 'id_tahun',tahun_akademik,status FROM tbl_tahun_akademik where flag_tahun=0  order by status "

	db.Raw(sql).Scan(&returnListData)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", returnListData)
	c.JSON(http.StatusOK, response)
}
