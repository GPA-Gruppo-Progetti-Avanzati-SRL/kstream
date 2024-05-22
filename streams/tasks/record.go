package tasks

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/kstream/v2/kafka"
)

type Record struct {
	kafka.Record
	ignore bool
}

func NewTaskRecord(record kafka.Record) *Record {
	return &Record{
		Record: record,
		ignore: false,
	}
}
