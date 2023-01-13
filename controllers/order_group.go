package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
)

func GetOrderGroup(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	if c.Query("id") == "" {
		var OrderGroup []models.OrderGroup
		if err := configs.Store.Preload("User").Preload("Affcode").Preload("Customer").Preload("OrderGroupType").Where("is_active=?", true).Find(&OrderGroup).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(r)
		}

		r.Message = "แสดงข้อมูลทั้งหมด"
		r.Data = &OrderGroup
		return c.Status(r.StatusCode).JSON(&r)
	}

	var OrderGroup models.OrderGroup
	if err := configs.Store.Preload("User").Preload("Affcode").Preload("Customer").Preload("OrderGroupType").First(&OrderGroup, "id=?", c.Query("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("แสดงข้อมูล %s", c.Query("id"))
	r.Data = &OrderGroup
	return c.Status(r.StatusCode).JSON(&r)
}

func CreateOrderGroup(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusCreated

	var frm models.OrderGroup
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var user models.User
	if err := configs.Store.First(&user, "username", frm.UserID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.UserID)
		return c.Status(r.StatusCode).JSON(r)
	}

	var affcode models.Affcode
	if err := configs.Store.First(&affcode, "title", frm.AffcodeID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.AffcodeID)
		return c.Status(r.StatusCode).JSON(r)
	}

	var customer models.Customer
	if err := configs.Store.First(&customer, "title", frm.CustomerID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.CustomerID)
		return c.Status(r.StatusCode).JSON(r)
	}

	var orderGroupType models.OrderGroupType
	if err := configs.Store.First(&orderGroupType, "title", frm.OrderGroupTypeID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.OrderGroupTypeID)
		return c.Status(r.StatusCode).JSON(r)
	}

	var OrderGroup models.OrderGroup
	OrderGroup.UserID = &user.ID
	OrderGroup.AffcodeID = &affcode.ID
	OrderGroup.CustomerID = &customer.ID
	OrderGroup.OrderGroupTypeID = &orderGroupType.ID
	OrderGroup.SubOrder = frm.SubOrder
	OrderGroup.Description = frm.Description
	OrderGroup.IsActive = frm.IsActive
	if err := configs.Store.Create(&OrderGroup).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("บันทึกข้อมูล %s เรียบร้อยแล้ว", OrderGroup.Affcode.Title)
	OrderGroup.User = &user
	OrderGroup.Affcode = &affcode
	OrderGroup.Customer = &customer
	OrderGroup.OrderGroupType = &orderGroupType
	r.Data = &OrderGroup
	return c.Status(r.StatusCode).JSON(&r)
}

func UpdateOrderGroup(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	var frm models.OrderGroup
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var OrderGroup models.OrderGroup
	if err := configs.Store.First(&OrderGroup, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	var user models.User
	if err := configs.Store.First(&user, "username", frm.UserID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.UserID)
		return c.Status(r.StatusCode).JSON(r)
	}

	var affcode models.Affcode
	if err := configs.Store.First(&affcode, "title", frm.AffcodeID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.AffcodeID)
		return c.Status(r.StatusCode).JSON(r)
	}

	var customer models.Customer
	if err := configs.Store.First(&customer, "title", frm.CustomerID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.CustomerID)
		return c.Status(r.StatusCode).JSON(r)
	}

	var orderGroupType models.OrderGroupType
	if err := configs.Store.First(&orderGroupType, "title", frm.OrderGroupTypeID).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("%s %s", err.Error(), *frm.OrderGroupTypeID)
		return c.Status(r.StatusCode).JSON(r)
	}

	OrderGroup.UserID = &user.ID
	OrderGroup.AffcodeID = &affcode.ID
	OrderGroup.CustomerID = &customer.ID
	OrderGroup.OrderGroupTypeID = &orderGroupType.ID
	OrderGroup.SubOrder = frm.SubOrder
	OrderGroup.Description = frm.Description
	OrderGroup.IsActive = frm.IsActive
	if err := configs.Store.Save(&OrderGroup).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}
	r.Message = fmt.Sprintf("อัพเดท %s เรียบร้อยแล้ว", c.Params("id"))
	OrderGroup.User = &user
	OrderGroup.Affcode = &affcode
	OrderGroup.Customer = &customer
	OrderGroup.OrderGroupType = &orderGroupType
	r.Data = &OrderGroup
	return c.Status(r.StatusCode).JSON(&r)
}

func DeleteOrderGroup(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK
	var OrderGroup models.OrderGroup
	if err := configs.Store.First(&OrderGroup, "id=?", c.Params("id")).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	if err := configs.Store.Delete(&OrderGroup).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(r)
	}

	r.Message = fmt.Sprintf("ลบข้อมูล %s เรียบร้อยแล้ว", c.Params("id"))
	return c.Status(r.StatusCode).JSON(&r)
}

func GenerateOrderGroup(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusOK

	file, _ := ioutil.ReadFile("./public/data/order_group.json")

	data := models.SeedOrderGroup{}

	if err := json.Unmarshal([]byte(file), &data); err != nil {
		panic(err)
	}

	for i := 0; i < len(data.Data); i++ {
		obj := data.Data[i]
		var user models.User
		if err := configs.Store.First(&user, "username", obj.UserName).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = fmt.Sprintf("%s %s", err.Error(), obj.UserName)
			return c.Status(r.StatusCode).JSON(r)
		}

		var affcode models.Affcode
		if err := configs.Store.First(&affcode, "title", obj.AffCode).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = fmt.Sprintf("%s %s", err.Error(), obj.AffCode)
			return c.Status(r.StatusCode).JSON(r)
		}

		var customer models.Customer
		if err := configs.Store.First(&customer, "title", obj.CustCode).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = fmt.Sprintf("%s %s", err.Error(), obj.CustCode)
			return c.Status(r.StatusCode).JSON(r)
		}

		var orderGroupType models.OrderGroupType
		if err := configs.Store.First(&orderGroupType, "title", obj.GroupOrder).Error; err != nil {
			r.StatusCode = fiber.StatusNotFound
			r.Message = fmt.Sprintf("%s %s", err.Error(), obj.GroupOrder)
			return c.Status(r.StatusCode).JSON(r)
		}

		var OrderGroup models.OrderGroup
		OrderGroup.UserID = &user.ID
		OrderGroup.AffcodeID = &affcode.ID
		OrderGroup.CustomerID = &customer.ID
		OrderGroup.OrderGroupTypeID = &orderGroupType.ID
		OrderGroup.SubOrder = obj.SubOrder
		OrderGroup.Description = orderGroupType.Description
		OrderGroup.IsActive = true
		if err := configs.Store.FirstOrCreate(&OrderGroup, &models.OrderGroup{
			AffcodeID:        &affcode.ID,
			CustomerID:       &customer.ID,
			OrderGroupTypeID: &orderGroupType.ID,
			SubOrder:         obj.SubOrder,
		}).Error; err != nil {
			r.StatusCode = fiber.StatusInternalServerError
			r.Message = err.Error()
			return c.Status(r.StatusCode).JSON(r)
		}
	}
	return c.Status(r.StatusCode).JSON(&r)
}
