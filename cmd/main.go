package main

import (
	"oapi-codegen-cultibio/api"
	"oapi-codegen-cultibio/bootstrap"
	"oapi-codegen-cultibio/cmd/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize the application's bootstrap components.
	app := bootstrap.NewInitializeBootsrap()

	// Create a new Fiber instance.
	f := fiber.New()

	// Serve the Swagger UI at the "/swagger" endpoint.
	f.Static("/swagger", "cmd")

	// Define the API versioned group.
	f.Group("/api/v1.0")

	// Initialize the application-specific handlers.
	serve := handlers.NewServiceInitial(app)
	checkController := serve.CheckHandler()

	// Create a wrapper for the controller interfaces.
	wrapper := &handlers.ServerInterfaceWrapper{
		CheckHandler: checkController,
	}

	// Register API handlers.
	api.RegisterHandlers(f, wrapper)

	// Start the server on port 3000.
	f.Listen(":3000")
}
