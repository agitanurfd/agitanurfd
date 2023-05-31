package main

import (
	"log"

	"github.com/agitanurfd/agitanurfd/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/agitanurfd/agitanurfd/url"

	"github.com/gofiber/fiber/v2"
	_ "github.com/agitanurfd/agitanurfd/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server

// @contact.name API Support
// @contact.url https://github.com/agitanurfd
// @contact.email 1214029@std.ulbi.ac.id

// @host agita.herokuapp.com
// @BasePath /
// @schemes http https

func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
