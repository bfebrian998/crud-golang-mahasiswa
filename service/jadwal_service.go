package service

import (
	"mahasiswa/models/request"
	"mahasiswa/models/response"
	"mahasiswa/repository"
)

func GetJadwalByName(name request.JadwalDto) ([]response.JadwalRespDto, error) {
	return repository.GetJadwalUniversitas(name)
}
