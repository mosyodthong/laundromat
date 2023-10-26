package main

import (
	"fmt"
	"laundromat/route"

	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	initTimeZone()
	route.Router(app)

	app.Listen(fmt.Sprintf(":%v", os.Getenv("API_PORT")))
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err.Error())
	}
	time.Local = ict
}
