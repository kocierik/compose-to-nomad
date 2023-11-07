package utils

import (
	"log"
	"os"
	"strings"
)

func CreateOutputDirectory(path string) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}
}

func SplitN(s, sep string, n int) []string {
	return strings.SplitN(s, sep, n)
}

func LoadTemplate(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
