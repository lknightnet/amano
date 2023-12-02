package main

import (
	"log"
	"time"
	"yukiteru-amano/config"
	"yukiteru-amano/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Println(err)
	}

	go app.RunAmano(cfg)
	go app.RunBot(cfg)

	time.Sleep(8760 * time.Hour)
}
