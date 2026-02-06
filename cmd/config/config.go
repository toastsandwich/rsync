package config

import "github.com/spf13/cobra"

var configCmd = cobra.Command{
	Use:   "config",
	Short: "configuration for hosts",
}

func Command() *cobra.Command {
	configCmd.AddCommand(SetConfigCmd())
	configCmd.AddCommand(GetConfigCmd())
	return &configCmd
}
