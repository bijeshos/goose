package cmd

import (
	"github.com/bijeshos/goose/dirops"
	"github.com/bijeshos/goose/gooseinit"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func init() {
	rootCmd.AddCommand(moveCmd)
	moveCmd.AddCommand(moveDirCommand)
}

var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "move related commands",
	Long:  `move related commands`,
}

var moveDirCommand = &cobra.Command{
	Use:   "dir",
	Short: "move:dir related commands",
	Long:  `move:dir related commands`,
	Run: func(cmd *cobra.Command, args []string) {
		gooseinit.ZapLogger(viper.GetString("common.log.dir"), viper.GetString("common.log.file"))
		for _, k := range viper.AllKeys() {
			zap.S().Debugw("config", k, viper.GetString(k))
			//fmt.Println(viper.GetString(k))
		}
		zap.S().Infow("executing move:dir")
		srcDir := viper.GetString("move.dir.src")
		targetDir := viper.GetString("move.dir.target")
		dirops.Move(srcDir, targetDir)
		zap.S().Infow("completed execution")
	},
}
