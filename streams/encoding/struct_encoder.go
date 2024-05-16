package encoding

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type StructEncoder struct {
	Struct reflect.Type
}

func NewStructEncoder(aStruct reflect.Type) *StructEncoder {
	return &StructEncoder{aStruct}
}

func (s StructEncoder) Encode(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, errors.New("struct is nil")
	}
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t != s.Struct {
		errS := fmt.Sprintf("invalid struct type %s vs %s", t.String(), s.Struct.String())
		return nil, errors.New(errS)
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
