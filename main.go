package main

import (
	"fmt"
	"net/http"
	"rest_api_bendahara/connection"
	"rest_api_bendahara/helper"
	"rest_api_bendahara/master_conf_biaya_kategori"
	"rest_api_bendahara/master_conf_spp_ppdb"
	"rest_api_bendahara/master_group_kategori"
	"rest_api_bendahara/master_jenis_trans"
	"rest_api_bendahara/master_kategori_uang"
	"rest_api_bendahara/master_kelas_akademik"
	"rest_api_bendahara/master_sett_periode"
	"rest_api_bendahara/master_siswa_akademik"
	"rest_api_bendahara/master_sub_kategori_uang"
	"rest_api_bendahara/master_tahun_akademik"
	"rest_api_bendahara/transaksi_uang_keluar_pra"
	"rest_api_bendahara/transaksi_uang_keluar_pra_act"
	"rest_api_bendahara/transaksi_uang_masuk_lainlain"
	"rest_api_bendahara/transaksi_uang_masuk_ppdb"
	"rest_api_bendahara/transaksi_uang_masuk_siswa"
	"rest_api_bendahara/transaksi_uang_masuk_spp"
	"rest_api_bendahara/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("api/v1/")

	db := connection.SetupConnection()

	api.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	api.GET("/", func(c *gin.Context) {
		response := helper.APIResponse("Aplikasi Bendahara Sekolah SMA AL-KHAIRIYAH...", http.StatusOK, "success", nil)
		c.JSON(http.StatusOK, response)
	})

	// userInput := models.GroupKategoriInput{}
	// userInput.Kd_jenis = ""
	// userInput.Nm_group = "Biaya Pembayaran PPDB"
	// if err := ctx.ShouldBindJSON(&dataInput); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	api.POST("/users/signup", user.SignUp)
	api.POST("/users/login", user.Login)
	api.GET("/users/getuser", authMiddleware(), user.FetchUser)

	api.GET("/masterjenistrans/listjenistrans", authMiddleware(), master_jenis_trans.ListJenisTrans)
	api.GET("/masterjenistrans/showjenistrans", authMiddleware(), master_jenis_trans.ShowJenisTrans)
	api.POST("/masterjenistrans/insertjenistrans", authMiddleware(), master_jenis_trans.InsertJenisTrans)
	api.PUT("/masterjenistrans/updatejenistrans/:kdjenis", authMiddleware(), master_jenis_trans.UpdateJenisTrans)
	api.PUT("/masterjenistrans/deletejenistrans/:kdjenis", authMiddleware(), master_jenis_trans.DeleteJenisTrans)

	api.GET("/mastergroupkategori/listgroupkategori", authMiddleware(), master_group_kategori.ListGroupKategori)
	api.GET("/mastergroupkategori/showgroupkategori", authMiddleware(), master_group_kategori.ShowGroupKategori)
	api.POST("/mastergroupkategori/insertgroupkategori", authMiddleware(), master_group_kategori.InsertGroupKategori)
	api.PUT("/mastergroupkategori/updategroupkategori/:kdgroup", authMiddleware(), master_group_kategori.UpdateGroupKategori)
	api.PUT("/mastergroupkategori/deletegroupkategori/:kdgroup", authMiddleware(), master_group_kategori.DeleteGroupKategori)

	api.GET("/masterkategoriuang/listkategoriuang", authMiddleware(), master_kategori_uang.ListKategoriUang)
	api.GET("/masterkategoriuang/showkategoriuang", authMiddleware(), master_kategori_uang.ShowKategoriUang)
	api.POST("/masterkategoriuang/insertkategoriuang", authMiddleware(), master_kategori_uang.InsertKategoriUang)
	api.PUT("/masterkategoriuang/updatekategoriuang/:kdkategori", authMiddleware(), master_kategori_uang.UpdateKategoriUang)
	api.PUT("/masterkategoriuang/deletekategoriuang/:kdkategori", authMiddleware(), master_kategori_uang.DeleteKategoriUang)

	api.GET("/masterkategoriuang/listsubkategoriuang", authMiddleware(), master_sub_kategori_uang.ListSubKategoriUang)
	api.GET("/masterkategoriuang/showsubkategoriuang", authMiddleware(), master_sub_kategori_uang.ShowSubKategoriUang)
	api.POST("/masterkategoriuang/insertsubkategoriuang", authMiddleware(), master_sub_kategori_uang.InsertSubKategoriUang)
	api.PUT("/masterkategoriuang/updatesubkategoriuang/:kdsubkategori", authMiddleware(), master_sub_kategori_uang.UpdateSubKategoriUang)
	api.PUT("/masterkategoriuang/deletesubkategoriuang/:kdsubkategori", authMiddleware(), master_sub_kategori_uang.DeleteSubKategoriUang)

	api.GET("/settingperiode/listconfperiode", authMiddleware(), master_sett_periode.ListConfPeriode)
	api.GET("/settingperiode/showconfperiode", authMiddleware(), master_sett_periode.ShowConfPeriode)
	api.POST("/settingperiode/insertconfperiode", authMiddleware(), master_sett_periode.InsertConfPeriode)
	//api.PUT("/settingperiode/updateconfperiode/:idconf", authMiddleware(), master_sett_periode.UpdateConfPeriode)
	api.PUT("/settingperiode/updateconfperiodeall", authMiddleware(), master_sett_periode.UpdateConfPeriodeAll)
	api.POST("/settingperiode/deleteconfperiode", authMiddleware(), master_sett_periode.DeleteConfPeriode)

	api.GET("/masterconfsppdb/showconfspppdb", authMiddleware(), master_conf_spp_ppdb.ShowConfSppPPDB)
	api.PUT("/masterconfsppdb/updateconfspppdb/:idlink", authMiddleware(), master_conf_spp_ppdb.UpdateConfSppPPDB)

	api.GET("/masterbiayakategori/listbiayakategori", authMiddleware(), master_conf_biaya_kategori.ListBiayaKategori)
	api.GET("/masterbiayakategori/showbiayakategori", authMiddleware(), master_conf_biaya_kategori.ShowBiayaKategori)
	api.POST("/masterbiayakategori/insertbiayakategori", authMiddleware(), master_conf_biaya_kategori.InsertBiayaKategori)
	api.PUT("/masterbiayakategori/updatebiayakategori/:kdbiayakategori", authMiddleware(), master_conf_biaya_kategori.UpdateBiayaKategori)
	api.DELETE("/masterbiayakategori/deletebiayakategori/:kdbiayakategori", authMiddleware(), master_conf_biaya_kategori.DeleteBiayaKategori)

	// api.GET("/settingperiode/listsettperiode", authMiddleware(), master_sett_spp.ListSettPeriode)
	// api.GET("/settingperiode/showsettperiode", authMiddleware(), master_sett_spp.ShowSettPeriode)
	// api.POST("/settingperiode/insertsettperiode", authMiddleware(), master_sett_spp.InsertSettPeriode)
	// api.PUT("/settingperiode/updatesettperiode/:kdsettspp", authMiddleware(), master_sett_spp.UpdateSettPeriode)
	// api.PUT("/settingperiode/deletesettperiode/:kdsettspp", authMiddleware(), master_sett_spp.DeleteSettPeriode)

	//UANG MASUK SPP
	api.GET("/transaksi/uangmasukspp/listgroupkategori", authMiddleware(), transaksi_uang_masuk_spp.ListGroupKategori)
	api.GET("/transaksi/uangmasukspp/listkategoriuang", authMiddleware(), transaksi_uang_masuk_spp.ListKategoriUang)
	api.GET("/transaksi/uangmasukspp/listkelas", authMiddleware(), transaksi_uang_masuk_spp.ListKelas)
	api.POST("/transaksi/uangmasukspp/listsiswa", authMiddleware(), transaksi_uang_masuk_spp.ListSiswa)
	api.POST("/transaksi/uangmasukspp/listdata", authMiddleware(), transaksi_uang_masuk_spp.ListData)
	api.POST("/transaksi/uangmasukspp/createuangmasukspp", authMiddleware(), transaksi_uang_masuk_spp.CreateUangMasukSpp)
	api.PUT("/transaksi/uangmasukspp/updateuangmasukspp/:idhead/:iddetail", authMiddleware(), transaksi_uang_masuk_spp.UpdateUangMasukSpp)
	api.PUT("/transaksi/uangmasukspp/deletealluangmasuk/:idhead", authMiddleware(), transaksi_uang_masuk_spp.DeleteAllUangMasuk)

	//UANG MASUK PPDB
	api.GET("/transaksi/uangmasukppdb/listgroupkategori", authMiddleware(), transaksi_uang_masuk_ppdb.ListGroupKategori)
	api.GET("/transaksi/uangmasukppdb/listkategoriuang", authMiddleware(), transaksi_uang_masuk_ppdb.ListKategoriUang)
	api.GET("/transaksi/uangmasukppdb/listkelas", authMiddleware(), transaksi_uang_masuk_ppdb.ListKelas)
	api.POST("/transaksi/uangmasukppdb/listsiswa", authMiddleware(), transaksi_uang_masuk_ppdb.ListSiswa)
	api.POST("/transaksi/uangmasukppdb/listdata", authMiddleware(), transaksi_uang_masuk_ppdb.ListData)
	api.POST("/transaksi/uangmasukppdb/createuangmasukppdb", authMiddleware(), transaksi_uang_masuk_ppdb.CreateUangMasukPPdb)
	api.PUT("/transaksi/uangmasukppdb/updateuangmasukppdb/:idhead/:iddetail", authMiddleware(), transaksi_uang_masuk_ppdb.UpdateUangMasukPPdb)
	api.PUT("/transaksi/uangmasukppdb/deletealluangmasuk/:idhead", authMiddleware(), transaksi_uang_masuk_ppdb.DeleteAllUangMasuk)

	//UANG MASUK SISWA
	api.GET("/transaksi/uangmasuksiswa/listgroupkategori", authMiddleware(), transaksi_uang_masuk_siswa.ListGroupKategori)
	api.POST("/transaksi/uangmasuksiswa/listkategoriuang", authMiddleware(), transaksi_uang_masuk_siswa.ListKategoriUang)
	api.POST("/transaksi/uangmasuksiswa/listdataaddsiswa", authMiddleware(), transaksi_uang_masuk_siswa.ListDataAddSiswa)
	api.POST("/transaksi/uangmasuksiswa/createuangmasuksiswa", authMiddleware(), transaksi_uang_masuk_siswa.CreateUangMasukSiswa)
	api.PUT("/transaksi/uangmasuksiswa/edituangmasuksiswa/:idhead", authMiddleware(), transaksi_uang_masuk_siswa.EditUangMasukSiswa)
	api.PUT("/transaksi/uangmasuksiswa/updateuangmasuksiswadetail/:idhead/:iddetail", authMiddleware(), transaksi_uang_masuk_siswa.UpdateUangMasukSiswaDetail)
	api.POST("/transaksi/uangmasuksiswa/createuangmasuksiswadetail", authMiddleware(), transaksi_uang_masuk_siswa.CreateUangMasukSiswaDetail)
	api.POST("/transaksi/uangmasuksiswa/listdata", authMiddleware(), transaksi_uang_masuk_siswa.ListData)
	api.PUT("/transaksi/uangmasuksiswa/deleteuangmasuksiswadetail", authMiddleware(), transaksi_uang_masuk_siswa.DeleteUangMasukSiswaDetail)
	api.PUT("/transaksi/uangmasuksiswa/deletealluangmasuk/:idhead", authMiddleware(), transaksi_uang_masuk_siswa.DeleteAllUangMasuk)

	//UANG MASUK LAIN
	api.GET("/transaksi/uangmasuklainlain/listgroupkategori", authMiddleware(), transaksi_uang_masuk_lainlain.ListGroupKategori)
	api.POST("/transaksi/uangmasuklainlain/listkategoriuang", authMiddleware(), transaksi_uang_masuk_lainlain.ListKategoriUang)
	//api.POST("/transaksi/uangmasuklainlain/listdataaddsiswa", authMiddleware(), transaksi_uang_masuk_lainlain.ListDataAddSiswa)
	api.POST("/transaksi/uangmasuklainlain/createuangmasuklain", authMiddleware(), transaksi_uang_masuk_lainlain.CreateUangMasukLain)
	api.PUT("/transaksi/uangmasuklainlain/edituangmasuklain/:idhead", authMiddleware(), transaksi_uang_masuk_lainlain.EditUangMasukLain)
	api.PUT("/transaksi/uangmasuklainlain/updateuangmasuklaindetail/:idhead/:iddetail", authMiddleware(), transaksi_uang_masuk_lainlain.UpdateUangMasukLainetail)
	api.POST("/transaksi/uangmasuklainlain/createuangmasuklaindetail", authMiddleware(), transaksi_uang_masuk_lainlain.CreateUangMasukLainDetail)
	api.POST("/transaksi/uangmasuklainlain/listdata", authMiddleware(), transaksi_uang_masuk_lainlain.ListData)
	api.PUT("/transaksi/uangmasuklainlain/deleteuangmasuklaindetail", authMiddleware(), transaksi_uang_masuk_lainlain.DeleteUangMasukLainDetail)
	api.PUT("/transaksi/uangmasuklainlain/deletealluangmasuk/:idhead", authMiddleware(), transaksi_uang_masuk_lainlain.DeleteAllUangMasuk)

	//UANG Keluar PRA
	api.GET("/transaksi/uangkeluarpra/listgroupkategori", authMiddleware(), transaksi_uang_keluar_pra.ListGroupKategori)
	api.POST("/transaksi/uangkeluarpra/listkategoriuang", authMiddleware(), transaksi_uang_keluar_pra.ListKategoriUang)
	api.POST("/transaksi/uangkeluarpra/createuangkeluar", authMiddleware(), transaksi_uang_keluar_pra.CreateUangKeluar)
	api.PUT("/transaksi/uangkeluarpra/edituangkeluar/:idhead", authMiddleware(), transaksi_uang_keluar_pra.EditUangKeluar)
	api.PUT("/transaksi/uangkeluarpra/updateuangkeluardetail/:idhead/:iddetail", authMiddleware(), transaksi_uang_keluar_pra.UpdateUangKeluarDetail)
	api.POST("/transaksi/uangkeluarpra/createuangkeluardetail", authMiddleware(), transaksi_uang_keluar_pra.CreateUangKeluarDetail)
	api.POST("/transaksi/uangkeluarpra/listdata", authMiddleware(), transaksi_uang_keluar_pra.ListData)
	api.PUT("/transaksi/uangkeluarpra/deleteuangkeluardetail", authMiddleware(), transaksi_uang_keluar_pra.DeleteUangKeluarDetail)
	api.PUT("/transaksi/uangkeluarpra/deletealluangkeluar/:idhead", authMiddleware(), transaksi_uang_keluar_pra.DeleteAllUangKeluar)

	//UANG Keluar PRA ACT
	api.GET("/transaksi/uangkeluarpraact/listgroupkategori", authMiddleware(), transaksi_uang_keluar_pra_act.ListGroupKategori)
	api.POST("/transaksi/uangkeluarpraact/listkategoriuang", authMiddleware(), transaksi_uang_keluar_pra_act.ListKategoriUang)
	api.POST("/transaksi/uangkeluarpraact/createuangkeluar", authMiddleware(), transaksi_uang_keluar_pra_act.CreateUangKeluar)
	api.PUT("/transaksi/uangkeluarpraact/edituangkeluar/:idhead", authMiddleware(), transaksi_uang_keluar_pra_act.EditUangKeluar)
	api.PUT("/transaksi/uangkeluarpraact/updateuangkeluardetail/:idhead/:iddetail", authMiddleware(), transaksi_uang_keluar_pra_act.UpdateUangKeluarDetail)
	api.POST("/transaksi/uangkeluarpraact/createuangkeluardetail", authMiddleware(), transaksi_uang_keluar_pra_act.CreateUangKeluarDetail)
	api.POST("/transaksi/uangkeluarpraact/listdata", authMiddleware(), transaksi_uang_keluar_pra_act.ListData)
	api.PUT("/transaksi/uangkeluarpraact/deleteuangkeluardetail", authMiddleware(), transaksi_uang_keluar_pra_act.DeleteUangKeluarDetail)
	api.PUT("/transaksi/uangkeluarpraact/deletealluangkeluar/:idhead", authMiddleware(), transaksi_uang_keluar_pra_act.DeleteAllUangKeluar)

	//api.POST("/transaksi/uangmasuksiswa/listsiswa", authMiddleware(), transaksi_uang_masuk_siswa.ListSiswa)
	//api.GET("/transaksi/uangmasuksiswa/listkelas", authMiddleware(), transaksi_uang_masuk_spp.ListKelas)

	// api.POST("/transaksi/uangmasukspp/listdata", authMiddleware(), transaksi_uang_masuk_spp.ListData)
	// api.POST("/transaksi/uangmasukspp/createuangmasukspp", authMiddleware(), transaksi_uang_masuk_spp.CreateUangMasukSpp)
	// api.PUT("/transaksi/uangmasukspp/updateuangmasukspp/:idhead/:iddetail", authMiddleware(), transaksi_uang_masuk_spp.UpdateUangMasukSpp)

	//AKADEMIK SIA
	api.GET("/akademik/listtahunakademik", authMiddleware(), master_tahun_akademik.ListTahunAkademik)
	api.GET("/akademik/listkelasakademik", authMiddleware(), master_kelas_akademik.ListKelasAkademik)
	api.GET("/akademik/listsiswaakademik", authMiddleware(), master_siswa_akademik.ListSiswaAkademik)

	r.Run(":2022")
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Hak Akses Tidak Ditemukan (Error1) ...", http.StatusUnauthorized, "error", "Token Bermasalah...")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)

			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := user.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Hak Akses Tidak Ditemukan (Error2) ...", http.StatusUnauthorized, "error", "Token Bermasalah...")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			fmt.Println("Error ", err)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Hak Akses Tidak Ditemukan (Error3) ...", http.StatusUnauthorized, "error", "Token Bermasalah...")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)

			return
		}

		Id_user := int(claim["Id_user"].(float64))
		Username := claim["Username"]

		db := c.MustGet("db").(*gorm.DB)
		var datauser user.Tbl_user
		checkUser := db.Find(&datauser, "id_user = ? and Username= ?", Id_user, Username)
		if checkUser.RowsAffected == 0 {
			response := helper.APIResponse("Hak Akses Tidak Ditemukan (Error4) ...", http.StatusUnauthorized, "error", checkUser.Error)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)

			return
		}
		c.Set("currentUser", Username)

	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
