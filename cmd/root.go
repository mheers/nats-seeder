package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// // Config holds the read config
	// Config *models.Config

	rootCmd = &cobra.Command{
		Use:   "nats-seeder",
		Short: "create seeds for nats",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(mqCmd)
}
