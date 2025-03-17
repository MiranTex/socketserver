package utils

import (
	"github.com/google/uuid"
)

func Contains(arr []string, target string) bool {
	for _, value := range arr {

		if value == target {
			return true
		}
	}
	return false
}

func GenerateUUID() string {
	return uuid.New().String()
}
