package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/samber/do/v2"

	"smartfarm-be/internal/adapter/inbound/web/farm"
	"smartfarm-be/internal/di"
)

func main() {
	setupLogger()

	injector, err := di.InitializeInjector()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize injector")
	}
	defer injector.Shutdown()

	app := do.MustInvoke[*fiber.App](injector)

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	registerRoutes(app, injector)

	log.Info().Msg("🚀 Server is starting...")
	go startServer(app)

	waitForSignal()

	if err := app.ShutdownWithTimeout(5 * time.Second); err != nil {
		log.Fatal().Err(err).Msg("Server failed to shutdown")
	}
}

func setupLogger() {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05",
		NoColor:    false,
	}
	log.Logger = zerolog.New(output).
		With().
		Timestamp().
		Str("service", "smartfarm-be").
		Caller().
		Logger().
		Level(zerolog.DebugLevel)
}

func startServer(app *fiber.App) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Info().Msgf("Server is listening on port %s", port)
	if err := app.Listen(":" + port); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("Server failed to start")
	}
}

func waitForSignal() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	log.Info().Msg("🛑 Server is shutting down...")
}

func registerRoutes(app *fiber.App, injector do.Injector) {
	// handlers
	farmHandler := do.MustInvoke[*farm.Handler](injector)

	// routes
	farmHandler.RegisterRoutes(app)
}
