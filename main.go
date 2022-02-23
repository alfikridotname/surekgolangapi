package main

import (
	"log"
	"surekapi/auth"
	"surekapi/handler"
	"surekapi/middleware"
	"surekapi/naskah"
	"surekapi/pejabat"
	"surekapi/unitkerja"
	"surekapi/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=surekdb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// User dan Login
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)
	// Kategori Penerima
	kategoriPenerimaHandler := handler.GetKategoriPenerima
	// Master Naskah
	masterNaskahRepository := naskah.NewRepository(db)
	masterNaskahService := naskah.NewService(masterNaskahRepository)
	masterNaskahHandler := handler.NewMasterNaskahHandler(masterNaskahService)
	// Pejabat
	pejabatRepository := pejabat.NewRepository(db)
	pejabatService := pejabat.NewService(pejabatRepository)
	pejabatHandler := handler.NewPejabatHandler(pejabatService)
	// Unit Kerja
	unitKerjaRepository := unitkerja.NewRepository(db)
	unitKerjaService := unitkerja.NewService(unitKerjaRepository)
	unitKerjaHandler := handler.NewUnitKerjaHandler(unitKerjaService)
	// Keamanan
	keamananHandler := handler.GetKeamanan
	// Kecepatan
	kecepatanHandler := handler.GetKecepatan
	// Pemeriksa Konsep Surat
	pemeriksaKonsepSuratHandler := handler.GetPemeriksa

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"information": "Surek API is running...",
		})
	})

	api := r.Group("/api/v1")
	api.POST("/login", userHandler.Login)
	api.GET("/naskah", middleware.AuthMiddleware(authService, userService), masterNaskahHandler.GetAll)
	api.GET("/pejabat", middleware.AuthMiddleware(authService, userService), pejabatHandler.GetByUnitKerjaID)
	api.GET("/kategori-penerima", middleware.AuthMiddleware(authService, userService), kategoriPenerimaHandler)
	api.GET("/unit-kerja", middleware.AuthMiddleware(authService, userService), unitKerjaHandler.GetAll)
	api.GET("/keamanan", middleware.AuthMiddleware(authService, userService), keamananHandler)
	api.GET("/kecepatan", middleware.AuthMiddleware(authService, userService), kecepatanHandler)
	api.GET("/pemeriksa", middleware.AuthMiddleware(authService, userService), pemeriksaKonsepSuratHandler)
	r.Run(":8080")
}
