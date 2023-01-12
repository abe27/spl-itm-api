package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetUnit(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	if c.Query("id") == "" {
		var unit []models.Unit
		if err := configs.Store.Where("is_active=?", true).Find(&unit).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(r)
		}

		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &unit
		return c.Status(r.StatusCode).JSON(&r)
	}

	var unit models.Unit
	if err := configs.Store.First(&unit, "id=?", c.Query("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("แสดงข้อมูล %s", c.Query("id"))
	r.Data = &unit
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateUnit(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated

	var frm models.Unit
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var unit models.Unit
	unit.Title = strings.ToUpper(frm.Title)
	unit.Description = frm.Description
	unit.IsActive = frm.IsActive
	if err := configs.Store.Create(&unit).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("บันทึกข้อมูล %s เรียบร้อยแล้ว", unit.Title)
	r.Data = &unit
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateUnit(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	var frm models.Unit
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var unit models.Unit
	if err := configs.Store.First(&unit, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	unit.Title = strings.ToUpper(frm.Title)
	unit.Description = frm.Description
	unit.IsActive = frm.IsActive
	if err := configs.Store.Save(&unit).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}
	r.Message = fmt.Sprintf("อัพเดท %s เรียบร้อยแล้ว", c.Params("id"))
	r.Data = &unit
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteUnit(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var unit models.Unit
	if err := configs.Store.First(&unit, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	if err := configs.Store.Delete(&unit).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("ลบข้อมูล %s เรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}
