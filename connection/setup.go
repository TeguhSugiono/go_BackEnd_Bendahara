package connection

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// func SetupConnection() *gorm.DB {
// 	dsn := "roots:roots@tcp(127.0.0.1:3306)/cuangshu_db_bendahara?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	return db
// }

func SetupConnection() *gorm.DB {
	dsn := "roots:roots@tcp(127.0.0.1:3306)/cuangshu_db_bendahara?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "smaalkha_db_bendahara:Db_Bendahara_Alkhairiyah@JKT2023@tcp(127.0.0.1:3306)/smaalkha_db_bendahara?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

func SetupConnectionSIA() *gorm.DB {
	dsn := "roots:roots@tcp(127.0.0.1:3306)/smaalkha_akademik?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "smaalkha_montong:smaAlkhairiyah@12345@tcp(127.0.0.1:3306)/smaalkha_akademik?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
