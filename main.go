package main

import (
	connectionhelper "api_golang/db"
	"api_golang/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the root endpoint",
		})
	})

	api := app.Group("/api")

	routes.RecordsRoute(api.Group("/records"))
}

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())
	// dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectionhelper.ConnectDB()

	setupRoutes(app)

	port := os.Getenv("PORT")
	app.Listen(":" + port)
	// app.Listen(":" + port)
	if app.Listen(":" + port) != nil {
		log.Fatal("Error app failed to start")
		panic(err)
	}
}
