//go:build prod
// +build prod

package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"server/api"
	"server/database"
)

//go:embed spa spa/_nuxt
var spaFs embed.FS

// getSpa loads the embedded filesystem from memory
func getSpa() http.FileSystem {
	spa, err := fs.Sub(spaFs, "spa")
	if err != nil {
		panic(err)
	}
	return http.FS(spa)
}

func main() {
	fmt.Println("Running Server in PROD Mode")
	dbErr := database.AutoMigration()
	if dbErr != nil {
		return
	}
	router := api.BuildRouter(getSpa())
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		panic("we're fucked")
	}
}
