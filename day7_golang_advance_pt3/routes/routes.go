package routes

import (
	"user-management/handler"

	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router, userHandler *handler.UserHandler) {
	r := app.Group("api/v1")

	r.Get("/users", userHandler.GetList)
	r.Post("/users", userHandler.CreateUser)
	r.Put("/users", userHandler.UpdateUser)
	r.Delete("/users", userHandler.DeleteUser)
}
