package controllers

import (
	"fmt"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetLastInvoice(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	if c.Query("id") == "" {
		var LastInvoice []models.LastInvoice
		if err := configs.Store.Preload("Factory").Preload("Factory").Where("is_active =?", true).Find(&LastInvoice).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(&r)
		}
		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &LastInvoice
		return c.Status(r.StatusCode).JSON(&r)
	}

	var LastInvoice models.LastInvoice
	if err := configs.Store.Preload("Factory").Preload("Factory").Where("is_active=?", true).Where("id=?", c.Query("id")).First(&LastInvoice).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}
	r.Message = fmt.Sprintf("แสดง %s ข้อมูล", c.Query("id"))
	r.Data = &LastInvoice
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateLastInvoice(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated

	var frm models.LastInvoice
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	var fac models.Factory
	if err := configs.Store.First(&fac, "title", frm.FactoryID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s id %s!", err.Error(), *frm.FactoryID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	var affcode models.Affcode
	if err := configs.Store.First(&affcode, "title", frm.AffcodeID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s id %s!", err.Error(), *frm.AffcodeID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	var LastInvoice models.LastInvoice
	LastInvoice.FactoryID = &fac.ID
	LastInvoice.AffcodeID = &affcode.ID
	LastInvoice.OnYear = frm.OnYear
	LastInvoice.LastRunning = frm.LastRunning
	if err := configs.Store.Create(&LastInvoice).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}
	r.Message = fmt.Sprintf("บันทึก %s ข้อมูลเรียบร้อยแล้ว", LastInvoice.ID)
	LastInvoice.Factory = fac
	LastInvoice.Affcode = affcode
	r.Data = &LastInvoice
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateLastInvoice(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	var frm models.LastInvoice
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	var LastInvoice models.LastInvoice
	if err := configs.Store.Where("id=?", c.Params("id")).First(&LastInvoice).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s id %s!", err.Error(), c.Params("id"))
		return c.Status(r.StatusCode).JSON(&r)
	}

	var fac models.Factory
	if err := configs.Store.First(&fac, "title", frm.FactoryID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s id %s!", err.Error(), *frm.FactoryID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	var affcode models.Affcode
	if err := configs.Store.First(&affcode, "title", frm.AffcodeID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s id %s!", err.Error(), *frm.AffcodeID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	LastInvoice.FactoryID = &fac.ID
	LastInvoice.AffcodeID = &affcode.ID
	LastInvoice.OnYear = frm.OnYear
	LastInvoice.LastRunning = frm.LastRunning
	LastInvoice.IsActive = frm.IsActive
	if err := configs.Store.Save(&LastInvoice).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	r.Message = fmt.Sprintf("อัพเดท %s เรียบร้อยแล้ว", c.Params("id"))
	LastInvoice.Factory = fac
	LastInvoice.Affcode = affcode
	r.Data = &LastInvoice
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteLastInvoice(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	var LastInvoice models.LastInvoice
	if err := configs.Store.First(&LastInvoice, "id", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	if err := configs.Store.Delete(&LastInvoice).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	r.Message = fmt.Sprintf("ลบ %s ข้อมูลเรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}
