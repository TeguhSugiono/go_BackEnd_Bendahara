package user

import (
	"net/http"

	"rest_api_bendahara/helper"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SignUp(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var dataInput SignUpInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validation ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var datauser Tbl_user
	checkUser := db.Select("*").Where("username = ?", dataInput.Username).Find(&datauser)
	if checkUser.RowsAffected > 0 {

		errorMessage := gin.H{"errors": "Username Already Exists ..."}
		response := helper.APIResponse("Checking Username ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//var datenow string = time.Now().Format("2006-01-02 15:04:05")
	var datenows string = time.Now().UTC().Format("2006-01-02 15:04:05")
	date := "2006-01-02 15:04:05"
	datenowx, err := time.Parse(date, datenows)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors, "date": datenowx}
		response := helper.APIResponse("Wrong Date Format ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(dataInput.Password), bcrypt.MinCost)

	data := Tbl_user{
		Username:   dataInput.Username,
		Password:   string(passwordHash),
		Created_on: datenowx,
	}

	err = db.Create(&data).Error
	if err != nil {
		response := helper.APIResponse("Save Data Failed ...", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var result DataTokenInput
	db.Raw("SELECT Id_user,Username FROM tbl_users WHERE Username = ?", dataInput.Username).Scan(&result)

	token, err := GenerateToken(result)
	if err != nil {
		response := helper.APIResponse("Generate Token Failed ...", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := FormatUser(data, token)
	response := helper.APIResponse("Save Data Successfully ...", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

//db.Table("tbl_users").Select("Id_user", "Username").Where("Username = ?", dataInput.Username).Scan(&result)
// claim := jwt.MapClaims{}
// claim["Id_user"] = result.Id_user
// claim["Username"] = result.Username

// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

// signedToken, _ := token.SignedString(SECRET_KEY)
// if err != nil {
// 	return signedToken, err
// }

//var inputToken result

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var dataInput LoginInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error Validation ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//cek dulu datanya
	var datauser Tbl_user
	checkUser := db.Select("*").Where("username = ?", dataInput.Username).Find(&datauser)
	if checkUser.RowsAffected == 0 {
		errorMessage := gin.H{"errors": "Username Not Found ..."}
		response := helper.APIResponse("Login Failed ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//username := dataInput.Username
	//password := dataInput.Password

	var result DataTokenInput
	db.Raw("SELECT Id_user,Password,Username FROM tbl_users WHERE Username = ?", dataInput.Username).Scan(&result)

	err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(dataInput.Password))
	if err != nil {
		errorMessage := gin.H{"errors": "Wrong Password ..."}
		response := helper.APIResponse("Login Failed ...", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := GenerateToken(result)
	if err != nil {
		response := helper.APIResponse("Generate Token Failed ...", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := FormatUser(datauser, token)
	//response := helper.APIResponse("Login Berhasil ... "+result.Password+" -- "+dataInput.Password, http.StatusOK, "success", formatter)
	response := helper.APIResponse("Login Successfully ...", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func FetchUser(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(Tbl_user)
	formatter := FormatUser(currentUser, "")

	response := helper.APIResponse("Successfuly Fetch User Data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
