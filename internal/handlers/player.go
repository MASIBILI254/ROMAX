package handlers


import (
	"github.com/gofiber/fiber/v2"
	"ROMAX/internal/models"
	"ROMAX/internal/storage"
	"ROMAX/internal/game"
)


// GetPlayer handles 
func GetPlayer(c *fiber.Ctx) error {
playerID := c.Params("id")
if playerID == "" {
return fiber.NewError(fiber.StatusBadRequest, "player id required")
}

player, ok := storage.GetPlayer(playerID)
if !ok {
return fiber.NewError(fiber.StatusNotFound, "player not found")
}

return c.JSON(player)
}


// CreatePlayer 
func CreatePlayer(c *fiber.Ctx) error {
var req struct {
ID string `json:"id"`
}
if err := c.BodyParser(&req); err != nil {
return fiber.ErrBadRequest
}
if req.ID == "" {
return fiber.NewError(fiber.StatusBadRequest, "player id required")
}

player := storage.CreatePlayer(req.ID)
return c.JSON(player)
}


// PlayHandler 
func PlayHandler(c *fiber.Ctx) error {
	var req models.RoundRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	player, ok := storage.GetPlayer(req.PlayerID)
	if !ok {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "player not found"})
	}

	round, err := game.RunRound(player, req.Bet, req.UseFreeSpin)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(round)
}