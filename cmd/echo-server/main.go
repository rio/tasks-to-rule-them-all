package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"tasks-to-rule-them-all/pkg/config"
	"tasks-to-rule-them-all/pkg/server"
)

func main() {
	// setup default logger
	// slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("parsing config failed", "error", err.Error())
		os.Exit(1)
	}

	if err := run(cfg); err != nil {
		slog.Error("listening failed", "error", err.Error())
	}
}

func run(cfg config.Config) error {
	logger := slog.Default()

	logger.Info("server starting", "k8s.detected", cfg.RunningOnK8S)
	defer logger.Info("server stopped")

	echoServer := server.NewServer(cfg)

	router := http.NewServeMux()
	router.HandleFunc("/echo", echoServer.Echo)
	router.HandleFunc("/healthz", echoServer.Healthz)

	logger.Info("routes mounted")

	logger.Info("server listening", "server.port", cfg.Port)

	return http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), router)
}
