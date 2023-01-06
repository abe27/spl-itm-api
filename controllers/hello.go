package controllers

import (
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	var r models.Response

	r.StatusCode = fiber.StatusOK
	r.Success = true
	r.Message = "Hello World!"
	return c.Status(r.StatusCode).JSON(&r)
}
