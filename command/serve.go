package command

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"github.com/MilesChou/namer/facade"
)

var (
	ServeCommand = cli.Command{
		Name:  "serve",
		Usage: "啟動伺服器",
		Action: func(c *cli.Context) error {
			return serve(c)
		},
	}
)

func serve(c *cli.Context) error {
	num := 10000000

	server := gin.Default()
	server.GET(`/generate`, func(g *gin.Context) {
		s := make([]string, num)

		facade.Generate(c.GlobalString("provider"), num, func(item string, index int) {
			s[index] = item
		})

		g.JSON(200, s)
	})

	server.Run()

	return nil
}
