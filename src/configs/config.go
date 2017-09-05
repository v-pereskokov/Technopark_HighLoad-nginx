package configs

import "fmt"

type Config struct {
	Server *Server
	Dir    *Dir
}

func (config *Config) GetDir() string {
	config.Dir.FindDir()

	fmt.Println(config.Dir.Path)

	return config.Dir.Path
}
