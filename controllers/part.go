package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/abe/erp.api/services"
	"github.com/gofiber/fiber/v2"
)

func GetPart(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	if c.Query("id") == "" {
		var Part []models.Part
		if err := configs.Store.Order("slug").Where("is_active=?", true).Find(&Part).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(r)
		}

		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &Part
		return c.Status(r.StatusCode).JSON(&r)
	}

	var Part models.Part
	if err := configs.Store.First(&Part, "id=?", c.Query("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("แสดงข้อมูล %s", c.Query("id"))
	r.Data = &Part
	return c.Status(r.StatusCode).JSON(&r)
}

func CreatePart(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated

	var frm models.Part
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var Part models.Part
	Part.Slug = strings.ToUpper(frm.Slug)
	Part.Title = strings.ToUpper(frm.Title)
	Part.Kinds = frm.Kinds
	Part.Size = frm.Size
	Part.Color = frm.Color
	Part.Description = frm.Description
	Part.IsActive = frm.IsActive
	if err := configs.Store.Create(&Part).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("บันทึกข้อมูล %s เรียบร้อยแล้ว", Part.Title)
	r.Data = &Part
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdatePart(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	var frm models.Part
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var Part models.Part
	if err := configs.Store.First(&Part, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	Part.Slug = strings.ToUpper(frm.Slug)
	Part.Title = strings.ToUpper(frm.Title)
	Part.Kinds = frm.Kinds
	Part.Size = frm.Size
	Part.Color = frm.Color
	Part.Description = frm.Description
	Part.IsActive = frm.IsActive
	if err := configs.Store.Save(&Part).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}
	r.Message = fmt.Sprintf("อัพเดท %s เรียบร้อยแล้ว", c.Params("id"))
	r.Data = &Part
	return c.Status(r.StatusCode).JSON(&r)
}

func DeletePart(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var Part models.Part
	if err := configs.Store.First(&Part, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	if err := configs.Store.Delete(&Part).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("ลบข้อมูล %s เรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}

func SeedPart(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var Area models.Area
	if err := configs.Store.First(&Area, "title", "CK").Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}
	var Factory models.Factory
	if err := configs.Store.First(&Factory, "title", c.Query("factory")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	if c.Query("factory") == "INJ" {
		configFile, err := os.Open(fmt.Sprintf(".%s/data/part_inj.json", configs.APP_PUBLIC_DIRS))
		if err != nil {
			fmt.Printf("opening config file %s\n", err.Error())
		}

		var part models.SdPart
		jsonParser := json.NewDecoder(configFile)
		if err = jsonParser.Decode(&part); err != nil {
			fmt.Printf("opening config file %s\n", err.Error())
		}

		var whsName [2]string
		whsName[0] = "COM"
		whsName[1] = "DOM"

		var ItemType models.ItemType
		if err := configs.Store.First(&ItemType, "title", "PART").Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(&r)
		}
		var Unit models.Unit
		if err := configs.Store.First(&Unit, "title", "BOX").Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(&r)
		}

		for i := range part.Data {
			p := part.Data[i]
			var partData models.Part
			partData.Slug = strings.ReplaceAll(p.Title, "-", "")
			partData.Title = p.Title
			partData.Kinds = "-"
			partData.Size = "-"
			partData.Color = "-"
			partData.Description = p.Description
			partData.IsActive = true

			if err := configs.Store.Create(&partData).Error; err != nil {
				fmt.Printf("saving part %s\n", err.Error())
			}

			for x := range whsName {
				var whs models.Whs
				if err := configs.Store.First(&whs, "title", whsName[x]).Error; err != nil {
					fmt.Printf("saving part %s\n", err.Error())
				}
				var ledger models.Ledger
				ledger.WhsID = &whs.ID
				ledger.AreaID = &Area.ID
				ledger.FactoryID = &Factory.ID
				ledger.PartID = &partData.ID
				ledger.ItemTypeID = &ItemType.ID
				ledger.UnitID = &Unit.ID
				ledger.IsActive = true
				if err := configs.Store.Create(&ledger).Error; err != nil {
					fmt.Printf("saving part %s\n", err.Error())
				}
			}
		}
	} else {
		configFile, err := os.Open(fmt.Sprintf(".%s/data/part_aw.json", configs.APP_PUBLIC_DIRS))
		if err != nil {
			fmt.Printf("opening config file %s\n", err.Error())
		}

		var part models.SdPart
		jsonParser := json.NewDecoder(configFile)
		if err = jsonParser.Decode(&part); err != nil {
			fmt.Printf("opening config file %s\n", err.Error())
		}

		var whs models.Whs
		if err := configs.Store.First(&whs, "title", "COM").Error; err != nil {
			fmt.Printf("saving part %s\n", err.Error())
		}

		var ItemType models.ItemType
		if err := configs.Store.First(&ItemType, "title", "WIRE").Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(&r)
		}
		var Unit models.Unit
		if err := configs.Store.First(&Unit, "title", "COIL").Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(&r)
		}

		for i := range part.Data {
			p := part.Data[i]
			var partData models.Part
			partData.Slug = strings.ReplaceAll(p.Title, "-", "")
			partData.Title = p.Title
			partData.Kinds, partData.Size, partData.Color = services.SubStringWire(p.Description)
			partData.Description = p.Description
			partData.IsActive = true

			if err := configs.Store.Create(&partData).Error; err != nil {
				fmt.Printf("saving part %s\n", err.Error())
			}

			var ledger models.Ledger
			ledger.WhsID = &whs.ID
			ledger.AreaID = &Area.ID
			ledger.FactoryID = &Factory.ID
			ledger.PartID = &partData.ID
			ledger.ItemTypeID = &ItemType.ID
			ledger.UnitID = &Unit.ID
			ledger.IsActive = true
			if err := configs.Store.Create(&ledger).Error; err != nil {
				fmt.Printf("saving part %s\n", err.Error())
			}
		}
	}
	r.Message = "บันทึกข้อมูลเรียบร้อยแล้ว"
	return c.Status(r.StatusCode).JSON(&r)
}
