package main

import (
	"os"
	"github.com/MilesChou/namer/command"
	"github.com/urfave/cli"
)

func main() {
	os.Chdir(".")

	app := cli.NewApp()
	app.Name = "Namer"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "provider",
			Value: "names.yml",
			Usage: "名字倉庫",
		},
	}
	app.Commands = []cli.Command{
		command.GenerateCommand,
		command.InitCommand,
		command.StatusCommand,
		command.ServeCommand,
		command.QueryCommand,
	}

	app.Run(os.Args)
}
