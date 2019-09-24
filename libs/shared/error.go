package shared

type Error interface {
	Code() string
	Msg() string
	Data() interface{}
	Err() error
	Error() string
}
