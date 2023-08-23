package handlers

import (
	"fmt"
	"github.com/artarts36/service-navigator/internal/services"
	"net/http"
)

type MainPageHandler struct {
	monitor *services.Monitor
}

func NewMainPageHandler(monitor *services.Monitor) *MainPageHandler {
	return &MainPageHandler{monitor: monitor}
}

func (h *MainPageHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	statuses, err := h.monitor.Show(req.Context())

	if err != nil {
		w.Write([]byte(err.Error()))
	}

	for _, st := range statuses {
		var line string

		if st.WebUrl == nil {
			line = fmt.Sprintf("%s:non", st.Name)
		} else {
			line = fmt.Sprintf("%s:%s", st.Name, *st.WebUrl)
		}

		w.Write([]byte(line))
		w.Write([]byte("\n"))
	}
}
