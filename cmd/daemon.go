package cmd

import (
	"github.com/urfave/cli"
	"os"
	"whypfs-gateway/api"
	"whypfs-gateway/core"
)

func DaemonCmd() {
	// add a command to run API node
	app := &cli.App{
		Name:  "daemon",
		Usage: "A light version of Estuary that allows users to upload and download data from the Filecoin network.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config", // config file
				Usage: "config file",
			},
		},
		Action: func(c *cli.Context) error {

			ln, err := core.NewLightNode(c)
			if err != nil {
				return err
			}

			// launch the API node
			api.InitializeEchoRouterConfig(ln)
			api.LoopForever()
			return nil
		},
	}

	app.Run(os.Args)
}
