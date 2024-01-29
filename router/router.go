package router

import (
	"github.com/agustfricke/crud-sql-fiber/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
  app.Get("/all", handlers.GetAll)
  app.Delete("/delete/:id", handlers.DeleteAlbum)
  app.Put("/put/:id", handlers.UpdateAlbum)
  app.Get("/name/:artist", handlers.GetByName)
  app.Get("/id/:id", handlers.GetByID)
  app.Post("/add", handlers.CreateAlbum)
}
