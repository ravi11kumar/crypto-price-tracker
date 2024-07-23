package utils

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var ExpiryDuration time.Duration

var configData map[string]interface{}

func LoadConfig() {
	jsonFile, err := os.ReadFile("configs/config.json")
	if err != nil {
		log.Fatalf("Error reading config.json: %v", err)
	}

	err = json.Unmarshal(jsonFile, &configData)
	if err != nil {
		log.Fatalf("Error parsing config.json: %v", err)
	}

	log.Printf("Loaded configuration: %+v", configData) // Debug log

	expiry, err := time.ParseDuration(Get(CacheExpiry))
	if err != nil {
		log.Fatalf("Invalid cache.expiry value: %v", err)
	}

	ExpiryDuration = expiry
}

func Get(key ConfigKey) string {
	val, err := getConfigValue(key)
	if err != nil {
		panic(err)
	}
	return val
}

func GetInt(key ConfigKey) int {
	val := Get(key)
	num, err := strconv.Atoi(val)
	if err != nil {
		panic(errors.New("configuration value not int [" + string(key) + "]"))
	}
	return num
}

func GetBool(key ConfigKey) bool {
	val := Get(key)
	retVal, err := strconv.ParseBool(val)
	if err != nil {
		panic(errors.New("configuration value not bool [" + string(key) + "]"))
	}
	return retVal
}

func GetOptional(key ConfigKey) (string, bool) {
	val, err := getConfigValue(key)
	if err != nil {
		return "", false
	}
	return val, true
}

func getConfigValue(key ConfigKey) (string, error) {
	keys := strings.Split(string(key), ".")
	var val interface{} = configData
	for _, k := range keys {
		log.Printf("Accessing key: %s", k) // Debug log
		m, ok := val.(map[string]interface{})
		if !ok {
			return "", errors.New("configuration value not found [" + string(key) + "]")
		}
		val, ok = m[k]
		if !ok {
			return "", errors.New("configuration value not found [" + string(key) + "]")
		}
	}

	strVal, ok := val.(string)
	if !ok {
		return "", errors.New("configuration value not a string [" + string(key) + "]")
	}

	return strVal, nil
}
