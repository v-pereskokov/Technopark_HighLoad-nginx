package main

import (
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/configs"
	"log"
)

func main() {
	config, err := configs.FromFile("config.json")
	if err != nil {
		log.Panicf("can not init config: %v", err)
	}

}
