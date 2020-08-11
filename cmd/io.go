package cmd

import (
	"fmt"
	"github.com/bijeshos/goose/util/fileutil"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ioCmd)
	ioCmd.AddCommand(fileCmd)
	ioCmd.AddCommand(dirCmd)
	fileCmd.AddCommand(fileReadCmd)
}

var ioCmd = &cobra.Command{
	Use:   "util",
	Short: "util related commands",
	Long:  `util related commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing util")
	},
}

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "util:file related commands",
	Long:  `util:file related commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing util:file")
	},
}

var fileReadCmd = &cobra.Command{
	Use:   "read",
	Short: "util:file:read related commands",
	Long:  `util:file:read related commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing util:file:read")
		fileutil.Read("/home/bos/1-bos/tmp/go-test")
	},
}

var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "util:dir related commands",
	Long:  `util:dir related commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing util:dir")
	},
}
