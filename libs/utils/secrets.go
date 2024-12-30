package utils

import (
	"os"
	"strings"
	"github.com/joho/godotenv"
)

func GetEnv() map[string]string {
	mergedConfig := make(map[string]string)

	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) == 2 {
			mergedConfig[pair[0]] = pair[1]
		}
	}
	envFile, _ := godotenv.Read(".env")
	for key, value := range envFile {
		if value != "" {
			mergedConfig[key] = value
		} else {
			mergedConfig[key] = os.Getenv(key)
		}
	}
	return mergedConfig
}