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

	storeDirCmd := &cli.Command{
		Name:  "store-dir",
		Usage: "Store a directory on the Filecoin network.",
		Action: func(c *cli.Context) error {
			fmt.Println("store")
			return nil
		},
	}

	storeCarCmd := &cli.Command{
		Name:  "store-car",
		Usage: "Store a car file on the Filecoin network.",
		Action: func(c *cli.Context) error {
			fmt.Println("store")
			return nil
		},
	}

	storeCommands = append(storeCommands, storeFileCmd, storeDirCmd, storeCarCmd)
	return storeCommands
}
