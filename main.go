package main

import (
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/server"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/utils"
	"log"
)

func main() {
	config, err := utils.FromFile("config.json")
	if err != nil {
		log.Panicf("can not init config: %v", err)
	}

	httpServer := server.Server{}
	httpServer.Start(config)
}
