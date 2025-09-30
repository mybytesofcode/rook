package values

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Values map[string]interface{}

func ValuesRead(path string) (Values, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("can't read values: %w", err)
	}

	result := make(Values)
	err = yaml.Unmarshal(file, result)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshal values: %w", err)
	}

	return result, nil
}
