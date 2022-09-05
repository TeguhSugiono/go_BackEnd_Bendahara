package master_sett_spp

import (
	"net/http"
	"rest_api_bendahara/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListSettPeriode(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var master []ListData

	sql := " SELECT a.*,b.nm_sett FROM tbl_sett_periode_spps a " +
		" INNER JOIN tbl_conf_periode_spps b on a.kd_periode_spp=b.kd_periode_spp " +
		" where a.flag_aktif=0 and b.flag_aktif=0 "

	db.Raw(sql).Scan(&master)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", FormatShowData(master))
	c.JSON(http.StatusOK, response)
}
