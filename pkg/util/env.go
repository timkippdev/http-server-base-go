package util

import (
	"os"
	"strconv"
)

func GetEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func GetEnvInt(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		intVal, _ := strconv.Atoi(val)
		return intVal
	}
	return defaultVal
}
