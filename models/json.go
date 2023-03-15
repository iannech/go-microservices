package models

import (
	"encoding/json"
	"io"
)

// using Encoder() is much faster and better compared to Marshall
// as it doesn't have to buffer the output into an in memory slice
// of bytes. This reduces allocations and the overheads of the service
func ToJSON(i interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(i)
}

func FromJSON(i interface{}, r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(i)
}
