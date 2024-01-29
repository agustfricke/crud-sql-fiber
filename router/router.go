package router

import (
	"github.com/agustfricke/crud-sql-fiber/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
  app.Get("/hello/:artist", handlers.Hello)
  app.Post("/add", handlers.CreateAlbum)
}
