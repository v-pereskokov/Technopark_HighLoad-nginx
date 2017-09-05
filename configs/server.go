package configs

type Server struct {
	Network  string `json:"network"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}
