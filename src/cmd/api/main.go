package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"push-service/internal/app"
	"push-service/internal/config"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		Handler:           app.New(app.Deps{Log: log}),
		ReadHeaderTimeout: 10 * time.Second,
	}
	log.Info("API đang chạy", "addr", srv.Addr)
	return srv.ListenAndServe()
}
