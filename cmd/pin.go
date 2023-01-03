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

// call the local APIs instead.
func PinCmd() []*cli.Command {
	var pinCommands []*cli.Command

	pinDealCmd := &cli.Command{
		Name:  "pin-deal",
		Usage: "Pin a file and make a deal",
		Action: func(c *cli.Context) error {
			////lightNode, _ := core.NewLightNode(c) // light node now
			//filePath := c.Args().Get(0)
			//// get the file and pin it.
			//fileNode, err := lightNode.Node.AddPinFile(context.Background(), r, nil)
			//if err != nil {
			//	return err
			//}
			//// make a deal.
			//
			return nil
		},
	}
	pinCmd := &cli.Command{
		Name:  "pin",
		Usage: "Pin a File.",
		Action: func(c *cli.Context) error {
			//lightNode, _ := core.NewLightNode(c) // light node now
			value := c.Args().Get(0)
			r, err := os.Open(value)
			if err != nil {
				return nil
			}

			fileNode, err := lightNode.Node.AddPinFile(context.Background(), r, nil)
			size, err := fileNode.Size()
			content := core.Content{
				Name:       r.Name(),
				Size:       int64(size),
				Cid:        fileNode.Cid().String(),
				Created_at: time.Now(),
				Updated_at: time.Now(),
			}
			lightNode.DB.Create(&content)
			return nil
		},
	}
	pinFileCmd := &cli.Command{
		Name:  "pin-file",
		Usage: "Pin a file on the Filecoin network.",
		Action: func(c *cli.Context) error {
			//lightNode, _ := core.NewLightNode(c) // light node now
			value := c.Args().Get(0)
			r, err := os.Open(value)
			if err != nil {
				return nil
			}

			fileNode, err := lightNode.Node.AddPinFile(context.Background(), r, nil)
			size, err := fileNode.Size()
			content := core.Content{
				Name:       r.Name(),
				Size:       int64(size),
				Cid:        fileNode.Cid().String(),
				Created_at: time.Now(),
				Updated_at: time.Now(),
			}
			lightNode.DB.Create(&content)
			return nil
		},
	}

	pinDirCmd := &cli.Command{
		Name:  "pin-dir",
		Usage: "Pin a directory on the Filecoin network.",
		Action: func(c *cli.Context) error {
			//lightNode, _ := core.NewLightNode(c) // light node now
			valuePath := c.Args().Get(0)
			fileNode, _ := lightNode.Node.AddPinDirectory(context.Background(), valuePath)
			fmt.Println(fileNode.Cid().String())
			return nil
		},
	}

	pinCarCmd := &cli.Command{
		Name:  "pin-car",
		Usage: "Pin a car file on the Filecoin network.",
		Action: func(c *cli.Context) error {
			//lightNode, _ := core.NewLightNode(c) // light node now
			fmt.Println(&lightNode.Node.Host)
			return nil
		},
	}

	pinCidCmd := &cli.Command{
		Name:  "pin-cid",
		Usage: "Pull a CID and store a CID on this light estuary node",
		Action: func(c *cli.Context) error {
			//lightNode, _ := core.NewLightNode(c) // light node now
			cid, err := cid2.Decode(c.Args().Get(0))
			if err != nil {
				return nil
			}
			fileNode, err := lightNode.Node.Get(context.Background(), cid)
			size, err := fileNode.Size()
			content := core.Content{
				Name:       fileNode.Cid().String(),
				Size:       int64(size),
				Cid:        fileNode.Cid().String(),
				Created_at: time.Now(),
				Updated_at: time.Now(),
			}
			lightNode.DB.Create(&content)
			return nil
		},
	}

	pinCommands = append(pinCommands, pinCmd, pinFileCmd, pinDirCmd, pinCarCmd, pinCidCmd, pinDealCmd)
	return pinCommands
}
