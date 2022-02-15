package controller

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/igeargeek/igg-golang-api-response/response"
	"github.com/otiai10/gosseract/v2"
)

var TEMPFILE_OCR_DIRECTORY = "./ocr-image"

type ocrController interface {
	ProcessImage(c *fiber.Ctx) error
}

type OCRController struct{}

func NewOCR() ocrController {
	return &OCRController{}
}

func (ocr *OCRController) ProcessImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		status, resData := response.BadRequest("")
		return c.Status(status).JSON(resData)
	}
	tNow := time.Now()
	directory := fmt.Sprintf("./%s/%d-%s", TEMPFILE_OCR_DIRECTORY, tNow.Unix(), file.Filename)

	ch := make(chan string)
	client := gosseract.NewClient()

	go func() {
		c.SaveFile(file, directory)
		client.SetImage(directory)
		text, err := client.Text()
		if err != nil {
			ch <- "Error"
		} else {
			ch <- text
		}
	}()

	go func() {
		<-ch
		os.Remove(directory)
		client.Close()
	}()

	text := <-ch
	ch <- "remove"

	status, resData := response.Item(text, "")
	return c.Status(status).JSON(resData)
}
