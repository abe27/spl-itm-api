package controllers

import (
	"fmt"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetLedger(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	if c.Query("part_no") == "" {
		var ledger []models.Ledger
		if err := configs.Store.
			Preload("Whs").
			Preload("Area").
			Preload("Factory").
			Preload("Part").
			Preload("ItemType").
			Preload("Unit").
			Find(&ledger).Error; err != nil {
			r.StatusCode = fiber.StatusInternalServerError
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(&r)
		}

		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &ledger
		return c.Status(r.StatusCode).JSON(&r)
	}

	var part models.Part
	if err := configs.Store.First(&part, "title", c.Query("part_no")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	var ledger []models.Ledger
	if err := configs.Store.
		Preload("Whs").
		Preload("Area").
		Preload("Factory").
		Preload("Part").
		Preload("ItemType").
		Preload("Unit").
		Find(&ledger).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	r.Message = fmt.Sprintf("แสดงข้อมูล %s", c.Query("part_no"))
	r.Data = &ledger
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateLedger(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated
	var frm models.Ledger
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	// WhsID
	var whs models.Whs
	if err := configs.Store.First(&whs, "title", frm.WhsID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.WhsID)
		return c.Status(r.StatusCode).JSON(&r)
	}
	// AreaID
	var area models.Area
	if err := configs.Store.First(&area, "title", frm.AreaID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.AreaID)
		return c.Status(r.StatusCode).JSON(&r)
	}
	// FactoryID
	var factory models.Factory
	if err := configs.Store.First(&factory, "title", frm.FactoryID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.FactoryID)
		return c.Status(r.StatusCode).JSON(&r)
	}
	// PartID
	var part models.Part
	if err := configs.Store.First(&part, "title", frm.PartID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.PartID)
		return c.Status(r.StatusCode).JSON(&r)
	}
	// ItemTypeID
	var itemType models.ItemType
	if err := configs.Store.First(&itemType, "title", frm.ItemTypeID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.ItemTypeID)
		return c.Status(r.StatusCode).JSON(&r)
	}
	// UnitID
	var unit models.Unit
	if err := configs.Store.First(&unit, "title", frm.UnitID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.UnitID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	var ledger models.Ledger
	if err := configs.Store.
		Where("whs_id=?", &whs.ID).
		Where("area_id=?", &area.ID).
		Where("factory_id=?", &factory.ID).
		Where("part_id=?", &part.ID).
		First(&ledger).Error; err == nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = fmt.Sprintf("บันทึกข้อมูลซ้ำ %s", ledger.ID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	ledger.WhsID = &whs.ID
	ledger.AreaID = &area.ID
	ledger.FactoryID = &factory.ID
	ledger.PartID = &part.ID
	ledger.ItemTypeID = &itemType.ID
	ledger.UnitID = &unit.ID
	ledger.DimWidth = frm.DimWidth
	ledger.DimLength = frm.DimLength
	ledger.DimHeight = frm.DimHeight
	ledger.GrossWeight = frm.GrossWeight
	ledger.NetWeight = frm.NetWeight
	ledger.Qty = frm.Qty
	ledger.Ctn = frm.Ctn
	ledger.IsActive = frm.IsActive

	if err := configs.Store.Create(&ledger).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	ledger.Whs = whs
	ledger.Area = area
	ledger.Factory = factory
	ledger.Part = part
	ledger.ItemType = itemType
	ledger.Unit = unit
	r.Data = &ledger
	r.Message = "บันทึกข้อมูลเรียบร้อยแล้ว"
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateLedger(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var frm models.Ledger
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	var ledger models.Ledger
	if err := configs.Store.First(&ledger, "id", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	// WhsID
	var whs models.Whs
	if err := configs.Store.First(&whs, "title", frm.WhsID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.WhsID)
		return c.Status(r.StatusCode).JSON(&r)
	}
	// AreaID
	var area models.Area
	if err := configs.Store.First(&area, "title", frm.AreaID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.AreaID)
		return c.Status(r.StatusCode).JSON(&r)
	}
	// FactoryID
	var factory models.Factory
	if err := configs.Store.First(&factory, "title", frm.FactoryID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.FactoryID)
		return c.Status(r.StatusCode).JSON(&r)
	}
	// PartID
	var part models.Part
	if err := configs.Store.First(&part, "title", frm.PartID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.PartID)
		return c.Status(r.StatusCode).JSON(&r)
	}
	// ItemTypeID
	var itemType models.ItemType
	if err := configs.Store.First(&itemType, "title", frm.ItemTypeID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.ItemTypeID)
		return c.Status(r.StatusCode).JSON(&r)
	}
	// UnitID
	var unit models.Unit
	if err := configs.Store.First(&unit, "title", frm.UnitID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.UnitID)
		return c.Status(r.StatusCode).JSON(&r)
	}

	ledger.WhsID = &whs.ID
	ledger.AreaID = &area.ID
	ledger.FactoryID = &factory.ID
	ledger.PartID = &part.ID
	ledger.ItemTypeID = &itemType.ID
	ledger.UnitID = &unit.ID
	ledger.DimWidth = frm.DimWidth
	ledger.DimLength = frm.DimLength
	ledger.DimHeight = frm.DimHeight
	ledger.GrossWeight = frm.GrossWeight
	ledger.NetWeight = frm.NetWeight
	ledger.Qty = frm.Qty
	ledger.Ctn = frm.Ctn
	ledger.IsActive = frm.IsActive

	if err := configs.Store.Save(&ledger).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	ledger.Whs = whs
	ledger.Area = area
	ledger.Factory = factory
	ledger.Part = part
	ledger.ItemType = itemType
	ledger.Unit = unit
	r.Data = &ledger
	r.Message = "อัพเดทข้อมูลเรียบร้อยแล้ว"
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteLedger(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var ledger models.Ledger
	if err := configs.Store.First(&ledger, "id", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}

	if err := configs.Store.Delete(&ledger).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}
	r.Message = "ลบข้อมูลเรียบร้อยแล้ว"
	return c.Status(r.StatusCode).JSON(&r)
}
