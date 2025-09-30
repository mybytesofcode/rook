package main

import (
	"github.com/spf13/cobra"

	executepkg "github.com/mybytesofcode/rook/internal/execute"
	helperpkg "github.com/mybytesofcode/rook/internal/helper"
)

var applyRemoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "remote",
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

		host, err := cmd.Flags().GetString("ssh.host")
		if err != nil {
			return err
		}

		port, err := cmd.Flags().GetInt("ssh.port")
		if err != nil {
			return err
		}

		user, err := cmd.Flags().GetString("ssh.user")
		if err != nil {
			return err
		}

		return applyRemoteCmdImpl(values, scripts, host, port, user)
	},
}

func applyRemoteCmdImpl(valuesPaths []string, scriptsPaths []string, host string, port int, user string) error {
	tmpPaths, err := Render(valuesPaths, scriptsPaths)
	if err != nil {
		return err
	}

	for _, tmpPath := range tmpPaths {
		err := executepkg.RemoteExecute(tmpPath, host, port, user)
		if err != nil {
			return err
		}
	}

	return nil
}
