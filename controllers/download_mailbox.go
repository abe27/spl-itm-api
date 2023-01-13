package controllers

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/abe/erp.api/services"
	"github.com/gofiber/fiber/v2"
)

func GetDownloadMailBox(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	db := configs.Store
	if c.Query("id") == "" {
		var data []models.DownloadMailBox
		if err := db.Order("updated_at desc").Preload("MailBox.Area").Preload("MailType.Factory").Where("is_active=?", true).Find(&data).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(&r)
		}

		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &data
		return c.Status(r.StatusCode).JSON(&r)
	}

	var data models.DownloadMailBox
	if err := db.Order("updated_at desc").Preload("MailBox.Area").Preload("MailType.Factory").Where("is_active=?", true).First(&data, "id", c.Query("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	r.Message = fmt.Sprintf("แสดงข้อมูล %s", c.Query("id"))
	r.Data = &data
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateDownloadMailBox(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated

	var frm models.DownloadMailBox
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	db := configs.Store
	facLength := strings.TrimSpace(frm.BatchID[:len("OES.WHAE.32T5")])
	var mailType models.MailType
	if err := db.Preload("Factory").First(&mailType, "prefix", facLength).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), facLength)
		return c.Status(r.StatusCode).JSON(&r)
	}

	var mailBox models.MailBox
	if err := db.Preload("Area").First(&mailBox, "mail_id", frm.MailBoxID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), frm.MailBoxID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	file, err := c.FormFile("file_upload")
	if err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}
	if file == nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = "ไม่พบไฟล์ที่ต้องการอัพโหลด"
		return c.Status(r.StatusCode).JSON(&r)
	}

	fileDirs := fmt.Sprintf(".%s/upload/edi/%s", configs.APP_PUBLIC_DIRS, time.Now().Format("20060102"))
	if err := os.MkdirAll(fileDirs, os.ModePerm); err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	if err := c.SaveFile(file, fmt.Sprintf("%s/%s.%s", fileDirs, frm.BatchNo, file.Filename)); err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	var data models.DownloadMailBox
	data.MailBoxID = mailBox.ID
	data.MailTypeID = mailType.ID
	data.BatchNo = frm.BatchNo
	data.Size = frm.Size
	data.BatchID = frm.BatchID
	data.CreationDate = frm.CreationDate
	data.Flags = frm.Flags
	data.Format = frm.Format
	data.Originator = frm.Originator
	data.IsDownload = true //frm.IsDownload
	data.IsActive = frm.IsActive
	data.FilePath = fmt.Sprintf("upload/edi/%s/%s.%s", time.Now().Format("20060102"), frm.BatchNo, file.Filename)
	// if err := db.Create(&data).Error; err != nil {
	// 	r.StatusCode = fiber.StatusInternalServerError
	// 	r.Message = err.Error()
	// 	return c.Status(r.StatusCode).JSON(&r)
	// }
	data.MailBox = mailBox
	data.MailType = mailType
	r.Data = &data
	r.Message = fmt.Sprintf("อัพโหลด %s เรียบร้อยแล้ว", file.Filename)
	userID := services.GetUserID(c)
	services.ReadEDI(&data, &userID)
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateDownloadMailBox(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var frm models.DownloadMailBox
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	db := configs.Store
	var data models.DownloadMailBox
	if err := db.First(&data, "id", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	data.IsActive = frm.IsActive
	if err := db.Save(&data).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}
	r.Message = fmt.Sprintf("อัพเดท %s เรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteDownloadMailBox(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	db := configs.Store
	var data models.DownloadMailBox
	if err := db.First(&data, "id", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	if err := db.Delete(&data).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	r.Message = fmt.Sprintf("ลบ %s เรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}
