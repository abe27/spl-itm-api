package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetWhs(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	if c.Query("id") == "" {
		var whs []models.Whs
		if err := configs.Store.Where("is_active =?", true).Find(&whs).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(&r)
		}
		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &whs
		return c.Status(r.StatusCode).JSON(&r)
	}

	var whs models.Whs
	if err := configs.Store.Where("is_active=?", true).Where("id=?", c.Query("id")).First(&whs).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}
	r.Message = fmt.Sprintf("แสดง %s ข้อมูล", c.Query("id"))
	r.Data = &whs
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateWhs(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated

	var frm models.Whs
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	var whs models.Whs
	whs.Prefix = strings.ToUpper(frm.Prefix)
	whs.Title = strings.ToUpper(frm.Title)
	whs.Value = frm.Value
	whs.Description = frm.Description
	whs.IsActive = frm.IsActive
	if err := configs.Store.Create(&whs).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}
	r.Message = fmt.Sprintf("บันทึก %s ข้อมูลเรียบร้อยแล้ว", whs.ID)
	r.Data = &whs
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateWhs(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	var frm models.Whs
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	var whs models.Whs
	if err := configs.Store.Where("id=?", c.Params("id")).First(&whs).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s id %s!", err.Error(), c.Params("id"))
		return c.Status(r.StatusCode).JSON(&r)
	}

	whs.Prefix = strings.ToUpper(frm.Prefix)
	whs.Title = strings.ToUpper(frm.Title)
	whs.Value = frm.Value
	whs.Description = frm.Description
	whs.IsActive = frm.IsActive
	if err := configs.Store.Save(&whs).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	r.Message = fmt.Sprintf("อัพเดท %s เรียบร้อยแล้ว", c.Params("id"))
	r.Data = &whs
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteWhs(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	var whs models.Whs
	if err := configs.Store.First(&whs, "id", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	if err := configs.Store.Delete(&whs).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	r.Message = fmt.Sprintf("ลบ %s ข้อมูลเรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}
