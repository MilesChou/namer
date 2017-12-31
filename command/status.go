package command

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/MilesChou/namer/provider"
)

var (
	StatusCommand = cli.Command{
		Name:  "status",
		Usage: "狀態",
		Action: func(c *cli.Context) error {
			t, _ := provider.ParseFile(c.GlobalString("provider"))

			fmt.Println(t)

			return nil
		},
	}
)
