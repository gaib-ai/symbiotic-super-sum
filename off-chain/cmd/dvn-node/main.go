package main

import (
	"context"
	"dvn-node/internal/dvn"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/go-errors/errors"
	"github.com/spf13/cobra"
)

type config struct {
	relayApiURL string
	evmRpcURLs  []string
	privateKey  string
	logLevel    string
}

var cfg config

func main() {
	if err := run(); err != nil && !errors.Is(err, context.Canceled) {
		slog.Error("Error executing command", "error", err)
		os.Exit(1)
	}
	slog.Info("DVN node completed successfully")
}

func run() error {
	rootCmd.PersistentFlags().StringVarP(&cfg.relayApiURL, "relay-api-url", "r", "", "Relay API URL")
	rootCmd.PersistentFlags().StringSliceVarP(&cfg.evmRpcURLs, "evm-rpc-urls", "e", []string{}, "EVM RPC URLs, format: eid:url,eid:url (e.g., '31337:http://localhost:8545,31338:http://localhost:8546')")
	rootCmd.PersistentFlags().StringVarP(&cfg.privateKey, "private-key", "p", "", "DVN worker private key for submitting transactions")
	rootCmd.PersistentFlags().StringVarP(&cfg.logLevel, "log-level", "l", "info", "Log level (debug, info, warn, error)")

	if err := rootCmd.MarkPersistentFlagRequired("relay-api-url"); err != nil {
		return errors.Errorf("failed to mark relay-api-url as required: %w", err)
	}
	if err := rootCmd.MarkPersistentFlagRequired("evm-rpc-urls"); err != nil {
		return errors.Errorf("failed to mark evm-rpc-urls as required: %w", err)
	}
	if err := rootCmd.MarkPersistentFlagRequired("private-key"); err != nil {
		return errors.Errorf("failed to mark private-key as required: %w", err)
	}

	return rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:           "dvn-node",
	Short:         "A LayerZero DVN off-chain worker powered by Symbiotic.",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Set up logger
		var logLevel slog.Level
		switch cfg.logLevel {
		case "debug":
			logLevel = slog.LevelDebug
		case "info":
			logLevel = slog.LevelInfo
		case "warn":
			logLevel = slog.LevelWarn
		case "error":
			logLevel = slog.LevelError
		default:
			logLevel = slog.LevelInfo
		}
		logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
		slog.SetDefault(logger)

		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()

		// Parse EVM RPC URLs
		evmRpcURLs := make(map[uint32]string)
		for _, urlPair := range cfg.evmRpcURLs {
			parts := strings.Split(urlPair, ":")
			if len(parts) < 2 {
				return errors.Errorf("invalid evm-rpc-url format: %s", urlPair)
			}
			eidInt, err := strconv.Atoi(parts[0])
			if err != nil {
				return errors.Errorf("invalid EID in evm-rpc-url: %s", parts[0])
			}
			// Re-join the rest of the parts in case the URL has a colon
			url := strings.Join(parts[1:], ":")
			evmRpcURLs[uint32(eidInt)] = url
		}

		processor, err := dvn.NewProcessor(ctx, cfg.relayApiURL, evmRpcURLs, cfg.privateKey, logger)
		if err != nil {
			return errors.Errorf("failed to setup processor: %w", err)
		}

		return processor.Listen(ctx)
	},
}
