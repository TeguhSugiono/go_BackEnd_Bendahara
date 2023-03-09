package master_tipe_pembayaran

import (
	"net/http"
	"rest_api_bendahara/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListTipePembayaran(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var listData []ListData
	sql := " SELECT kd_pembayaran,tipe_pembayaran,`default` FROM tbl_tipe_pembayarans where flag_aktif=0  order by `default` desc  "

	db.Raw(sql).Scan(&listData)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", listData)
	c.JSON(http.StatusOK, response)
}
