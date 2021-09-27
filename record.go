package godb

import "encoding/json"

type Record struct {
	ID      int64           `json:"id"`
	Payload json.RawMessage `json:"payload"`
	_bytes  []byte
}

func (record *Record) bytes() []byte {
	if record._bytes == nil {
		data, err := json.Marshal(record)
		if err != nil {
			panic(err)
		}
		record._bytes = data
	}
	return record._bytes
}
