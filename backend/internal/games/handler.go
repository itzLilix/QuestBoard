package games

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
)

type Handler struct {
    db *pgx.Conn
}

func NewHandler(db *pgx.Conn) *Handler {
	return &Handler{db: db}
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	games := app.Group("/games")
	games.Get("/", h.getGames)
	games.Get("/:id", h.getGameById)
	games.Post("/", h.createGame)
	games.Patch("/:id", h.editGame)
	games.Delete("/:id", h.deleteGameById)
	games.Post("/:id/join", h.addPlayerToGame)
}

func (h *Handler) getGames(c fiber.Ctx) error {
	return nil;
}

func (h *Handler) getGameById(c fiber.Ctx) error {
	return nil;
}

func (h *Handler) createGame(c fiber.Ctx) error {
	return nil;
}

func (h *Handler) editGame(c fiber.Ctx) error {
	return nil;
}

func (h *Handler) deleteGameById(c fiber.Ctx) error {
	return nil;
}

func (h *Handler) addPlayerToGame(c fiber.Ctx) error {
	return nil;
}