package main

import (
	"log"
	"net/http"
	"time"

	"github.com/artarts36/service-navigator/internal/config"
)

const httpReadTimeout = 3 * time.Second

func main() {
	cont := config.InitContainer()

	go func() {
		cont.Services.Poller.Poll()
	}()

	go func() {
		cont.Images.Poller.Poll(cont.Images.PollRequestsChannel)
	}()

	mux := http.NewServeMux()

	bindRoutes(mux, cont)
	bindFileServer(mux)

	hServer := &http.Server{
		Addr:        ":8080",
		Handler:     mux,
		ReadTimeout: httpReadTimeout,
	}

	log.Print("Listening...")

	err := hServer.ListenAndServe()
	if err != nil {
		log.Printf("Failed listeing: %s", err)

		return
	}
}

func bindRoutes(mux *http.ServeMux, cont *config.Container) {
	mux.Handle("/", cont.HTTP.Handlers.HomePageHandler)
	mux.Handle("/containers/kill", cont.HTTP.Handlers.ContainerKillHandler)
	mux.Handle("/images", cont.HTTP.Handlers.ImageListHandler)
	mux.Handle("/images/remove", cont.HTTP.Handlers.ImageRemoveHandler)
	mux.Handle("/images/refresh", cont.HTTP.Handlers.ImageRefreshHandler)
}

func bindFileServer(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("/app/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
}
