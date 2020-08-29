package cmd

import (
	"github.com/bijeshos/goose/dirsync"
	"github.com/bijeshos/goose/logwrap"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func init() {
	rootCmd.AddCommand(syncCmd)
	syncCmd.AddCommand(syncDirCommand)
}

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "sync related commands",
	Long:  `sync related commands`,
}

var syncDirCommand = &cobra.Command{
	Use:   "dir",
	Short: "sync:dir related commands",
	Long:  `sync:dir related commands`,
	Run: func(cmd *cobra.Command, args []string) {
		logwrap.InitZapLogger(viper.GetString("common.log.dir"), viper.GetString("common.log.file"))
		for _, k := range viper.AllKeys() {
			zap.S().Debugw("config", k, viper.GetString(k))
			//fmt.Println(viper.GetString(k))
		}

		zap.S().Infow("executing sync:dir")
		dirsync.Perform(viper.GetString("sync.dir.src"), viper.GetString("sync.dir.target"))
	},
}
