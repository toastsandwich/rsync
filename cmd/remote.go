package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/toastsandwich/rsync/utils"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "use to create rsync config dir default path home dir",
	RunE: func(cmd *cobra.Command, args []string) error {
		fullpath, dir, err := utils.ConfigPath()
		if err != nil {
			return err
		}
		if err := os.Mkdir(dir, 0777); err != nil {
			return err
		}
		if _, err := os.Create(fullpath); err != nil {
			return err
		}
		fmt.Println("init done")
		return nil
	},
}
