// Package server ...
package server

import (
	"attendance/bootstrap"
	"attendance/config"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// EventServer ...
type EventServer struct {
	cfg *config.Config
	ctx context.Context
}

// NewEventServer ...
func NewEventServer(ctx context.Context, cfg *config.Config) *EventServer {
	return &EventServer{ctx: ctx, cfg: cfg}
}

// Run ...
func (s *EventServer) Run(ctx context.Context) {
	// Start consumers
	ConsumerRouter(ctx, bootstrap.NewBootstrap(s.cfg), s.cfg)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case <-quit:
		log.Println("🛑 OS Signal received. Shutting down worker...")
	case <-ctx.Done():
		log.Println("🛑 Parent context cancelled. Shutting down worker...")
	}

	log.Println("Stopping event consumers...")
	// Optional: beri waktu consumers shutdown
	time.Sleep(2 * time.Second)
	log.Println("✅ Event consumer shutdown completed.")
}
