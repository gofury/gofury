package gofury

import "encoding/json"

type Model interface {
	json.Marshaler
	json.Unmarshaler
	Reset()
}
