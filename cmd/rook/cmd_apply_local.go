package main

import (
	"github.com/spf13/cobra"

	executepkg "github.com/mybytesofcode/rook/internal/execute"
	helperpkg "github.com/mybytesofcode/rook/internal/helper"
)

var applyLocalCmd = &cobra.Command{
	Use:   "local",
	Short: "local",
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

		return applyLocalCmdImpl(values, scripts)
	},
}

func applyLocalCmdImpl(valuesPaths []string, scriptsPaths []string) error {
	tmpPaths, err := Render(valuesPaths, scriptsPaths)
	if err != nil {
		return err
	}

	for _, tmpPath := range tmpPaths {
		err := executepkg.LocalExecute(tmpPath)
		if err != nil {
			return err
		}
	}

	return nil
}
