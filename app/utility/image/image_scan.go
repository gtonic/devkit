package image

import (
	"context"

	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
)

var scanCommand = &cli.Command{
	Name:  "scan",
	Usage: "scan image vulnerabilies using trivy",

	Flags: []cli.Flag{
		ImageFlag,
	},

	Action: func(c *cli.Context) error {
		image := MustImage(c)
		return runTrivy(c.Context, image)
	},
}

func runTrivy(ctx context.Context, image string) error {
	options := docker.RunOptions{
		Env: map[string]string{},

		Volumes: map[string]string{
			"trivy-cache": "/root/.cache/",
		},
	}

	return docker.RunInteractive(ctx, "aquasec/trivy", options, image)
}