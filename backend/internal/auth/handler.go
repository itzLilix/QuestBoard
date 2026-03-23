package auth

import (
	"github.com/gofiber/fiber/v3"
)

type Handler interface {
	RegisterRoutes(app *fiber.App)
}

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) RegisterRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/login", h.login)
	auth.Post("/signup", h.signup)
	auth.Post("/logout", h.logout)
	auth.Get("/activate/:link", h.activate)
	auth.Get("/refresh", h.refresh)
	auth.Get("/me", h.restoreSession)
}

func (h *handler) login(c fiber.Ctx) error {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var req request
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	user, token, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *handler) signup(c fiber.Ctx) error {
	type request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req request
	if err := c.Bind().Body(&req); err != nil {
        return err
    }
	user, token, err := h.service.Register(req.Username, req.Email, req.Password)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *handler) logout(c fiber.Ctx) error {
	c.ClearCookie("access_token")
	return c.SendStatus(fiber.StatusOK)
}

func (h *handler) activate(c fiber.Ctx) error {
	return nil;
}

func (h *handler) refresh(c fiber.Ctx) error {
	return nil;
}

func (h *handler) restoreSession(c fiber.Ctx) error {
	tokenString := c.Cookies("access_token")

	if tokenString == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	user, err := h.service.ValidateToken(tokenString)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(user)
}