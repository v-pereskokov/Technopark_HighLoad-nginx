package main

import (
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/handler"
	modelConfig "github.com/vladpereskokov/Technopark_HighLoad-nginx/src/models/configs"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/server"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/utils"
	"log"
)

const CONFIG = "configs/server.json"

func main() {
	config := new(modelConfig.Config)

	err := utils.FromFile(CONFIG, &config)
	if err != nil {
		log.Panicf("can not init config: %v", err)
	}

	httpServer := server.CreateServer(*config.GetServer())
	httpServer.Start(handler.CreateHandler(config.Dir.Path))
}
