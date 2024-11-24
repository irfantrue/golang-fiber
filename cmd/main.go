package main

import (
	"golang-fiber/internal/middleware"
	"golang-fiber/pkg/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type User struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Age      uint   `json:"age"`
}

func main() {
	app := config.FiberConfig()
	app.Use(middleware.Cors())
	app.Use(middleware.RequestId())
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.Compress())
	app.Use(middleware.Limiter())
	app.Use(middleware.Helmet())

	app.Get("/metrics", monitor.New())
	app.Get("/api", func(c *fiber.Ctx) error {
		dataUser := User{
			Email:    "irfansusilo88@gmail.com",
			FullName: "Irfan Nurul Susilo",
			Age:      23,
		}

		if dataUser.Age == 23 {
			return fiber.NewError(400, middleware.Panic())
		}

		return c.JSON(dataUser)
	})

	log.Fatal(app.Listen(":8888"))
}
