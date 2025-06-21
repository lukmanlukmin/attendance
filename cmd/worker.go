// Package cmd ...
package cmd

import (
	"attendance/config"
	"attendance/server"
	"context"
)

// StartWorker ...
func StartWorker(ctx context.Context, cfg *config.Config) {
	worker := server.NewEventServer(ctx, cfg)
	worker.Run(ctx)
}
