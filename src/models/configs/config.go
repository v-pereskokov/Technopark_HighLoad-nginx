package configs

import modelServer "github.com/vladpereskokov/Technopark_HighLoad-nginx/src/models/server"

type Config struct {
	Server   *Server
	Dir      *Dir
	Statuses *modelServer.Statuses
}

func (config *Config) GetServer() *Server {
	return config.Server
}

func (config *Config) GetDir() *Dir {
	return config.Dir
}
