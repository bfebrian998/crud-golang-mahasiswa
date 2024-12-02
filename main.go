package main

import (
	"log"
	"mahasiswa/config"
	"mahasiswa/controlllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading.env file")
	}

	config.InitDb()
	router := gin.Default()

	controlllers.SetRoutes(router)
	port := config.Routes()

	router.Run(":" + port)

}
