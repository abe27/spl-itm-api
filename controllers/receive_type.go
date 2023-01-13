package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetReceiveType(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	if c.Query("id") == "" {
		var ReceiveType []models.ReceiveType
		if err := configs.Store.Preload("Whs").Where("is_active=?", true).Find(&ReceiveType).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(r)
		}

		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &ReceiveType
		return c.Status(r.StatusCode).JSON(&r)
	}

	var ReceiveType models.ReceiveType
	if err := configs.Store.Preload("Whs").First(&ReceiveType, "id=?", c.Query("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("แสดงข้อมูล %s", c.Query("id"))
	r.Data = &ReceiveType
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateReceiveType(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated

	var frm models.ReceiveType
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var whs models.Whs
	if err := configs.Store.First(&whs, "title", frm.WhsID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), frm.WhsID)
		return c.Status(r.StatusCode).JSON(r)
	}

	var ReceiveType models.ReceiveType
	ReceiveType.WhsID = whs.ID
	ReceiveType.Prefix = strings.ToUpper(frm.Prefix)
	ReceiveType.Title = strings.ToUpper(frm.Title)
	ReceiveType.Description = frm.Description
	ReceiveType.IsActive = frm.IsActive
	if err := configs.Store.Create(&ReceiveType).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("บันทึกข้อมูล %s เรียบร้อยแล้ว", ReceiveType.Title)
	ReceiveType.Whs = whs
	r.Data = &ReceiveType
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateReceiveType(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	var frm models.ReceiveType
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var ReceiveType models.ReceiveType
	if err := configs.Store.First(&ReceiveType, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var whs models.Whs
	if err := configs.Store.First(&whs, "title", frm.WhsID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), frm.WhsID)
		return c.Status(r.StatusCode).JSON(r)
	}

	ReceiveType.WhsID = whs.ID
	ReceiveType.Prefix = strings.ToUpper(frm.Prefix)
	ReceiveType.Title = strings.ToUpper(frm.Title)
	ReceiveType.Description = frm.Description
	ReceiveType.IsActive = frm.IsActive
	if err := configs.Store.Save(&ReceiveType).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}
	r.Message = fmt.Sprintf("อัพเดท %s เรียบร้อยแล้ว", c.Params("id"))
	ReceiveType.Whs = whs
	r.Data = &ReceiveType
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteReceiveType(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var ReceiveType models.ReceiveType
	if err := configs.Store.First(&ReceiveType, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	if err := configs.Store.Delete(&ReceiveType).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("ลบข้อมูล %s เรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}
