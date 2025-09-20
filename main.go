package main

import (
	"absensiTU-Backend/Database"
	"absensiTU-Backend/Models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {
	migrasi := godotenv.Load()
	if migrasi != nil {
		log.Fatal("Error loading .env file")
	}

	Database.Connect()
	//Migrasi Database
	migrasi = Database.DB.AutoMigrate(
		&Models.TUAccessPasswords{},
		&Models.TUAttendances{},
		&Models.TUEmbed{},
		&Models.TUPermissions{},
		&Models.TURole{},
		&Models.TUSchedules{},
		&Models.TUUsers{},
	)

	if migrasi != nil {
		log.Fatal("Gagal Migrasi Database", migrasi)
	}

	router := gin.Default()

	//Cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"}, // Ganti dengan domain frontend Anda
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	AuthRoutes := router.Group("/auth")
	AuthRoutes.POST("/login")

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "Welcome to !",
			"version":  "1.0.0",
			"Author":   "Rizky Budiarto",
			"Username": "KingRovs",
		})
	})

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	log.Printf("Server berjalan di http://localhost:%s", PORT)
	if err := router.Run(":" + PORT); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
