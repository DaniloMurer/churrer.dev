package telemetryapi

import (
	"github.com/DaniloMurer/churrer.xyz/internal/api/model"
	"github.com/DaniloMurer/churrer.xyz/internal/database"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"os"
)

const (
	errorStringFormat = "ERROR: %+v\n"
)

var logger = log.New(os.Stdout, "[XYZ] - ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

func GetTelemetries(c *fiber.Ctx) error {
	telemetries := database.GetAllTelemetry()
	return c.Status(http.StatusOK).JSON(telemetries)
}

func CreateTelemetry(c *fiber.Ctx) error {
	var newTelemetry dto.TelemetryDto
	if err := c.BodyParser(&newTelemetry); err != nil {
		logger.Printf(errorStringFormat, err)
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	database.CreateTelemetry(newTelemetry.ToTelemetry())
	return c.Status(http.StatusCreated).JSON(newTelemetry)
}
