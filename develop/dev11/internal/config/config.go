package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Port string `json:"port"`
}

func NewConfig(filename string) *Config {

	var conf Config

	jsonFile, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(jsonFile, &conf)

	return &conf
}
