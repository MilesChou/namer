package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Namer"
	app.Commands = commands()

	app.Run(os.Args)
}

func commands() []cli.Command {
	return []cli.Command{
		{
			Name:  "generate",
			Usage: "產生假名",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "num",
					Value: "10",
					Usage: "產生數量",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("Hello Generate")
				fmt.Println("Generate " + c.String("num"))

				return nil
			},
		},
		{
			Name:  "status",
			Usage: "狀態",
			Action: func(c *cli.Context) error {
				fmt.Println("Hello Status")

				return nil
			},
		},
	}
}
