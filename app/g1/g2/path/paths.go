// Content managed by Project Forge, see [projectforge.md] for details.
package path

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Paths []*Path

func (p Paths) Get(id uuid.UUID) *Path {
	return lo.FindOrElse(p, nil, func(x *Path) bool {
		return x.ID == id
	})
}

func (p Paths) GetByIDs(ids ...uuid.UUID) Paths {
	return lo.Filter(p, func(x *Path, _ int) bool {
		return lo.Contains(ids, x.ID)
	})
}

func (p Paths) IDs() []uuid.UUID {
	return lo.Map(p, func(x *Path, _ int) uuid.UUID {
		return x.ID
	})
}

func (p Paths) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(p)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(p, func(x *Path, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (p Paths) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(p)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(p, func(x *Path, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (p Paths) Clone() Paths {
	return slices.Clone(p)
}
