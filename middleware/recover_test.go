package middleware

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRecover(t *testing.T) {
	t.Run("test recover middleware", func(t *testing.T) {
		errorMessage := "Error panic test"
		app := fiber.New()
		app.Use(Recover)
		app.Get("/", func(c *fiber.Ctx) error {
			panic(errorMessage)
		})
		request := httptest.NewRequest("GET", "/", nil)
		response, err := app.Test(request, -1)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
		body, err := io.ReadAll(response.Body)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "{\"status\":\"error\",\"message\":\""+errorMessage+"\"}", string(body))
	})
}
