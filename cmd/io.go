package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ioCmd)
}

var ioCmd = &cobra.Command{
	Use:   "io",
	Short: "io related commands",
	Long:  `io related commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing io")
	},
}
