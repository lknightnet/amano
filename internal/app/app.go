package app

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
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

func RunBot(cfg *config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := bot.New(cfg.TG.Token)
	if err != nil {
		log.Println(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/logs", bot.MatchTypeExact, logs)

	b.Start(ctx)
}

func logs(ctx context.Context, b *bot.Bot, u *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: u.Message.Chat.ID,
		Text:   "i am work!",
	})
}
