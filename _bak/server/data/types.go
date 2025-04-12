package data

import (
	"encoding/base64"
	"time"

	"gorm.io/gorm"
)

// Model base struct for database models
type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Telemetry represents telemetry data for country tracking purposes.
type Telemetry struct {
	Model
	CountryName string    `json:"countryName"`
	CountryISO  string    `json:"countryIso"`
	TimeStamp   time.Time `json:"timestamp"`
}

// TelemetryDto represents the telemetry data transfer object for telemetry information
type TelemetryDto struct {
	CountryName string    `json:"countryName"`
	CountryISO  string    `json:"countryIso"`
	TimeStamp   time.Time `json:"timestamp"`
}

// ToTelemetry converts a TelemetryDto model to a Telemetry database model
func (dto TelemetryDto) ToTelemetry() *Telemetry {
	return &Telemetry{
		CountryName: dto.CountryName,
		CountryISO:  dto.CountryISO,
		TimeStamp:   dto.TimeStamp,
	}
}

// Experience represents a database model for work experience
type Experience struct {
	Model
	Company          string `json:"company"`
	Position         string `json:"position"`
	TimeFrame        string `json:"timeFrame"`
	Responsibilities string `json:"responsibilities"`
}

// ExperienceDto represents a data transfer object for work experience information.
type ExperienceDto struct {
	Company          string `json:"company"`
	Position         string `json:"position"`
	TimeFrame        string `json:"timeFrame"`
	Responsibilities string `json:"responsibilities"`
}

// ToExperience converts an ExperienceDto model to an Experience database model
func (dto ExperienceDto) ToExperience() *Experience {
	return &Experience{
		Company:          dto.Company,
		Position:         dto.Position,
		TimeFrame:        dto.TimeFrame,
		Responsibilities: dto.Responsibilities,
	}
}

// Technology represents a database model for known technologies
type Technology struct {
	Model
	Name        string `json:"name"`
	Experience  string `json:"experience"`
	Description string `json:"description"`
	LogoClass   string `json:"logoClass"`
}

// TechnologyDto represents a data transfer object for known technology.
type TechnologyDto struct {
	Name        string `json:"name"`
	Experience  string `json:"experience"`
	Description string `json:"description"`
	LogoClass   string `json:"logoClass"`
}

// ToTechnology converts an TechnologyDto model to an Technology database model
func (dto TechnologyDto) ToTechnology() *Technology {
	return &Technology{
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
