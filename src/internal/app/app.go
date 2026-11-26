package app

import (
	"log/slog"
	"net/http"
	"push-service/internal/health"
	"push-service/internal/httpx"
)

type Deps struct {
	Log *slog.Logger
}

func New(deps Deps) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", health.Check)
	return httpx.LogRequests(deps.Log, mux)
}
