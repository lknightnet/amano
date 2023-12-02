package main

import (
	"log"
	"yukiteru-amano/config"
	"yukiteru-amano/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Println(err)
	}

	app.RunAmano(cfg)
}
