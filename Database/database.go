package Database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() {
	checking := os.Getenv("DATABASE_URL")

	if checking == "" {
		log.Fatal("DATABASE URL Environment Variable not set")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(checking), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		log.Fatal("Gagal Terhubung ke database", err)
	}

	log.Println("Koneksi Database Berhasil Terhubung")
}
