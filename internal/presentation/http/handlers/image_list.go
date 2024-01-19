package handlers

import (
	"github.com/artarts36/service-navigator/internal/presentation/config"
	"net/http"
	"sort"

	log "github.com/sirupsen/logrus"
	"github.com/tyler-sommer/stick"

	"github.com/artarts36/service-navigator/internal/infrastructure/repository"
	"github.com/artarts36/service-navigator/internal/presentation/view"
)

type ImageListHandler struct {
	images     *repository.ImageRepository
	renderer   *view.Renderer
	pageConfig *config.ImagePage
}

func NewImageListHandler(
	images *repository.ImageRepository,
	renderer *view.Renderer,
	pageConfig *config.ImagePage,
) *ImageListHandler {
	return &ImageListHandler{images: images, renderer: renderer, pageConfig: pageConfig}
}

func (h *ImageListHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	images := h.images.All()

	sort.SliceStable(images, func(i, j int) bool {
		if images[i].Unknown {
			return false
		}

		return images[i].Name < images[j].Name ||
			(images[i].Name == images[j].Name && images[i].ID < images[j].ID)
	})

	err := h.renderer.Render("pages/images.twig.html", w, map[string]stick.Value{
		"images":     images,
		"pageConfig": h.pageConfig,
	})

	if err != nil {
		log.Errorf("Failed to render: %s", err)

		w.WriteHeader(http.StatusInternalServerError)
	}
}
