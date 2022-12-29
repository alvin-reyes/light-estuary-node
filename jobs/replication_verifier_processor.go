package jobs

// check replication if exists
type ReplicationVerifierProcessor struct {
	Processor
}

func NewReplicationVerifierProcessor() ReplicationVerifierProcessor {
	return ReplicationVerifierProcessor{}
}
