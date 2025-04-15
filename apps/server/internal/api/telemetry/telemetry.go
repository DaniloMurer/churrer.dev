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

// GetTelemetries godoc
// @Summary Get telemetries
// @Description Gets telemetries
// @Tags telemetry
// @Accept json
// @Produce json
// @Success 200 {array} dto.TelemetryDto "Telemetries"
// @Failure 500 "Internal Server Error"
// @Router /api/telemetry [get]
func GetTelemetries(c *fiber.Ctx) error {
	telemetries := database.GetAllTelemetry()
	return c.Status(http.StatusOK).JSON(telemetries)
}

// CreateTelemetry godoc
// @Summary Create telemetry
// @Description Creates telemetry
// @Tags telemetry
// @Accept json
// @Produce json
// @Param user body dto.TelemetryDto true "Telemetry"
// @Success 201 {object} dto.ExperienceDto "Created telemetry"
// @Failure 500 "Internal Server Error"
// @Router /api/telemetry [post]
func CreateTelemetry(c *fiber.Ctx) error {
	var newTelemetry dto.TelemetryDto
	if err := c.BodyParser(&newTelemetry); err != nil {
		logger.Printf(errorStringFormat, err)
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	database.CreateTelemetry(newTelemetry.ToTelemetry())
	return c.Status(http.StatusCreated).JSON(newTelemetry)
}
