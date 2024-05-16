package encoding

import (
	"encoding/json"
	"reflect"
)

type StructEncoder struct {
	Struct reflect.Type
}

func NewStructEncoder(aStruct reflect.Type) *StructEncoder {
	return &StructEncoder{aStruct}
}

func (s StructEncoder) Encode(v interface{}) ([]byte, error) {
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
