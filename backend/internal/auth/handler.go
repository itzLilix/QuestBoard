package auth

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
	auth := app.Group("/auth")
	auth.Post("/login", h.login)
	auth.Post("/signup", h.signup)
	auth.Post("/logout", h.logout)
	auth.Get("/activate/:link", h.activate)
	auth.Get("/refresh", h.refresh)

}

func (h *Handler) login(c fiber.Ctx) error {
	return nil;
}

func (h *Handler) signup(c fiber.Ctx) error {
	return nil;
}

func (h *Handler) logout(c fiber.Ctx) error {
	return nil;
}

func (h *Handler) activate(c fiber.Ctx) error {
	return nil;
}

func (h *Handler) refresh(c fiber.Ctx) error {
	return nil;
}