package main

// mengimport gin
import (
	"gin-quickstart/config"
	"gin-quickstart/models"
	"gin-quickstart/routers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// membuat sebuah function
func main() {
	// cek file env apakah ada
	if err := godotenv.Load(); err != nil {
		log.Println("File .env tidak di temukan")
	}
	// inisialisasi gin
	r := gin.Default()
	//   membuat router dengan method get
	config.ConnectDB()
	config.ConnectRedis()
	config.DB.AutoMigrate(&models.Anime{}, &models.User{})
	routers.SetupRouters(r)
	//   menjalankan servernya
	r.Run(":3000") // listens on 0.0.0.0:8080 by default
}
