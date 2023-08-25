package handlers

import (
	"log"
	"net/http"
	"sort"

	"github.com/artarts36/service-navigator/internal/presentation"
	"github.com/artarts36/service-navigator/internal/service/monitor"
	"github.com/tyler-sommer/stick"
)

type HomePageHandler struct {
	monitor  *monitor.Monitor
	renderer *presentation.Renderer
}

func NewHomePageHandler(monitor *monitor.Monitor, renderer *presentation.Renderer) *HomePageHandler {
	return &HomePageHandler{monitor: monitor, renderer: renderer}
}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	statuses, err := h.monitor.Show(req.Context())

	if err != nil {
		log.Printf("Failed to fetch services: %s", err)

		w.WriteHeader(serverError)

		return
	}

	sort.SliceStable(statuses, func(i, j int) bool {
		return statuses[i].Name < statuses[j].Name
	})

	err = h.renderer.Render("pages/home.twig.html", w, map[string]stick.Value{
		"services": statuses,
	})

	if err != nil {
		log.Printf("Failed to render: %s", err)

		w.WriteHeader(serverError)
	}
}
