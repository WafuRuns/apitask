package main

import "github.com/gofiber/fiber/v2"

// Response if client entered invalid data
func clientError(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
		"status":  fiber.StatusBadRequest,
		"success": false,
		"error":   message,
	})
}

// Response if server couldn't process request
func serverError(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
		"status":  fiber.StatusInternalServerError,
		"success": false,
		"error":   message,
	})
}
