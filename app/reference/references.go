// Package reference - Content managed by Project Forge, see [projectforge.md] for details.
package reference

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type References []*Reference

func (r References) Get(id uuid.UUID) *Reference {
	return lo.FindOrElse(r, nil, func(x *Reference) bool {
		return x.ID == id
	})
}

func (r References) GetByIDs(ids ...uuid.UUID) References {
	return lo.Filter(r, func(xx *Reference, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (r References) GetByID(id uuid.UUID) References {
	return lo.Filter(r, func(xx *Reference, _ int) bool {
		return xx.ID == id
	})
}

func (r References) IDs() []uuid.UUID {
	return lo.Map(r, func(x *Reference, _ int) uuid.UUID {
		return x.ID
	})
}

func (r References) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(r, func(x *Reference, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (r References) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(r)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(r, func(x *Reference, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (r References) Clone() References {
	return slices.Clone(r)
}
