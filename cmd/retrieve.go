package cmd

import "github.com/urfave/cli/v2"

func RetrieveCmds() []*cli.Command {

	var retrieveCommands []*cli.Command

	retrieveFileCmd := &cli.Command{
		Name:  "retrieve-file",
		Usage: "Retrieve a file from the Filecoin network.",
		Action: func(c *cli.Context) error {
			return nil
		},
	}

	retrieveDirCmd := &cli.Command{
		Name:  "retrieve-dir",
		Usage: "Retrieve a directory from the Filecoin network.",
		Action: func(c *cli.Context) error {
			return nil
		},
	}

	retrieveCarCmd := &cli.Command{
		Name:  "retrieve-car",
		Usage: "Retrieve a car file from the Filecoin network.",
		Action: func(c *cli.Context) error {
			return nil
		},
	}

	retrieveCommands = append(retrieveCommands, retrieveFileCmd, retrieveDirCmd, retrieveCarCmd)
	return retrieveCommands
}
