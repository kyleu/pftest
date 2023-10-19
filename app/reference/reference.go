// Package reference - Content managed by Project Forge, see [projectforge.md] for details.
package reference

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/foo"
	"github.com/kyleu/pftest/app/util"
)

type Reference struct {
	ID      uuid.UUID   `json:"id,omitempty"`
	Custom  *foo.Custom `json:"custom,omitempty"`
	Self    *SelfCustom `json:"self,omitempty"`
	Created time.Time   `json:"created,omitempty"`
}

func New(id uuid.UUID) *Reference {
	return &Reference{ID: id}
}

func Random() *Reference {
	return &Reference{
		ID:      util.UUID(),
		Custom:  nil,
		Self:    nil,
		Created: util.TimeCurrent(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Reference, error) {
	ret := &Reference{}
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
	tmpCustom, err := m.ParseString("custom", true, true)
	if err != nil {
		return nil, err
	}
	customArg := &foo.Custom{}
	err = util.FromJSON([]byte(tmpCustom), customArg)
	if err != nil {
		return nil, err
	}
	ret.Custom = customArg
	tmpSelf, err := m.ParseString("self", true, true)
	if err != nil {
		return nil, err
	}
	selfArg := &SelfCustom{}
	err = util.FromJSON([]byte(tmpSelf), selfArg)
	if err != nil {
		return nil, err
	}
	ret.Self = selfArg
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (r *Reference) Clone() *Reference {
	return &Reference{r.ID, r.Custom.Clone(), r.Self.Clone(), r.Created}
}

func (r *Reference) String() string {
	return r.ID.String()
}

func (r *Reference) TitleString() string {
	return r.String()
}

func (r *Reference) WebPath() string {
	return "/reference/" + r.ID.String()
}

func (r *Reference) Diff(rx *Reference) util.Diffs {
	var diffs util.Diffs
	if r.ID != rx.ID {
		diffs = append(diffs, util.NewDiff("id", r.ID.String(), rx.ID.String()))
	}
	diffs = append(diffs, util.DiffObjects(r.Custom, rx.Custom, "custom")...)
	diffs = append(diffs, util.DiffObjects(r.Self, rx.Self, "self")...)
	if r.Created != rx.Created {
		diffs = append(diffs, util.NewDiff("created", r.Created.String(), rx.Created.String()))
	}
	return diffs
}

func (r *Reference) ToData() []any {
	return []any{r.ID, r.Custom, r.Self, r.Created}
}
