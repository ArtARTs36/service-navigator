package handlers

import (
	"log"
	"net/http"
	"sort"

	"github.com/tyler-sommer/stick"

	"github.com/artarts36/service-navigator/internal/infrastructure/repository"
	"github.com/artarts36/service-navigator/internal/presentation/view"
)

type ImageListHandler struct {
	images   *repository.ImageRepository
	renderer *view.Renderer
}

func NewImageListHandler(images *repository.ImageRepository, renderer *view.Renderer) *ImageListHandler {
	return &ImageListHandler{images: images, renderer: renderer}
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
		"images": images,
	})

	if err != nil {
		log.Printf("Failed to render: %s", err)

		w.WriteHeader(serverError)
	}
}
