package amqp

type CustomMessage struct {
	Numerator    int
	Denominator int
}

func NewCustomMessage(numerator int, denominator int) *CustomMessage {
	return &CustomMessage{Numerator: numerator, Denominator: denominator}
}
