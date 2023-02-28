package api

import "github.com/gofiber/fiber/v2"

func LaunchServer() {
    app := fiber.New()

    app.Get("/api/upcoming", func(c *fiber.Ctx) error {
        contests := g 
    })

    app.Listen(":3000")
}

func main

