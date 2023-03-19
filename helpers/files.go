package helpers

import (
	"os"
	"path/filepath"
	"strings"
)

func ExpandHomePath(path string) string {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}

		return filepath.Join(home, path[1:])
	}
	return path
}
