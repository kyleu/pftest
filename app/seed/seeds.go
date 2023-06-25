// Content managed by Project Forge, see [projectforge.md] for details.
package seed

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Seeds []*Seed

func (s Seeds) Get(id uuid.UUID) *Seed {
	return lo.FindOrElse(s, nil, func(x *Seed) bool {
		return x.ID == id
	})
}

func (s Seeds) GetByIDs(ids ...uuid.UUID) Seeds {
	return lo.Filter(s, func(x *Seed, _ int) bool {
		return lo.Contains(ids, x.ID)
	})
}

func (s Seeds) IDs() []uuid.UUID {
	return lo.Map(s, func(x *Seed, _ int) uuid.UUID {
		return x.ID
	})
}

func (s Seeds) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *Seed, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (s Seeds) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(s, func(x *Seed, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (s Seeds) Clone() Seeds {
	return slices.Clone(s)
}
