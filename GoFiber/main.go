package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Ninja struct {
	Name   string
	Weapon string
}

var ninja Ninja

func getNinja(ctx *fiber.Ctx) error {
	if ninja.Name == "" {
		return ctx.Status(fiber.StatusNotFound).SendString("Ninja not found")
	} else {
		return ctx.Status(fiber.StatusOK).JSON(ninja)
	}
}

func CreateNinja(ctx *fiber.Ctx) error {
	body := new(Ninja)
	err := ctx.BodyParser(body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString("Error parsing body")
		return err
	}
	ninja = Ninja{
		Name:   body.Name,
		Weapon: body.Weapon,
	}

	return ctx.Status(fiber.StatusOK).JSON(ninja)
}

func main() {
	fmt.Println("Hello, Learning how to use Go Fiber")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("<h1>Hello, World!</h1>")
	})

	// app.Get("/ninja", getNinja)

	// app.Post("/ninja", CreateNinja)

	// OR we can group also

	// These re middlewares which intercep the request for specific taks and the pont it will work for reuests below this nor above one's are effected.
	app.Use(logger.New())
	app.Use(requestid.New())
	ninjaGroup := app.Group("/ninja")
	ninjaGroup.Get("/", getNinja)
	ninjaGroup.Post("/", CreateNinja)

	app.Listen(":3000")

}
