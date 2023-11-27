package middlewares

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	handler http.Handler
}

func NewLogMiddleware(handler http.Handler) http.Handler {
	return &LogMiddleware{handler: handler}
}

func (h *LogMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Debugf("[HTTP] handling %s request %s", req.Method, req.RequestURI)

	h.handler.ServeHTTP(w, req)

	log.Debugf("[HTTP] %s request %s handled", req.Method, req.RequestURI)
}
