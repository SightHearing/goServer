package config

type MySqlConfig struct {
	Endpoint string `json:"Endpoint"`
	UserName string `json:"UserName"`
	Password string `json:"Password"`
	Port     string `json:"Port"`
}
