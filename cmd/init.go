package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize goose",
	Long:  `Initialize goose with options`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("initializing goose")
	},
}
