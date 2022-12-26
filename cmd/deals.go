package cmd

import "github.com/urfave/cli/v2"

func DealsCmd() []*cli.Command {
	var dealsCommands []*cli.Command

	dealsListCmd := &cli.Command{
		Name:  "deals",
		Usage: "List all deals.",
		Action: func(c *cli.Context) error {
			return nil
		},
	}

	dealsCommands = append(dealsCommands, dealsListCmd)
	return dealsCommands
}
