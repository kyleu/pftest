// Package seed - Content managed by Project Forge, see [projectforge.md] for details.
package seed

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Seeds []*Seed

func (s Seeds) Get(id uuid.UUID) *Seed {
	return lo.FindOrElse(s, nil, func(x *Seed) bool {
		return x.ID == id
	})
}

func (s Seeds) IDs() []uuid.UUID {
	return lo.Map(s, func(xx *Seed, _ int) uuid.UUID {
		return xx.ID
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

func (s Seeds) GetByID(id uuid.UUID) Seeds {
	return lo.Filter(s, func(xx *Seed, _ int) bool {
		return xx.ID == id
	})
}

func (s Seeds) GetByIDs(ids ...uuid.UUID) Seeds {
	return lo.Filter(s, func(xx *Seed, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (s Seeds) Clone() Seeds {
	return slices.Clone(s)
}
