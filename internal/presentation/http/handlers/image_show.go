package handlers

import (
	"fmt"
	"net/http"

	"github.com/artarts36/service-navigator/internal/infrastructure/repository"
	"github.com/artarts36/service-navigator/internal/presentation/view"
	log "github.com/sirupsen/logrus"
	"github.com/tyler-sommer/stick"
)

type ImageShowHandler struct {
	renderer *view.Renderer
	images   *repository.ImageRepository
}

func NewImageShowHandler(renderer *view.Renderer, images *repository.ImageRepository) *ImageShowHandler {
	return &ImageShowHandler{
		renderer: renderer,
		images:   images,
	}
}

func (h *ImageShowHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	imageID := req.PathValue("id")
	image, err := h.images.FindByID(imageID)
	if err != nil {
		writeErrorMessage(h.renderer, w, fmt.Sprintf("Image %s not found", imageID))
	}

	err = h.renderer.Render("pages/image_show.twig.html", w, map[string]stick.Value{
		"image": image,
	})

	if err != nil {
		log.Errorf("Failed to write success response: %s", err)
	}
}
