// Package cmd ...
package cmd

import (
	"attendance/config"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

// Start ...
func Start() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("🛑 Caught OS signal. Canceling context...")
		cancel()
	}()

	rootCmd := &cobra.Command{}
	serveHTTPCmd := &cobra.Command{
		Use:   "serve-http",
		Short: "Run HTTP Server",
		Run: func(_ *cobra.Command, _ []string) {
			cfg := &config.Config{}
			err := config.ReadModuleConfig(cfg, "") /// set default config // working on it later
			if err == nil {
				StartHTTP(ctx, cfg)
			}
		},
	}
	rootCmd.AddCommand(serveHTTPCmd)

	serveWorkerCmd := &cobra.Command{
		Use:   "serve-worker",
		Short: "Run Worker",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := &config.Config{}
			err := config.ReadModuleConfig(cfg, "") /// set default config // working on it later
			if err == nil {
				StartWorker(ctx, cfg)
			}
		},
	}
	rootCmd.AddCommand(serveWorkerCmd)

	serveAllCmd := &cobra.Command{
		Use:   "serve-all",
		Short: "Run All",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := &config.Config{}
			err := config.ReadModuleConfig(cfg, "") /// set default config // working on it later
			if err == nil {
				go StartHTTP(ctx, cfg)
				time.Sleep(10 * time.Second)
				go StartWorker(ctx, cfg)
				<-ctx.Done()
				log.Println("✅ Exiting main process gracefully...")
			}
		},
	}
	rootCmd.AddCommand(serveAllCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
