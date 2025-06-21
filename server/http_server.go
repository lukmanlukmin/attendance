// Package server ...
package server

import (
	"attendance/bootstrap"
	"attendance/config"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
)

// HTTPApi ...
type HTTPApi struct {
	cfg *config.Config
}

// NewHTTPApi ...
func NewHTTPApi(cfg *config.Config) *HTTPApi {
	return &HTTPApi{
		cfg: cfg,
	}
}

// Run ...
func (h *HTTPApi) Run(ctx context.Context) {
	app := fiber.New(fiber.Config{})
	h.HTTPRouter(app, bootstrap.NewBootstrap(h.cfg), h.cfg)
	go func() {
		if err := app.Listen(h.cfg.Server.HTTPPort); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()
	waitForContextShutdown(ctx, app)
}

func waitForContextShutdown(ctx context.Context, app *fiber.App) {
	<-ctx.Done()
	log.Println("🛑 shutting down Http Server...")

	// timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("🛑 Http Server forced to shutdown: %v", err)
	}
	log.Println("✅ Http Server gracefully stopped.")
}
