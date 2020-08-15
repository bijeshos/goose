package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of goose",
	Long:  `Print the version number of goose application`,
	Run: func(cmd *cobra.Command, args []string) {
		zap.S().Infow("goose v0.0.1-alpha")
	},
}
