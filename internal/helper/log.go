package helper

import (
	"os"

	"go.uber.org/zap"
)

func LogCreate() (*zap.Logger, error) {
	if os.Getenv("DEBUG") == "1" {
		return zap.NewDevelopment(zap.Development())
	}

	return zap.NewProduction()
}
