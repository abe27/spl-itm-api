package controllers

import (
	"fmt"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetMailBox(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	if c.Query("id") == "" {
		var MailBox []models.MailBox
		if err := configs.Store.Preload("Area").Where("is_active=?", true).Find(&MailBox).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(r)
		}

		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &MailBox
		return c.Status(r.StatusCode).JSON(&r)
	}

	var MailBox models.MailBox
	if err := configs.Store.Preload("Area").First(&MailBox, "id=?", c.Query("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("แสดงข้อมูล %s", c.Query("id"))
	r.Data = &MailBox
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateMailBox(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated

	var frm models.MailBox
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var area models.Area
	if err := configs.Store.Where("title=?", frm.AreaID).First(&area).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), frm.AreaID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	var MailBox models.MailBox
	MailBox.AreaID = area.ID
	MailBox.MailID = frm.MailID
	MailBox.Password = frm.Password
	MailBox.Url = frm.Url
	MailBox.IsActive = frm.IsActive

	if err := configs.Store.Create(&MailBox).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("บันทึกข้อมูล %s เรียบร้อยแล้ว", MailBox.MailID)
	MailBox.Area = area
	r.Data = &MailBox
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateMailBox(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	var frm models.MailBox
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var MailBox models.MailBox
	if err := configs.Store.First(&MailBox, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var area models.Area
	if err := configs.Store.Where("title=?", frm.AreaID).First(&area).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), frm.AreaID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	MailBox.AreaID = area.ID
	MailBox.MailID = frm.MailID
	MailBox.Password = frm.Password
	MailBox.Url = frm.Url
	MailBox.IsActive = frm.IsActive
	if err := configs.Store.Save(&MailBox).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}
	r.Message = fmt.Sprintf("อัพเดท %s เรียบร้อยแล้ว", c.Params("id"))
	MailBox.Area = area
	r.Data = &MailBox
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteMailBox(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var MailBox models.MailBox
	if err := configs.Store.First(&MailBox, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	if err := configs.Store.Delete(&MailBox).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("ลบข้อมูล %s เรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}
