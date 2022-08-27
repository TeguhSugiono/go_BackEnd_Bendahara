package main

import (
	"fmt"
	"net/http"

	"rest_api_bendahara/connection"
	"rest_api_bendahara/helper"
	"rest_api_bendahara/master_group_kategori"
	"rest_api_bendahara/master_jenis_trans"
	"rest_api_bendahara/master_kategori_uang"
	"rest_api_bendahara/master_sub_kategori_uang"
	"rest_api_bendahara/user"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/dgrijalva/jwt-go"
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

	r.Run(":2022")
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Hak Akses Tidak Ditemukan ...", http.StatusUnauthorized, "error", nil)
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
			response := helper.APIResponse("Hak Akses Tidak Ditemukan ...", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			fmt.Println("Error ", err)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Hak Akses Tidak Ditemukan ...", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)

			return
		}

		Id_user := int(claim["Id_user"].(float64))
		Username := claim["Username"]

		db := c.MustGet("db").(*gorm.DB)
		var datauser user.Tbl_user
		checkUser := db.Find(&datauser, "id_user = ? and Username= ?", Id_user, Username)
		if checkUser.RowsAffected == 0 {
			response := helper.APIResponse("Hak Akses Tidak Ditemukan ...", http.StatusUnauthorized, "error", checkUser.Error)
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
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
