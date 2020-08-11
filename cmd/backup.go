package cmd

import (
	"fmt"
	"github.com/bijeshos/goose/io/fileutil"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(backup)
	backup.AddCommand(primaryCmd)
	backup.AddCommand(dirCmd)
	primaryCmd.AddCommand(secondaryCmd)
}

var backup = &cobra.Command{
	Use:   "backup",
	Short: "backup related commands",
	Long:  `backup related commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing backup")
	},
}

var primaryCmd = &cobra.Command{
	Use:   "primary",
	Short: "backup:primary related commands",
	Long:  `backup:primary related commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing backup:primary")
	},
}

var secondaryCmd = &cobra.Command{
	Use:   "secondary",
	Short: "backup:secondary:read related commands",
	Long:  `backup:secondary:read related commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing io:file:secondary")
		fileutil.Read("/home/bos/1-bos/tmp/go-test")
	},
}
