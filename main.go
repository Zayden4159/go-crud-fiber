package main

import (
	"fmt"

	"github.com/Zayden4159/crud-fiber/todo"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/users", todo.GetUsers)
	app.Get("/user/:id", todo.GetUser)
	app.Post("/user", todo.CreateUser)
	app.Put("/user/:id", todo.UpdateUser)
	app.Delete("/user/:id", todo.DeleteUser)
	app.Post("/upload", todo.UploadFile)

	fmt.Println("Server listening on port 8080")
	app.Listen(":8080")
}
