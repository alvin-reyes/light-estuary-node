package jobs

// check replication if exists
type ReplicationVerifierProcessor struct {
	Processor
}

func NewReplicationVerifierProcessor() ReplicationVerifierProcessor {
	return ReplicationVerifierProcessor{}
}

func (r *ReplicationVerifierProcessor) Run() {
	// check the content deal table and check replications.
}
