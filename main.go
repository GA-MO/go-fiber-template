package main

import (
	"golang-template/app/handlers"
	"golang-template/app/repositories"
	"golang-template/app/services"
	"golang-template/database"
	"golang-template/logger"
	"golang-template/middleware"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	logger := logger.NewLogger()

	db, err := database.GetSQLite()
	if err != nil {
		logger.Fatal(err)
	}

	// Create new Fiber app
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(compress.New())
	app.Use(healthcheck.New())
	app.Use(middleware.Recover)
	app.Use(middleware.NewRequestLog(logger))
	app.Use(middleware.NewResponseLog(logger))

	api := app.Group("/api")

	// Register routes
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	handlers.RegisterUserRoutes(api.Group("/v1/user"), userHandler)

	// Start server
	app.Listen(":9090")
}
