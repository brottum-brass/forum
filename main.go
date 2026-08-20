package main

import (
	"log"

	"github.com/brottum-brass/forum/src/api"
	"github.com/brottum-brass/forum/src/api/middleware"
	"github.com/brottum-brass/forum/src/api/router"
	"github.com/brottum-brass/forum/src/utils"
)

func main() {
	appCtx := utils.GetAppContext()

	counter := middleware.NewCounter()
	logger := middleware.NewLogger()
	preference := middleware.NewPreference()
	router := router.NewRouter()

	mux := counter.Next(logger.Next(preference.Next(router)))

	server := api.NewServer(appCtx.Config.Port, mux)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
