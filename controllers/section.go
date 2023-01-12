package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetSection(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	if c.Query("id") == "" {
		var Section []models.Section
		if err := configs.Store.Where("is_active=?", true).Find(&Section).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(r)
		}

		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &Section
		return c.Status(r.StatusCode).JSON(&r)
	}

	var Section models.Section
	if err := configs.Store.First(&Section, "id=?", c.Query("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("แสดงข้อมูล %s", c.Query("id"))
	r.Data = &Section
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateSection(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated

	var frm models.Section
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var Section models.Section
	Section.Title = strings.ToUpper(frm.Title)
	Section.Description = frm.Description
	Section.IsActive = frm.IsActive
	if err := configs.Store.Create(&Section).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("บันทึกข้อมูล %s เรียบร้อยแล้ว", Section.Title)
	r.Data = &Section
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateSection(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	var frm models.Section
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var Section models.Section
	if err := configs.Store.First(&Section, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	Section.Title = strings.ToUpper(frm.Title)
	Section.Description = frm.Description
	Section.IsActive = frm.IsActive
	if err := configs.Store.Save(&Section).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}
	r.Message = fmt.Sprintf("อัพเดท %s เรียบร้อยแล้ว", c.Params("id"))
	r.Data = &Section
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteSection(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var Section models.Section
	if err := configs.Store.First(&Section, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	if err := configs.Store.Delete(&Section).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("ลบข้อมูล %s เรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}
