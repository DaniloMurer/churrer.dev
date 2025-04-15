package technologyapi

import (
	dto "github.com/DaniloMurer/churrer.xyz/internal/api/model"
	"github.com/DaniloMurer/churrer.xyz/internal/database"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	errorStringFormat = "ERROR: %+v\n"
)

var logger = log.New(os.Stdout, "[XYZ] - ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

// GetTechnologies godoc
// @Summary Get technologies
// @Description Gets technologies
// @Tags technology
// @Accept json
// @Produce json
// @Success 200 {array} dto.TechnologyDto "Technologies"
// @Failure 500 "Internal Server Error"
// @Router /api/technology [get]
func GetTechnologies(c *fiber.Ctx) error {
	technologies := database.GetAllTechnology()
	return c.Status(http.StatusOK).JSON(technologies)
}

// CreateTechnology godoc
// @Summary Create technology
// @Description Creates technologies
// @Tags technology
// @Accept json
// @Produce json
// @Param technology body dto.TechnologyDto true "Technology"
// @Param Authorization header string true "BasicAuth token"
// @Success 201 {object} dto.TechnologyDto "Created technology"
// @Failure 500 "Internal Server Error"
// @Router /api/technology [post]
func CreateTechnology(c *fiber.Ctx) error {
	var newTechnology dto.TechnologyDto
	if err := c.BodyParser(&newTechnology); err != nil {
		logger.Printf(errorStringFormat, err)
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	database.CreateTechnology(newTechnology.ToTechnology())
	return c.Status(http.StatusCreated).JSON(newTechnology)
}

// UpdateTechnology godoc
// @Summary Update technology
// @Description Updates technologies
// @Tags technology
// @Accept json
// @Produce json
// @Param technology body dto.TechnologyDto true "Technology"
// @Param Authorization header string true "BasicAuth token"
// @Success 200 {object} dto.TechnologyDto "Updated technology"
// @Failure 500 "Internal Server Error"
// @Router /api/technology [put]
func UpdateTechnology(c *fiber.Ctx) error {
	var technology dto.TechnologyDto
	if err := c.BodyParser(&technology); err != nil {
		logger.Printf(errorStringFormat, err)
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	database.UpdateTechnology(technology.ToTechnology())
	return c.Status(http.StatusOK).JSON(technology)
}

// DeleteTechnology godoc
// @Summary Delete technology
// @Description Deletes technology
// @Tags technology
// @Accept json
// @Produce json
// @Param id path number true "Technology id"
// @Param Authorization header string true "BasicAuth token"
// @Success 200 {object} dto.ResponseDto "Success message"
// @Failure 500 "Internal Server Error"
// @Router /api/technology [delete]
func DeleteTechnology(c *fiber.Ctx) error {
	technologyId := c.Params("id")

	// convert experienceId to int
	if convertedId, err := strconv.Atoi(technologyId); err != nil {
		logger.Printf(errorStringFormat, err)
		return fiber.NewError(http.StatusInternalServerError, "Provided id not a valid integer")
	} else {
		database.DeleteTechnology(uint(convertedId))
		return c.Status(http.StatusOK).JSON(dto.ResponseDto{Message: "Technology deleted"})
	}
}
