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

func GetTechnologies(c *fiber.Ctx) error {
	technologies := database.GetAllTechnology()
	return c.Status(http.StatusOK).JSON(technologies)
}

func CreateTechnology(c *fiber.Ctx) error {
	var newTechnology dto.TechnologyDto
	if err := c.BodyParser(&newTechnology); err != nil {
		logger.Printf(errorStringFormat, err)
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	database.CreateTechnology(newTechnology.ToTechnology())
	return c.Status(http.StatusCreated).JSON(newTechnology)
}

func UpdateTechnology(c *fiber.Ctx) error {
	var technology dto.TechnologyDto
	if err := c.BodyParser(&technology); err != nil {
		logger.Printf(errorStringFormat, err)
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	database.UpdateTechnology(technology.ToTechnology())
	return c.Status(http.StatusOK).JSON(technology)
}

func DeleteTechnology(c *fiber.Ctx) error {
	technologyId := c.Params("id")

	// convert experienceId to int
	if convertedId, err := strconv.Atoi(technologyId); err != nil {
		logger.Printf(errorStringFormat, err)
		return fiber.NewError(http.StatusInternalServerError, "Provided id not a valid integer")
	} else {
		database.DeleteTechnology(uint(convertedId))
		return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Technology deleted"})
	}
}
