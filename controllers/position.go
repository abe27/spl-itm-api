package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetPosition(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	if c.Query("id") == "" {
		var Position []models.Position
		if err := configs.Store.Where("is_active=?", true).Find(&Position).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(r)
		}

		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &Position
		return c.Status(r.StatusCode).JSON(&r)
	}

	var Position models.Position
	if err := configs.Store.First(&Position, "id=?", c.Query("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("แสดงข้อมูล %s", c.Query("id"))
	r.Data = &Position
	return c.Status(r.StatusCode).JSON(&r)
}

func CreatePosition(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated

	var frm models.Position
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var Position models.Position
	Position.Title = strings.ToUpper(frm.Title)
	Position.Description = frm.Description
	Position.IsActive = frm.IsActive
	if err := configs.Store.Create(&Position).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("บันทึกข้อมูล %s เรียบร้อยแล้ว", Position.Title)
	r.Data = &Position
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdatePosition(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	var frm models.Position
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var Position models.Position
	if err := configs.Store.First(&Position, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	Position.Title = strings.ToUpper(frm.Title)
	Position.Description = frm.Description
	Position.IsActive = frm.IsActive
	if err := configs.Store.Save(&Position).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}
	r.Message = fmt.Sprintf("อัพเดท %s เรียบร้อยแล้ว", c.Params("id"))
	r.Data = &Position
	return c.Status(r.StatusCode).JSON(&r)
}

func DeletePosition(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var Position models.Position
	if err := configs.Store.First(&Position, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	if err := configs.Store.Delete(&Position).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("ลบข้อมูล %s เรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}
