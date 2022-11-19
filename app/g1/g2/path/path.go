// Content managed by Project Forge, see [projectforge.md] for details.
package path

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/util"
)

type Path struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Status  string    `json:"status"`
	Created time.Time `json:"created"`
}

func New(id uuid.UUID) *Path {
	return &Path{ID: id}
}

func Random() *Path {
	return &Path{
		ID:      util.UUID(),
		Name:    util.RandomString(12),
		Status:  util.RandomString(12),
		Created: time.Now(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Path, error) {
	ret := &Path{}
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

func (p *Path) Clone() *Path {
	return &Path{
		ID:      p.ID,
		Name:    p.Name,
		Status:  p.Status,
		Created: p.Created,
	}
}

func (p *Path) String() string {
	return p.ID.String()
}

func (p *Path) TitleString() string {
	return p.Name
}

func (p *Path) WebPath() string {
	return "/g1/g2/path/" + p.ID.String()
}

func (p *Path) Diff(px *Path) util.Diffs {
	var diffs util.Diffs
	if p.ID != px.ID {
		diffs = append(diffs, util.NewDiff("id", p.ID.String(), px.ID.String()))
	}
	if p.Name != px.Name {
		diffs = append(diffs, util.NewDiff("name", p.Name, px.Name))
	}
	if p.Status != px.Status {
		diffs = append(diffs, util.NewDiff("status", p.Status, px.Status))
	}
	if p.Created != px.Created {
		diffs = append(diffs, util.NewDiff("created", p.Created.String(), px.Created.String()))
	}
	return diffs
}

func (p *Path) ToData() []any {
	return []any{p.ID, p.Name, p.Status, p.Created}
}
