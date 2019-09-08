package value_object

type String struct {
	value string
}

func NewString(value string) *String {
	return &String{value: value}
}

func (vo *String) Value() string {
	return vo.value
}
