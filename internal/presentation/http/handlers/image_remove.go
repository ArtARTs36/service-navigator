package handlers

import (
	"github.com/artarts36/service-navigator/internal/infrastructure/image/monitor"
	"github.com/artarts36/service-navigator/internal/presentation/view"
	"github.com/tyler-sommer/stick"
	"log"
	"net/http"
)

type ImageRemoveHandler struct {
	monitor  *monitor.Monitor
	renderer *view.Renderer
}

func NewImageRemoveHandler(monitor *monitor.Monitor, renderer *view.Renderer) *ImageRemoveHandler {
	return &ImageRemoveHandler{monitor: monitor, renderer: renderer}
}

func (h *ImageRemoveHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	if err != nil {
		writeWarningMessage(h.renderer, w, err.Error())

		return
	}

	imageID := req.Form.Get("imageId")
	force := req.Form.Get("force") == "1"

	statuses, removeErr := h.monitor.Remove(req.Context(), imageID, force)

	if removeErr != nil {
		err = h.renderer.Render("forms/image_remove_fail.twig.html", w, map[string]stick.Value{
			"imageID": imageID,
			"error":   removeErr,
		})

		if err != nil {
			log.Printf("Failed to write success response: %s", err)
		}

		return
	}

	err = h.renderer.Render("forms/image_remove_success.twig.html", w, map[string]stick.Value{
		"imageID":  imageID,
		"statuses": statuses,
	})

	if err != nil {
		log.Printf("Failed to write success response: %s", err)
	}
}
