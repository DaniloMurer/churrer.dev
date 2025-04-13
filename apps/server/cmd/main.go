package main

import (
	"github.com/DaniloMurer/churrer.xyz/internal/api/authentication"
	"github.com/DaniloMurer/churrer.xyz/internal/api/experience"
	"github.com/DaniloMurer/churrer.xyz/internal/api/middleware"
	"github.com/DaniloMurer/churrer.xyz/internal/api/technology"
	"github.com/DaniloMurer/churrer.xyz/internal/api/telemetry"
	"github.com/DaniloMurer/churrer.xyz/internal/database"
	"github.com/gofiber/fiber/v2"
	"log"
)

func createApp() *fiber.App {
	app := fiber.New()

	api := app.Group("/api")

	api.Get("/telemetry", telemetryapi.GetTelemetries)
	api.Post("/telemetry", telemetryapi.CreateTelemetry)

	api.Get("/experience", experienceapi.GetExperiences)
	api.Post("/experience", middleware.Protected(), experienceapi.CreateExperience)
	api.Put("/experience", middleware.Protected(), experienceapi.UpdateExperience)
	api.Delete("/experience/:id", middleware.Protected(), experienceapi.DeleteExperience)

	api.Get("/technology", technologyapi.GetTechnologies)
	api.Post("/technology", middleware.Protected(), technologyapi.CreateTechnology)
	api.Put("/technology", middleware.Protected(), technologyapi.UpdateTechnology)
	api.Delete("/technology/:id", middleware.Protected(), technologyapi.DeleteTechnology)

	api.Post("/authentication", authenticationapi.AuthenticateUser)

	return app
}

func main() {
	err := database.AutoMigration()
	if err != nil {
		log.Fatal("Error while migrating database", err)
	}
	app := createApp()

	log.Fatal(app.Listen("0.0.0.0:8080"))
}
