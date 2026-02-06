package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/toastsandwich/rsync/cmd/config"
	"github.com/toastsandwich/rsync/rsync"
)

// rsync host:home/toastsandwich/src.txt ./src.txt -p <password>
// rsync config set <host> <username> <password>
var (
	rootCmd = cobra.Command{
		Use:   "rsync",
		Short: "rsync will synchronise",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from rsync")
			rsync.Rsync(rsync.Option{
				Args: args,
			})

		},
	}

	subCommands = []*cobra.Command{
		InitCmd,
		config.Command(),
	}
)

func Root() {
	rootCmd.AddCommand(subCommands...)
	rootCmd.Execute()
}
