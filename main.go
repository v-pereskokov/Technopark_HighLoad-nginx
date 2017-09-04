package main

import (
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/utils"
	"log"
)

func main() {
	_, err := utils.FromFile("config.json")
	if err != nil {
		log.Panicf("can not init config: %v", err)
	}
}
