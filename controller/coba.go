package controller

import (
	"net/http"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	inimodel "github.com/agitanurfd/undanganRapat/model"
	inimodule "github.com/agitanurfd/undanganRapat/module"
	inimodullatihan "github.com/indrariksa/be_presensi/module"
	inimodul "github.com/indrariksa/be_presensi/model"
	// inimodultugas "github.com/agitanurfd/undanganRapat/module"
	"github.com/agitanurfd/agitanurfd/config"
	"github.com/aiteung/musik"
	cek "github.com/aiteung/presensi"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"github_repo": "https://github.com/agitanurfd/agitanurfd",
		"message":     "You are at the root endpoint 😉",
		"success":     true,
	})
}

func Homepage(c *fiber.Ctx) error {
ipaddr := musik.GetIPaddress()
return c.JSON(ipaddr)
}

func GetPresensi(c *fiber.Ctx) error {
	ps := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
	return c.JSON(ps)
}

func GetAllUndanganRapatFromNamaTamu(c *fiber.Ctx) error {
	ps := inimodule.GetUndanganRapatFromNamaTamu(config.Ulbimongoconn, "Jaemin", "undanganrapat")
	return c.JSON(ps)
}

func GetAllTamuFromJabatan(c *fiber.Ctx) error {
	ps := inimodule.GetTamuFromJabatan(config.Ulbimongoconn, "Dosen", "tamu")
	return c.JSON(ps)
}

func GetAllJamRapatFromDurasi (c *fiber.Ctx) error {
	ps := inimodule.GetJamRapatFromDurasi(config.Ulbimongoconn, "1 jam", "jamrapat")
	return c.JSON(ps)
}

func GetAllUniversitasFromJurusan (c *fiber.Ctx) error {
	ps := inimodule.GetUniversitasFromJurusan(config.Ulbimongoconn, "Teknik Informatika", "universitas")
	return c.JSON(ps)
}

func GetAllRuanganFromNoRuangan (c *fiber.Ctx) error {
	ps := inimodule.GetRuanganFromNoRuangan(config.Ulbimongoconn, "315", "ruangan")
	return c.JSON(ps)
}

func InsertUndanganRapat(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var undanganrapat inimodel.UndanganRapat
	if err := c.BodyParser(&undanganrapat); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodule.InsertUndanganRapat(db, "undanganrapat",
	undanganrapat.Location,
	undanganrapat.Phone_number,
	undanganrapat.Biodata,
	undanganrapat.Prodi,
	)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// func GetAllTamuFromNama(c *fiber.Ctx) error {
// 	ps := inimodule.GetTamuFromJabatan(config.Ulbimongoconn,"tamu", "Dosen")
// 	return c.JSON(ps)
// }

// GetAllPresensi godoc
// @Summary Get All Data Presensi.
// @Description Mengambil semua data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Success 200 {object} Presensi
// @Router /presensi [get]
func GetAllPresensi(c *fiber.Ctx) error {
	ps := inimodullatihan.GetAllPresensi(config.Ulbimongoconn, "presensi")
	return c.JSON(ps)
}

// GetPresensiID godoc
// @Summary Get By ID Data Presensi.
// @Description Ambil per ID data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /presensi/{id} [get]
func GetPresensiID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := inimodullatihan.GetPresensiFromID(objID, config.Ulbimongoconn, "presensi")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

func GetAll(c *fiber.Ctx) error {
	ps := inimodule.GetAllUndanganRapat(config.Ulbimongoconn, "undanganrapat")
	return c.JSON(ps)
}

func GetAllJamRapat(c *fiber.Ctx) error {
	ps := inimodule.GetAllJamRapat(config.Ulbimongoconn, "jamrapat")
	return c.JSON(ps)
}
func GetAllTamu(c *fiber.Ctx) error {
	ps := inimodule.GetAllTamu(config.Ulbimongoconn, "tamu")
	return c.JSON(ps)
}
func GetAllRuangan(c *fiber.Ctx) error {
	ps := inimodule.GetAllRuangan(config.Ulbimongoconn, "ruangan")
	return c.JSON(ps)
}
func GetAllLokasi(c *fiber.Ctx) error {
	ps := inimodule.GetAllLokasi(config.Ulbimongoconn, "lokasi")
	return c.JSON(ps)
}
func GetAllUniversitas(c *fiber.Ctx) error {
	ps := inimodule.GetAllUniversitas(config.Ulbimongoconn, "universitas")
	return c.JSON(ps)
}

// InsertData godoc
// @Summary Insert data presensi.
// @Description Input data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param request body Presensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /ins [post]
func InsertData(c *fiber.Ctx) error { // presensi
	db := config.Ulbimongoconn
	var presensi inimodul.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodullatihan.InsertPresensi(db, "presensi",
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertDataUndanganRapat(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var undanganrapat inimodel.UndanganRapat
	if err := c.BodyParser(&undanganrapat); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodule.InsertUndanganRapat(db, "undanganrapat",
		undanganrapat.Location,
		undanganrapat.Phone_number,
		undanganrapat.Biodata,
		undanganrapat.Prodi)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertDataTamu(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var tamu inimodel.Tamu
	if err := c.BodyParser(&tamu); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodule.InsertTamu(db, "tamu",
			tamu.Nama,
			tamu.Phone_number,
			tamu.Jabatan,
			tamu.Jam_kerja,
			tamu.Hari_kerja)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertDataJamRapat(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var jamrapat inimodel.JamRapat
	if err := c.BodyParser(&jamrapat); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodule.InsertJamRapat(db, "jamrapat",
			jamrapat.Durasi,
			jamrapat.Jam_rapat,
			jamrapat.Hari,
			jamrapat.Tanggal)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertDataUniversitas(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var universitas inimodel.Universitas
	if err := c.BodyParser(&universitas); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodule.InsertUniversitas(db, "universitas",
			universitas.Jurusan)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertDataLokasi(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var lokasi inimodel.Lokasi
	if err := c.BodyParser(&lokasi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodule.InsertLokasi(db, "lokasi",
			lokasi.Nama,
			lokasi.Alamat)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertDataRuangan(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var ruangan inimodel.Ruangan
	if err := c.BodyParser(&ruangan); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodule.InsertRuangan(db, "ruangan",
			ruangan.No_ruangan)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}
// UpdateData godoc
// @Summary Update data presensi.
// @Description Ubah data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Presensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /upd/{id} [put]
func UpdateData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var presensi inimodul.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = inimodullatihan.UpdatePresensi(db, "presensi",
		objectID,
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// DeletePresensiByID godoc
// @Summary Delete data presensi.
// @Description Hapus data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
func DeletePresensiByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = inimodullatihan.DeletePresensiByID(objID, config.Ulbimongoconn, "presensi")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}

// func GetAllUndanganRapatFromTamu(c *fiber.Ctx) error {
// 	ps := inimodule.GetUndanganRapatFromNamaTamu(config.Ulbimongoconn, "Jaemin", "undanganrapat")
// 	return c.JSON(ps)
// }

// func GetTamuFromJabatan(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	if id == "" {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": "Wrong parameter",
// 		})
// 	}
// 	// objID, err := primitive.ObjectIDFromHex(id)
// 	// if err != nil {
// 	// 	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 	// 		"status":  http.StatusBadRequest,
// 	// 		"message": "Invalid id parameter",
// 	// 	})
// 	// }
// 	ps, err := inimodule.GetTamuFromJabatan(config.Ulbimongoconn, "tamu")
// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return c.Status(http.StatusNotFound).JSON(fiber.Map{
// 				"status":  http.StatusNotFound,
// 				"message": fmt.Sprintf("No data found for id %s", id),
// 			})
// 		}
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": fmt.Sprintf("Error retrieving data for id %s", id),
// 		})
// 	}
// 	return c.JSON(ps)
// }

