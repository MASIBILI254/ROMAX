package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"ROMAX/internal/api"
	"ROMAX/internal/storage"
)

func main(){
	app:=fiber.New()

	storage.Init()

	api.RegisterRoutes(app)

	log.Println("starting Roamx on port:8080")
	log.Fatal(app.Listen(":8080"))
}