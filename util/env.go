package util

import "os"

func GetEnvWithDefault(key, defaultValue string) string {
	if value, found := os.LookupEnv(key); found {
		return value
	}

	return defaultValue
}
