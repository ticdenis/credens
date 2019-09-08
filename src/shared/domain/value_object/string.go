package value_object

type String struct {
	value string
}

func (vo *String) Value() string {
	return vo.value
}
