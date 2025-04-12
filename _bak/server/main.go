//go:build !prod

package main

import (
	"fmt"
	"server/api"
	"server/database"
)

func main() {
	fmt.Println("Running Server in DEV Mode")
	dbErr := database.AutoMigration()
	if dbErr != nil {
		return
	}
	router := api.BuildRouter(nil)
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		panic("we're fucked")
	}
}
