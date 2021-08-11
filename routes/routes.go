package routes

import (
	"github.com/gofiber/fiber/v2"
	"api_golang/controllers"
)

func RecordsRoute(route fiber.Router) {
	route.Post("/addRecord", controllers.AddRecord)
    route.Delete("/:id", controllers.DeleteRecord)
    route.Get("/:id", controllers.GetRecord)
    route.Put("/:id", controllers.UpdateRecord)

}