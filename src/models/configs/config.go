package configs

import modelServer "github.com/vladpereskokov/Technopark_HighLoad-nginx/src/models/server"

type Config struct {
	Server   *Server
	Dir      *Dir
	Statuses *modelServer.Statuses
	Methods  *modelServer.Methods
}

func (config *Config) GetServer() *Server {
	return config.Server
}

func (config *Config) GetDir() *Dir {
	return config.Dir
}

func (config *Config) GetStatuses() *modelServer.Statuses {
	return config.Statuses
}

func (config *Config) GetMethods() *modelServer.Methods {
	return config.Methods
}
