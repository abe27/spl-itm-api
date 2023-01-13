package controllers

import (
	"fmt"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetReceiveEnt(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	if c.Query("id") == "" {
		var data []models.Receive
		if err := configs.Store.
			Preload("DownloadMailBox.MailBox.Area").
			Preload("DownloadMailBox.MailType.Factory").
			Preload("ReceiveType.Whs").
			Preload("ReceiveDetail.Ledger.Part").
			Preload("ReceiveDetail.Ledger.ItemType").
			Preload("ReceiveDetail.Ledger.Unit").
			Find(&data).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(&r)
		}
		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &data
		return c.Status(r.StatusCode).JSON(&r)
	}

	var data models.Receive
	if err := configs.Store.Order("updated_at desc").Preload("MailBox.Area").Preload("MailType.Factory").Where("is_active=?", true).First(&data, "id", c.Query("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	r.Message = fmt.Sprintf("แสดงข้อมูล %s", c.Query("id"))
	r.Data = &data
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateReceiveEnt(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateReceiveEnt(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteReceiveEnt(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	return c.Status(r.StatusCode).JSON(&r)
}
