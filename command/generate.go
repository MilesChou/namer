package command

import (
	"fmt"
	"strconv"
	"github.com/urfave/cli"
	"github.com/MilesChou/namer/facade"
)

var (
	GenerateCommand = cli.Command{
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

			return facade.Generate(c.GlobalString("provider"), num, func(item string, index int) {
				fmt.Println(item)
			})
		},
	}
)
