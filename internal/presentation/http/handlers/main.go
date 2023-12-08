package handlers

import (
	"log"
	"net/http"
	"sort"

	"github.com/tyler-sommer/stick"

	"github.com/artarts36/service-navigator/internal/infrastructure/repository"
	"github.com/artarts36/service-navigator/internal/presentation/view"
)

type HomePageHandler struct {
	services *repository.ServiceRepository
	renderer *view.Renderer
}

func NewHomePageHandler(services *repository.ServiceRepository, renderer *view.Renderer) *HomePageHandler {
	return &HomePageHandler{services: services, renderer: renderer}
}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	services := h.services.All()

	sort.SliceStable(services, func(i, j int) bool {
		return services[i].Name < services[j].Name ||
			(services[i].Name == services[j].Name && services[i].ContainerID < services[j].ContainerID)
	})

	err := h.renderer.Render("pages/home.twig.html", w, map[string]stick.Value{
		"services": services,
	})

	if err != nil {
		log.Printf("Failed to render: %s", err)

		w.WriteHeader(http.StatusInternalServerError)
	}
}
