package cmd

import (
	"fmt"
	"github.com/bijeshos/goose/util/fileutil"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(backup)
	backup.AddCommand(primaryCmd)
	backup.AddCommand(dirCmd)
	primaryCmd.AddCommand(secondaryCmd)
}

var backup = &cobra.Command{
	Use:   "file-sync",
	Short: "files-ync related commands",
	Long:  `files-ync related commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing file-sync")
	},
}

var primaryCmd = &cobra.Command{
	Use:   "primary",
	Short: "file-sync:primary related commands",
	Long:  `file-sync:primary related commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing file-sync:primary")
	},
}

var secondaryCmd = &cobra.Command{
	Use:   "secondary",
	Short: "filesync:secondary:read related commands",
	Long:  `filesync:secondary:read related commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing util:file:secondary")
		fileutil.Read("/home/bos/1-bos/tmp/go-test")
	},
}
