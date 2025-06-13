package middleware

import (
	"fmt"
	"golang-template/app"

	"github.com/gofiber/fiber/v2"
)

func Recover(c *fiber.Ctx) error {
	defer func() {
		if err := recover(); err != nil {
			// Handle the error and respond with an error message
			c.Status(fiber.StatusInternalServerError).JSON(app.Response{
				Status:  "error",
				Message: fmt.Sprintf("%v", err),
			})
		}
	}()
	return c.Next()
}
