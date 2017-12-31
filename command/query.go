package command

import (
	"fmt"
	"github.com/MilesChou/namer/provider"
	"github.com/urfave/cli"
)

var (
	QueryCommand = cli.Command{
		Name:  "query",
		Usage: "查詢字典",
		Action: func(c *cli.Context) error {
			return query(c)
		},
	}
)

func query(c *cli.Context) error {
	str := c.Args().First()

	dict, err := provider.Query(str)

	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, heteronym := range dict.Heteronyms {
		for _, definition := range heteronym.Definitions {
			fmt.Println(definition)
		}
	}

	return nil
}
