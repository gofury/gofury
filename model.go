package gofury

type Model interface {
	UnmarshalJSON([]byte) error
	MarshalJSON() ([]byte, error)
	Reset()
}
