package configs

//import modelServer "github.com/vladpereskokov/Technopark_HighLoad-nginx/src/models/server"

type Config struct {
	Server   *Server
	Dir      *Dir
	Statuses *Statuses
}

func (config *Config) GetServer() *Server {
	return config.Server
}

func (config *Config) GetDir() *Dir {
	return config.Dir
}

type Statuses struct {
	Status []struct {
		Message string
		Code    int
	} `json:"status"`
}
