package queue

type Consumer interface {
	Consume(key string) ([]byte, error)
}
