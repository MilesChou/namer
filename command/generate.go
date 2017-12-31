package command

import (
	"fmt"
	"strconv"
	"github.com/MilesChou/namer/provider"
	"github.com/urfave/cli"
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
			t, _ := provider.ParseFile(c.GlobalString("provider"))

			fmt.Println(t)

			return generate(c)
		},
	}
)

func generate(c *cli.Context) error {
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
}
