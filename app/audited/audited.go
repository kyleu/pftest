// Content managed by Project Forge, see [projectforge.md] for details.
package audited

import (
	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/util"
)

type Audited struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func New(id uuid.UUID) *Audited {
	return &Audited{ID: id}
}

func Random() *Audited {
	return &Audited{
		ID:   util.UUID(),
		Name: util.RandomString(12),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Audited, error) {
	ret := &Audited{}
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
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (a *Audited) Clone() *Audited {
	return &Audited{
		ID:   a.ID,
		Name: a.Name,
	}
}

func (a *Audited) String() string {
	return a.ID.String()
}

func (a *Audited) TitleString() string {
	return a.Name
}

func (a *Audited) WebPath() string {
	return "/audited/" + a.ID.String()
}

func (a *Audited) Diff(ax *Audited) util.Diffs {
	var diffs util.Diffs
	if a.ID != ax.ID {
		diffs = append(diffs, util.NewDiff("id", a.ID.String(), ax.ID.String()))
	}
	if a.Name != ax.Name {
		diffs = append(diffs, util.NewDiff("name", a.Name, ax.Name))
	}
	return diffs
}

func (a *Audited) ToData() []any {
	return []any{a.ID, a.Name}
}
