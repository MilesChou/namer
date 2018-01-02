package command

import (
	"github.com/urfave/cli"
	"github.com/MilesChou/namer/provider"
)

var (
	InitCommand = cli.Command{
		Name:  "init",
		Usage: "初始化名字倉庫",
		Action: func(c *cli.Context) error {
			err := provider.InitFileDefault(c.GlobalString("provider"))

			return err
		},
	}
)
