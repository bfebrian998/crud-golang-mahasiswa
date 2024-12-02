package repository

import (
	"log"
	"mahasiswa/config"
	"mahasiswa/models/request"
	"mahasiswa/models/response"
)

func GetJadwalUniversitas(requestDto request.JadwalDto) ([]response.JadwalRespDto, error) {
	db := config.ConfigDatabases()

	var result []response.JadwalRespDto

	query := `
    SELECT
        c."name" AS mahasiswa,
        c.address AS alamat_mahasiswa,
        a."name" AS nama_dosen,
        a.alumni,
        b.matkul,
        b.jam_matkul,
        b.kelas
    FROM
        universitas.dosen a
    INNER JOIN universitas.mata_pelajaran b ON a.id = b.id_dosen
    INNER JOIN universitas.mahasiswa c ON b.id_dosen = c.id
    WHERE c."name" = $1;`

	rows, err := db.Query(query, requestDto.NameMahasiswa)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var resp response.JadwalRespDto
		if err := rows.Scan(
			&resp.Mahasiswa,
			&resp.AlamatMahasiswa,
			&resp.NamaDosen,
			&resp.Alumni,
			&resp.Matkul,
			&resp.JamMatkul,
			&resp.Kelas); err != nil {
			return nil, err
		}
		result = append(result, resp)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("failed to get jadwal", err)
		return nil, err

	}
	return result, nil
}
