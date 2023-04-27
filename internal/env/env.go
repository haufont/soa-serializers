package env

import (
	"os"
	"strconv"
)

func GetStringEnv(env string, def string) string {
	envValue, ok := os.LookupEnv(env)
	if !ok {
		return def
	}
	return envValue
}

func GetIntEnv(env string, def int) int {
	envValueStr, ok := os.LookupEnv(env)
	if !ok {
		return def
	}
	envValueInt, err := strconv.Atoi(envValueStr)
	if err != nil {
		return def
	}
	return envValueInt
}
