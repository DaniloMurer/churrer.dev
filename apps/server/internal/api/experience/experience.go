package experienceapi

import (
	"github.com/DaniloMurer/churrer.xyz/internal/api/model"
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

func GetExperiences(c *fiber.Ctx) error {
	experiences := database.GetAllExperience()
	return c.Status(http.StatusOK).JSON(experiences)
}

func CreateExperience(c *fiber.Ctx) error {
	var newExperience dto.ExperienceDto
	if err := c.BodyParser(&newExperience); err != nil {
		logger.Printf(errorStringFormat, err)
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	database.CreateExperience(newExperience.ToExperience())
	return c.Status(http.StatusCreated).JSON(newExperience)
}

func UpdateExperience(c *fiber.Ctx) error {
	var experience dto.ExperienceDto
	if err := c.BodyParser(&experience); err != nil {
		logger.Printf(errorStringFormat, err)
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	database.UpdateExperience(experience.ToExperience())
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": experience})
}

func DeleteExperience(c *fiber.Ctx) error {
	experienceId := c.Params("id")

	// convert experienceId to int
	if convertedId, err := strconv.Atoi(experienceId); err != nil {
		logger.Printf(errorStringFormat, err)
		c.Status(http.StatusInternalServerError)
	} else {
		database.DeleteExperience(uint(convertedId))
		return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Experience deleted"})
	}
	return nil
}
