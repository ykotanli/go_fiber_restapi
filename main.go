package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/ykotanli/app"
	"github.com/ykotanli/configs"
	"github.com/ykotanli/repository"
	"github.com/ykotanli/services"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDB()
	dbClient := configs.GetCollection(configs.DB, "todos")

	TodoRepositoryDb := repository.NewTodoRepositoryDB(dbClient)

	td := app.TodoHandler{Service: services.NewTodoService(TodoRepositoryDb)}
	appRoute.Use(logger.New())
	appRoute.Use(requestid.New())
	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todo", td.GetAllTodo)
	appRoute.Delete("/api/todo/:id", td.DeleteTodo)
	appRoute.Delete("/api/todo/", td.DeleteAll)
	appRoute.Listen(":8080")
}
