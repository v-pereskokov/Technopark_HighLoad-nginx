package main

import (
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/configs"
	modelConfig "github.com/vladpereskokov/Technopark_HighLoad-nginx/src/models/configs"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/server"
	"log"
)

const CONFIG = "configs/server.json"

func main() {
	config := new(modelConfig.Config)

	err := configs.FromFile(CONFIG, &config)
	if err != nil {
		log.Panicf("can not init config: %v", err)
	}

	httpServer := server.Server{}

	httpServer.CreateServer(*config.GetServer())
	httpServer.Start()
}
