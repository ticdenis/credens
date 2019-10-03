package queue

type Publisher interface {
	Publish(message []byte) error
}
