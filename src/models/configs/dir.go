package configs

import (
	"fmt"
	"os"
)

type Dir struct {
	Path string `json:"path"`
}

func (dir *Dir) findDir() {
	if dir.Path == "pwd" {
		pwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		dir.Path = pwd
	}
}

func (dir *Dir) GetDir() string {
	dir.findDir()

	fmt.Println(dir.Path)

	return dir.Path
}
