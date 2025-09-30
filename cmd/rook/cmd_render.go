package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	helperpkg "github.com/mybytesofcode/rook/internal/helper"
)

var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "render",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		valuesDenormalized, err := cmd.Flags().GetStringSlice("values")
		if err != nil {
			return err
		}

		values := []string{}
		for _, valueDenormalized := range valuesDenormalized {
			values = append(values, helperpkg.PathNormalize(valueDenormalized))
		}

		scriptsDenormalized, err := cmd.Flags().GetStringSlice("scripts")
		if err != nil {
			return err
		}

		scripts := []string{}
		for _, valueDenormalized := range scriptsDenormalized {
			scripts = append(scripts, helperpkg.PathNormalize(valueDenormalized))
		}

		return renderCmdImpl(values, scripts)
	},
}

func renderCmdImpl(valuesPaths []string, scriptsPaths []string) error {
	tmpPaths, err := Render(valuesPaths, scriptsPaths)
	if err != nil {
		return err
	}

	for _, tmpPath := range tmpPaths {
		buffer, err := os.ReadFile(tmpPath)
		if err != nil {
			return err
		}

		fmt.Println(string(buffer))
	}

	return nil
}
