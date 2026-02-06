package config

import (
	"github.com/spf13/cobra"
	"github.com/toastsandwich/rsync/rsync"
	"github.com/toastsandwich/terror"
)

func SetConfigCmd() *cobra.Command {
	setConfigCmd := &cobra.Command{
		Use:   "set",
		Short: "sets configuration to host for rsync",
		RunE:  execSetConfig,
	}
	setConfigCmd.Flags().StringP("alias", "a", "", "set alias")
	setConfigCmd.Flags().StringP("host", "H", "", "set hostname")
	setConfigCmd.Flags().StringP("pass", "p", "", "set password")
	setConfigCmd.Flags().StringP("user", "u", "", "set username")

	setConfigCmd.MarkFlagRequired("alias")
	setConfigCmd.MarkFlagRequired("host")
	setConfigCmd.MarkFlagRequired("user")
	setConfigCmd.MarkFlagRequired("pass")

	return setConfigCmd
}

func execSetConfig(cmd *cobra.Command, args []string) error {
	alias, err := cmd.Flags().GetString("alias")
	if err != nil {
		return terror.Wrap(err, "getting alias from config command")
	}

	host, err := cmd.Flags().GetString("host")
	if err != nil {
		return terror.Wrap(err, "getting hostname from config command")
	}
	user, err := cmd.Flags().GetString("user")
	if err != nil {
		return terror.Wrap(err, "getting username from config command")
	}
	pass, err := cmd.Flags().GetString("pass")
	if err != nil {
		return terror.Wrap(err, "getting password from config command")
	}

	return rsync.SetConfig(rsync.SetConfigOptions{
		Alias:    alias,
		Hostname: host,
		Username: user,
		Password: pass,
	})
}
