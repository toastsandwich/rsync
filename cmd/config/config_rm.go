package config

import (
	"github.com/spf13/cobra"
	"github.com/toastsandwich/rsync/rsync"
	terr "github.com/toastsandwich/terror"
)

func RemoveConfigCmd() *cobra.Command {
	rmCmd := &cobra.Command{
		Use:  "rm",
		RunE: execRmConfigCmd,
	}
	return rmCmd
}

func execRmConfigCmd(_ *cobra.Command, args []string) error {
	if len(args) != 1 {
		return terr.Newf("alias not provided. use rsync config -h")
	}
	return rsync.RemoveConfig(args[0])
}
