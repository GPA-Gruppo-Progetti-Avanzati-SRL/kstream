package librd

const (
	AcksPropertyName                             = "acks"
	AutoOffsetResetPropertyName                  = "auto.offset.reset"
	BootstrapServersPropertyName                 = "bootstrap.servers"
	CommitModeAuto                               = "auto"
	CommitModeManual                             = "manual"
	CommitModeTransaction                        = "tx"
	ConnectionsMaxIdleMs                         = "connections.max.idle.ms"
	DeliveryTimeoutMs                            = "delivery.timeout.ms"
	EnableAutoCommitPropertyName                 = "enable.auto.commit"
	EnablePartitionEOFPropertyName               = "enable.partition.eof"
	EnableSSLCertificateVerificationPropertyName = "enable.ssl.certificate.verification"
	GoApplicationRebalanceEnablePropertyName     = "go.application.rebalance.enable"
	GroupIdPropertyName                          = "group.id"
	HeartBeatIntervalMs                          = "heartbeat.interval.ms"
	IsolationLevelPropertyName                   = "isolation.level"
	LingerMs                                     = "linger.ms"
	MaxPollIntervalMs                            = "max.poll.interval.ms"
	MessageSendMaxRetries                        = "message.send.max.retries"
	MetadataMaxAgeMs                             = "metadata.max.age.ms" // 180000
	MetadataMaxIdleMs                            = "metadata.max.idle.ms"
	RequestTimeoutMs                             = "request.timeout.ms" //60000
	Retries                                      = "retries"
	SASLMechanismPropertyName                    = "sasl.mechanism"
	SASLPasswordPropertyName                     = "sasl.password"
	SASLUsernamePropertyName                     = "sasl.username"
	SSLCaLocationPropertyName                    = "ssl.ca.location"
	SecurityProtocolPropertyName                 = "security.protocol"
	SessionTimeOutMsPropertyName                 = "session.timeout.ms"
	SocketKeepaliveEnable                        = "socket.keepalive.enable" // true
	TransactionalIdPropertyName                  = "transactional.id"
	TransactionalTimeoutMsPropertyName           = "transaction.timeout.ms"
)
