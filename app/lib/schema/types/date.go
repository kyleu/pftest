package types

import "time"

const KeyDate = "date"

type Date struct{}

var _ Type = (*Date)(nil)

func (x *Date) Key() string {
	return KeyDate
}

func (x *Date) String() string {
	return x.Key()
}

func (x *Date) Sortable() bool {
	return true
}

func (x *Date) From(v interface{}) interface{} {
	return invalidInput(x.Key(), x)
}

func (x *Date) Default(string) interface{} {
	return time.Now()
}

func NewDate() *Wrapped {
	return Wrap(&Date{})
}
