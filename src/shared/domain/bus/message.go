package bus

type Message interface {
	MessageId() string
	MessageType() string
}
