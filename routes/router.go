package routes

import (
	"fmt"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/controllers"
	"github.com/gofiber/fiber/v2"
)

func Router(c *fiber.App) {
	c.Get("", controllers.Hello)
	// Prefix Api
	r := c.Group(fmt.Sprintf("/api/%s", configs.APP_VERSION))
	r.Get("", controllers.Hello)

	log := r.Group("/logger")
	log.Get("", controllers.GetSystemLogger)
	log.Post("", controllers.CreateSystemLogger)
	log.Put("/:id", controllers.UpdateSystemLogger)
}
