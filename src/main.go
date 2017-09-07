package main

import (
	"fmt"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/handler"
	modelConfig "github.com/vladpereskokov/Technopark_HighLoad-nginx/src/models/configs"
	modelServer "github.com/vladpereskokov/Technopark_HighLoad-nginx/src/models/server"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/server"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/utils"
	"log"
)

const (
	SERVER_CONFIG         = "configs/server.json"
	HTTP_CONSTANTS_CONFIG = "configs/http.json"
)

func main() {
	serverConfig := new(modelConfig.Config)

	err := utils.FromFile(SERVER_CONFIG, &serverConfig)
	if err != nil {
		log.Panicf("can not init server config: %v", err)
	}

	httpConstantsConfig := new(modelServer.Constants)

	err = utils.FromFile(HTTP_CONSTANTS_CONFIG, &httpConstantsConfig)
	if err != nil {
		log.Panicf("can not init http config: %v", err)
	}

	httpServer := server.CreateServer(*serverConfig.GetServer())
	httpServer.Start(handler.CreateHandler(serverConfig.Dir.Path))
}
