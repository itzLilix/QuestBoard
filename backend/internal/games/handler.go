package games

import (
	"github.com/gofiber/fiber/v3"
)

func Handler(app *fiber.App) {
	games := app.Group("/games")
	games.Get("/", getGames)
	games.Get("/:id", getGameById)
	games.Post("/", createGame)
	games.Patch("/:id", editGame)
	games.Delete("/:id", deleteGameById)
	games.Post("/:id/join", addPlayerToGame)
}

func getGames(c fiber.Ctx) error {
	return nil;
}

func getGameById(c fiber.Ctx) error {
	return nil;
}

func createGame(c fiber.Ctx) error {
	return nil;
}

func editGame(c fiber.Ctx) error {
	return nil;
}

func deleteGameById(c fiber.Ctx) error {
	return nil;
}

func addPlayerToGame(c fiber.Ctx) error {
	return nil;
}