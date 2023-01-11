package controllers

import (
	"time"

	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetWhs(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateWhs(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateWhs(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteWhs(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	return c.Status(r.StatusCode).JSON(&r)
}
