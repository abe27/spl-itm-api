package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetOrderGroupType(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	if c.Query("id") == "" {
		var OrderGroupType []models.OrderGroupType
		if err := configs.Store.Where("is_active=?", true).Find(&OrderGroupType).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(r)
		}

		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &OrderGroupType
		return c.Status(r.StatusCode).JSON(&r)
	}

	var OrderGroupType models.OrderGroupType
	if err := configs.Store.First(&OrderGroupType, "id=?", c.Query("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("แสดงข้อมูล %s", c.Query("id"))
	r.Data = &OrderGroupType
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateOrderGroupType(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated

	var frm models.OrderGroupType
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var OrderGroupType models.OrderGroupType
	OrderGroupType.Title = strings.ToUpper(frm.Title)
	OrderGroupType.Description = frm.Description
	OrderGroupType.IsActive = frm.IsActive
	if err := configs.Store.Create(&OrderGroupType).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("บันทึกข้อมูล %s เรียบร้อยแล้ว", OrderGroupType.Title)
	r.Data = &OrderGroupType
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateOrderGroupType(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	var frm models.OrderGroupType
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var OrderGroupType models.OrderGroupType
	if err := configs.Store.First(&OrderGroupType, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	OrderGroupType.Title = strings.ToUpper(frm.Title)
	OrderGroupType.Description = frm.Description
	OrderGroupType.IsActive = frm.IsActive
	if err := configs.Store.Save(&OrderGroupType).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}
	r.Message = fmt.Sprintf("อัพเดท %s เรียบร้อยแล้ว", c.Params("id"))
	r.Data = &OrderGroupType
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteOrderGroupType(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var OrderGroupType models.OrderGroupType
	if err := configs.Store.First(&OrderGroupType, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	if err := configs.Store.Delete(&OrderGroupType).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("ลบข้อมูล %s เรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}
