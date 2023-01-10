package controllers

import (
	"fmt"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetSystemLogger(c *fiber.Ctx) error {
	var r models.Response
	r.StatusCode = fiber.StatusOK
	var obj []models.SystemLogger
	if err := configs.Store.Order("updated_at desc").Find(&obj).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}
	r.Message = "Show All System Logs."
	r.Data = &obj
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateSystemLogger(c *fiber.Ctx) error {
	var r models.Response
	r.StatusCode = fiber.StatusCreated
	var obj models.SystemLogger
	if err := c.BodyParser(&obj); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	if err := configs.Store.Create(&obj).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}
	r.Message = fmt.Sprintf("Create System Logger %s", obj.Title)
	r.Data = &obj
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateSystemLogger(c *fiber.Ctx) error {
	var r models.Response
	r.StatusCode = fiber.StatusOK
	var frm models.SystemLogger
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	var obj models.SystemLogger
	if err := configs.Store.First(&obj, "id", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	obj.Title = frm.Title
	obj.Description = frm.Description
	obj.IsSuccess = frm.IsSuccess
	if err := configs.Store.Save(&obj).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	r.Message = fmt.Sprintf("Update System Logger %s", c.Params("id"))
	r.Data = &obj
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteSystemLogger(c *fiber.Ctx) error {
	var r models.Response
	r.StatusCode = fiber.StatusOK
	var obj models.SystemLogger
	if err := configs.Store.First(&obj, "id", c.Params("id")).Error; err != nil {
		r.Message = err.Error()
		r.StatusCode = fiber.StatusNotFound
		return c.Status(r.StatusCode).JSON(&r)
	}

	if err := configs.Store.Delete(&obj).Error; err != nil {
		r.Message = err.Error()
		r.StatusCode = fiber.StatusInternalServerError
		return c.Status(r.StatusCode).JSON(&r)
	}

	r.Message = "Delete System Logger " + obj.Title
	r.StatusCode = fiber.StatusOK
	return c.Status(r.StatusCode).JSON(&r)
}
