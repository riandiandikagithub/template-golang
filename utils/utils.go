package utils

import (
	"fmt"
	"os"
)

func GetPathNow() (string, error) {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(mydir)
	return mydir, nil
}

// Simple helper function to read an environment or return a default value
func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
