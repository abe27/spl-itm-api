package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetItemType(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	if c.Query("id") == "" {
		var ItemType []models.ItemType
		if err := configs.Store.Where("is_active=?", true).Find(&ItemType).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(r)
		}

		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &ItemType
		return c.Status(r.StatusCode).JSON(&r)
	}

	var ItemType models.ItemType
	if err := configs.Store.First(&ItemType, "id=?", c.Query("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("แสดงข้อมูล %s", c.Query("id"))
	r.Data = &ItemType
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateItemType(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated

	var frm models.ItemType
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var ItemType models.ItemType
	ItemType.Title = strings.ToUpper(frm.Title)
	ItemType.Description = frm.Description
	ItemType.IsActive = frm.IsActive
	if err := configs.Store.Create(&ItemType).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("บันทึกข้อมูล %s เรียบร้อยแล้ว", ItemType.Title)
	r.Data = &ItemType
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateItemType(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	var frm models.ItemType
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var ItemType models.ItemType
	if err := configs.Store.First(&ItemType, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	ItemType.Title = strings.ToUpper(frm.Title)
	ItemType.Description = frm.Description
	ItemType.IsActive = frm.IsActive
	if err := configs.Store.Save(&ItemType).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}
	r.Message = fmt.Sprintf("อัพเดท %s เรียบร้อยแล้ว", c.Params("id"))
	r.Data = &ItemType
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteItemType(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var ItemType models.ItemType
	if err := configs.Store.First(&ItemType, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	if err := configs.Store.Delete(&ItemType).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("ลบข้อมูล %s เรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}
