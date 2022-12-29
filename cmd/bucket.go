package cmd

import (
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
	"light-estuary-node/core"
	"time"
)

func init() {
	bucketDefault1 := core.CreateDefaultBucket("default-1")
	bucketDefault2 := core.CreateDefaultBucket("default-2")
	bucketDefault3 := core.CreateDefaultBucket("default-3")
	bucketDefault4 := core.CreateDefaultBucket("default-4")
	bucketDefault5 := core.CreateDefaultBucket("default-5")

	lightNode.DB.Create(&bucketDefault1)
	lightNode.DB.Create(&bucketDefault2)
	lightNode.DB.Create(&bucketDefault3)
	lightNode.DB.Create(&bucketDefault4)
	lightNode.DB.Create(&bucketDefault5)
}

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