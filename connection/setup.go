package connection

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupConnection() *gorm.DB {

	//dsn := "teguh:Bendahara@Montong54321X@tcp(127.0.0.1:3306)/smaalkha_db_bendahara?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "teguh:teguh@tcp(127.0.0.1:3314)/smaalkha_db_bendahara?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// return db

	if err != nil {
		log.Fatal(err.Error())
	}

	// Set the maximum number of open connections and idle connections
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Adjust these values based on your requirements
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	//sqlDB.SetConnMaxLifetime(time.Second * 28800)

	// Check if the database connection is successful
	err = sqlDB.Ping()
	if err != nil {
		if isMaxConnectionsError(err) {
			log.Println("Error: Maximum connections reached. Retrying in 10 seconds...")
			time.Sleep(10 * time.Second) // Add a delay before retrying
			err = sqlDB.Ping()
			if err != nil {
				log.Fatal("Failed to reconnect after delay:", err.Error())
			}

		} else {
			log.Fatal(err.Error())
		}
	}

	log.Println("Connected to the database!")

	return db
}

func SetupConnectionSIA() *gorm.DB {
	//dsn := "root:Bendahara@Montong54321X@tcp(127.0.0.1:3307)/smaalkha_akademik?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "teguh:teguh@tcp(127.0.0.1:3314)/smaalkha_akademik?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// return db

	if err != nil {
		log.Fatal(err.Error())
	}

	// Set the maximum number of open connections and idle connections
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Adjust these values based on your requirements
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	//sqlDB.SetConnMaxLifetime(time.Second * 28800)

	// Check if the database connection is successful
	err = sqlDB.Ping()
	if err != nil {
		if isMaxConnectionsError(err) {
			log.Println("Error: Maximum connections reached. Retrying in 10 seconds...")
			time.Sleep(10 * time.Second) // Add a delay before retrying
			err = sqlDB.Ping()
			if err != nil {
				log.Fatal("Failed to reconnect after delay:", err.Error())
			}

		} else {
			log.Fatal(err.Error())
		}
	}

	log.Println("Connected to the database!")

	return db
}

func isMaxConnectionsError(err error) bool {
	// Adjust this condition based on the specific error message or code for "Too many connections" in your MySQL server.
	return err.Error() == "Error 1040: Too many connections"
}

// func SetupConnection() *gorm.DB {
// 	dsn := "roots:roots@tcp(127.0.0.1:3306)/cuangshu_db_bendahara?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	return db
// }

// func SetupConnection() *gorm.DB {
// 	//dsn := "root:root@1234569@tcp(127.0.0.1:3314)/smaalkha_db_bendahara?charset=utf8mb4&parseTime=True&loc=Local"
// 	dsn := "smaalkha_db_bendahara:Db_Bendahara_Alkhairiyah@JKT2023@tcp(127.0.0.1:3306)/smaalkha_db_bendahara?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	sqlDB.SetMaxOpenConns(10)
// 	sqlDB.SetMaxIdleConns(5)

// 	// SetConnMaxLifetime menentukan berapa lama suatu koneksi boleh hidup sebelum ditutup dan digantikan oleh koneksi baru.
// 	sqlDB.SetConnMaxLifetime(time.Second * 25) // Set lebih pendek dari `wait_timeout`

// 	// Membuat goroutine untuk menjalankan aktivitas periodik setiap 20 detik.
// 	go func() {
// 		ticker := time.NewTicker(20 * time.Second)
// 		defer ticker.Stop()

// 		for {
// 			select {
// 			case <-ticker.C:
// 				// Melakukan aktivitas periodik untuk menjaga koneksi tetap hidup.
// 				if err := performPeriodicActivity(db); err != nil {
// 					handleConnectionError(db, err)
// 				}
// 			}
// 		}
// 	}()

// 	// Memeriksa koneksi setelah inisialisasi
// 	if err := sqlDB.Ping(); err != nil {
// 		handleConnectionError(db, err)
// 	}

// 	log.Println("Connected to the database!")

// 	return db
// }

// func performPeriodicActivity(db *gorm.DB) error {
// 	// Implementasikan aktivitas periodik di sini.
// 	// Sebagai contoh, kita hanya menjalankan kueri sederhana untuk menjaga koneksi tetap hidup.
// 	if err := db.Exec("SELECT 1").Error; err != nil {
// 		return fmt.Errorf("failed to perform periodic activity: %w", err)
// 	}
// 	return nil
// }

// func handleConnectionError(db *gorm.DB, err error) {
// 	if isMaxConnectionsError(err) {
// 		log.Println("Error: Maximum connections reached. Retrying in 10 seconds...")
// 		time.Sleep(10 * time.Second)
// 		err = db.Exec("SELECT 1").Error
// 		if err != nil {
// 			log.Fatal("Failed to reconnect after delay:", err.Error())
// 		}
// 	} else {
// 		log.Fatal(err.Error())
// 	}
// }

// func SetupConnectionSIA() *gorm.DB {
// 	//dsn := "root:root@1234569@tcp(127.0.0.1:3314)/smaalkha_akademik?charset=utf8mb4&parseTime=True&loc=Local"
// 	dsn := "smaalkha_montong:smaAlkhairiyah@12345@tcp(127.0.0.1:3306)/smaalkha_akademik?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	sqlDB.SetMaxOpenConns(10)
// 	sqlDB.SetMaxIdleConns(5)

// 	// SetConnMaxLifetime menentukan berapa lama suatu koneksi boleh hidup sebelum ditutup dan digantikan oleh koneksi baru.
// 	sqlDB.SetConnMaxLifetime(time.Second * 25) // Set lebih pendek dari `wait_timeout`

// 	// Membuat goroutine untuk menjalankan aktivitas periodik setiap 20 detik.
// 	go func() {
// 		ticker := time.NewTicker(20 * time.Second)
// 		defer ticker.Stop()

// 		for {
// 			select {
// 			case <-ticker.C:
// 				// Melakukan aktivitas periodik untuk menjaga koneksi tetap hidup.
// 				if err := performPeriodicActivity(db); err != nil {
// 					handleConnectionError(db, err)
// 				}
// 			}
// 		}
// 	}()

// 	// Memeriksa koneksi setelah inisialisasi
// 	if err := sqlDB.Ping(); err != nil {
// 		handleConnectionError(db, err)
// 	}

// 	log.Println("Connected to the database!")

// 	return db
// }

// isMaxConnectionsError checks if the given error indicates a maximum connections error.
