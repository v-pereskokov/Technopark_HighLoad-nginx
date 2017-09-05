package configs

import "github.com/vladpereskokov/Technopark_HighLoad-nginx/utils"

type Dir struct {
	Path string `json:"path"`
}

func (dir *Dir) FindDir() {
	if dir.Path == "pwd" {
		dir.Path = utils.Pwd()
	}
}
