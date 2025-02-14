package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/data"
	"server/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	warningAttemptedWrongCredentials = "WARN: Attempted login with wrong credentials"
	errorStringFormat                = "ERROR: %+v\n"
	errorMessageInvalidCredentials   = "Invalid credentials"
)

var (
	adminUsername = os.Getenv("ADMIN_USERNAME")
	adminPassword = os.Getenv("ADMIN_PASSWORD")
)

var logger = log.New(os.Stdout, "[XYZ] - ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

// Authorize checks if the provided username and password match the admin credentials.
// Returns true if authorized, otherwise false. If there is an error, returns an error object.
func Authorize(username string, password string, ok bool) (bool, error) {
	if !ok {
		logger.Println("WARN: Non compliant BasicAuth header received")
		return false, fmt.Errorf("non compliant BasicAuth Header found")
	}
	if username != adminUsername || password != adminPassword {
		return false, nil
	}
	return true, nil
}

// AuthenticateUser handles the post request to authenticate a login request
func AuthenticateUser(c *gin.Context) {
	var user data.UserDto
	if err := c.BindJSON(&user); err != nil {
		logger.Printf(errorStringFormat, err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	authorized, err := Authorize(user.Username, user.Password, true)
	if err != nil {
		logger.Printf(errorStringFormat, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if !authorized {
		logger.Println(warningAttemptedWrongCredentials)
		c.JSON(http.StatusUnauthorized, gin.H{"message": errorMessageInvalidCredentials})
		return
	}
	user.CreateToken()
	c.JSON(http.StatusAccepted, user)
}

// GetTelemetries handles the get request to retrieve all telemetries
func GetTelemetries(c *gin.Context) {
	telemetries := database.GetAllTelemetry()
	c.JSON(http.StatusOK, telemetries)
}

// CreateTelemetry handles a post request to create a new telemetry
func CreateTelemetry(c *gin.Context) {
	var newTelemetry data.TelemetryDto
	if err := c.BindJSON(&newTelemetry); err != nil {
		logger.Printf(errorStringFormat, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	database.CreateTelemetry(newTelemetry.ToTelemetry())
	c.JSON(http.StatusCreated, newTelemetry)
}

// GetExperiences handles the get request to retrieve all experiences
func GetExperiences(c *gin.Context) {
	experiences := database.GetAllExperience()
	c.JSON(http.StatusOK, experiences)
}

// CreateExperience handles a post request to create a new experience
func CreateExperience(c *gin.Context) {
	authorized, err := Authorize(c.Request.BasicAuth())
	if err != nil {
		logger.Printf(errorStringFormat, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if !authorized {
		logger.Println(warningAttemptedWrongCredentials)
		c.JSON(http.StatusUnauthorized, gin.H{"message": errorMessageInvalidCredentials})
		return
	}

	var newExperience data.ExperienceDto
	if err := c.BindJSON(&newExperience); err != nil {
		logger.Printf(errorStringFormat, err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	database.CreateExperience(newExperience.ToExperience())
	c.JSON(http.StatusCreated, newExperience)
}

func DeleteExperience(c *gin.Context) {
	authorized, err := Authorize(c.Request.BasicAuth())
	if err != nil {
		logger.Printf(errorStringFormat, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if !authorized {
		logger.Println(warningAttemptedWrongCredentials)
		c.JSON(http.StatusUnauthorized, gin.H{"message": errorMessageInvalidCredentials})
		return
	}
	experienceId := c.Param("id")

	// convert experienceId to int
	if convertedId, err := strconv.Atoi(experienceId); err != nil {
		logger.Printf(errorStringFormat, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Provided id not a valid integer"})
	} else {
		database.DeleteExperience(uint(convertedId))
		c.JSON(http.StatusOK, gin.H{"message": "Experience deleted"})
	}
}

func UpdateExperience(c *gin.Context) {
	authorized, err := Authorize(c.Request.BasicAuth())
	if err != nil {
		logger.Printf(errorStringFormat, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if !authorized {
		logger.Println(warningAttemptedWrongCredentials)
		c.JSON(http.StatusUnauthorized, gin.H{"message": errorMessageInvalidCredentials})
		return
	}

	var experience data.ExperienceDto
	if err := c.BindJSON(&experience); err != nil {
		logger.Printf(errorStringFormat, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	database.UpdateExperience(experience.ToExperience())
	c.JSON(http.StatusOK, gin.H{"message": experience})
}

// GetTechnologies handles the get request to retrieve all technologies
func GetTechnologies(c *gin.Context) {
	technologies := database.GetAllTechnology()
	c.JSON(http.StatusOK, technologies)
}

// CreateTechnology handles the post request to create a new technology
func CreateTechnology(c *gin.Context) {
	authorized, err := Authorize(c.Request.BasicAuth())
	if err != nil {
		logger.Printf(errorStringFormat, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	if !authorized {
		logger.Println(warningAttemptedWrongCredentials)
		c.JSON(http.StatusUnauthorized, gin.H{"message": errorMessageInvalidCredentials})
		return
	}

	var newTechnology data.TechnologyDto
	if err := c.BindJSON(&newTechnology); err != nil {
		logger.Printf(errorStringFormat, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	database.CreateTechnology(newTechnology.ToTechnology())
	c.JSON(http.StatusCreated, newTechnology)
}

func DeleteTechnology(c *gin.Context) {
	logger.Println("deleting technology....")
	authorized, err := Authorize(c.Request.BasicAuth())
	if err != nil {
		logger.Printf(errorStringFormat, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if !authorized {
		logger.Println(warningAttemptedWrongCredentials)
		c.JSON(http.StatusUnauthorized, gin.H{"message": errorMessageInvalidCredentials})
		return
	}
	technologyId := c.Param("id")

	// convert experienceId to int
	if convertedId, err := strconv.Atoi(technologyId); err != nil {
		logger.Printf(errorStringFormat, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Provided id not a valid integer"})
	} else {
		database.DeleteTechnology(uint(convertedId))
		c.JSON(http.StatusOK, gin.H{"message": "Experience deleted"})
	}
}

func UpdateTechnology(c *gin.Context) {
	authorized, err := Authorize(c.Request.BasicAuth())
	if err != nil {
		logger.Printf(errorStringFormat, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if !authorized {
		logger.Println(warningAttemptedWrongCredentials)
		c.JSON(http.StatusUnauthorized, gin.H{"message": errorMessageInvalidCredentials})
		return
	}

	var technology data.TechnologyDto
	if err := c.BindJSON(&technology); err != nil {
		logger.Printf(errorStringFormat, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	database.UpdateTechnology(technology.ToTechnology())
	c.JSON(http.StatusOK, gin.H{"message": technology})
}
