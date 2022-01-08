package softdel

import (
	"time"

	"github.com/kyleu/pftest/app/util"
)

type Softdel struct {
	ID      string     `json:"id"`
	Created time.Time  `json:"created"`
	Updated *time.Time `json:"updated,omitempty"`
	Deleted *time.Time `json:"deleted,omitempty"`
}

func New(id string) *Softdel {
	return &Softdel{ID: id}
}

func FromMap(m util.ValueMap, setPK bool) (*Softdel, error) {
	ret := &Softdel{}
	var err error
	if setPK {
		ret.ID, err = m.ParseString("id", true, true)
		if err != nil {
			return nil, err
		}
	}
	ret.Deleted, err = m.ParseTime("deleted", true, true)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *Softdel) String() string {
	return s.ID
}

func (s *Softdel) WebPath() string {
	return "/softdel" + "/" + s.ID
}

func (s *Softdel) ToData() []interface{} {
	return []interface{}{s.ID, s.Created, s.Updated, s.Deleted}
}

type Softdels []*Softdel
