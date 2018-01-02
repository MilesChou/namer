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
				Name:  "amount",
				Value: "10",
				Usage: "產生數量",
			},
			cli.StringFlag{
				Name:  "gender",
				Usage: "性別，可用參數： male, female",
			},
			cli.StringFlag{
				Name:  "first-name-num",
				Value: "0",
				Usage: "名的數量， 0 為 1 ~ 2 的亂數，單名給 1 ，雙名給 2",
			},
		},
		Action: func(c *cli.Context) error {
			num, err := strconv.Atoi(c.String("amount"))

			if err != nil {
				return err
			}

			firstNameNum, err := strconv.Atoi(c.String("first-name-num"))

			if err != nil {
				return err
			}

			facade.GenerateFirstNameNum = firstNameNum
			facade.GenerateGender = c.String("gender")

			return facade.Generate(c.GlobalString("provider"), num, func(item string, index int) {
				fmt.Println(item)
			})
		},
	}
)
