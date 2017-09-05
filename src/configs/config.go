package configs

type Config struct {
	Server *Server
	Dir    *Dir
}

func (config *Config) GetServer() *Server {
	return config.Server
}

func (config *Config) GetDir() *Dir {
	return config.Dir
}
