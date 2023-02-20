package dashboard

import (
	"net/http"
	"rest_api_bendahara/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DataDashboard(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	SetData_Transaksi := []Data_Transaksi{}
	arrayData_Transaksi := Data_Transaksi{}

	SetTotal_Detail_Transaksi := []Total_Detail_Transaksi{}
	arrayTotal_Detail_Transaksi := Total_Detail_Transaksi{}

	// Sisa Uang Bulan Lalu
	var keterangan_uang_bulan_lalu string
	sql := " SELECT concat('Sisa Saldo',' ',date_format(LAST_DAY((CURRENT_DATE - INTERVAL 1 MONTH)),'%d %M %Y')) 'keterangan_uang_bulan_lalu'  "
	db.Raw(sql).Scan(&keterangan_uang_bulan_lalu)
	arrayTotal_Detail_Transaksi.Keterangan = keterangan_uang_bulan_lalu

	var uang_masuk_bulan_lalu float64
	sql = " SELECT sum(total_bayar) FROM vw_report_umsiswa_dll " +
		" where DATE_FORMAT(tgl_bayar,'%Y-%m') <= DATE_FORMAT((CURRENT_DATE() - INTERVAL 1 MONTH),'%Y-%m')  "
	db.Raw(sql).Scan(&uang_masuk_bulan_lalu)

	var uang_keluar_bulan_lalu float64
	sql = " SELECT sum(total_bayar) FROM vw_report_uksiswa_dll " +
		" where DATE_FORMAT(tgl_bayar,'%Y-%m') <= DATE_FORMAT((CURRENT_DATE() - INTERVAL 1 MONTH),'%Y-%m')  "
	db.Raw(sql).Scan(&uang_keluar_bulan_lalu)

	arrayTotal_Detail_Transaksi.Total = uang_masuk_bulan_lalu - uang_keluar_bulan_lalu

	SetTotal_Detail_Transaksi = append(SetTotal_Detail_Transaksi, arrayTotal_Detail_Transaksi)
	arrayData_Transaksi.DetailData = SetTotal_Detail_Transaksi

	// End Sisa Uang Bulan Lalu

	// dapetin saldo masuk per tanggal 1 s/d tanggal kemarin
	arrayTotal_Detail_Transaksi = Total_Detail_Transaksi{}

	var keterangan_tgl_1_sd_tgl_kemarin string
	sql = " SELECT concat('Saldo Masuk Tgl 1',' s/d ',date_format((CURRENT_DATE - INTERVAL 1 DAY),'%d %M %Y')) 'keterangan_tgl_1_sd_tgl_kemarin' "
	db.Raw(sql).Scan(&keterangan_tgl_1_sd_tgl_kemarin)
	arrayTotal_Detail_Transaksi.Keterangan = keterangan_tgl_1_sd_tgl_kemarin

	var saldo_tgl_1_sd_tgl_kemarin float64
	sql = " SELECT sum(total_bayar) FROM vw_report_umsiswa_dll " +
		" where tgl_bayar >= DATE(concat(date_format(CURRENT_DATE,'%Y-%m'),'-','01')) " +
		" and tgl_bayar <= (CURRENT_DATE - INTERVAL 1 DAY)  "
	db.Raw(sql).Scan(&saldo_tgl_1_sd_tgl_kemarin)

	arrayTotal_Detail_Transaksi.Total = saldo_tgl_1_sd_tgl_kemarin
	SetTotal_Detail_Transaksi = append(SetTotal_Detail_Transaksi, arrayTotal_Detail_Transaksi)
	arrayData_Transaksi.DetailData = SetTotal_Detail_Transaksi

	// end dapetin saldo masuk per tanggal 1 s/d tanggal kemarin

	// dapetin saldo keluar per tanggal 1 s/d tanggal kemarin
	arrayTotal_Detail_Transaksi = Total_Detail_Transaksi{}
	keterangan_tgl_1_sd_tgl_kemarin = ""
	sql = " SELECT concat('Saldo Keluar Tgl 1',' s/d ',date_format((CURRENT_DATE - INTERVAL 1 DAY),'%d %M %Y')) 'keterangan_tgl_1_sd_tgl_kemarin' "
	db.Raw(sql).Scan(&keterangan_tgl_1_sd_tgl_kemarin)
	arrayTotal_Detail_Transaksi.Keterangan = keterangan_tgl_1_sd_tgl_kemarin

	saldo_tgl_1_sd_tgl_kemarin = 0
	sql = " SELECT sum(total_bayar) FROM vw_report_uksiswa_dll " +
		" where tgl_bayar >= DATE(concat(date_format(CURRENT_DATE,'%Y-%m'),'-','01')) " +
		" and tgl_bayar <= (CURRENT_DATE - INTERVAL 1 DAY)  "
	db.Raw(sql).Scan(&saldo_tgl_1_sd_tgl_kemarin)

	arrayTotal_Detail_Transaksi.Total = saldo_tgl_1_sd_tgl_kemarin
	SetTotal_Detail_Transaksi = append(SetTotal_Detail_Transaksi, arrayTotal_Detail_Transaksi)
	arrayData_Transaksi.DetailData = SetTotal_Detail_Transaksi

	// end dapetin saldo keluar per tanggal 1 s/d tanggal kemarin

	// dapetin saldo masuk hari ini
	arrayTotal_Detail_Transaksi = Total_Detail_Transaksi{}
	keterangan_tgl_1_sd_tgl_kemarin = ""
	sql = " SELECT concat('Saldo Masuk Tgl ','',date_format(CURRENT_DATE,'%d %M %Y')) 'keterangan_tgl_1_sd_tgl_kemarin' "
	db.Raw(sql).Scan(&keterangan_tgl_1_sd_tgl_kemarin)
	arrayTotal_Detail_Transaksi.Keterangan = keterangan_tgl_1_sd_tgl_kemarin

	saldo_tgl_1_sd_tgl_kemarin = 0
	sql = " SELECT sum(total_bayar) FROM vw_report_umsiswa_dll " +
		" where tgl_bayar >= CURRENT_DATE " +
		" and tgl_bayar <= CURRENT_DATE  "
	db.Raw(sql).Scan(&saldo_tgl_1_sd_tgl_kemarin)

	arrayTotal_Detail_Transaksi.Total = saldo_tgl_1_sd_tgl_kemarin
	SetTotal_Detail_Transaksi = append(SetTotal_Detail_Transaksi, arrayTotal_Detail_Transaksi)
	arrayData_Transaksi.DetailData = SetTotal_Detail_Transaksi

	// end dapetin saldo masuk hari ini

	// dapetin saldo keluar hari ini
	arrayTotal_Detail_Transaksi = Total_Detail_Transaksi{}
	keterangan_tgl_1_sd_tgl_kemarin = ""
	sql = " SELECT concat('Saldo Keluar Tgl ','',date_format(CURRENT_DATE,'%d %M %Y')) 'keterangan_tgl_1_sd_tgl_kemarin' "
	db.Raw(sql).Scan(&keterangan_tgl_1_sd_tgl_kemarin)
	arrayTotal_Detail_Transaksi.Keterangan = keterangan_tgl_1_sd_tgl_kemarin

	saldo_tgl_1_sd_tgl_kemarin = 0
	sql = " SELECT sum(total_bayar) FROM vw_report_uksiswa_dll " +
		" where tgl_bayar >= CURRENT_DATE " +
		" and tgl_bayar <= CURRENT_DATE  "
	db.Raw(sql).Scan(&saldo_tgl_1_sd_tgl_kemarin)

	arrayTotal_Detail_Transaksi.Total = saldo_tgl_1_sd_tgl_kemarin
	SetTotal_Detail_Transaksi = append(SetTotal_Detail_Transaksi, arrayTotal_Detail_Transaksi)
	arrayData_Transaksi.DetailData = SetTotal_Detail_Transaksi

	// end dapetin saldo keluar hari ini

	SetData_Transaksi = append(SetData_Transaksi, arrayData_Transaksi)

	response := helper.APIResponse("List Data ...", http.StatusOK, "success", SetData_Transaksi)
	c.JSON(http.StatusOK, response)
}
