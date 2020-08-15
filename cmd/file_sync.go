package cmd

import (
	"github.com/bijeshos/goose/dirutil"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	rootCmd.AddCommand(fileSyncCmd)
	fileSyncCmd.AddCommand(primaryCmd)
	primaryCmd.AddCommand(secondaryCmd)
}

var fileSyncCmd = &cobra.Command{
	Use:   "file-sync",
	Short: "files-sync related commands",
	Long:  `files-sync related commands`,
	Run: func(cmd *cobra.Command, args []string) {
		zap.S().Infow("executing file-sync")
	},
}

var primaryCmd = &cobra.Command{
	Use:   "primary",
	Short: "file-sync:primary related commands",
	Long:  `file-sync:primary related commands`,
	Run: func(cmd *cobra.Command, args []string) {
		zap.S().Infow("executing file-sync:primary")
	},
}

var secondaryCmd = &cobra.Command{
	Use:   "secondary",
	Short: "file-sync:secondary:read related commands",
	Long:  `file-sync:secondary:read related commands`,
	Run: func(cmd *cobra.Command, args []string) {
		zap.S().Infow("executing util:file:secondary")
		dirutil.Read("/home/bos/1-bos/tmp/go-test")
	},
}
