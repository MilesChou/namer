package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/MilesChou/namer/provider"
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
				num, err := strconv.Atoi(c.String("num"))

				if err != nil {
					return err
				}

				fmt.Println("Generate " + strconv.Itoa(num))

				generator := provider.Create()

				for i := 0; i < num; i++ {
					fmt.Println(generator.Name())
				}

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
