package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// App struct to hold refs and database info
type App struct {
	server  *fiber.App
	url     *string
	address *string
	refresh *int
	Data    *[]OUI
}

// Initialize app struct with database info
func (a *App) Initialize(url string, address string, refresh int) {
	a.url = &url
	a.address = &address
	a.refresh = &refresh

	Log.Logger.Info().Msg("API running ...")

	a.Refresh()

	config := fiber.Config{
		ServerHeader:          "MAC-API",
		DisableStartupMessage: true,
	}

	a.server = fiber.New(config)
	a.server.Get("/v1/oui/:oui", a.GetOUI)
	a.server.Get("/v1/mac/:mac", a.GetMAC)
	a.server.All("*", a.NotFound)
}

func (a *App) Refresh() {
	Log.Logger.Info().Msg("Starting data refresh ... ")
	// Get MAC data
	ouisData, err := a.GetMACData(*a.url)
	if err != nil {
		Log.Logger.Warn().Str("error", err.Error()).Msg("Error while getting the data from the IEEE.")
	}
	// Parse MAC data
	ouis, err := a.ParseMACData(&ouisData)
	if err != nil {
		Log.Logger.Warn().Str("error", err.Error()).Msg("Error while parsing the data from the IEEE.")
	}
	a.Data = &ouis
	Log.Logger.Info().Msg("Finished data refresh.")
}

func (a *App) Run() {
	// Run server
	go func() {
		Log.Logger.Info().Msg("API Webserver running ...")
		err := a.server.Listen(*a.address)
		if err != nil {
			// It is fine to use Fatal here because it is not main goroutine
			Log.Logger.Fatal().Str("error", err.Error()).Msg("API Webserver server error.")
		}
	}()

	// Setup signal catching
	sigs := make(chan os.Signal, 1)
	// Catch all signals since not explicitly listing
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGHUP)
	// Method invoked upon seeing signal
	go func() {
		s := <-sigs
		Log.Logger.Info().Str("reason", s.String()).Msg("API shutting down ...")
		err := a.server.ShutdownWithTimeout(5 * time.Second)
		if err != nil {
			Log.Logger.Warn().Str("error", err.Error()).Msg("Error while stopping API Webserver.")
			return
		}
		Log.Logger.Info().Msg("API Webserver stopped.")
		Log.Logger.Info().Msg("API stopped.")
		// Exit
		if err != nil {
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}()

	for {
		// Wait the provided time to before refreshing the data
		d := time.Second * time.Duration(*a.refresh)
		Log.Logger.Info().Interface("duration", d).Msg("Waiting for data refresh.")
		time.Sleep(d)
		// Do data refresh
		a.Refresh()
	}
}
