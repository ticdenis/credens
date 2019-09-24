package value_object

import (
	uuidV4 "github.com/google/uuid"
)

type UUID struct {
	value string
}

// value nil|string
func NewUuid(value interface{}) *UUID {
	if nil == value {
		uuid, _ := uuidV4.NewRandom()

		return &UUID{uuid.String()}
	} else {
		uuid, _ := uuidV4.Parse(value.(string))

		return &UUID{uuid.String()}
	}
}

func (vo UUID) Value() string {
	return vo.value
}
