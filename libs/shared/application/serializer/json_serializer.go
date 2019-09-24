package serializer

type JSONSerializer interface {
	Serialize(input interface{}) ([]byte, error)
	Deserialize(input []byte, data interface{}) error
}
