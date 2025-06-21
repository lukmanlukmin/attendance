// Package cmd ...
package cmd

import (
	"attendance/config"
	"attendance/server"
	"context"
)

// StartHTTP ...
func StartHTTP(ctx context.Context, cfg *config.Config) {
	defer ctx.Done()
	api := server.NewHTTPApi(cfg)
	api.Run(ctx)
}
