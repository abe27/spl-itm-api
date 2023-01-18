package controllers

import (
	"fmt"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetLedgerCarton(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	if c.Query("id") == "" {
		var carton []models.Carton
		if err := configs.Store.
			Preload("Ledger.Whs").
			Preload("Ledger.Area").
			Preload("Ledger.Factory").
			Preload("Ledger.Part").
			Preload("Ledger.ItemType").
			Preload("Ledger.Unit").
			Preload("ReceiveDetail.Receive.DownloadMailBox.MailBox.Area").
			Preload("ReceiveDetail.Receive.DownloadMailBox.MailType.Factory").
			Preload("ReceiveDetail.Receive.ReceiveType").
			Preload("Shelve").
			Find(&carton).Error; err != nil {
			r.StatusCode = fiber.StatusInternalServerError
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(r)
		}

		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &carton
		return c.Status(r.StatusCode).JSON(&r)
	}

	if c.Query("ledger_id") != "" {
		var carton []models.Carton
		if err := configs.Store.
			Preload("Ledger.Whs").
			Preload("Ledger.Area").
			Preload("Ledger.Factory").
			Preload("Ledger.Part").
			Preload("Ledger.ItemType").
			Preload("Ledger.Unit").
			Preload("ReceiveDetail.Receive.DownloadMailBox.MailBox.Area").
			Preload("ReceiveDetail.Receive.DownloadMailBox.MailType.Factory").
			Preload("ReceiveDetail.Receive.ReceiveType").
			Preload("Shelve").
			Find(&carton, "ledger_id", c.Query("ledger_id")).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(r)
		}

		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &carton
		return c.Status(r.StatusCode).JSON(&r)
	}

	var carton models.Carton
	if err := configs.Store.
		Preload("Ledger.Whs").
		Preload("Ledger.Area").
		Preload("Ledger.Factory").
		Preload("Ledger.Part").
		Preload("Ledger.ItemType").
		Preload("Ledger.Unit").
		Preload("ReceiveDetail.Receive.DownloadMailBox.MailBox.Area").
		Preload("ReceiveDetail.Receive.DownloadMailBox.MailType.Factory").
		Preload("ReceiveDetail.Receive.ReceiveType").
		Preload("Shelve").
		First(&carton, "id", c.Query("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("แสดง %s ข้อมูล", c.Query("id"))
	r.Data = &carton
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateLedgerCarton(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateLedgerCarton(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteLedgerCarton(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	return c.Status(r.StatusCode).JSON(&r)
}
