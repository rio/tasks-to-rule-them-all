package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"tasks-to-rule-them-all/pkg/server"
)

func main() {
	// setup default logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := run(logger); err != nil {
		logger.Error("listening failed", "error", err.Error())
	}
}

func run(logger *slog.Logger) error {
	logger.Info("server starting")
	defer logger.Info("server stopped")

	port := os.Getenv("ECHO_SERVER_LISTEN_PORT")
	if port == "" {
		port = "8080"
	}

	if _, err := strconv.Atoi(port); err != nil {
		return fmt.Errorf("`ECHO_SERVER_LISTEN_PORT` must be an integer. Received %q", port)
	}

	router := http.NewServeMux()
	router.HandleFunc("/echo", server.Echo)
	router.HandleFunc("/healthz", server.Healthz)

	logger.Info("routes mounted")

	logger.Info("server listening", "server.port", port)

	return http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
