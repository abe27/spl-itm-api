package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/abe/erp.api/services"
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
	var frm models.FrmCarton
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var whs models.Whs
	if err := configs.Store.First(&whs, "title", frm.WhsID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), frm.WhsID)
		return c.Status(r.StatusCode).JSON(r)
	}

	var part models.Part
	if err := configs.Store.First(&part, "title", frm.LedgerID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), frm.LedgerID)
		return c.Status(r.StatusCode).JSON(r)
	}

	var ledger models.Ledger
	if err := configs.Store.
		Where("whs_id=?", whs.ID).
		Where("part_id=?", part.ID).
		First(&ledger).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), frm.LedgerID)
		return c.Status(r.StatusCode).JSON(r)
	}

	var shelve models.Shelve
	if err := configs.Store.First(&shelve, "title", frm.ShelveID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.ShelveID)
		return c.Status(r.StatusCode).JSON(r)
	}

	var carton models.Carton
	carton.LedgerID = &ledger.ID
	carton.SerialNo = frm.SerialNo
	carton.LotNo = frm.LotNo
	carton.LineNo = frm.LineNo
	carton.ReviseNo = frm.ReviseNo
	carton.PalletNo = frm.PalletNo
	carton.ShelveID = &shelve.ID
	carton.Qty = frm.Qty
	carton.IsActive = frm.IsActive
	if err := configs.Store.Create(&carton).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	ledger.Ctn += 1
	if err := configs.Store.Save(&ledger).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = fmt.Sprintf("Update Ledger %s", err.Error())
		return c.Status(r.StatusCode).JSON(r)
	}

	// Create Carton History
	var cartonHistory models.CartonHistory
	cartonHistory.CartonID = carton.SerialNo
	cartonHistory.PalletNo = frm.PalletNo
	cartonHistory.ShelveID = shelve.Title
	cartonHistory.Qty = frm.Qty
	cartonHistory.IpAddress = c.Context().RemoteIP()
	cartonHistory.EmpID = services.GetEmpID(c)
	cartonHistory.Description = "CREATED"
	cartonHistory.IsActive = true
	if err := configs.Store.Create(&cartonHistory).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = fmt.Sprintf("%s %s", err.Error(), frm.SerialNo)
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("บันทึกข้อมูล %s เรียบร้อยแล้ว", frm.SerialNo)
	ledger.Whs = whs
	ledger.Part = part
	carton.Ledger = ledger
	carton.Shelve = shelve
	r.Data = &carton
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateLedgerCarton(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var frm models.FrmCarton
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var carton models.Carton
	if err := configs.Store.First(&carton, "serial_no", frm.SerialNo).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), frm.SerialNo)
		return c.Status(r.StatusCode).JSON(r)
	}

	var whs models.Whs
	if err := configs.Store.First(&whs, "title", frm.WhsID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), frm.WhsID)
		return c.Status(r.StatusCode).JSON(r)
	}

	var part models.Part
	if err := configs.Store.First(&part, "title", frm.LedgerID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), frm.LedgerID)
		return c.Status(r.StatusCode).JSON(r)
	}

	var ledger models.Ledger
	if err := configs.Store.
		Where("whs_id=?", whs.ID).
		Where("part_id=?", part.ID).
		First(&ledger).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), frm.LedgerID)
		return c.Status(r.StatusCode).JSON(r)
	}

	var shelve models.Shelve
	if err := configs.Store.First(&shelve, "title", frm.ShelveID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.ShelveID)
		return c.Status(r.StatusCode).JSON(r)
	}

	carton.LedgerID = &ledger.ID
	carton.SerialNo = frm.SerialNo
	carton.LotNo = frm.LotNo
	carton.LineNo = frm.LineNo
	carton.ReviseNo = frm.ReviseNo
	carton.PalletNo = frm.PalletNo
	carton.ShelveID = &shelve.ID
	carton.Qty = frm.Qty
	carton.IsActive = frm.IsActive
	if err := configs.Store.Save(&carton).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	// Create Carton History
	var cartonHistory models.CartonHistory
	cartonHistory.CartonID = carton.SerialNo
	cartonHistory.PalletNo = frm.PalletNo
	cartonHistory.ShelveID = shelve.Title
	cartonHistory.Qty = frm.Qty
	cartonHistory.IpAddress = c.Context().RemoteIP()
	cartonHistory.EmpID = services.GetEmpID(c)
	cartonHistory.Description = "UPDATED"
	cartonHistory.IsActive = true
	if err := configs.Store.Create(&cartonHistory).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = fmt.Sprintf("%s %s", err.Error(), frm.SerialNo)
		return c.Status(r.StatusCode).JSON(r)
	}

	var ctnLedger int64
	if err := configs.Store.Where("qty > ?", 0).Where("ledger_id=?", ledger.ID).Find(&models.Carton{}).Count(&ctnLedger).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = fmt.Sprintf("%s %s", err.Error(), frm.SerialNo)
		return c.Status(r.StatusCode).JSON(r)
	}

	ledger.Ctn, _ = strconv.ParseFloat(strconv.Itoa(int(ctnLedger)), 64)
	if err := configs.Store.Save(&ledger).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = fmt.Sprintf("Update Ledger %s", err.Error())
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("อัพดทข้อมูล %s เรียบร้อยแล้ว", frm.SerialNo)
	ledger.Whs = whs
	ledger.Part = part
	carton.Ledger = ledger
	carton.Shelve = shelve
	r.Data = &carton
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteLedgerCarton(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var carton models.Carton
	if err := configs.Store.First(&carton, "id", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), c.Params("id"))
		return c.Status(r.StatusCode).JSON(r)
	}

	var ctnLedger int64
	if err := configs.Store.Where("qty > ?", 0).Where("ledger_id=?", carton.LedgerID).Find(&models.Carton{}).Count(&ctnLedger).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = fmt.Sprintf("%s %s", err.Error(), *carton.LedgerID)
		return c.Status(r.StatusCode).JSON(r)
	}

	var ledger models.Ledger
	if err := configs.Store.First(&ledger, "id", carton.LedgerID).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = fmt.Sprintf("%s %s", err.Error(), *carton.LedgerID)
		return c.Status(r.StatusCode).JSON(r)
	}

	ledger.Ctn, _ = strconv.ParseFloat(strconv.Itoa(int(ctnLedger)), 64)
	if err := configs.Store.Save(&ledger).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = fmt.Sprintf("Update Ledger %s", err.Error())
		return c.Status(r.StatusCode).JSON(r)
	}

	if err := configs.Store.Delete(&carton).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = fmt.Sprintf("%s %s", err.Error(), c.Params("id"))
	}

	// Create Carton History
	var cartonHistory models.CartonHistory
	cartonHistory.CartonID = carton.SerialNo
	cartonHistory.PalletNo = "-"
	cartonHistory.ShelveID = "DELETED"
	cartonHistory.Qty = 0
	cartonHistory.IpAddress = c.Context().RemoteIP()
	cartonHistory.EmpID = services.GetEmpID(c)
	cartonHistory.Description = "DELETED"
	cartonHistory.IsActive = true
	if err := configs.Store.Create(&cartonHistory).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = fmt.Sprintf("%s %s", err.Error(), c.Params("id"))
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("ลบข้อมูล %s เรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}
