// Package basic - Content managed by Project Forge, see [projectforge.md] for details.
package basic

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/util"
)

type Basic struct {
	ID      uuid.UUID `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Status  string    `json:"status,omitempty"`
	Created time.Time `json:"created,omitempty"`
}

func New(id uuid.UUID) *Basic {
	return &Basic{ID: id}
}

func Random() *Basic {
	return &Basic{
		ID:      util.UUID(),
		Name:    util.RandomString(12),
		Status:  util.RandomString(12),
		Created: util.TimeCurrent(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Basic, error) {
	ret := &Basic{}
	var err error
	if setPK {
		retID, e := m.ParseUUID("id", true, true)
		if e != nil {
			return nil, e
		}
		if retID != nil {
			ret.ID = *retID
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Name, err = m.ParseString("name", true, true)
	if err != nil {
		return nil, err
	}
	ret.Status, err = m.ParseString("status", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (b *Basic) Clone() *Basic {
	return &Basic{b.ID, b.Name, b.Status, b.Created}
}

func (b *Basic) String() string {
	return b.ID.String()
}

func (b *Basic) TitleString() string {
	return b.Name
}

func (b *Basic) WebPath() string {
	return "/basic/" + b.ID.String()
}

func (b *Basic) Diff(bx *Basic) util.Diffs {
	var diffs util.Diffs
	if b.ID != bx.ID {
		diffs = append(diffs, util.NewDiff("id", b.ID.String(), bx.ID.String()))
	}
	if b.Name != bx.Name {
		diffs = append(diffs, util.NewDiff("name", b.Name, bx.Name))
	}
	if b.Status != bx.Status {
		diffs = append(diffs, util.NewDiff("status", b.Status, bx.Status))
	}
	if b.Created != bx.Created {
		diffs = append(diffs, util.NewDiff("created", b.Created.String(), bx.Created.String()))
	}
	return diffs
}

func (b *Basic) ToData() []any {
	return []any{b.ID, b.Name, b.Status, b.Created}
}
