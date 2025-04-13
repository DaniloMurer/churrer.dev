package authenticationapi

import (
	"github.com/DaniloMurer/churrer.xyz/internal/api/middleware"
	"github.com/DaniloMurer/churrer.xyz/internal/api/model"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"os"
)

const (
	errorStringFormat = "ERROR: %+v\n"
)

var logger = log.New(os.Stdout, "[XYZ] - ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

func AuthenticateUser(c *fiber.Ctx) error {
	var user dto.UserDto
	if err := c.BodyParser(&user); err != nil {
		logger.Printf(errorStringFormat, err)
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}
	authorized, err := middleware.Authorize(user.Username, user.Password, true)
	if err != nil {
		logger.Printf(errorStringFormat, err)
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	if !authorized {
		logger.Println("Invalid credentials")
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
	}
	user.CreateToken()
	return c.Status(http.StatusAccepted).JSON(user)
}
