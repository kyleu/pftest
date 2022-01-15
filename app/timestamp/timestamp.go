package timestamp

import (
	"time"

	"github.com/kyleu/pftest/app/util"
)

type Timestamp struct {
	ID      string     `json:"id"`
	Created time.Time  `json:"created"`
	Updated *time.Time `json:"updated,omitempty"`
	Deleted *time.Time `json:"deleted,omitempty"`
}

func New(id string) *Timestamp {
	return &Timestamp{ID: id}
}

func FromMap(m util.ValueMap, setPK bool) (*Timestamp, error) {
	ret := &Timestamp{}
	var err error
	if setPK {
		ret.ID, err = m.ParseString("id", true, true)
		if err != nil {
			return nil, err
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Deleted, err = m.ParseTime("deleted", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (t *Timestamp) String() string {
	return t.ID
}

func (t *Timestamp) WebPath() string {
	return "/timestamp" + "/" + t.ID
}

func (t *Timestamp) ToData() []interface{} {
	return []interface{}{t.ID, t.Created, t.Updated, t.Deleted}
}

type Timestamps []*Timestamp
