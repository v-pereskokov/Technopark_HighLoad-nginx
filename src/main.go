package main

import (
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/handler"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/models/configs"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/server"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/utils"
	"log"
	"runtime"
)

const (
	SERVER_CONFIG = "configs/server.json"
	MAX_PROCS     = 1
)

func main() {
	serverConfig := new(configs.Config)

	err := utils.FromFile(SERVER_CONFIG, &serverConfig)
	if err != nil {
		log.Panicf("can not init server config: %v", err)
	}

	serverConfig.Dir.SetDir()

	log.Printf("cpu: %v\n", MAX_PROCS)
	runtime.GOMAXPROCS(MAX_PROCS)

	httpServer := server.CreateServer(*serverConfig.GetServer())
	httpServer.Start(handler.CreateHandler(serverConfig.Dir.Path))
}
