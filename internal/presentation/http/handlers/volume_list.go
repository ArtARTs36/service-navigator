package handlers

import (
	"log"
	"net/http"
	"sort"

	"github.com/tyler-sommer/stick"

	"github.com/artarts36/service-navigator/internal/infrastructure/repository"
	"github.com/artarts36/service-navigator/internal/presentation/view"
)

type VolumeListHandler struct {
	volumes  *repository.VolumeRepository
	renderer *view.Renderer
}

func NewVolumeListHandler(volumes *repository.VolumeRepository, renderer *view.Renderer) *VolumeListHandler {
	return &VolumeListHandler{volumes: volumes, renderer: renderer}
}

func (h *VolumeListHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	volumes := h.volumes.All()

	sort.SliceStable(volumes, func(i, j int) bool {
		return volumes[i].Name < volumes[j].Name || volumes[i].Name == volumes[j].Name
	})

	err := h.renderer.Render("pages/volumes.twig.html", w, map[string]stick.Value{
		"volumes": volumes,
	})

	if err != nil {
		log.Printf("Failed to render: %s", err)

		w.WriteHeader(http.StatusInternalServerError)
	}
}
