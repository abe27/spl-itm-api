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
	log.Patch("", controllers.TestSendMail)
	log.Put("/:id", controllers.UpdateSystemLogger)
	log.Delete("/:id", controllers.DeleteSystemLogger)

	// Prefix User
	user := r.Group("/user")
	user.Post("/register", controllers.MemberRegister)
	user.Post("/login", controllers.MemberAuth)
	user.Post("/admin", controllers.CreateAdmin)

	// Begin Use Middleware
	appRouter := r.Use(services.AuthorizationRequired)
	// User Authentication
	auth := appRouter.Group("/auth")
	auth.Get("/me", controllers.MemberProfile)

	// Master Group
	area := appRouter.Group("/area")
	area.Get("", controllers.GetArea)
	area.Post("", controllers.CreateArea)
	area.Put("/:id", controllers.UpdateArea)
	area.Delete("/:id", controllers.DeleteArea)
	// <!----->
	whs := appRouter.Group("/whs")
	whs.Get("", controllers.GetWhs)
	whs.Post("", controllers.CreateWhs)
	whs.Put("/:id", controllers.UpdateWhs)
	whs.Delete("/:id", controllers.DeleteWhs)
	// <!----->
	factory := appRouter.Group("/factory")
	factory.Get("", controllers.GetFactory)
	factory.Post("", controllers.CreateFactory)
	factory.Put("/:id", controllers.UpdateFactory)
	factory.Delete("/:id", controllers.DeleteFactory)
	// <!----->
	unit := appRouter.Group("/unit")
	unit.Get("", controllers.GetUnit)
	unit.Post("", controllers.CreateUnit)
	unit.Put("/:id", controllers.UpdateUnit)
	unit.Delete("/:id", controllers.DeleteUnit)
	// <!----->
	itemType := appRouter.Group("/itemType")
	itemType.Get("", controllers.GetItemType)
	itemType.Post("", controllers.CreateItemType)
	itemType.Put("/:id", controllers.UpdateItemType)
	itemType.Delete("/:id", controllers.DeleteItemType)
	// <!----->
	position := appRouter.Group("/position")
	position.Get("", controllers.GetPosition)
	position.Post("", controllers.CreatePosition)
	position.Put("/:id", controllers.UpdatePosition)
	position.Delete("/:id", controllers.DeletePosition)
	// <!----->
	section := appRouter.Group("/section")
	section.Get("", controllers.GetSection)
	section.Post("", controllers.CreateSection)
	section.Put("/:id", controllers.UpdateSection)
	section.Delete("/:id", controllers.DeleteSection)
	// <!----->
	department := appRouter.Group("/department")
	department.Get("", controllers.GetPosition)
	department.Post("", controllers.CreatePosition)
	department.Put("/:id", controllers.UpdatePosition)
	department.Delete("/:id", controllers.DeletePosition)
	// <!----->
	shipment := appRouter.Group("/shipment")
	shipment.Get("", controllers.GetShipment)
	shipment.Post("", controllers.CreateShipment)
	shipment.Put("/:id", controllers.UpdateShipment)
	shipment.Delete("/:id", controllers.DeleteShipment)
	// <!----->
	mailType := appRouter.Group("/mailType")
	mailType.Get("", controllers.GetMailType)
	mailType.Post("", controllers.CreateMailType)
	mailType.Put("/:id", controllers.UpdateMailType)
	mailType.Delete("/:id", controllers.DeleteMailType)
	// <!----->
	mailBox := appRouter.Group("/mailbox")
	mailBox.Get("", controllers.GetMailBox)
	mailBox.Post("", controllers.CreateMailBox)
	mailBox.Put("/:id", controllers.UpdateMailBox)
	mailBox.Delete("/:id", controllers.DeleteMailBox)
}
