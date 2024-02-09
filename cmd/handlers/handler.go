package handlers

import (
	"oapi-codegen-cultibio/bootstrap"
	"oapi-codegen-cultibio/controller"
	"oapi-codegen-cultibio/dal"
	"oapi-codegen-cultibio/models"

	"github.com/gofiber/fiber/v2"
)

// MyHandler represents the handler struct.
type MyHandler struct {
	Application bootstrap.Application
}

// NewServiceInitial creates a new instance of MyHandler.
func NewServiceInitial(app bootstrap.Application) MyHandler {
	return MyHandler{
		Application: app,
	}
}

// ServerInterfaceWrapper wraps controller interfaces.
type ServerInterfaceWrapper struct {
	CheckHandler controller.ICheckController
	TimescaleDB  *dal.TimescaleDB
}

// AddSensorData implements api.ServerInterface.
func (h *ServerInterfaceWrapper) AddSensorData(c *fiber.Ctx) error {
	// Parse the request body into a SensorData object
	var requestData models.SensorData
	if err := c.BodyParser(&requestData); err != nil {
		// Handle parsing error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Store the received sensor data in TimescaleDB
	err := h.TimescaleDB.AddSensorData(requestData)
	if err != nil {
		// Handle database storage error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to store sensor data in the database",
		})
	}

	// You can now use the requestData values as needed in your application logic

	// Return a response if necessary
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Sensor data received and stored successfully",
		"data":    requestData,
	})
}

// Check implements api.ServerInterface.
func (h *ServerInterfaceWrapper) Check(c *fiber.Ctx) error {
	return h.CheckHandler.Execute(c)
}
