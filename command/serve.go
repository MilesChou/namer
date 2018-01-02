package command

import (
	"strconv"
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
	server := gin.Default()
	server.GET(`/generate`, func(g *gin.Context) {
		num, err := strconv.Atoi(g.DefaultQuery("amount", "10"))

		if err != nil {
			g.JSON(400, err)
			return
		}

		firstNameNum, err := strconv.Atoi(g.DefaultQuery("firstNameNum", "0"))

		if err != nil {
			g.JSON(400, err)
			return
		}

		facade.GenerateFirstNameNum = firstNameNum
		facade.GenerateGender = g.DefaultQuery("gender", "")

		s := make([]string, num)
		facade.Generate(c.GlobalString("provider"), num, func(item string, index int) {
			s[index] = item
		})

		g.JSON(200, s)
	})

	server.Run()

	return nil
}
