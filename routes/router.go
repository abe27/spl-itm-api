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
	auth.Get("/logout", controllers.MemberLogOut)

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
	// <!----->
	part := appRouter.Group("/part")
	part.Get("", controllers.GetPart)
	part.Post("", controllers.CreatePart)
	part.Put("/:id", controllers.UpdatePart)
	part.Delete("/:id", controllers.DeletePart)
	// <!----->
	receiveType := appRouter.Group("/receiveType")
	receiveType.Get("", controllers.GetReceiveType)
	receiveType.Post("", controllers.CreateReceiveType)
	receiveType.Put("/:id", controllers.UpdateReceiveType)
	receiveType.Delete("/:id", controllers.DeleteReceiveType)
	// <!----->
	affcode := appRouter.Group("/affcode")
	affcode.Get("", controllers.GetAffcode)
	affcode.Post("", controllers.CreateAffcode)
	affcode.Put("/:id", controllers.UpdateAffcode)
	affcode.Delete("/:id", controllers.DeleteAffcode)
	// <!----->
	customer := appRouter.Group("/customer")
	customer.Get("", controllers.GetCustomer)
	customer.Post("", controllers.CreateCustomer)
	customer.Put("/:id", controllers.UpdateCustomer)
	customer.Delete("/:id", controllers.DeleteCustomer)
	// <!----->
	reviseOrder := appRouter.Group("/reviseOrder")
	reviseOrder.Get("", controllers.GetReviseOrder)
	reviseOrder.Post("", controllers.CreateReviseOrder)
	reviseOrder.Put("/:id", controllers.UpdateReviseOrder)
	reviseOrder.Delete("/:id", controllers.DeleteReviseOrder)
	// <!----->
	pc := appRouter.Group("/pc")
	pc.Get("", controllers.GetPc)
	pc.Post("", controllers.CreatePc)
	pc.Put("/:id", controllers.UpdatePc)
	pc.Delete("/:id", controllers.DeletePc)
	// <!----->
	commercial := appRouter.Group("/commercial")
	commercial.Get("", controllers.GetCommercial)
	commercial.Post("", controllers.CreateCommercial)
	commercial.Put("/:id", controllers.UpdateCommercial)
	commercial.Delete("/:id", controllers.DeleteCommercial)
	// <!----->
	sampleFlg := appRouter.Group("/sampleFlg")
	sampleFlg.Get("", controllers.GetSampleFlg)
	sampleFlg.Post("", controllers.CreateSampleFlg)
	sampleFlg.Put("/:id", controllers.UpdateSampleFlg)
	sampleFlg.Delete("/:id", controllers.DeleteSampleFlg)
	// <!----->
	orderType := appRouter.Group("/orderType")
	orderType.Get("", controllers.GetOrderType)
	orderType.Post("", controllers.CreateOrderType)
	orderType.Put("/:id", controllers.UpdateOrderType)
	orderType.Delete("/:id", controllers.DeleteOrderType)
	// <!----->
	orderZone := appRouter.Group("/orderZone")
	orderZone.Get("", controllers.GetOrderZone)
	orderZone.Post("", controllers.CreateOrderZone)
	orderZone.Put("/:id", controllers.UpdateOrderZone)
	orderZone.Delete("/:id", controllers.DeleteOrderZone)

	// <!----->
	orderGroupType := appRouter.Group("/orderGroupType")
	orderGroupType.Get("", controllers.GetOrderGroupType)
	orderGroupType.Post("", controllers.CreateOrderGroupType)
	orderGroupType.Put("/:id", controllers.UpdateOrderGroupType)
	orderGroupType.Delete("/:id", controllers.DeleteOrderGroupType)
	// <!----->
	orderGroup := appRouter.Group("/orderGroup")
	orderGroup.Get("", controllers.GetOrderGroup)
	orderGroup.Post("", controllers.CreateOrderGroup)
	orderGroup.Put("/:id", controllers.UpdateOrderGroup)
	orderGroup.Delete("/:id", controllers.DeleteOrderGroup)
	orderGroup.Patch("", controllers.GenerateOrderGroup)

	// <!----->
	lastInvoice := appRouter.Group("/lastInvoice")
	lastInvoice.Get("", controllers.GetLastInvoice)
	lastInvoice.Post("", controllers.CreateLastInvoice)
	lastInvoice.Put("/:id", controllers.UpdateLastInvoice)
	lastInvoice.Delete("/:id", controllers.DeleteLastInvoice)

	// Route Process Upload EDI
	edi := appRouter.Group("/edi")
	edi.Get("", controllers.GetDownloadMailBox)
	edi.Post("", controllers.CreateDownloadMailBox)
	edi.Put("/:id", controllers.UpdateDownloadMailBox)
	edi.Delete("/:id", controllers.DeleteDownloadMailBox)

	// Route Process Receive
	receiveEnt := appRouter.Group("/receiveEnt")
	receiveEnt.Get("", controllers.GetReceiveEnt)
	receiveEnt.Post("", controllers.CreateReceiveEnt)
	receiveEnt.Put("/:id", controllers.UpdateReceiveEnt)
	receiveEnt.Delete("/:id", controllers.DeleteReceiveEnt)
}
