package utils

import (
	"os"
)

// Env 读取环境变量，若
func Env(key string, d string) string {
	value := os.Getenv(key)

	if value == "" {
		return d
	}

	return value
}
