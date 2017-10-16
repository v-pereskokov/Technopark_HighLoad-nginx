package configs

import "os"

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

func (dir *Dir) SetDir() string {
	dir.findDir()

	return dir.Path
}
