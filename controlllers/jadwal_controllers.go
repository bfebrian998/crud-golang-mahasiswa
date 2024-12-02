package controlllers

import (
	"mahasiswa/models/request"
	"mahasiswa/service"

	"github.com/gin-gonic/gin"
)

func GetJadwalByName(e *gin.Context) {
	var req request.JadwalDto
	if err := e.ShouldBindJSON(&req); err != nil {
		e.JSON(400, gin.H{"error": err.Error()})
		return
	}

	jadwalService, err := service.GetJadwalByName(req)

	if err != nil {
		e.JSON(500, gin.H{"error": err.Error()})
		return
	}

	e.JSON(200, jadwalService)
}

func SetRoutes(router *gin.Engine) {
	router.GET("/jadwal-universitas/getAll", GetJadwalByName)
}
