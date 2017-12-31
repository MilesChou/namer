package command

import (
	"fmt"
	"github.com/urfave/cli"
)

var (
	StatusCommand = cli.Command{
		Name:  "status",
		Usage: "狀態",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello Status")

			return nil
		},
	}
)
