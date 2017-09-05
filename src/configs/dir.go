package configs

import (
	"fmt"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/utils"
)

type Dir struct {
	Path string `json:"path"`
}

func (dir *Dir) findDir() {
	if dir.Path == "pwd" {
		dir.Path = utils.Pwd()
	}
}

func (config *Config) GetDir() string {
	config.Dir.findDir()

	fmt.Println(config.Dir.Path)

	return config.Dir.Path
}
