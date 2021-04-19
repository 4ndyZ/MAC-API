package main

import (
	"context"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// App struct to hold refs and database info
type App struct {
	router *mux.Router
	url    *string
	address   *string
	refresh *int
	Data   *[]OUI
}

// Initialize app struct with database info
func (a *App) Initialize(url string, address string, refresh int) {
	a.url = &url
	a.address = &address
	a.refresh = &refresh

	Log.Logger.Info().Msg("API running ...")

	a.Refresh()

	a.router = mux.NewRouter()
	a.router.HandleFunc("/v1/oui/{oui}", a.GetOUI).Methods("GET")
	a.router.HandleFunc("/v1/mac/{mac}", a.GetMAC).Methods("GET")
	a.router.NotFoundHandler = http.HandlerFunc(a.NotFound)
}

func (a *App) Refresh() {
	Log.Logger.Info().Msg("Starting data refresh ... ")
	// Get MAC data
	ouisdata, err := a.GetMACData(*a.url)
	if err != nil {
		Log.Logger.Warn().Str("error", err.Error()).Msg("Error while getting the data from the IEEE.")
	}
	// Parse MAC data
	ouis, err := a.ParseMACData(&ouisdata)
	if err != nil {
		Log.Logger.Warn().Str("error", err.Error()).Msg("Error while parsing the data from the IEEE.")
	}
	a.Data = &ouis
	Log.Logger.Info().Msg("Finshed data refresh.")
}

func (a *App) Run() {

	ctx, cancel := context.WithCancel(context.Background())

	httpServer := &http.Server{
		Addr:        *a.address,
		Handler:     a.router,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}

	// Run server
	go func() {
		Log.Logger.Info().Msg("API Webserver running ...")
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			// it is fine to use Fatal here because it is not main gorutine
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

		gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelShutdown()

		if err := httpServer.Shutdown(gracefullCtx); err != nil {
			Log.Logger.Warn().Str("error", err.Error()).Msg("Error while stopping API Webserver.")
			return
		} else {
			Log.Logger.Info().Msg("API Webserver stopped.")
		}

		cancel()

		Log.Logger.Info().Msg("API stopped.")

		Log.Rotate()
		os.Exit(1)
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
