package cmd

import (
	"fmt"
	"github.com/bijeshos/goose/fileutil"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(fileSyncCmd)
	fileSyncCmd.AddCommand(primaryCmd)
	fileSyncCmd.AddCommand(dirCmd)
	primaryCmd.AddCommand(secondaryCmd)
}

var fileSyncCmd = &cobra.Command{
	Use:   "file-sync",
	Short: "files-sync related commands",
	Long:  `files-sync related commands`,
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
	Short: "file-sync:secondary:read related commands",
	Long:  `file-sync:secondary:read related commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing util:file:secondary")
		fileutil.Read("/home/bos/1-bos/tmp/go-test")
	},
}
