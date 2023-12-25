package main

import (
	"backend/go-fiber/data"
	"backend/go-fiber/service"

	_ "backend/go-fiber/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

func main() {
	db := data.New()

	host := fiber.New()
	{
		host.Get("/docs/*", swagger.HandlerDefault)

		host.Use(cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
			AllowMethods: "GET,POST,PUT,DELETE",
		}))
	}

	service := service.New(db)
	{
		service.Registor(host)
	}

	host.Listen(":3000")
}
