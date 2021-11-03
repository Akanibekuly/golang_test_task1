package config

import (
	"os"
	"strconv"
	"strings"
)

func getEnvAsStr(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	valueStr := getEnvAsStr(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func getEnvAsBool(key string, defaultVal bool) bool {
	valStr := getEnvAsStr(key, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

func getEnvAsSlice(key string, defaultVal []string) []string {
	if key == "" {
		return defaultVal
	}

	return strings.Split(key, ",")
}
