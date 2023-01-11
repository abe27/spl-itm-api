package controllers

import (
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

type FrmAdmin struct {
	UserName string `form:"username"`
}

func CreateAdmin(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated
	var frm FrmAdmin
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	var admin models.Administrator
	admin.UserID = &frm.UserName
	admin.IsActive = true
	if err := configs.Store.Create(&admin).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	r.Message = "Create Admin Success!"
	r.Data = &admin
	return c.Status(r.StatusCode).JSON(&r)
}
