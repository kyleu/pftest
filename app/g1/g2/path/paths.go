package path

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

type Paths []*Path

func (p Paths) Get(id uuid.UUID) *Path {
	return lo.FindOrElse(p, nil, func(x *Path) bool {
		return x.ID == id
	})
}

func (p Paths) IDs() []uuid.UUID {
	return lo.Map(p, func(xx *Path, _ int) uuid.UUID {
		return xx.ID
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

func (p Paths) GetByID(id uuid.UUID) Paths {
	return lo.Filter(p, func(xx *Path, _ int) bool {
		return xx.ID == id
	})
}

func (p Paths) GetByIDs(ids ...uuid.UUID) Paths {
	return lo.Filter(p, func(xx *Path, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (p Paths) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), lo.Map(p, func(x *Path, _ int) []string {
		return x.Strings()
	})
}

func (p Paths) Random() *Path {
	return util.RandomElement(p)
}

func (p Paths) Clone() Paths {
	return slices.Clone(p)
}
