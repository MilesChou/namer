package main

import (
	"os"
	"github.com/MilesChou/namer/command"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Namer"
	app.Commands = []cli.Command{
		command.GenerateCommand,
		command.StatusCommand,
	}

	app.Run(os.Args)
}
