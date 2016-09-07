package fastmvc

type Model interface {
	UnmarshalJSON([]byte) error
	MarshalJSON() ([]byte, error)
	Reset()
}
