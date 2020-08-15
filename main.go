package main

import (
	"github.com/bijeshos/goose/cmd"
	"github.com/bijeshos/goose/logwrap"
	"go.uber.org/zap"
)

func main() {
	logwrap.InitZapLogger("./.logs", "goose.log")
	zap.S().Infow("starting goose")
	cmd.Execute()
}
