package controllers

import (
	"fmt"
	"strings"

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
	if err := configs.Store.Where("username=?", frm.UserName).First(&userData).Error; err != nil {
		r.StatusCode = fiber.StatusNotFound
		r.Message = err.Error()
		return c.Status(r.StatusCode).JSON(&r)
	}
	r.StatusCode = fiber.StatusCreated
	r.Message = fmt.Sprintf("%s login is successfully!", frm.UserName)
	r.Data = services.CreateToken(&userData)
	return c.Status(r.StatusCode).JSON(&r)
}

func MemberProfile(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(r.StatusCode).JSON(&r)
}
