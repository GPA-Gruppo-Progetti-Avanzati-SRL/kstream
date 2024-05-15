package kafka

import (
	"github.com/tryfix/log"
	"github.com/tryfix/metrics"
	"go.opentelemetry.io/otel/trace"
)

type ProducerConfig struct {
	Id               string
	BootstrapServers []string
	PartitionerFunc  PartitionerFunc
	Acks             RequiredAcks
	Transactional    struct {
		Enabled bool
		Id      string
	}
	Idempotent      bool
	Logger          log.Logger
	MetricsReporter metrics.Reporter
	TracerProvider  trace.TracerProvider
}

func (conf *ProducerConfig) Copy() *ProducerConfig {
	return &ProducerConfig{
		Id:               conf.Id,
		BootstrapServers: conf.BootstrapServers,
		PartitionerFunc:  conf.PartitionerFunc,
		Acks:             conf.Acks,
		Transactional:    conf.Transactional,
		Idempotent:       conf.Idempotent,
		Logger:           conf.Logger,
		MetricsReporter:  conf.MetricsReporter,
		TracerProvider:   conf.TracerProvider,
	}
}

func NewProducerConfig() *ProducerConfig {
	return &ProducerConfig{
		Acks:            WaitForAll,
		Logger:          log.NewNoopLogger(),
		MetricsReporter: metrics.NoopReporter(),
	}
}
