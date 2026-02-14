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
	rootCmd = &cobra.Command{
		Use:   "rsync",
		Short: "rsync will synchronise",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			err := rsync.Rsync(args[0], args[1])
			if err != nil {
				fmt.Println(err)
				fmt.Println(err.Trace())
				return
			}
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
