package parser

import (
	"encoding/json"
	"io"
	"strconv"
)

func ParseId(id string) (uint64, error) {
	return strconv.ParseUint(id, 10, 32)
}

func ParseBody(r io.Reader, v interface{}) error {
	if err := json.NewDecoder(r).Decode(v); err != nil {
		return err
	}
	return nil
}
