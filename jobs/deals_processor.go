package jobs

type DealsProcessor struct {
	Processor
}

func NewDealsProcessor() DealsProcessor {
	return DealsProcessor{}
}

func (r *DealsProcessor) Run() {

	// get the cid of the bucket

	//r.LightNode.Filclient.MakeDeal(context.Background())
}
