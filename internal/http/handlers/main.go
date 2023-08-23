package handlers

import (
	"github.com/artarts36/service-navigator/internal/services"
	"github.com/tyler-sommer/stick"
	"github.com/tyler-sommer/stick/twig"
	"net/http"
)

type MainPageHandler struct {
	monitor *services.Monitor
}

func NewMainPageHandler(monitor *services.Monitor) *MainPageHandler {
	return &MainPageHandler{monitor: monitor}
}

func (h *MainPageHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	statuses, err := h.monitor.Show(req.Context())

	if err != nil {
		w.WriteHeader(500)

		return
	}

	loader := stick.NewFilesystemLoader("/app/templates")
	renderer := twig.New(loader)

	err = renderer.Execute("main.twig.html", w, map[string]stick.Value{
		"services": statuses,
	})

	if err != nil {
		w.WriteHeader(500)
	}
}
