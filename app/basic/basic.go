// Content managed by Project Forge, see [projectforge.md] for details.
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

func New(id uuid.UUID) *Basic {
	return &Basic{ID: id}
}

func Random() *Basic {
	return &Basic{
		ID:      util.UUID(),
		Name:    util.RandomString(12),
		Created: time.Now(),
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
		ret.ID = *retID
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Name, err = m.ParseString("name", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (b *Basic) Clone() *Basic {
	return &Basic{
		ID:      b.ID,
		Name:    b.Name,
		Created: b.Created,
	}
}

func (b *Basic) String() string {
	return b.ID.String()
}

func (b *Basic) WebPath() string {
	return "/basic" + "/" + b.ID.String()
}

func (b *Basic) Diff(bx *Basic) util.Diffs {
	var diffs util.Diffs
	if b.ID != bx.ID {
		diffs = append(diffs, util.NewDiff("id", b.ID.String(), bx.ID.String()))
	}
	if b.Name != bx.Name {
		diffs = append(diffs, util.NewDiff("name", b.Name, bx.Name))
	}
	if b.Created != bx.Created {
		diffs = append(diffs, util.NewDiff("created", fmt.Sprint(b.Created), fmt.Sprint(bx.Created)))
	}
	return diffs
}

func (b *Basic) ToData() []interface{} {
	return []interface{}{b.ID, b.Name, b.Created}
}

type Basics []*Basic
