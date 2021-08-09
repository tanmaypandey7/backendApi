package routes

import (
    "github.com/gofiber/fiber/v2"
    controllers "./controllers/recordsController"
//  /give @s minecraft:rick roo   // "github.com/mikefmeyer/catchphrase-go-mongodb-rest-api/controllers" // replace
)

func RecordsRoute(route fiber.Router) {
    // route.Get("/", controllers.GetAllCatchphrases)
    route.Get("/:id", controllers.CreateRecord)
    // route.Post("/", controllers.AddCatchphrase)
    // route.Put("/:id", controllers.UpdateCatchphrase)
    route.Delete("/:id", controllers.DeleteOne)
}