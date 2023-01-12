package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetMailType(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	if c.Query("id") == "" {
		var MailType []models.MailType
		if err := configs.Store.Where("is_active=?", true).Find(&MailType).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(r)
		}

		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &MailType
		return c.Status(r.StatusCode).JSON(&r)
	}

	var MailType models.MailType
	if err := configs.Store.First(&MailType, "id=?", c.Query("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("แสดงข้อมูล %s", c.Query("id"))
	r.Data = &MailType
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateMailType(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated

	var frm models.MailType
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var MailType models.MailType
	MailType.Prefix = strings.ToUpper(frm.Prefix)
	MailType.Title = strings.ToUpper(frm.Title)
	MailType.Description = frm.Description
	MailType.IsActive = frm.IsActive
	if err := configs.Store.Create(&MailType).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("บันทึกข้อมูล %s เรียบร้อยแล้ว", MailType.Title)
	r.Data = &MailType
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateMailType(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	var frm models.MailType
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var MailType models.MailType
	if err := configs.Store.First(&MailType, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	MailType.Prefix = strings.ToUpper(frm.Prefix)
	MailType.Title = strings.ToUpper(frm.Title)
	MailType.Description = frm.Description
	MailType.IsActive = frm.IsActive
	if err := configs.Store.Save(&MailType).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}
	r.Message = fmt.Sprintf("อัพเดท %s เรียบร้อยแล้ว", c.Params("id"))
	r.Data = &MailType
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteMailType(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var MailType models.MailType
	if err := configs.Store.First(&MailType, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	if err := configs.Store.Delete(&MailType).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("ลบข้อมูล %s เรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}
