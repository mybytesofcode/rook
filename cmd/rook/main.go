package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rook",
	Short: "rook",
	Long:  "TBD",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func main() {
	// render
	renderCmd.Flags().StringSlice("values", []string{"values.yaml"}, "Path to values.yaml")
	renderCmd.Flags().StringSlice("scripts", []string{"script.sh"}, "Path to script.sh")
	rootCmd.AddCommand(renderCmd)

	// apply local
	applyLocalCmd.Flags().StringSlice("values", []string{"values.yaml"}, "Path to values.yaml")
	applyLocalCmd.Flags().StringSlice("scripts", []string{"script.sh"}, "Path to script.sh")
	applyCmd.AddCommand(applyLocalCmd)

	// apply remote
	applyRemoteCmd.Flags().StringSlice("values", []string{"values.yaml"}, "Path to values.yaml")
	applyRemoteCmd.Flags().StringSlice("scripts", []string{"script.sh"}, "Path to script.sh")
	applyRemoteCmd.Flags().String("ssh.host", "127.0.0.1", "ssh host")
	applyRemoteCmd.Flags().Int("ssh.port", 22, "ssh port")
	applyRemoteCmd.Flags().String("ssh.user", "user", "ssh user")
	applyCmd.AddCommand(applyRemoteCmd)

	// apply
	rootCmd.AddCommand(applyCmd)

	// .
	rootCmd.Execute()
}
