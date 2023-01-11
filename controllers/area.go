package controllers

import (
	"fmt"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetArea(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	if c.Query("id") == "" {
		var data []models.Area
		if err := configs.Store.Order("created_at").Find(&data).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(&r)
		}

		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &data
		return c.Status(r.StatusCode).JSON(&r)
	}

	var data models.Area
	if err := configs.Store.First(&data, c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	r.Message = fmt.Sprintf("แสดงข้อมูล %s", c.Query("id"))
	r.Data = &data
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateArea(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated
	// Body Parse
	var frm models.Area
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusNotAcceptable
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	var area models.Area
	area.Title = frm.Title
	area.Description = frm.Description
	area.IsActive = frm.IsActive
	if err := configs.Store.Create(&area).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	r.Message = fmt.Sprintf("บันทึก %s เรียบร้อยแล้ว", frm.Title)
	r.Data = &area
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateArea(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	// Body Parse
	var frm models.Area
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusNotAcceptable
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	var area models.Area
	if err := configs.Store.First(&area, "id", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	area.Title = frm.Title
	area.Description = frm.Description
	area.IsActive = frm.IsActive
	if err := configs.Store.Save(&area).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}
	r.Message = fmt.Sprintf("อัพเดท %s เรียบร้อยแล้ว", frm.Title)
	r.Data = &area
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteArea(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	var area models.Area
	if err := configs.Store.First(&area, "id", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	if err := configs.Store.Delete(&area).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	r.Message = fmt.Sprintf("ลบ %s เรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}
