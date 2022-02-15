package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/noppadollanngain/tesseract-ocr-golang/controller"
)

func main() {
	app := fiber.New()

	ocr := controller.NewOCR()

	app.Post("/processimage", ocr.ProcessImage)

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("server is running")
	})

	app.Listen(":3000")
}
