package main

import (
	"fmt"
	"log"
	"os"

	"github.com/danielwiratman/go-url-shortener/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}

func main() {
	app := fiber.New(fiber.Config{
		Views: django.New("./views", ".html"),
	})
	app.Static("/", "./static")

	app.Get("/:shortURL", handlers.GetShortUrl)
	app.Post("/create", handlers.PostNewUrl)
	app.Get("/", handlers.HandleIndex)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
	// because app.Listen returns error, might as well just print it
	// This logging happen when error actually happens anyway
}
