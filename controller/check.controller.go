package controller

import (
	"oapi-codegen-cultibio/api"

	"github.com/gofiber/fiber/v2"
)

// checkController handles the logic for checking the application's status.
type checkController struct {
	// Add any necessary initialization for use cases or other components here.
}

// ICheckController defines the interface for the checkController.
type ICheckController interface {
	Execute(c *fiber.Ctx) error
}

// NewCheckController creates a new instance of checkController.
func NewCheckController() ICheckController {
	return &checkController{}
}

// Execute implements the ICheckController interface.
func (*checkController) Execute(c *fiber.Ctx) error {
	var apiRes api.GlobalResponses
	apiRes.ResponseCode = "200"
	apiRes.ResponseMessage = "Go is running!"
	return c.JSON(apiRes)
}
