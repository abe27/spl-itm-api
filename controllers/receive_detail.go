package controllers

import (
	"fmt"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetReceiveDetail(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	if c.Query("id") == "" {
		if c.Query("recID") == "" {
			r.StatusCode = fiber.StatusMethodNotAllowed
			r.Message = "กรุณาระบุเลขที่เอกสารด้วย"
			return c.Status(r.StatusCode).JSON(&r)
		}

		var receiveDetail []models.ReceiveDetail
		if err := configs.Store.
			Preload("Receive.DownloadMailBox.MailBox.Area").
			Preload("Receive.DownloadMailBox.MailType.Factory").
			Preload("Receive.ReceiveType.Whs").
			Preload("Ledger.Area").
			Preload("Ledger.Factory").
			Preload("Ledger.Part").
			Preload("Ledger.ItemType").
			Preload("Ledger.Unit").
			Find(&receiveDetail, "receive_id", c.Query("recID")).
			Error; err != nil {
			r.StatusCode = fiber.StatusInternalServerError
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(&r)
		}
		r.Message = fmt.Sprintf("แสดงข้อมูล %s", c.Query("recID"))
		r.Data = &receiveDetail
		return c.Status(r.StatusCode).JSON(&r)
	}

	var receiveDetail models.ReceiveDetail
	if err := configs.Store.
		Preload("Receive.DownloadMailBox.MailBox.Area").
		Preload("Receive.DownloadMailBox.MailType.Factory").
		Preload("Receive.ReceiveType.Whs").
		Preload("Ledger.Area").
		Preload("Ledger.Factory").
		Preload("Ledger.Part").
		Preload("Ledger.ItemType").
		Preload("Ledger.Unit").
		Find(&receiveDetail, "id", c.Query("id")).
		Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	r.Message = fmt.Sprintf("แสดงข้อมูล %s", c.Query("id"))
	r.Data = &receiveDetail
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateReceiveDetail(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated
	var frm models.ReceiveDetail
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	var receiveEnt models.Receive
	if err := configs.Store.
		Preload("DownloadMailBox.MailBox.Area").
		Preload("DownloadMailBox.MailType.Factory").
		Preload("ReceiveType.Whs").
		First(&receiveEnt, "transfer_out_no", frm.ReceiveID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.ReceiveID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	var part models.Part
	if err := configs.Store.First(&part, "title", frm.LedgerID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.LedgerID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	var ledger models.Ledger
	if err := configs.Store.
		Where("area_id=?", receiveEnt.DownloadMailBox.MailBox.Area.ID).
		Where("factory_id=?", receiveEnt.DownloadMailBox.MailType.Factory.ID).
		Where("part_id=?", part.ID).
		First(&ledger).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.LedgerID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	var receiveBody models.ReceiveDetail
	if err := configs.Store.
		Where("receive_id=?", receiveEnt.ID).
		Where("ledger_id=?", ledger.ID).
		First(&receiveBody).Error; err == nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = fmt.Sprintf("บันทึกข้อมูลซ้ำ %s", receiveBody.ID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	receiveBody.ReceiveID = &receiveEnt.ID
	receiveBody.LedgerID = &ledger.ID
	receiveBody.PlanQty = frm.PlanQty
	receiveBody.PlanCtn = frm.PlanCtn
	receiveBody.IsActive = frm.IsActive
	if err := configs.Store.Create(&receiveBody).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}
	receiveBody.Receive = receiveEnt
	receiveBody.Ledger = ledger
	r.Data = &receiveBody
	r.Message = fmt.Sprintf("บันทึกข้อมูล %s เรียบร้อยแล้ว", receiveBody.ID)
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateReceiveDetail(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var frm models.ReceiveDetail
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	var receiveBody models.ReceiveDetail
	if err := configs.Store.First(&receiveBody, "id", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	var receiveEnt models.Receive
	if err := configs.Store.
		Preload("DownloadMailBox.MailBox.Area").
		Preload("DownloadMailBox.MailType.Factory").
		Preload("ReceiveType.Whs").
		First(&receiveEnt, "transfer_out_no", frm.ReceiveID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.ReceiveID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	var part models.Part
	if err := configs.Store.First(&part, "title", frm.LedgerID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.LedgerID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	var ledger models.Ledger
	if err := configs.Store.
		Where("area_id=?", receiveEnt.DownloadMailBox.MailBox.Area.ID).
		Where("factory_id=?", receiveEnt.DownloadMailBox.MailType.Factory.ID).
		Where("part_id=?", part.ID).
		First(&ledger).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.LedgerID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	receiveBody.ReceiveID = &receiveEnt.ID
	receiveBody.LedgerID = &ledger.ID
	receiveBody.PlanQty = frm.PlanQty
	receiveBody.PlanCtn = frm.PlanCtn
	receiveBody.IsActive = frm.IsActive
	if err := configs.Store.Save(&receiveBody).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}
	receiveBody.Receive = receiveEnt
	receiveBody.Ledger = ledger
	r.Data = &receiveBody
	r.Message = fmt.Sprintf("อัพเดทข้อมูล %s เรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteReceiveDetail(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var receiveBody models.ReceiveDetail
	if err := configs.Store.First(&receiveBody, "id", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	if err := configs.Store.Delete(&receiveBody).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}
	r.Message = fmt.Sprintf("ลบข้อมูล %s เรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}
