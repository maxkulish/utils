package cjson

import (
	"encoding/json"
	"errors"
)

const QUOTES_BYTE = 34

// CustomFloat64 is a wrapper around float64 that allows for unmarshaling
type CustomFloat64 struct {
	Float64 float64
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (cf *CustomFloat64) UnmarshalJSON(data []byte) error {
	if data[0] == QUOTES_BYTE {
		err := json.Unmarshal(data[1:len(data)-1], &cf.Float64)
		if err != nil {
			return errors.New("CustomFloat64: " + err.Error())
		}
	} else {
		err := json.Unmarshal(data, &cf.Float64)
		if err != nil {
			return errors.New("CustomFloat64: UnmarshalJSON: " + err.Error())
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (cf CustomFloat64) MarshalJSON() ([]byte, error) {
	return json.Marshal(cf.Float64)
}
