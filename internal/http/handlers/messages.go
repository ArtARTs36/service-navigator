package handlers

import (
	"log"
	"net/http"

	"github.com/artarts36/service-navigator/internal/presentation"
	"github.com/tyler-sommer/stick"
)

func writeSuccessMessage(renderer *presentation.Renderer, w http.ResponseWriter, message string) {
	err := renderer.Render("messages/success.twig.html", w, map[string]stick.Value{
		"message": message,
	})

	if err != nil {
		log.Printf("Failed to write success response: %s", err)
	}

	w.WriteHeader(ok)
}

func writeWarningMessage(renderer *presentation.Renderer, w http.ResponseWriter, message string) {
	err := renderer.Render("messages/warning.twig.html", w, map[string]stick.Value{
		"message": message,
	})

	if err != nil {
		log.Printf("Failed to write warning response: %s", err)
	}

	w.WriteHeader(userError)
}

func writeErrorMessage(renderer *presentation.Renderer, w http.ResponseWriter, message string) {
	err := renderer.Render("messages/error.twig.html", w, map[string]stick.Value{
		"message": message,
	})

	if err != nil {
		log.Printf("Failed to write error response: %s", err)
	}

	w.WriteHeader(serverError)
}
