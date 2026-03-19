package auth

import (
	"github.com/gofiber/fiber/v3"
)

func Handler(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/login", login)
	auth.Post("/signup", signup)
	auth.Post("/logout", logout)
	auth.Get("/activate/:link", activate)
	auth.Get("/refresh", refresh)

}

func login(c fiber.Ctx) error {
	return nil;
}

func signup(c fiber.Ctx) error {
	return nil;
}

func logout(c fiber.Ctx) error {
	return nil;
}

func activate(c fiber.Ctx) error {
	return nil;
}

func refresh(c fiber.Ctx) error {
	return nil;
}