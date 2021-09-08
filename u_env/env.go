package u_env

import (
	"os"
	"strconv"
	"strings"
)

func GetString(key string, defaultValue string) string {
	envValue := os.Getenv(key)
	if strings.TrimSpace(envValue) == "" {
		return defaultValue
	}
	return envValue
}

func GetInt(key string, defaultValue int64) int64 {
	envValue := os.Getenv(key)
	if strings.TrimSpace(envValue) == "" {
		return defaultValue
	}
	intValue, err := strconv.ParseInt(envValue, 10, 64)
	if err != nil {
		return defaultValue
	}
	return intValue
}

func GetFloat(key string, defaultValue float64) float64 {
	envValue := os.Getenv(key)
	if strings.TrimSpace(envValue) == "" {
		return defaultValue
	}
	floatValue, err := strconv.ParseFloat(envValue, 64)
	if err != nil {
		return defaultValue
	}
	return floatValue
}

func GetArray(key string, separator string, defaultValue []string) []string {
	envValue := os.Getenv(key)
	if strings.TrimSpace(envValue) == "" {
		return defaultValue
	}
	return strings.Split(envValue, separator)
}

func IsDev() bool {
	return GetString("ENV", "") == "dev"
}

func IsProd() bool {
	return GetString("ENV", "") == "prod"
}

func IsTest() bool {
	return GetString("TESTING", "") == "yes"
}
