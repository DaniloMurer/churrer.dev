package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"log"
	"os"
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

func Protected() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			adminUsername: adminPassword,
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusUnauthorized)
		},
	})
}
