// Content managed by Project Forge, see [projectforge.md] for details.
package seed

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/util"
)

type Seed struct {
	ID   uuid.UUID     `json:"id"`
	Name string        `json:"name"`
	Size int           `json:"size"`
	Obj  util.ValueMap `json:"obj"`
}

func New(id uuid.UUID) *Seed {
	return &Seed{ID: id}
}

func Random() *Seed {
	return &Seed{
		ID:   util.UUID(),
		Name: util.RandomString(12),
		Size: util.RandomInt(10000),
		Obj:  util.RandomValueMap(4),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Seed, error) {
	ret := &Seed{}
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
	ret.Size, err = m.ParseInt("size", true, true)
	if err != nil {
		return nil, err
	}
	ret.Obj, err = m.ParseMap("obj", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (s *Seed) Clone() *Seed {
	return &Seed{s.ID, s.Name, s.Size, s.Obj.Clone()}
}

func (s *Seed) String() string {
	return s.ID.String()
}

func (s *Seed) TitleString() string {
	return s.Name
}

func (s *Seed) WebPath() string {
	return "/seed/" + s.ID.String()
}

func (s *Seed) Diff(sx *Seed) util.Diffs {
	var diffs util.Diffs
	if s.ID != sx.ID {
		diffs = append(diffs, util.NewDiff("id", s.ID.String(), sx.ID.String()))
	}
	if s.Name != sx.Name {
		diffs = append(diffs, util.NewDiff("name", s.Name, sx.Name))
	}
	if s.Size != sx.Size {
		diffs = append(diffs, util.NewDiff("size", fmt.Sprint(s.Size), fmt.Sprint(sx.Size)))
	}
	diffs = append(diffs, util.DiffObjects(s.Obj, sx.Obj, "obj")...)
	return diffs
}

func (s *Seed) ToData() []any {
	return []any{s.ID, s.Name, s.Size, s.Obj}
}
