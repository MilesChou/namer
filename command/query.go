package command

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/MilesChou/namer/facade"
)

var (
	QueryCommand = cli.Command{
		Name:  "query",
		Usage: "查詢字典",
		Action: func(c *cli.Context) error {
			str := c.Args().First()

			return facade.Query(str, func(item string) {
				fmt.Println(item)
			})
		},
	}
)
