package util

import (
	"log"
	"os"
	"path/filepath"
)

// GetEnv retrieves a variable from environment or some default value.
func GetEnv(name, defaultValue string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}

	return defaultValue
}

// GetCwd returns current working directory
func GetCwd() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	return dir
}
