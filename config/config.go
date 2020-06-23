package config

import (
	"encoding/json"
	"os"
)

var config *Config

func init() {
	config = ReadConfig("./config.json")
}

func GetConfig() Config {
	return *config
}

type Config struct {
	URL    string `json:"url"`
	Cookie string `json:"cookie"`
}

func ReadConfig(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	config := &Config{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		panic(err)
	}
	return config
}
