package handlers

import (
	"net/http"
)

type VolumeRefreshHandler struct {
	pollRequestChannel chan bool
}

func NewVolumeRefreshHandler(pollRequestChannel chan bool) *VolumeRefreshHandler {
	return &VolumeRefreshHandler{pollRequestChannel}
}

func (h *VolumeRefreshHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	h.pollRequestChannel <- true

	w.Header().Add("Location", "/volumes")
	w.WriteHeader(http.StatusFound)
}
