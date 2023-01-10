package routes

import (
	"fmt"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/controllers"
	"github.com/abe/erp.api/services"
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
	log.Delete("/:id", controllers.DeleteSystemLogger)

	// Prefix User
	user := r.Group("/user")
	user.Post("/register", controllers.MemberRegister)
	user.Post("/login", controllers.MemberAuth)
	// Begin Use Middleware
	appRouter := r.Use(services.AuthorizationRequired)

	// User Authentication
	auth := appRouter.Group("/auth")
	auth.Get("/me", controllers.MemberProfile)
}
