package bash

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"

	templatepkg "github.com/mybytesofcode/rook/internal/template"
	valuespkg "github.com/mybytesofcode/rook/internal/values"
)

func BashPreprocess(prefix string, scriptLines []string, values valuespkg.Values) (string, error) {
	scriptPath := fmt.Sprintf("/tmp/%s.sh", uuid.New().String())
	file, err := os.Create(scriptPath)
	if err != nil {
		return "", fmt.Errorf("can't create script: %w", err)
	}
	defer file.Close()

	for _, scriptLine := range scriptLines {
		trimmedLine := strings.Trim(scriptLine, " ")

		// Split line into macro and arguments
		parts := strings.Split(trimmedLine, " ")
		if len(parts) == 0 {
			continue
		}

		macro := parts[0]

		switch macro {
		case "render":
			if len(parts) > 3 {
				return "", fmt.Errorf("render macro has one or two arguments: %w", err)
			}

			var fromPath string
			var toPath string

			if len(parts) == 2 {
				fromPath = parts[1]
				toPath = parts[1]
			} else {
				fromPath = parts[1]
				toPath = parts[2]
			}

			// Macro replace line
			configPath := strings.Replace(fromPath, "/", "", 1)
			_, err := file.WriteString(fmt.Sprintf("cat > %s << 'EOF'\n", toPath))
			if err != nil {
				return "", err
			}

			configLines, err := templatepkg.TemplateRender(
				fmt.Sprintf("%s/%s", prefix, configPath),
				values,
			)
			if err != nil {
				return "", err
			}

			for _, configLine := range configLines {
				_, err = file.WriteString(configLine + "\n")
				if err != nil {
					return "", fmt.Errorf("can't write line to script: %w", err)
				}
			}

			_, err = file.WriteString("EOF\n\n")
			if err != nil {
				return "", fmt.Errorf("can't write line to script: %w", err)
			}
		case "copy":
			if len(parts) > 3 {
				return "", fmt.Errorf("copy macro has one or two arguments: %w", err)
			}

			var fromPath string
			var toPath string

			if len(parts) == 2 {
				fromPath = parts[1]
				toPath = parts[1]
			} else {
				fromPath = parts[1]
				toPath = parts[2]
			}

			binaryPath := strings.Replace(fromPath, "/", "", 1)
			buffer, err := os.ReadFile(
				fmt.Sprintf("%s/%s", prefix, binaryPath),
			)
			if err != nil {
				return "", err
			}

			counter := 0
			binaryEncoded := []string{}
			for _, char := range hex.EncodeToString(buffer) {
				if counter == 0 {
					binaryEncoded = append(binaryEncoded, "\\x")
				}

				binaryEncoded = append(binaryEncoded, string(char))
				counter += 1

				if counter == 2 {
					counter = 0
				}
			}

			_, err = file.WriteString("echo -n -e '" + strings.Join(binaryEncoded, "") + "' > " + toPath + "")
			if err != nil {
				return "", err
			}
		default:
			// Copy line from src to dst
			if !strings.HasPrefix(trimmedLine, "render /") {
				_, err := file.WriteString(scriptLine + "\n")
				if err != nil {
					return "", fmt.Errorf("can't write line to script: %w", err)
				}

				continue
			}
		}
	}

	return scriptPath, err
}
