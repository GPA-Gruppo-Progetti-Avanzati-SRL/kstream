package librd

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/kstream/v2/kafka"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/kstream/v2/pkg/errors"
	librdKafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/rs/zerolog/log"
	"strings"
)

type ProducerConfig struct {
	Librd *librdKafka.ConfigMap
	*kafka.ProducerConfig
}

func NewProducerConfig() *ProducerConfig {
	return &ProducerConfig{
		Librd: &librdKafka.ConfigMap{
			`partitioner`: string(PartitionerConsistentMurmur2),
		},
		ProducerConfig: kafka.NewProducerConfig(),
	}
}

func (conf *ProducerConfig) validate() error {
	return nil
}

func (conf *ProducerConfig) setUp() error {
	if err := conf.Librd.SetKey(`client.id`, conf.Id); err != nil {
		panic(err)
	}

	if err := conf.Librd.SetKey(`bootstrap.servers`, strings.Join(conf.BootstrapServers, `,`)); err != nil {
		panic(err)
	}

	if err := conf.Librd.SetKey(`go.logs.channel.enable`, true); err != nil {
		return errors.New(err.Error())
	}

	if err := conf.Librd.SetKey(`log_level`, toLibrdLogLevel("INFO")); err != nil {
		return errors.New(err.Error())
	}

	// Making sure transactional properties are set
	if conf.Transactional.Enabled {
		if err := conf.Librd.SetKey(`enable.idempotence`, true); err != nil {
			panic(err)
		}

		// For transactional producers, delivery success is
		// acknowledged by producer batch commit, so we don't need
		// to listen to individual delivery reports
		if err := conf.Librd.SetKey(`go.delivery.reports`, false); err != nil {
			panic(err)
		}

		//TODO use this to recreate producer transaction timeout scenarios
		//if err := conf.Librd.SetKey(`transaction.timeout.ms`, 2000); err != nil {
		//	panic(err)
		//}

		if err := conf.Librd.SetKey(`transactional.id`, conf.Transactional.Id); err != nil {
			panic(err)
		}

		if err := conf.Librd.SetKey(`max.in.flight.requests.per.connection`, 1); err != nil {
			panic(err)
		}

		if err := conf.Librd.SetKey(`acks`, `all`); err != nil {
			panic(err)
		}
	}

	switch conf.SecurityProtocol {
	case "SSL":
		if conf.SSL.CaLocation != "" {
			if err := conf.Librd.SetKey(SecurityProtocolPropertyName, "SSL"); err != nil {
				panic(err)
			}
			if err := conf.Librd.SetKey(SSLCaLocationPropertyName, conf.SSL.CaLocation); err != nil {
				panic(err)
			}
			if err := conf.Librd.SetKey(EnableSSLCertificateVerificationPropertyName, !conf.SSL.SkipVerify); err != nil {
				return errors.New(err.Error())
			}
		} else {
			if err := conf.Librd.SetKey(EnableSSLCertificateVerificationPropertyName, false); err != nil {
				return errors.New(err.Error())
			}
			log.Info().Str(SecurityProtocolPropertyName, conf.SecurityProtocol).Msg(" ca-location not configured")
		}
	case "SASL_SSL":
		fallthrough
	case "SASL":
		if err := conf.Librd.SetKey(SecurityProtocolPropertyName, "SASL_SSL"); err != nil {
			return errors.New(err.Error())
		}
		if err := conf.Librd.SetKey(SASLMechanismPropertyName, conf.SASL.Mechanisms); err != nil {
			return errors.New(err.Error())
		}
		if err := conf.Librd.SetKey(SASLUsernamePropertyName, conf.SASL.Username); err != nil {
			return errors.New(err.Error())
		}
		if err := conf.Librd.SetKey(SASLPasswordPropertyName, conf.SASL.Password); err != nil {
			return errors.New(err.Error())
		}
		if conf.SASL.CaLocation != "" {
			if err := conf.Librd.SetKey(SSLCaLocationPropertyName, conf.SASL.CaLocation); err != nil {
				return errors.New(err.Error())
			}
			if err := conf.Librd.SetKey(EnableSSLCertificateVerificationPropertyName, !conf.SASL.SkipVerify); err != nil {
				return errors.New(err.Error())
			}
		} else {
			log.Info().Str(SecurityProtocolPropertyName, conf.SecurityProtocol).Msg("ca-location not configured")
			if err := conf.Librd.SetKey(EnableSSLCertificateVerificationPropertyName, false); err != nil {
				return errors.New(err.Error())
			}
		}
	default:
		log.Info().Str(SecurityProtocolPropertyName, conf.SecurityProtocol).Msg(" skipping security-protocol settings")
	}

	return nil
}

func (conf *ProducerConfig) copy() *ProducerConfig {
	librdCopy := librdKafka.ConfigMap{}
	for key, val := range *conf.Librd {
		librdCopy[key] = val
	}

	return &ProducerConfig{
		Librd:          &librdCopy,
		ProducerConfig: conf.ProducerConfig.Copy(),
	}
}
