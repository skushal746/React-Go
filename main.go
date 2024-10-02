package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {

	fmt.Println("hello world ")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "Hello World"})
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}
	})

	log.Fatal(app.Listen(":4000"))

}
