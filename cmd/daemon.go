package cmd

import (
	"github.com/urfave/cli/v2"
	"light-estuary-node/api"
	"light-estuary-node/core"
)

func DaemonCmd() []*cli.Command {
	// add a command to run API node
	var daemonCommands []*cli.Command

	daemonCmd := &cli.Command{
		Name:  "daemon",
		Usage: "A light version of Estuary that allows users to upload and download data from the Filecoin network.",
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

	// add commands.
	daemonCommands = append(daemonCommands, daemonCmd)

	return daemonCommands

}
