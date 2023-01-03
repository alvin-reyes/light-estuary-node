package cmd

import (
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
	"light-estuary-node/core"
	"time"
)

func BucketCmds() []*cli.Command {

	var bucketCmds []*cli.Command
	createBucket := &cli.Command{
		Name:  "create-bucket",
		Usage: "Create a new bucket",
		Action: func(c *cli.Context) error {
			//lightNode, _ := core.NewLightNode(c) // light node now
			name := c.Args().Get(0)
			uuid, err := uuid.NewUUID()
			if err != nil {
				return nil
			}
			bucket := &core.Bucket{
				Name:       name,
				UUID:       uuid.String(),
				Status:     "open",
				Cid:        "",
				Created_at: time.Now(),
				Updated_at: time.Now(),
			}
			lightNode.DB.Create(&bucket)
			return nil
		},
	}

	bucketCmds = append(bucketCmds, createBucket)
	return bucketCmds
}
