package utils

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

var ExpiryDuration time.Duration

var configData map[string]interface{}

func LoadConfig() {
	yamlFile, err := ioutil.ReadFile("configs/config.yml")
	if err != nil {
		log.Fatalf("Error reading config.yml: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &configData)
	if err != nil {
		log.Fatalf("Error parsing config.yml: %v", err)
	}

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
	keys := string(key)
	val, ok := configData[keys]
	if !ok {
		return "", errors.New("configuration value not found [" + keys + "]")
	}

	strVal, ok := val.(string)
	if !ok {
		return "", errors.New("configuration value not a string [" + keys + "]")
	}

	return strVal, nil
}
