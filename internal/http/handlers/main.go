package handlers

import (
	"log"
	"net/http"

	"github.com/artarts36/service-navigator/internal/presentation"
	"github.com/artarts36/service-navigator/internal/service/monitor"
	"github.com/tyler-sommer/stick"
)

type MainPageHandler struct {
	monitor  *monitor.Monitor
	renderer *presentation.Renderer
}

func NewMainPageHandler(monitor *monitor.Monitor, renderer *presentation.Renderer) *MainPageHandler {
	return &MainPageHandler{monitor: monitor, renderer: renderer}
}

func (h *MainPageHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	statuses, err := h.monitor.Show(req.Context())

	if err != nil {
		log.Printf("Failed to fetch services: %s", err)

		w.WriteHeader(serverError)

		return
	}

	err = h.renderer.Render("main.twig.html", w, map[string]stick.Value{
		"services": statuses,
	})

	if err != nil {
		log.Printf("Failed to render: %s", err)

		w.WriteHeader(serverError)
	}
}
