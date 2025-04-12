package dto

import (
	"encoding/base64"
	"github.com/DaniloMurer/churrer.xyz/internal/database/model"
	"time"
)

// TelemetryDto represents the telemetry data transfer object for telemetry information
type TelemetryDto struct {
	CountryName string    `json:"countryName"`
	CountryISO  string    `json:"countryIso"`
	TimeStamp   time.Time `json:"timestamp"`
}

// ToTelemetry converts a TelemetryDto model to a Telemetry database model
func (dto TelemetryDto) ToTelemetry() *entities.Telemetry {
	return &entities.Telemetry{
		CountryName: dto.CountryName,
		CountryISO:  dto.CountryISO,
		TimeStamp:   dto.TimeStamp,
	}
}

// ExperienceDto represents a data transfer object for work experience information.
type ExperienceDto struct {
	Company          string `json:"company"`
	Position         string `json:"position"`
	TimeFrame        string `json:"timeFrame"`
	Responsibilities string `json:"responsibilities"`
}

// ToExperience converts an ExperienceDto model to an Experience database model
func (dto ExperienceDto) ToExperience() *entities.Experience {
	return &entities.Experience{
		Company:          dto.Company,
		Position:         dto.Position,
		TimeFrame:        dto.TimeFrame,
		Responsibilities: dto.Responsibilities,
	}
}

// TechnologyDto represents a data transfer object for known technology.
type TechnologyDto struct {
	Name        string `json:"name"`
	Experience  string `json:"experience"`
	Description string `json:"description"`
	LogoClass   string `json:"logoClass"`
}

// ToTechnology converts an TechnologyDto model to a Technology database model
func (dto TechnologyDto) ToTechnology() *entities.Technology {
	return &entities.Technology{
		Name:        dto.Name,
		Experience:  dto.Experience,
		Description: dto.Description,
		LogoClass:   dto.LogoClass,
	}
}

// UserDto represents a data transfer object of a login request
type UserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// CreateToken generates an authentication token by encoding the username and password and fills the Token field
func (dto *UserDto) CreateToken() {
	dto.Token = base64.StdEncoding.EncodeToString([]byte(dto.Username + ":" + dto.Password))
}
