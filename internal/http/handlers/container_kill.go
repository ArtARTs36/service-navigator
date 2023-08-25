package handlers

import (
	"fmt"
	"net/http"

	"github.com/artarts36/service-navigator/internal/presentation"
	"github.com/artarts36/service-navigator/internal/service/monitor"
)

type ContainerKillHandler struct {
	monitor  *monitor.Monitor
	renderer *presentation.Renderer
}

func NewContainerKillHandler(monitor *monitor.Monitor, renderer *presentation.Renderer) *ContainerKillHandler {
	return &ContainerKillHandler{monitor: monitor, renderer: renderer}
}

func (h *ContainerKillHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	if err != nil {
		writeWarningMessage(h.renderer, w, err.Error())

		return
	}

	contID := req.Form.Get("containerId")

	err = h.monitor.KillContainer(req.Context(), contID)

	if err != nil {
		writeErrorMessage(h.renderer, w, err.Error())

		return
	}

	writeSuccessMessage(h.renderer, w, fmt.Sprintf("Container \"%s\" was killed", contID))
}
