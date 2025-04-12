package entities

import (
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

// Experience represents a database model for work experience
type Experience struct {
	Model
	Company          string `json:"company"`
	Position         string `json:"position"`
	TimeFrame        string `json:"timeFrame"`
	Responsibilities string `json:"responsibilities"`
}

// Technology represents a database model for known technologies
type Technology struct {
	Model
	Name        string `json:"name"`
	Experience  string `json:"experience"`
	Description string `json:"description"`
	LogoClass   string `json:"logoClass"`
}
