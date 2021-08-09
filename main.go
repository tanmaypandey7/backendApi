package main

import (
	"log"
	"net/http"

	"example.com/hello/routes"
	"github.com/gorilla/mux"


	"github.com/gofiber/fiber/v2"
)

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	// myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/returnAllArticles", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/article/{title}", returnSingleArticle).Methods("GET")
	myRouter.HandleFunc("/postArticles", postArticles).Methods("POST")
	myRouter.HandleFunc("/createNewArticle", createNewArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

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

}
