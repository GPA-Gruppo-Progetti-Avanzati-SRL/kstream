package streams

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/kstream/v2/streams/encoding"
	"reflect"
)

type Serde struct {
	Key   reflect.Type
	Value reflect.Type
}

func NewSerde(key, value interface{}) *Serde {

	k := reflect.TypeOf(key)
	v := reflect.TypeOf(value)
	return &Serde{k, v}
}

func (c *Serde) KeyEncoder() encoding.Encoder {

	return encoding.NewStructEncoder(c.Key)

}

func (c *Serde) ValueEncoder() encoding.Encoder {

	return encoding.NewStructEncoder(c.Value)

}
