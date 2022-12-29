package jobs

import "light-estuary-node/core"

type Processor struct {
	ProcessorInterface
	LightNode *core.LightNode
}

type ProcessorInterface interface {
	PreProcess()
	PostProcess()
	Run()
	Verify()
}
