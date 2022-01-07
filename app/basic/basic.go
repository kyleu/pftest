package basic

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/util"
)

type Basic struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}

func New(id uuid.UUID, name string) *Basic {
	return &Basic{ID: id, Name: name}
}

func FromMap(m util.ValueMap, setPK bool) (*Basic, error) {
	ret := &Basic{}
	var err error
	if setPK {
		ret.ID, err = m.ParseUUID("id")
		if err != nil {
			return nil, err
		}
		ret.Name, err = m.ParseString("name")
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func (b *Basic) String() string {
	return fmt.Sprintf("%s::%s", b.ID.String(), b.Name)
}

func (b *Basic) WebPath() string {
	return "/basic" + "/" + b.ID.String() + "/" + b.Name
}

func (b *Basic) ToData() []interface{} {
	return []interface{}{b.ID, b.Name, b.Created}
}

type Basics []*Basic
