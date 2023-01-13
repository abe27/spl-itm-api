package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetReviseOrder(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	if c.Query("id") == "" {
		var ReviseOrder []models.ReviseOrder
		if err := configs.Store.Where("is_active=?", true).Find(&ReviseOrder).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(r)
		}

		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &ReviseOrder
		return c.Status(r.StatusCode).JSON(&r)
	}

	var ReviseOrder models.ReviseOrder
	if err := configs.Store.First(&ReviseOrder, "id=?", c.Query("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("แสดงข้อมูล %s", c.Query("id"))
	r.Data = &ReviseOrder
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateReviseOrder(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated

	var frm models.ReviseOrder
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var ReviseOrder models.ReviseOrder
	ReviseOrder.Title = strings.ToUpper(frm.Title)
	ReviseOrder.Description = frm.Description
	ReviseOrder.IsActive = frm.IsActive
	if err := configs.Store.Create(&ReviseOrder).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("บันทึกข้อมูล %s เรียบร้อยแล้ว", ReviseOrder.Title)
	r.Data = &ReviseOrder
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateReviseOrder(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	var frm models.ReviseOrder
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var ReviseOrder models.ReviseOrder
	if err := configs.Store.First(&ReviseOrder, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	ReviseOrder.Title = strings.ToUpper(frm.Title)
	ReviseOrder.Description = frm.Description
	ReviseOrder.IsActive = frm.IsActive
	if err := configs.Store.Save(&ReviseOrder).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}
	r.Message = fmt.Sprintf("อัพเดท %s เรียบร้อยแล้ว", c.Params("id"))
	r.Data = &ReviseOrder
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteReviseOrder(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var ReviseOrder models.ReviseOrder
	if err := configs.Store.First(&ReviseOrder, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	if err := configs.Store.Delete(&ReviseOrder).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("ลบข้อมูล %s เรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}
