package controllers

import (
	"fmt"
	"strings"
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
			Preload("ReceiveDetail.Ledger.Area").
			Preload("ReceiveDetail.Ledger.Factory").
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
	if err := configs.Store.
		Preload("DownloadMailBox.MailBox.Area").
		Preload("DownloadMailBox.MailType.Factory").
		Preload("ReceiveType.Whs").
		Preload("ReceiveDetail.Ledger.Area").
		Preload("ReceiveDetail.Ledger.Factory").
		Preload("ReceiveDetail.Ledger.Part").
		Preload("ReceiveDetail.Ledger.ItemType").
		Preload("ReceiveDetail.Ledger.Unit").
		Where("is_active=?", true).
		First(&data, "id", c.Query("id")).Error; err != nil {
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

	var frm models.Receive
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	var ediType models.DownloadMailBox
	if err := configs.Store.First(&ediType, "batch_no", frm.DownloadID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.DownloadID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	var receiveType models.ReceiveType
	if err := configs.Store.First(&receiveType, "prefix", frm.ReceiveTypeID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.ReceiveTypeID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	var receiveEnt models.Receive
	receiveEnt.DownloadID = &ediType.ID
	receiveEnt.ReceiveTypeID = &receiveType.ID
	receiveEnt.ReceiveDate = frm.ReceiveDate
	receiveEnt.TransferOutNo = strings.ToUpper(frm.TransferOutNo)
	receiveEnt.TexNo = strings.ToUpper(frm.TexNo)
	receiveEnt.Item = frm.Item
	receiveEnt.PlanCtn = frm.PlanCtn
	receiveEnt.ReceiveCtn = frm.ReceiveCtn
	receiveEnt.IsSync = frm.IsSync
	receiveEnt.IsActive = frm.IsActive
	if err := configs.Store.Create(&receiveEnt).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}
	receiveEnt.DownloadMailBox = ediType
	receiveEnt.ReceiveType = receiveType
	r.Data = &receiveEnt
	r.Message = fmt.Sprintf("บันทึกข้อมูล %s เรียบร้อยแล้ว", frm.TransferOutNo)
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateReceiveEnt(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	var frm models.Receive
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	var receiveEnt models.Receive
	if err := configs.Store.First(&receiveEnt, "id", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	var ediType models.DownloadMailBox
	if err := configs.Store.First(&ediType, "batch_no", frm.DownloadID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.DownloadID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	var receiveType models.ReceiveType
	if err := configs.Store.First(&receiveType, "prefix", frm.ReceiveTypeID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.ReceiveTypeID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	receiveEnt.DownloadID = &ediType.ID
	receiveEnt.ReceiveTypeID = &receiveType.ID
	receiveEnt.ReceiveDate = frm.ReceiveDate
	receiveEnt.TransferOutNo = strings.ToUpper(frm.TransferOutNo)
	receiveEnt.TexNo = strings.ToUpper(frm.TexNo)
	receiveEnt.Item = frm.Item
	receiveEnt.PlanCtn = frm.PlanCtn
	receiveEnt.ReceiveCtn = frm.ReceiveCtn
	receiveEnt.IsSync = frm.IsSync
	receiveEnt.IsActive = frm.IsActive
	if err := configs.Store.Save(&receiveEnt).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}
	receiveEnt.DownloadMailBox = ediType
	receiveEnt.ReceiveType = receiveType
	r.Data = &receiveEnt
	r.Message = fmt.Sprintf("อัพเดทข้อมูล %s เรียบร้อยแล้ว", frm.TransferOutNo)
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteReceiveEnt(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var receiveEnt models.Receive
	if err := configs.Store.First(&receiveEnt, "id", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	if err := configs.Store.Delete(&receiveEnt).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}
	r.Message = fmt.Sprintf("ลบข้อมูล %s เรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}
