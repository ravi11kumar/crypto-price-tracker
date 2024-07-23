package utils

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"time"
)

var ExpiryDuration time.Duration

var Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Cache struct {
		Expiry string `yaml:"expiry"`
	} `yaml:"cache"`
}

func LoadConfig() {
	yamlFile, err := os.ReadFile("configs/config.yml")
	if err != nil {
		log.Fatalf("Error reading config.yml: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	expiry, err := time.ParseDuration(Config.Cache.Expiry)
	if err != nil {
		log.Fatalf("Invalid CACHE_EXPIRY value: %v", err)
	}

	ExpiryDuration = expiry
}
