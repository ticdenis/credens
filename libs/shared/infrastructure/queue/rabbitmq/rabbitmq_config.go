package rabbitmq

type RabbitMQConfig struct {
	Url       string
	QueueName string
}

func NewRabbitMQConfig(url string, queueName string) *RabbitMQConfig {
	return &RabbitMQConfig{Url: url, QueueName: queueName}
}
