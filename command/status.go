package command

import (
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
)

var (
	StatusCommand = cli.Command{
		Name:  "status",
		Usage: "狀態",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "provider",
				Value: "names.yml",
				Usage: "名字倉庫",
			},
		},
		Action: func(c *cli.Context) error {
			os.Chdir(".")

			r, _ := ioutil.ReadFile(c.String("provider"))

			fmt.Println(`------ File Content Start ------`)
			fmt.Printf("%s", r)
			fmt.Println(`------- File Content End -------`)

			return nil
		},
	}
)
