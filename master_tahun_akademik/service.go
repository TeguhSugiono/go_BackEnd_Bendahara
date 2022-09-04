package master_tahun_akademik

import (
	"net/http"
	"rest_api_bendahara/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListTahunAkademik(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var master []ListData
	//db.Find(&master)
	db.Where("flag_tahun = ?", 0).Find(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", FormatShowData(master))
	c.JSON(http.StatusOK, response)
}
