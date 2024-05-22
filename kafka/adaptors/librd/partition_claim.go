package librd

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/kstream/v2/kafka"
)

type partitionClaim struct {
	tp       kafka.TopicPartition
	messages chan kafka.Record
}

func (t *partitionClaim) TopicPartition() kafka.TopicPartition {
	return t.tp
}

func (t *partitionClaim) Records() <-chan kafka.Record {
	return t.messages
}
