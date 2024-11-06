package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type ToDo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	app := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env")
	}

	PORT := os.Getenv("PORT")

	todos := []ToDo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	//create Todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &ToDo{}
		err := c.BodyParser(todo)
		if err != nil {
			return err
		}
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"message": "body field is required"})
		}
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(fiber.Map{"message": "todo successfully created", "todo": todo})
	})

	//Update Todos

	app.Patch("api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true

				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON("todo not found")
	})

	//delete todo

	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"message": "Successfully deleted"})
			}
		}

		return c.Status(404).JSON("Not Found")
	})

	app.Listen(":" + PORT)
}
