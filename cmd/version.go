package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of goose",
	Long:  `Print the version number of goose application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("goose v0.0.1-alpha")
	},
}
