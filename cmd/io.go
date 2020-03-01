package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ioCmd)
	ioCmd.AddCommand(fileCmd)
	ioCmd.AddCommand(dirCmd)
}

var ioCmd = &cobra.Command{
	Use:   "io",
	Short: "io related commands",
	Long:  `io related commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing io")
	},
}

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "io:file related commands",
	Long:  `io:file related commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing io:file")
	},
}

var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "io:dir related commands",
	Long:  `io:dir related commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing io:dir")
	},
}
