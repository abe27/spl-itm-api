package controllers

import (
	"time"

	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetDownloadMailBox(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateDownloadMailBox(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateDownloadMailBox(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteDownloadMailBox(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	return c.Status(r.StatusCode).JSON(&r)
}
