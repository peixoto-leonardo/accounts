package env

import (
	"os"
	"strconv"
)

func GetEnvOrDefaultInt(env string, defaultValue int) int {
	value, err := strconv.Atoi(os.Getenv(env))
	if err != nil {
		return defaultValue
	}

	return value
}
