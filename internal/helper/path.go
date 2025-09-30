package helper

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func PathNormalize(path string) string {
	path = strings.Trim(path, " ")

	if len(path) == 0 {
		return path
	}

	if path[0] == '/' {
		return path
	}

	// Expand home path
	if path[0] == '~' {
		u, err := user.Current()
		if err != nil {
			// TODO: log
			return path
		}

		return fmt.Sprintf("%s%s", u.HomeDir, path[1:])
	}

	// Expand relative path
	d, err := os.Getwd()
	if err != nil {
		// TODO: log
		return path
	}

	return fmt.Sprintf("%s/%s", d, path)
}
