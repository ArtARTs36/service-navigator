package handlers

import (
	"log"
	"net/http"

	"github.com/tyler-sommer/stick"

	"github.com/artarts36/service-navigator/internal/presentation/view"
)

func writeSuccessMessage(renderer *view.Renderer, w http.ResponseWriter, message string) {
	err := renderer.Render("messages/success.twig.html", w, map[string]stick.Value{
		"message": message,
	})

	if err != nil {
		log.Printf("Failed to write success response: %s", err)
	}

	w.WriteHeader(http.StatusOK)
}

func writeWarningMessage(renderer *view.Renderer, w http.ResponseWriter, message string) {
	err := renderer.Render("messages/warning.twig.html", w, map[string]stick.Value{
		"message": message,
	})

	if err != nil {
		log.Printf("Failed to write warning response: %s", err)
	}

	w.WriteHeader(http.StatusUnprocessableEntity)
}

func writeErrorMessage(renderer *view.Renderer, w http.ResponseWriter, message string) {
	err := renderer.Render("messages/error.twig.html", w, map[string]stick.Value{
		"message": message,
	})

	if err != nil {
		log.Printf("Failed to write error response: %s", err)
	}

	w.WriteHeader(http.StatusInternalServerError)
}
