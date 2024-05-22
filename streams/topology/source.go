package topology

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/kstream/v2/kafka"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/kstream/v2/streams/encoding"
)

type SourceEncoder struct {
	Key, Value encoding.Encoder
}

type Source interface {
	Node
	NodeBuilder
	Encoder() SourceEncoder
	Topic() string
	ShouldCoPartitionedWith(source Source)
	TopicConfigs() kafka.TopicConfig
	CoPartitionedWith() Source
	RePartitionedAs() Source
	AutoCreate() bool
	Internal() bool
	InitialOffset() kafka.Offset
}
