package queue

type QueueMessage struct {
	Name        string
	MessageId   string
	MessageType string
	Data        interface{}
}

func NewQueueMessage(name string, messageId string, messageType string, data interface{}) *QueueMessage {
	return &QueueMessage{Name: name, MessageId: messageId, MessageType: messageType, Data: data}
}
