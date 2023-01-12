package controllers

import (
	"fmt"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetFactory(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	if c.Query("id") == "" {
		var factory []models.Factory
		if err := configs.Store.Where("is_active=?", true).Find(&factory).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(&r)
		}

		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &factory
		return c.Status(r.StatusCode).JSON(&r)
	}

	var factory models.Factory
	if err := configs.Store.First(&factory, "id", c.Query("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	r.Message = fmt.Sprintf("แสดงข้อมูล %s", c.Query("id"))
	r.Data = &factory
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateFactory(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated

	var frm models.Factory
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	var factory models.Factory
	factory.Title = frm.Title
	factory.Description = frm.Description
	factory.LabelPrefix = frm.LabelPrefix
	factory.InvPrefix = frm.InvPrefix
	factory.Value = frm.Value
	factory.IsActive = frm.IsActive
	if err := configs.Store.Create(&factory).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	r.Message = fmt.Sprintf("บันทึกข้อมูล %s เรียบร้อยแล้ว", factory.Title)
	r.Data = &factory
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateFactory(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	var frm models.Factory
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	var factory models.Factory
	if err := configs.Store.First(&factory, "id", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	factory.Title = frm.Title
	factory.Description = frm.Description
	factory.LabelPrefix = frm.LabelPrefix
	factory.InvPrefix = frm.InvPrefix
	factory.Value = frm.Value
	factory.IsActive = frm.IsActive
	if err := configs.Store.Save(&factory).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	r.Message = fmt.Sprintf("อัพเดท %s เรียบร้อยแล้ว", c.Params("id"))
	r.Data = &factory
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteFactory(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	var factory models.Factory
	if err := configs.Store.First(&factory, "id", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	if err := configs.Store.Delete(&factory).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}
	r.Message = fmt.Sprintf("ลบข้อมูล %s เรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}
