package util

import "os"

// GetEnv - get an os Env Var or use fallback value
func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
