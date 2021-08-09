package main

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	// "ioutil"
	"io/ioutil"
	"github.com/gofiber/fiber/v2"
)
type Article struct {
	Id      string `json:"Id"`
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

var Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func postArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: postArticles")

}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // return the string response containing the request body    
    reqBody, _ := ioutil.ReadAll(r.Body)
    fmt.Fprintf(w, "%+v", string(reqBody))
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    // key := vars["id"]
    title := vars["title"]

	for _, article := range Articles{
		if article.Title == title {
			json.NewEncoder(w).Encode(article)
		}
	}
}
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
            "success":     true,
            "message":     "You are at the root endpoint",
        })
    })

    api := app.Group("/api")

    routes.RecordsRoute(api.Group("/records"))
}

func main() {

}