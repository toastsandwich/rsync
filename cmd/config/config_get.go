package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/toastsandwich/rsync/rsync"
	terr "github.com/toastsandwich/terror"
)

func GetConfigCmd() *cobra.Command {
	getConfigCmd := &cobra.Command{
		Use:   "get",
		Short: "gets configuratino to a specific host from rsync config file",
		RunE:  execGetConfig,
	}
	getConfigCmd.Flags().Bool("all", false, "use to get all configs from config file")
	return getConfigCmd
}

func execGetConfig(cmd *cobra.Command, args []string) error {
	all, err := cmd.Flags().GetBool("all")
	if err != nil {
		return err
	}
	if all {
		return execGetAllConfig()
	}

	if len(args) != 1 {
		return terr.Newf("alias not provided. use rsync config -h")
	}

	conf, err := rsync.GetConfig(args[0])
	if err != nil {
		return err
	}
	fmt.Println(conf.String())
	return nil
}

func execGetAllConfig() error {
	confs, err := rsync.GetAllConfig()
	if err != nil {
		return err
	}
	for _, conf := range confs {
		fmt.Println(conf.String())
	}
	return nil
}
