package main

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
)

func main(){

	// Initialize database
	db, err := InitDatabase()
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}
	defer db.Close()

	// Initialize Fiber app
	app := fiber.New()

	// Build routes
	BuildRoutes(app, db)

	// Start server
	if err := app.Listen(":3000"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
