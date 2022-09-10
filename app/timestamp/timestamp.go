// Content managed by Project Forge, see [projectforge.md] for details.
package timestamp

import (
	"fmt"
	"time"

	"golang.org/x/exp/slices"

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

func Random() *Timestamp {
	return &Timestamp{
		ID:      util.RandomString(12),
		Created: time.Now(),
		Updated: util.NowPointer(),
		Deleted: nil,
	}
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

func (t *Timestamp) Clone() *Timestamp {
	return &Timestamp{
		ID:      t.ID,
		Created: t.Created,
		Updated: t.Updated,
		Deleted: t.Deleted,
	}
}

func (t *Timestamp) String() string {
	return t.ID
}

func (t *Timestamp) TitleString() string {
	return t.String()
}

func (t *Timestamp) WebPath() string {
	return "/timestamp" + "/" + t.ID
}

func (t *Timestamp) Diff(tx *Timestamp) util.Diffs {
	var diffs util.Diffs
	if t.ID != tx.ID {
		diffs = append(diffs, util.NewDiff("id", t.ID, tx.ID))
	}
	if t.Created != tx.Created {
		diffs = append(diffs, util.NewDiff("created", t.Created.String(), tx.Created.String()))
	}
	if (t.Deleted == nil && tx.Deleted != nil) || (t.Deleted != nil && tx.Deleted == nil) || (t.Deleted != nil && tx.Deleted != nil && *t.Deleted != *tx.Deleted) {
		diffs = append(diffs, util.NewDiff("deleted", fmt.Sprint(t.Deleted), fmt.Sprint(tx.Deleted))) //nolint:gocritic // it's nullable
	}
	return diffs
}

func (t *Timestamp) ToData() []any {
	return []any{t.ID, t.Created, t.Updated, t.Deleted}
}

type Timestamps []*Timestamp

func (t Timestamps) Get(id string) *Timestamp {
	for _, x := range t {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (t Timestamps) Clone() Timestamps {
	return slices.Clone(t)
}
