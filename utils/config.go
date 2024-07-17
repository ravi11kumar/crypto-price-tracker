package utils

import (
	"github.com/BurntSushi/toml"
	"log"
	"time"
)

var ExpiryDuration time.Duration

var Config struct {
	Server struct {
		Port string
	}
	Cache struct {
		Expiry string
	}
}

func LoadConfig() {
	if _, err := toml.DecodeFile("configs/config.toml", &Config); err != nil {
		log.Fatalf("Error loading config.toml: %v", err)
	}

	expiry, err := time.ParseDuration(Config.Cache.Expiry)
	if err != nil {
		log.Fatalf("Invalid CACHE_EXPIRY value err : %v", err)
	}

	ExpiryDuration = expiry
}
