package handlers

import (
	"golang-template/app"
	"golang-template/app/models"
	"golang-template/app/services"
	"golang-template/validator"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	Register(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
}

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{userService: userService}
}

func RegisterUserRoutes(route fiber.Router, handler UserHandler) {
	route.Post("/register", handler.Register)
	route.Put("/update", handler.Update)
	route.Get("/list", handler.List)
}

func (h *userHandler) Register(c *fiber.Ctx) error {
	var newUser models.UserRegister
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(app.NewResponseError(err))
	}

	if err := validator.ValidateStruct(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(app.NewResponseError(err))
	}

	err := h.userService.Register(&newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(app.NewResponseError(err))
	}

	return c.JSON(app.NewResponse("User registered successfully", nil))
}

func (h *userHandler) Update(c *fiber.Ctx) error {
	var userUpdate models.UserUpdatePassword
	if err := c.BodyParser(&userUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(app.NewResponseError(err))
	}

	if err := validator.ValidateStruct(&userUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(app.NewResponseError(err))
	}

	err := h.userService.Update(&userUpdate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(app.NewResponseError(err))
	}

	return c.JSON(app.NewResponse("User updated successfully", nil))
}

func (h *userHandler) List(c *fiber.Ctx) error {
	users, err := h.userService.List()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(app.NewResponseError(err))
	}

	return c.JSON(app.NewResponse("Users listed successfully", &users))
}
