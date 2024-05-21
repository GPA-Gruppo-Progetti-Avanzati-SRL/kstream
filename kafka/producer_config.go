package kafka

import (
	"github.com/tryfix/log"
	"github.com/tryfix/metrics"
	"go.opentelemetry.io/otel/trace"
)

type ProducerConfig struct {
	Id               string
	BootstrapServers []string
	SecurityProtocol string
	SSL              SSLCfg
	SASL             SaslCfg
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
	DltTopic        string
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
		SSL:              conf.SSL,
		SASL:             conf.SASL,
		MetricsReporter:  conf.MetricsReporter,
		TracerProvider:   conf.TracerProvider,
		SecurityProtocol: conf.SecurityProtocol,
		DltTopic:         conf.DltTopic,
	}
}

func NewProducerConfig() *ProducerConfig {
	return &ProducerConfig{
		Acks:            WaitForAll,
		Logger:          log.NewNoopLogger(),
		MetricsReporter: metrics.NoopReporter(),
	}
}
