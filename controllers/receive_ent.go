package controllers

import (
	"time"

	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetReceiveEnt(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateReceiveEnt(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateReceiveEnt(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteReceiveEnt(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	return c.Status(r.StatusCode).JSON(&r)
}
