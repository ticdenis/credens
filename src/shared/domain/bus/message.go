package bus

import "credens/src/shared/domain/value_object"

type Message struct {
	MessageId   string
	MessageType string
}

func NewMessage(messageType string) *Message {
	return &Message{MessageId: value_object.NewUuid(nil).Value(), MessageType: messageType}
}
