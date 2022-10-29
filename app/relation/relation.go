// Content managed by Project Forge, see [projectforge.md] for details.
package relation

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/util"
)

type Relation struct {
	ID      uuid.UUID `json:"id"`
	BasicID uuid.UUID `json:"basicID"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}

func New(id uuid.UUID) *Relation {
	return &Relation{ID: id}
}

func Random() *Relation {
	return &Relation{
		ID:      util.UUID(),
		BasicID: util.UUID(),
		Name:    util.RandomString(12),
		Created: time.Now(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Relation, error) {
	ret := &Relation{}
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
	retBasicID, e := m.ParseUUID("basicID", true, true)
	if e != nil {
		return nil, e
	}
	if retBasicID != nil {
		ret.BasicID = *retBasicID
	}
	ret.Name, err = m.ParseString("name", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (r *Relation) Clone() *Relation {
	return &Relation{
		ID:      r.ID,
		BasicID: r.BasicID,
		Name:    r.Name,
		Created: r.Created,
	}
}

func (r *Relation) String() string {
	return r.ID.String()
}

func (r *Relation) TitleString() string {
	return r.Name
}

func (r *Relation) WebPath() string {
	return "/relation" + "/" + r.ID.String()
}

func (r *Relation) Diff(rx *Relation) util.Diffs {
	var diffs util.Diffs
	if r.ID != rx.ID {
		diffs = append(diffs, util.NewDiff("id", r.ID.String(), rx.ID.String()))
	}
	if r.BasicID != rx.BasicID {
		diffs = append(diffs, util.NewDiff("basicID", r.BasicID.String(), rx.BasicID.String()))
	}
	if r.Name != rx.Name {
		diffs = append(diffs, util.NewDiff("name", r.Name, rx.Name))
	}
	if r.Created != rx.Created {
		diffs = append(diffs, util.NewDiff("created", r.Created.String(), rx.Created.String()))
	}
	return diffs
}

func (r *Relation) ToData() []any {
	return []any{r.ID, r.BasicID, r.Name, r.Created}
}
