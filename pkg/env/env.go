package env

import (
	"os"
	"strconv"
)

func MustGetEnvString(key string, defaultValue string) string {
	foundValue := os.Getenv(key)
	if foundValue == "" {
		return defaultValue
	}

	return foundValue
}

func MustGetEnvInt64(key string, defaultValue int64) int64 {
	foundValue := os.Getenv(key)
	if foundValue == "" {
		return defaultValue
	}

	parsedValue, err := strconv.ParseInt(foundValue, 10, 64)
	if err != nil {
		panic(err)
	}

	return parsedValue
}
