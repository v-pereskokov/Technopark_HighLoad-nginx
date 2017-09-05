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

func (dir *Dir) GetDir() string {
	dir.findDir()

	fmt.Println(dir.Path)

	return dir.Path
}
