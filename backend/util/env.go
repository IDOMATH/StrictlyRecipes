package util

import (
	"errors"
	"os"
)

func EnvOrDefault(envValue, defValue string) string {
	str := os.Getenv(envValue)
	if str == "" {
		return defValue
	}
	return str
}

func EnvMust(envValue string) (string, error) {
	str := os.Getenv(envValue)
	if str == "" {
		return "", errors.New("env variable not found")
	}
	return str, nil
}
