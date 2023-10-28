// Package basic - Content managed by Project Forge, see [projectforge.md] for details.
package basic

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

type Basics []*Basic

func (b Basics) Get(id uuid.UUID) *Basic {
	return lo.FindOrElse(b, nil, func(x *Basic) bool {
		return x.ID == id
	})
}

func (b Basics) IDs() []uuid.UUID {
	return lo.Map(b, func(xx *Basic, _ int) uuid.UUID {
		return xx.ID
	})
}

func (b Basics) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(b)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(b, func(x *Basic, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (b Basics) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(b)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(b, func(x *Basic, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (b Basics) GetByID(id uuid.UUID) Basics {
	return lo.Filter(b, func(xx *Basic, _ int) bool {
		return xx.ID == id
	})
}

func (b Basics) GetByIDs(ids ...uuid.UUID) Basics {
	return lo.Filter(b, func(xx *Basic, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (b Basics) Random() *Basic {
	if len(b) == 0 {
		return nil
	}
	return b[util.RandomInt(len(b))]
}

func (b Basics) Clone() Basics {
	return slices.Clone(b)
}
