package database

import (
	"github.com/DaniloMurer/churrer.xyz/internal/database/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	connectionString = os.Getenv("DB_CONNECTION")
	database         *gorm.DB
)

var logger = log.New(os.Stdout, "[XYZ] - ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

// openDatabaseConnection creates a postgresql connection
func openDatabaseConnection() {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		logger.Fatal("Error while opening database connection", err.Error())
	} else {
		database = db
	}
}

// AutoMigration migrates the schema defined in the data package to the database
func AutoMigration() error {
	openDatabaseConnection()
	return database.AutoMigrate(&entities.Telemetry{}, &entities.Experience{}, &entities.Technology{})
}

// GetAllTelemetry returns all telemetries from the database
func GetAllTelemetry() []entities.Telemetry {
	logger.Println("Retrieving all telemetries")
	var telemetries []entities.Telemetry
	database.Find(&telemetries)
	return telemetries
}

// CreateTelemetry creates new data.Telemetry entry in database
func CreateTelemetry(telemetry *entities.Telemetry) {
	logger.Println("Creating new telemetry entry")
	database.Create(telemetry)
}

// GetAllExperience returns all experiences from the database
func GetAllExperience() []entities.Experience {
	logger.Println("Retrieving all experiences")
	var experiences []entities.Experience
	database.Find(&experiences)
	return experiences
}

// CreateExperience creates new data.Experience entry in database
func CreateExperience(experience *entities.Experience) {
	logger.Println("Creating new experience entry")
	database.Create(experience)
}

func DeleteExperience(id uint) {
	logger.Println("Deleting experience entry")
	database.Delete(&entities.Experience{}, id)
}

func UpdateExperience(experience *entities.Experience) {
	logger.Println("Updating experience entry")
	database.Save(experience)
}

// GetAllTechnology return all technologies from the database
func GetAllTechnology() []entities.Technology {
	logger.Println("Retrieving all technologies")
	var technologies []entities.Technology
	database.Find(&technologies)
	return technologies
}

// CreateTechnology creates new data.Technology entry in database
func CreateTechnology(technology *entities.Technology) {
	logger.Println("Creating new technology entry")
	database.Create(technology)
}

func DeleteTechnology(id uint) {
	logger.Println("Deleting technology")
	database.Delete(&entities.Technology{}, id)
}

func UpdateTechnology(technology *entities.Technology) {
	logger.Println("Updating technology")
	database.Save(technology)
}
