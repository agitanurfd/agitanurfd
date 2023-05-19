package url

import (
	"github.com/agitanurfd/agitanurfd/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Homepage) //ujicoba panggil package musik
	// page.Get("/presensi", controller.GetPresensi)
	page.Get("/tamu", controller.GetAllUndanganRapatFromNamaTamu)
	page.Get("/jabatan", controller.GetAllTamuFromJabatan)
	page.Get("/jamrapat", controller.GetAllJamRapatFromDurasi)
	page.Get("/jurusan", controller.GetAllUniversitasFromJurusan)
	page.Get("/ruangan", controller.GetAllRuanganFromNoRuangan)
	page.Get("/insundangan", controller.InsertUndanganRapat)
	page.Get("/presensi", controller.GetAllPresensi) //menampilkan seluruh data presensi
	page.Get("/presensi/:id", controller.GetPresensiID) //menampilkan data presensi berdasarkan id
	page.Get("/all", controller.GetAll) //menampilkan seluruh data undangan rapat
	page.Get("/all-jamrapat", controller.GetAllJamRapat)
	page.Get("/all-tamu", controller.GetAllTamu)
	page.Get("/all-universitas", controller.GetAllUniversitas)
	page.Get("/all-ruangan", controller.GetAllRuangan)
	page.Get("/all-lokasi", controller.GetAllLokasi)
	page.Post("/ins", controller.InsertData)
	page.Put("/upd/:id", controller.UpdateData)
	page.Delete("/delete/:id", controller.DeletePresensiByID)
	// page.Get("/undanganrapat/:id", controller.GetUndanganRapatFromID) //menampilkan data presensi berdasarkan id
}
