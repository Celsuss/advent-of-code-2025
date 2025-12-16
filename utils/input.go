package utils

import (
	"os"
	"strings"
)

// Note: In Go, functions starting with a Capital Letter are "Exported" (public)
func ReadInput(path string)(string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(content)), nil
}
