package middleware

import (
	"encoding/json"
	"golang-template/logger"

	"github.com/gofiber/fiber/v2"
)

func getJsonBody(body string) map[string]interface{} {
	bodyMap := map[string]interface{}{}
	_ = json.Unmarshal([]byte(body), &bodyMap)
	return bodyMap
}

func NewRequestLog(logger logger.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := getJsonBody(string(c.Body()))

		logger.Request(map[string]interface{}{
			"request_id":   c.Locals("requestid"),
			"method":       c.Method(),
			"path":         c.Path(),
			"body":         body,
			"ip":           c.IP(),
			"user_agent":   c.Get("User-Agent"),
			"referer":      c.Get("Referer"),
			"host":         c.Hostname(),
			"protocol":     c.Protocol(),
			"headers":      c.Get("User-Agent"),
			"query":        c.Query("query"),
			"params":       c.Params("params"),
			"cookies":      c.Cookies("cookies"),
			"session":      c.Locals("session"),
			"session_id":   c.Locals("session_id"),
			"session_data": c.Locals("session_data"),
		})
		return c.Next()
	}
}

func NewResponseLog(logger logger.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		reponseData := getJsonBody(string(c.Response().Body()))

		logger.Response(map[string]interface{}{
			"request_id": c.Locals("requestid"),
			"status":     c.Response().StatusCode(),
			"method":     c.Method(),
			"path":       c.Path(),
			"ip":         c.IP(),
			"host":       c.Hostname(),
			"response":   reponseData,
		})

		return err
	}
}
