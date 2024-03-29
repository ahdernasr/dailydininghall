package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ahdernasr/dailydininghall/internal/db"
	"github.com/ahdernasr/dailydininghall/internal/routes"

	// "github.com/ahdernasr/dailydininghall/internal/scraper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Setup routes
	routes.SetupRoutes(app)

	// Load connection string frome .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file could not be loaded.")
	}

	connectionString := os.Getenv("CONNECTION_STRING")

	// Connect to db
	if err := db.Connect(connectionString); err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	} else {
		fmt.Println("Connected to db!")
	}

	log.Fatal(app.Listen(":4000"))
}
