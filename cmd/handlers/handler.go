package handlers

import (
	"fmt"
	"oapi-codegen-cultibio/api"
	"oapi-codegen-cultibio/bootstrap"
	"oapi-codegen-cultibio/controller"

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
}

// AddSensorData implements api.ServerInterface.
func (h *ServerInterfaceWrapper) AddSensorData(c *fiber.Ctx) error {
	// Parse the request body into a SensorData object
	var requestData api.SensorData
	if err := c.BodyParser(&requestData); err != nil {
		// Handle parsing error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Do something with the parsed data
	// For example, print the received values
	fmt.Printf("Received SensorID: %v, Timestamp: %v, Value: %v\n",
		requestData.SensorID, requestData.Timestamp, requestData.Value)

	// You can now use the requestData values as needed in your application logic

	// Return a response if necessary
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Sensor data received successfully",
		"data":    requestData,
	})
}

// Check implements api.ServerInterface.
func (h *ServerInterfaceWrapper) Check(c *fiber.Ctx) error {
	return h.CheckHandler.Execute(c)
}
