package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"tasks-to-rule-them-all/pkg/config"
)

type Server struct {
	cfg config.Config
}

func NewServer(cfg config.Config) Server {
	return Server{cfg: cfg}
}

func (s *Server) Echo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	logger := slog.With(slog.Group("http.request", "method", r.Method, "path", r.URL.Path))

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		logger.ErrorContext(ctx, "request failed", "http.response.status_code", http.StatusMethodNotAllowed, "error", http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	defer r.Body.Close()

	req := Request{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.ErrorContext(ctx, "unable to decode request", "http.response.status_code", http.StatusBadRequest, "error", err.Error())
		return

	}

	resp := Response{
		Message:       req.Message,
		Count:         countCharacters(req.Message),
		KubernetesEnv: s.cfg.Env,
	}

	jsonData, err := json.MarshalIndent(&resp, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.ErrorContext(ctx, "unable to encode response", "http.response.status_code", http.StatusInternalServerError, "error", err.Error())
		return
	}

	// w.Header().Add("Content-Type", "application/json")
	bodySize, _ := w.Write(jsonData)

	logger.InfoContext(ctx, "request handled", "http.response.status_code", http.StatusOK, "http.resonse.body.size", bodySize)
}

func countCharacters(input string) int {
	return len(input)
}

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Message       string            `json:"message"`
	Count         int               `json:"count"`
	KubernetesEnv map[string]string `json:"kubernetesEnv,omitempty"`
}

func (s *Server) Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
