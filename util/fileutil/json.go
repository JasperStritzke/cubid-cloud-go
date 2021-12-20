package fileutil

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io"
)

const (
	CodeTypeJSON = "json"
	CodeTypeGOB  = "gob"
)

type IEncoder interface {
	Encode(interface{}) error
}

type IDecoder interface {
	Decode(interface{}) error
}

func NewPrettyEncoder(w io.Writer, encodeType string) IEncoder {
	switch encodeType {
	case CodeTypeJSON:
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		return encoder
	case CodeTypeGOB:
		encoder := gob.NewEncoder(w)
		return encoder
	default:
		return nil
	}
}

//NewDecoder is just implemented for uniformed handling
func NewDecoder(r io.Reader, encodeType string) IDecoder {
	switch encodeType {
	case CodeTypeJSON:
		decoder := json.NewDecoder(r)
		return decoder
	case CodeTypeGOB:
		decoder := gob.NewDecoder(r)
		return decoder
	default:
		return nil
	}
}

func MarshalIndent(cfg interface{}, codeType string) ([]byte, error) {
	if codeType == CodeTypeGOB {
		var buf bytes.Buffer
		encoder := gob.NewEncoder(&buf)
		err := encoder.Encode(cfg)

		return buf.Bytes(), err
	}
	return json.MarshalIndent(cfg, "", "  ")
}
