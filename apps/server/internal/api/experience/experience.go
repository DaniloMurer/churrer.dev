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

// GetExperiences godoc
// @Summary Get experiences
// @Description Gets experiences
// @Tags experience
// @Accept json
// @Produce json
// @Success 200 {array} dto.ExperienceDto "Experiences"
// @Failure 500 "Internal Server Error"
// @Router /api/experience [get]
func GetExperiences(c *fiber.Ctx) error {
	experiences := database.GetAllExperience()
	return c.Status(http.StatusOK).JSON(experiences)
}

// CreateExperience godoc
// @Summary Create experience
// @Description Creates experience
// @Tags experience
// @Accept json
// @Produce json
// @Param user body dto.UserDto true "Experience"
// @Param Authorization header string true "BasicAuth token"
// @Success 201 {object} dto.ExperienceDto "Created experience"
// @Failure 500 "Internal Server Error"
// @Router /api/experience [post]
func CreateExperience(c *fiber.Ctx) error {
	var newExperience dto.ExperienceDto
	if err := c.BodyParser(&newExperience); err != nil {
		logger.Printf(errorStringFormat, err)
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	database.CreateExperience(newExperience.ToExperience())
	return c.Status(http.StatusCreated).JSON(newExperience)
}

// UpdateExperience godoc
// @Summary Update experiences
// @Description Updates experience
// @Tags experience
// @Accept json
// @Produce json
// @Param user body dto.ExperienceDto true "Experience"
// @Param Authorization header string true "BasicAuth token"
// @Success 200 {object} dto.ExperienceDto "Updated experience"
// @Failure 500 "Internal Server Error"
// @Router /api/experience [put]
func UpdateExperience(c *fiber.Ctx) error {
	var experience dto.ExperienceDto
	if err := c.BodyParser(&experience); err != nil {
		logger.Printf(errorStringFormat, err)
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	database.UpdateExperience(experience.ToExperience())
	return c.Status(http.StatusOK).JSON(experience)
}

// DeleteExperience godoc
// @Summary Delete experiences
// @Description Deletes experience
// @Tags experience
// @Accept json
// @Produce json
// @Param id path number true "Experience id"
// @Param Authorization header string true "BasicAuth token"
// @Success 200 {object} dto.ResponseDto "Success message"
// @Failure 500 "Internal Server Error"
// @Router /api/experience [delete]
func DeleteExperience(c *fiber.Ctx) error {
	experienceId := c.Params("id")

	// convert experienceId to int
	if convertedId, err := strconv.Atoi(experienceId); err != nil {
		logger.Printf(errorStringFormat, err)
		c.Status(http.StatusInternalServerError)
	} else {
		database.DeleteExperience(uint(convertedId))
		return c.Status(http.StatusOK).JSON(dto.ResponseDto{Message: "Experience deleted"})
	}
	return nil
}
