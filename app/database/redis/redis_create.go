package redis

import (
	"fmt"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"

	"github.com/sethvargo/go-password/password"
)

func CreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "create instance",

		Flags: []cli.Flag{
			app.PortFlag,
		},

		Action: func(c *cli.Context) error {
			ctx := c.Context
			image := "redis:6-bullseye"

			target := 6379
			port := app.MustPortOrRandom(c, target)

			password := password.MustGenerate(10, 4, 0, false, false)

			options := docker.RunOptions{
				Labels: map[string]string{
					common.KindKey: Redis,
				},

				Env: map[string]string{
					"REDIS_PASSWORD": password,
				},

				Ports: map[int]int{
					port: target,
				},

				// Volumes: map[string]string{
				// 	name: "/data",
				// },
			}

			if err := docker.Run(ctx, image, options); err != nil {
				return err
			}

			cli.Table([]string{"Name", "Value"}, [][]string{
				{"Host", fmt.Sprintf("localhost:%d", port)},
				{"Password", password},
				{"URL", fmt.Sprintf("redis://:%s@localhost:%d", password, port)},
			})

			return nil
		},
	}
}