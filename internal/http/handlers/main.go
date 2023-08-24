package handlers

import (
	"github.com/artarts36/service-navigator/internal/service/monitor"
	"github.com/tyler-sommer/stick"
	"github.com/tyler-sommer/stick/twig"
	"net/http"
)

type MainPageHandler struct {
	monitor     *monitor.Monitor
	appName     string
	networkName string
}

func NewMainPageHandler(monitor *monitor.Monitor, appName string, networkName string) *MainPageHandler {
	return &MainPageHandler{monitor: monitor, appName: appName, networkName: networkName}
}

func (h *MainPageHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	statuses, err := h.monitor.Show(req.Context())

	if err != nil {
		w.WriteHeader(500)

		return
	}

	loader := stick.NewFilesystemLoader("/app/templates")
	renderer := twig.New(loader)

	err = renderer.Execute("main.twig.html", w, map[string]stick.Value{
		"services": statuses,
		"_appName": h.appName,
	})

	if err != nil {
		w.WriteHeader(500)
	}
}
