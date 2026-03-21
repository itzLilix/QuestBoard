package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/itzLilix/QuestBoard/backend/internal/auth"
	"github.com/itzLilix/QuestBoard/backend/internal/games"
	"github.com/itzLilix/QuestBoard/backend/pkg/database"
	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	app := fiber.New()
	
	dbURL := os.Getenv("POSTGRES_URL")
	conn, err := database.Connect(dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	fmt.Println("Successfully connected to database")
	defer conn.Close(context.Background())

	err = database.RunMigrations(os.Getenv("MIGRATE_URL"))
	if err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}
	fmt.Println("Migrations ran successfully")

	authHandler := auth.NewHandler(conn)
	gamesHandler := games.NewHandler(conn)

	authHandler.RegisterRoutes(app)
	gamesHandler.RegisterRoutes(app)

	log.Fatal(app.Listen(":3000"))
}