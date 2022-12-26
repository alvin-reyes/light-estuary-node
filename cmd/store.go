package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func StoreCmd() []*cli.Command {
	var storeCommands []*cli.Command

	storeFileCmd := &cli.Command{
		Name:  "store-file",
		Usage: "Store a file on the Filecoin network.",
		Action: func(c *cli.Context) error {
			fmt.Println("store")
			return nil
		},
	}

	storeCommands = append(storeCommands, storeFileCmd)
	return storeCommands
}
