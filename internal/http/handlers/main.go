package handlers

import (
	"log"
	"net/http"
	"sort"

	"github.com/artarts36/service-navigator/internal/presentation"
	"github.com/artarts36/service-navigator/internal/service/repository"
	"github.com/tyler-sommer/stick"
)

type HomePageHandler struct {
	services *repository.ServiceRepository
	renderer *presentation.Renderer
}

func NewHomePageHandler(services *repository.ServiceRepository, renderer *presentation.Renderer) *HomePageHandler {
	return &HomePageHandler{services: services, renderer: renderer}
}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	statuses := h.services.All()

	sort.SliceStable(statuses, func(i, j int) bool {
		return statuses[i].Name < statuses[j].Name ||
			(statuses[i].Name == statuses[j].Name && statuses[i].ContainerID < statuses[j].ContainerID)
	})

	err := h.renderer.Render("pages/home.twig.html", w, map[string]stick.Value{
		"services": statuses,
	})

	if err != nil {
		log.Printf("Failed to render: %s", err)

		w.WriteHeader(serverError)
	}
}
