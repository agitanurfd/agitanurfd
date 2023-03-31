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
	insertedID := inimodule.InsertUndanganRapat(db, "undanganrapat",
	undanganrapat.Location,
	undanganrapat.Phone_number,
	undanganrapat.Biodata,
	undanganrapat.Prodi,
)
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

func GetAllPresensi(c *fiber.Ctx) error {
	ps := inimodullatihan.GetAllPresensi(config.Ulbimongoconn, "presensi")
	return c.JSON(ps)
}

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