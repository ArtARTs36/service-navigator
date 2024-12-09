package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/artarts36/service-navigator/internal/config"
)

const httpReadTimeout = 3 * time.Second

func main() {
	log.SetLevel(log.DebugLevel)

	cont := config.InitContainer()

	ctx, cancel := context.WithCancel(context.Background())

	hServer := createHTTPServer(cont)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	wg := startWorkers([]func(group *sync.WaitGroup){
		func(wg *sync.WaitGroup) {
			defer wg.Done()
			cont.Services.Poller.Poll(ctx)
		},
		func(wg *sync.WaitGroup) {
			defer wg.Done()
			cont.Images.Poller.Poll(ctx, cont.Images.PollRequestsChannel)
		},
		func(wg *sync.WaitGroup) {
			defer wg.Done()
			cont.Volumes.Poller.Poll(ctx, cont.Volumes.PollRequestsChannel)
		},
		func(wg *sync.WaitGroup) {
			defer wg.Done()

			log.Printf("[Http][Server] Listening on %s", hServer.Addr)

			if err := hServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("[Http][Server] Listen: %s\n", err)
			}

			log.Print("[Http][Server] Started")
		},
	})

	<-done
	if err := hServer.Shutdown(ctx); err != nil {
		log.Fatalf("[Http][Server] Server shutdown failed:%+v", err)
	}

	log.Print("Cancelling root context")
	cancel()

	wg.Wait()
}

func startWorkers(workers []func(wg *sync.WaitGroup)) *sync.WaitGroup {
	wg := &sync.WaitGroup{}

	for _, worker := range workers {
		wg.Add(1)

		go worker(wg)
	}

	return wg
}

func createHTTPServer(cont *config.Container) *http.Server {
	mux := http.NewServeMux()

	bindRoutes(mux, cont)
	bindFileServer(mux)

	hServer := &http.Server{
		Addr:        ":8080",
		Handler:     mux,
		ReadTimeout: httpReadTimeout,
	}

	return hServer
}

func bindRoutes(mux *http.ServeMux, cont *config.Container) {
	mux.Handle("/", cont.Presentation.HTTP.Handlers.HomePageHandler)
	mux.Handle("/containers/kill", cont.Presentation.HTTP.Handlers.ContainerKillHandler)
	mux.Handle("/images/", cont.Presentation.HTTP.Handlers.ImageListHandler)
	mux.Handle("/images/{id}/show", cont.Presentation.HTTP.Handlers.ImageShowHandler)
	mux.Handle("/images/remove", cont.Presentation.HTTP.Handlers.ImageRemoveHandler)
	mux.Handle("/images/refresh", cont.Presentation.HTTP.Handlers.ImageRefreshHandler)
	mux.Handle("/volumes", cont.Presentation.HTTP.Handlers.VolumeListHandler)
	mux.Handle("/volumes/refresh", cont.Presentation.HTTP.Handlers.VolumeRefreshHandler)
}

func bindFileServer(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("/app/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
}
