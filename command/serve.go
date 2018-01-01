package command

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"github.com/MilesChou/namer/provider"
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
	res, _ := provider.ParseFile(c.GlobalString("provider"))
	num := 10

	server := gin.Default()
	server.GET(`/generate`, func(c *gin.Context) {
		generator := provider.Create()
		generator.Resource = res

		s := []string{}
		for i := 0; i < num; i++ {
			s = append(s, generator.Name())
		}

		provider.Create()
		c.JSON(200, s)
	})

	server.Run()

	return nil
}
