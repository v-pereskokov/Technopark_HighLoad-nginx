package main

import (
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/configs"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/server"
	"log"
)

const CONFIG = "config.json"

func main() {
	config, err := configs.FromFile(CONFIG)
	if err != nil {
		log.Panicf("can not init config: %v", err)
	}

	httpServer := server.Server{}
	httpServer.Start(config)
}
