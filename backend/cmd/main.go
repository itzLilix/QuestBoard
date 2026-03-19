package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/itzLilix/QuestBoard/backend/internal/auth"
	"github.com/itzLilix/QuestBoard/backend/internal/games"
)


func main() {
	app := fiber.New()

	auth.Handler(app)
	games.Handler(app)

	log.Fatal(app.Listen(":3000"))
}