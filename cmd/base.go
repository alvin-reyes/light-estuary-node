package cmd

import (
	"context"
	"light-estuary-node/core"
)

var lightNode *core.LightNode

func init() {
	lightNode, _ = core.NewLightNode(context.Background()) // light node now
}
