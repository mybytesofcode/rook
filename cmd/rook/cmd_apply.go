package main

import (
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "apply",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
