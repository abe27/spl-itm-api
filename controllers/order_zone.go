package controllers

import (
	"fmt"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetOrderZone(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	if c.Query("id") == "" {
		var OrderZone []models.OrderZone
		if err := configs.Store.Preload("Whs").Preload("Factory").Where("is_active=?", true).Find(&OrderZone).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(r)
		}

		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &OrderZone
		return c.Status(r.StatusCode).JSON(&r)
	}

	var OrderZone models.OrderZone
	if err := configs.Store.Preload("Whs").Preload("Factory").First(&OrderZone, "id=?", c.Query("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("แสดงข้อมูล %s", c.Query("id"))
	r.Data = &OrderZone
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateOrderZone(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated

	var frm models.OrderZone
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var facData models.Factory
	if err := configs.Store.First(&facData, "title", frm.FactoryID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", *frm.FactoryID, err.Error())
		return c.Status(r.StatusCode).JSON(r)
	}

	var whsData models.Whs
	if err := configs.Store.First(&whsData, "title", frm.WhsID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", *frm.WhsID, err.Error())
		return c.Status(r.StatusCode).JSON(r)
	}

	var OrderZone models.OrderZone
	OrderZone.Value = frm.Value
	OrderZone.FactoryID = &facData.ID
	OrderZone.WhsID = &whsData.ID
	OrderZone.Description = frm.Description
	OrderZone.IsActive = frm.IsActive
	if err := configs.Store.Create(&OrderZone).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("บันทึกข้อมูล %s เรียบร้อยแล้ว", *frm.WhsID)
	OrderZone.Factory = facData
	OrderZone.Whs = whsData
	r.Data = &OrderZone
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateOrderZone(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	var frm models.OrderZone
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var OrderZone models.OrderZone
	if err := configs.Store.First(&OrderZone, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var facData models.Factory
	if err := configs.Store.First(&facData, "title", frm.FactoryID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", *frm.FactoryID, err.Error())
		return c.Status(r.StatusCode).JSON(r)
	}

	var whsData models.Whs
	if err := configs.Store.First(&whsData, "title", frm.WhsID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", *frm.WhsID, err.Error())
		return c.Status(r.StatusCode).JSON(r)
	}

	OrderZone.Value = frm.Value
	OrderZone.FactoryID = &facData.ID
	OrderZone.WhsID = &whsData.ID
	OrderZone.Description = frm.Description
	OrderZone.IsActive = frm.IsActive
	if err := configs.Store.Save(&OrderZone).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}
	r.Message = fmt.Sprintf("อัพเดท %s เรียบร้อยแล้ว", c.Params("id"))
	OrderZone.Factory = facData
	OrderZone.Whs = whsData
	r.Data = &OrderZone
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteOrderZone(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var OrderZone models.OrderZone
	if err := configs.Store.First(&OrderZone, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	if err := configs.Store.Delete(&OrderZone).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("ลบข้อมูล %s เรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}
