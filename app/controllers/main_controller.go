package controllers

import (
	"bufio"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func RenderHello(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	})
}

func RenderName(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"info":  "Rajesh Patil",
	})
}

func UploadFile(c *fiber.Ctx) error {
	// Get first file from form field "document":
	file, err := c.FormFile("document")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get Buffer from file
	buffer, err := file.Open()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	defer buffer.Close()

	scanner := bufio.NewScanner(buffer)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"info":  file.Filename,
	})
}
