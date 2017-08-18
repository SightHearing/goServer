package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	AwsSql MySqlConfig `json:"MySqlConfig"`
}

func (c *Config) InitConfig() {
	file, err := os.Open("config/sightHearingConfig.json")
	if err != nil {
		fmt.Println("error:", err)
		panic("Can not find sightHearingConfig.json file.\nThis file is critical for credentials")
		os.Exit(-1)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		fmt.Println("error:", err)
	}
	// fmt.Println(c)
}
