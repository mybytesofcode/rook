package template

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	valuespkg "github.com/mybytesofcode/rook/internal/values"
)

func TemplateRender(path string, values valuespkg.Values) ([]string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("can't read template: %w", err)
	}

	var buffer bytes.Buffer
	template, err := template.New("test").Parse(string(file))
	if err != nil {
		return nil, fmt.Errorf("can't create template: %w", err)
	}

	err = template.Execute(&buffer, values)
	if err != nil {
		return nil, fmt.Errorf("can't render template: %w", err)
	}

	return strings.Split(buffer.String(), "\n"), err
}
