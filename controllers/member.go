package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/abe/erp.api/services"
	"github.com/gofiber/fiber/v2"
)

func MemberRegister(c *fiber.Ctx) error {
	var r models.Response
	var frm models.User
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	// validate hashing password
	password := services.HashingPassword(frm.Password)
	isMatch := services.ComparePassword(frm.Password, password)
	if !isMatch {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = "Password does not match"
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	frm.Password = password
	frm.Email = strings.ToLower(frm.Email)

	if err := configs.Store.Create(&frm).Error; err != nil {
		r.StatusCode = fiber.StatusInternalServerError
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	// Send Mail After Register
	_, _, _ = services.SendMail(frm.Email, "ลงทะเบียนเข้าใช้งานระบบ", fmt.Sprintf("คุณได้ลงทะเบียนในชื่อ: %s\nลงทะเบียนเรียบร้อยแล้ว\nกรุณาติดต่อทางผู้ดูแลระบบเพื่อร้องขอเข้าใช้งาน", frm.UserName))
	r.Message = fmt.Sprintf("%s Registered!", frm.UserName)
	r.StatusCode = fiber.StatusCreated
	r.Data = &frm
	return c.Status(r.StatusCode).JSON(&r)
}

func MemberAuth(c *fiber.Ctx) error {
	var r models.Response
	var frm models.User
	if err := c.BodyParser(&frm); err != nil {
		r.StatusCode = fiber.StatusBadRequest
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	var userData models.User
	if err := configs.Store.Where("username=?", frm.UserName).Where("is_active=?", true).First(&userData).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = fmt.Sprintf("ไม่พบข้อมูลผู้ใช้งาน %s", frm.UserName)
		return c.Status(r.StatusCode).JSON(&r)
	}

	if !services.ComparePassword(frm.Password, userData.Password) {
		r.StatusCode = fiber.StatusUnauthorized
		r.Message = fmt.Sprintf("ขอ อภัย %s รหัสผ่านของท่านไม่ถูกต้อง!", frm.UserName)
		// Delete JWT
		configs.Store.Where("user_id=?", &userData.ID).Delete(&models.JwtToken{})
		return c.Status(r.StatusCode).JSON(&r)
	}

	r.StatusCode = fiber.StatusCreated
	r.Message = fmt.Sprintf("%s login is successfully!", frm.UserName)
	r.Data = services.CreateToken(&userData)
	return c.Status(r.StatusCode).JSON(&r)
}

func MemberProfile(c *fiber.Ctx) error {
	var r models.Response
	r.StatusCode = fiber.StatusOK
	r.At = time.Now()
	jwtToken := services.GetJWTToken(c)
	var userData models.JwtToken
	if err := configs.Store.
		Preload("User.Area").
		Preload("User.Whs").
		Preload("User.Factory").
		Preload("User.Position").
		Preload("User.Section").
		Preload("User.Department").
		Where("id=?", jwtToken).First(&userData).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}
	r.Message = "แสดงข้อมูลส่วนตัว"
	r.Data = &userData
	return c.Status(r.StatusCode).JSON(&r)
}
