package routes

import (
	"github.com/abe/erp.api/controllers"
	"github.com/gofiber/fiber/v2"
)

func Router(c *fiber.App) {
	c.Get("", controllers.Hello)
	// Prefix Api
	r := c.Group("/api/v1")
	r.Get("", controllers.Hello)
}
