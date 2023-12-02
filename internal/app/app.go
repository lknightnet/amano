package app

import (
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/signal"
	"syscall"
	"yukiteru-amano/config"
	"yukiteru-amano/internal/controller/http"
	"yukiteru-amano/internal/service"
	"yukiteru-amano/pkg/heapofshit"
	"yukiteru-amano/pkg/server"
)

func RunAmano(cfg *config.Config) {
	log.Println(cfg.Name + " is starting")
	deps := &service.Dependencies{
		Heap: heapofshit.NewHeap(10),
	}
	services := service.NewServices(deps)

	rout := mux.NewRouter()
	http.NewMessageRoutes(rout, services.MD)
	srv := server.NewServer(rout, server.Port(cfg.HTTP.Port), server.ReadTimeout(cfg.ReadTimeout),
		server.WriteTimeout(cfg.WriteTimeout), server.ShutdownTimeout(cfg.ShutdownTimeout))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("Run: " + s.String())
	case err := <-srv.Notify():
		log.Println(errors.Wrap(err, "Run: signal.Notify"))
	}

	err := srv.Shutdown()
	if err != nil {
		log.Println(errors.Wrap(err, "Run: server shutdown"))
	}
}
