package api


import (
"github.com/gofiber/fiber/v2"
"ROMAX/internal/handlers"
)


// RegisterRoutes sets up all API endpoints
func RegisterRoutes(app *fiber.App) {
api := app.Group("/api/v1")


// Player endpoints
api.Get("/player/:id", handlers.GetPlayer)
api.Post("/player", handlers.CreatePlayer)


// Game endpoints
api.Post("/play", handlers.PlayHandler)
}