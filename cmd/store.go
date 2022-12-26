package cmd

import (
	"context"
	"fmt"
	cid2 "github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"
	"light-estuary-node/core"
	"os"
	"time"
)

func StoreCmd() []*cli.Command {

	var storeCommands []*cli.Command

	storeFileCmd := &cli.Command{
		Name:  "store-file",
		Usage: "Store a file on the Filecoin network.",
		Action: func(c *cli.Context) error {
			lightNode, _ := core.NewLightNode(c) // light node now
			value := c.Args().Get(0)
			r, err := os.Open(value)
			if err != nil {
				return nil
			}

			fileNode, err := lightNode.Node.AddPinFile(context.Background(), r, nil)
			size, err := fileNode.Size()
			content := core.Content{
				Name:          r.Name(),
				Size:          int64(size),
				Cid:           fileNode.Cid().String(),
				StagingBucket: "",
				Created_at:    time.Now(),
				Updated_at:    time.Now(),
			}
			lightNode.DB.Create(&content)
			return nil
		},
	}

	storeDirCmd := &cli.Command{
		Name:  "store-dir",
		Usage: "Store a directory on the Filecoin network.",
		Action: func(c *cli.Context) error {
			lightNode, _ := core.NewLightNode(c) // light node now
			valuePath := c.Args().Get(0)
			fileNode, _ := lightNode.Node.AddPinDirectory(context.Background(), valuePath)
			fmt.Println(fileNode.Cid().String())
			return nil
		},
	}

	storeCarCmd := &cli.Command{
		Name:  "store-car",
		Usage: "Store a car file on the Filecoin network.",
		Action: func(c *cli.Context) error {
			lightNode, _ := core.NewLightNode(c) // light node now
			fmt.Println(&lightNode.Node.Host)
			return nil
		},
	}

	storeCidCmd := &cli.Command{
		Name:  "store-cid",
		Usage: "Pull a CID and store a CID on this light estuary node",
		Action: func(c *cli.Context) error {
			lightNode, _ := core.NewLightNode(c) // light node now
			cid, err := cid2.Decode(c.Args().Get(0))
			if err != nil {
				return nil
			}
			fileNode, err := lightNode.Node.Get(context.Background(), cid)
			size, err := fileNode.Size()
			content := core.Content{
				Name:          fileNode.Cid().String(),
				Size:          int64(size),
				Cid:           fileNode.Cid().String(),
				StagingBucket: "",
				Created_at:    time.Now(),
				Updated_at:    time.Now(),
			}
			lightNode.DB.Create(&content)
			return nil
		},
	}

	storeCommands = append(storeCommands, storeFileCmd, storeDirCmd, storeCarCmd, storeCidCmd)
	return storeCommands
}
