package json_iterator

import (
	infraError "credens/src/shared/infrastructure/error"
	"github.com/json-iterator/go"
)

type JSONIteratorJSONSerializer struct {
	json jsoniter.API
}

func NewJSONIteratorJSONSerializer() *JSONIteratorJSONSerializer {
	return &JSONIteratorJSONSerializer{json: jsoniter.ConfigCompatibleWithStandardLibrary}
}

func (serializer JSONIteratorJSONSerializer) Serialize(input interface{}) ([]byte, error) {
	data, err := serializer.json.Marshal(input)
	if err != nil {
		err = infraError.NewInfrastructureError("500", "Serialization error", input, err)
	}
	return data, err
}

func (serializer JSONIteratorJSONSerializer) Deserialize(input []byte, data interface{}) error {
	err := serializer.json.Unmarshal(input, data)
	if err != nil {
		err = infraError.NewInfrastructureError("500", "Deserialization error", input, err)
	}
	return err
}
