package encoding

import (
	"encoding/json"
	"errors"
	"reflect"
)

type StructEncoder struct {
	Struct reflect.Type
}

func NewStructEncoder(aStruct reflect.Type) *StructEncoder {
	return &StructEncoder{aStruct}
}

func (s StructEncoder) Encode(v interface{}) ([]byte, error) {
	if reflect.TypeOf(v) != s.Struct {
		return nil, errors.New("invalid struct type")
	}

	return json.Marshal(v)
}

func (s StructEncoder) Decode(data []byte) (interface{}, error) {
	instance := reflect.New(s.Struct).Interface()
	err := json.Unmarshal(data, &instance)
	if err != nil {
		return nil, err
	}
	return instance, nil
}
