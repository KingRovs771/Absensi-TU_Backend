package main

import (
	"absensiTU-Backend/Controllers"
	"absensiTU-Backend/Database"
	"absensiTU-Backend/Middleware"
	"absensiTU-Backend/Models"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	Database.Connect()

	if err := Database.DB.Migrator().AutoMigrate(
		&Models.TUAccessPasswords{},
		&Models.TUAttendances{},
		&Models.TUEmbed{},
		&Models.TUPermissions{},
		&Models.TURole{},
		&Models.TUSchedules{},
		&Models.TUUsers{},
	); err != nil {
		log.Fatal("Gagal Migrasi Database:", err)
	}

	router := gin.Default()

	// Logging
	router.Use(gin.Logger())

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup routes
	setupRoutes(router)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	log.Printf("Server berjalan di http://localhost:%s", PORT)
	if err := router.Run(":" + PORT); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}

func setupRoutes(r *gin.Engine) {
	// Authentication
	AuthRoutes := r.Group("/auth/tu")
	{
		AuthRoutes.POST("/loginTU", Controllers.LoginHandler)
		AuthRoutes.POST("/logoutTU", Controllers.LogoutHandler)
	}

	// Profile
	ProfileRoutes := r.Group("/profile")
	ProfileRoutes.Use(Middleware.RequireAuth)
	{
		ProfileRoutes.GET("/me", Controllers.GetAllRoles)
	}

	// Role
	RoleRoutes := r.Group("/role")
	{
		RoleRoutes.GET("/role", Controllers.GetAllRoles)
		RoleRoutes.POST("/createdRole", Controllers.CreateRole)
		RoleRoutes.GET("/getRoleById", Controllers.GetRoleById)
		RoleRoutes.PUT("/updatedRole", Controllers.UpdateRole)
		RoleRoutes.DELETE("/deleteRole", Controllers.DeleteRole)
	}

	// Home
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "Welcome to !",
			"version":  "1.0.0",
			"Author":   "Rizky Budiarto",
			"Username": "KingRovs",
		})
	})
}
