package main

import (
	"github.com/bijeshos/goose/cmd"
	"github.com/bijeshos/goose/logwrap"
)

func main() {
	logwrap.Infow("starting goose")
	cmd.Execute()
}
