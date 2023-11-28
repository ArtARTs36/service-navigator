package handlers

import (
	"net/http"
)

type ImageRefreshHandler struct {
	pollRequestChannel chan bool
}

func NewImageRefreshHandler(pollRequestChannel chan bool) *ImageRefreshHandler {
	return &ImageRefreshHandler{pollRequestChannel}
}

func (h *ImageRefreshHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	h.pollRequestChannel <- true

	w.Header().Add("Location", "/images")
	w.WriteHeader(http.StatusFound)
}
